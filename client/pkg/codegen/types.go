package codegen

import "strings"

const (
	TypeString  TypeName = "string"
	TypeInt64   TypeName = "int64"
	TypeBool    TypeName = "bool"
	TypeSlice   TypeName = "slice"
	TypePointer TypeName = "pointer"
	TypeUnknown TypeName = "unknown"
)

var (
	// typeIdentMap must contain all valid type identifiers.
	typeIdentMap = map[TypeName]string{
		TypeString:  "string",
		TypeInt64:   "int64",
		TypeBool:    "bool",
		TypeSlice:   "[]",
		TypePointer: "*",
	}
)

type (
	TypeName string

	Type interface {
		// Type returns a type name as string.
		// It must be stable for the same type.
		Type() string

		// Is checks whether two types are the same.
		Is(Type) bool

		// Append appends another type to the types chain.
		Append(Type) Type
	}

	base struct {
		typ   TypeName
		chain []Type
	}
)

func NewType(typ TypeName) Type {
	b := &base{
		typ:   TypeUnknown,
		chain: []Type{},
	}

	if _, ok := typeIdentMap[typ]; ok {
		b.typ = typ
	}

	return b
}

func NewTypeFromString(typ string) Type {
	return NewType(TypeNameFromString(typ))
}

func TypeNameFromString(typ string) TypeName {
	normalized := TypeName(strings.ToLower(typ))
	if _, ok := typeIdentMap[normalized]; ok {
		return normalized
	}

	return TypeUnknown
}

// Type returns a type name as string.
// It must be stable for the same type.
//
// Example:
//
//	Slice.Append(Pointer).Append(String) => []*string
func (b *base) Type() string {
	buf := strings.Builder{}
	buf.WriteString(identityOfType(b.typ))

	var cur Type = b
	for i := range b.chain {
		cur = b.chain[i]
		buf.WriteString(cur.Type())
	}

	return buf.String()
}

func (b *base) Is(t Type) bool {
	return b.Type() == t.Type()
}

func (b *base) Append(t Type) Type {
	b.chain = append(b.chain, t)

	return b
}

func identityOfType(typ TypeName) string {
	if ident, ok := typeIdentMap[typ]; ok {
		return ident
	}

	return string(TypeUnknown)
}
