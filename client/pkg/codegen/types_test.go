package codegen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypeIs(t *testing.T) {
	testCases := []struct {
		name     string
		type1    Type
		type2    Type
		expected bool
	}{
		{
			name:     "int==int",
			type1:    NewType(TypeInt64),
			type2:    NewType(TypeInt64),
			expected: true,
		},
		{
			name:  "int==string",
			type1: NewType(TypeInt64),
			type2: NewType(TypeString),
		},
		{
			name:  "bool==unknown",
			type1: NewType(TypeBool),
			type2: NewType(TypeUnknown),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.type1.Is(tc.type2)
			assert.Equal(t, tc.expected, actual, "type1(%q), type2(%q)", tc.type1.Type(), tc.type2.Type())
		})
	}
}

func TestType(t *testing.T) {
	testCases := []struct {
		name     string
		typ      Type
		expected string
	}{
		{
			name:     "int64",
			typ:      NewType(TypeInt64),
			expected: "int64",
		},
		{
			name:     "*int64",
			typ:      NewType(TypePointer).Append(NewType(TypeInt64)),
			expected: "*int64",
		},
		{
			name:     "[]string",
			typ:      NewType(TypeSlice).Append(NewType(TypeString)),
			expected: "[]string",
		},
		{
			name:     "[]*string",
			typ:      NewType(TypeSlice).Append(NewType(TypePointer)).Append(NewType(TypeString)),
			expected: "[]*string",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.typ.Type()
			assert.Equal(t, tc.expected, actual)
		})
	}
}
