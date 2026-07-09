package codegen

import (
	"errors"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type (
	// Struct holds information about parsed struct.
	Struct struct {
		// Name is a of parsed struct.
		Name string

		// Fields is a collection of field definitions in the parsed struct.
		Fields []*Field
	}

	// Field holds information about particular field in parsed struct.
	Field struct {
		// Name is the struct's field name.
		Name string

		// NameTarget is the target name to use during code generation.
		// This value comes from the struct's field tag.
		NameTarget string

		// // MikrotikName is a field name defined by struct tag (actual MikroTik field name).
		// MikrotikName string

		// Type holds a field type.
		Type Type

		Readonly bool
	}
)

const (
	codegenTagKey = "codegen"

	optReadonly = "readonly"
	optOmit     = "omit"
)

// ParseFile parses a .go file with struct declaration.
//
// This functions searches for struct definition `structName` and parses it.
// If `structName` is empty, function stops at first struct definition in the file right after `startLine`.
func ParseFile(filename string, startLine int, structName string) (*Struct, error) {
	_, err := os.Stat(filename)
	if err != nil {
		return nil, err
	}

	fSet := token.NewFileSet()
	aFile, err := parser.ParseFile(fSet, filename, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	if aFile == nil {
		return nil, errors.New("parsing of the file returned unexpected nil as *ast.File")
	}

	s, err := parse(fSet, aFile, startLine, structName)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func parse(fSet *token.FileSet, node ast.Node, startLine int, structName string) (*Struct, error) {
	structNode, foundName, err := findStruct(fSet, node, startLine, structName)
	if err != nil {
		return nil, err
	}

	parsedStruct, err := parseStruct(structNode)
	if err != nil {
		return nil, err
	}
	parsedStruct.Name = foundName

	return parsedStruct, nil
}

func findStruct(fSet *token.FileSet, node ast.Node, startLine int, structName string) (*ast.StructType, string, error) {
	var foundName string
	var structNode *ast.StructType

	ast.Inspect(node, func(n ast.Node) bool {
		if n == nil {
			return true
		}
		if n.Pos().IsValid() {
			pos := fSet.Position(n.Pos())
			if pos.Line < startLine {
				return true
			}
		}
		typeSpec, ok := n.(*ast.TypeSpec)
		if !ok {
			return true
		}
		if typeSpec.Type == nil {
			return true
		}
		// if struct name is provided, ignore other structs on the way
		if structName != "" && typeSpec.Name.Name != structName {
			return true
		}

		foundName = typeSpec.Name.Name
		t, ok := typeSpec.Type.(*ast.StructType)
		if !ok {
			return true
		}

		structNode = t

		// stop after first struct is found
		return false
	})

	if foundName == "" {
		return nil, "", errors.New("struct not found")
	}
	return structNode, foundName, nil
}

func parseStruct(structNode *ast.StructType) (*Struct, error) {
	result := &Struct{}

	for _, astField := range structNode.Fields.List {
		if astField.Tag == nil {
			continue
		}

		// always unquote tag literal, otherwise it is treated as '`key:"options,here"`'
		unquoted, err := strconv.Unquote(astField.Tag.Value)
		if err != nil {
			return nil, err
		}
		tag := reflect.StructTag(unquoted)
		tagKey := codegenTagKey
		tagValue, ok := tag.Lookup(tagKey)
		if !ok {
			continue
		}
		parts := strings.Split(tagValue, ",")
		name, opts := parts[0], parts[1:]

		if name == "-" {
			continue
		}

		// determine the type of the field
		fieldType := NewType(TypeUnknown)
		if exp, ok := astField.Type.(*ast.SelectorExpr); ok {
			// selector expression when type comes from another package, e.g. types.MikrotikList
			fieldType = NewType(TypeNameFromString(exp.Sel.Name))
		}
		if exp, ok := astField.Type.(*ast.Ident); ok {
			// identifier, when it is a builtin type, e.g. "string"
			fieldType = NewType(TypeNameFromString(exp.Name))
		}
		if exp, ok := astField.Type.(*ast.ArrayType); ok {
			fieldType = NewType(TypeSlice)
			switch el := exp.Elt.(type) {
			case *ast.Ident:
				fieldType.Append(NewTypeFromString(el.Name))
			case *ast.StarExpr:
				fieldType.Append(NewType(TypePointer))
				if ident, ok := el.X.(*ast.Ident); ok {
					fieldType.Append(NewTypeFromString(ident.Name))
				}
			}
		}
		field := Field{
			Name:       astField.Names[0].Name,
			NameTarget: name,
			// MikrotikName: name,
			Type: fieldType,
		}
		omit := false
		for _, o := range opts {
			switch {
			case o == optReadonly:
				field.Readonly = true
			case o == optOmit:
				omit = true
			}
		}
		if omit {
			continue
		}

		result.Fields = append(result.Fields, &field)
	}

	return result, nil
}
