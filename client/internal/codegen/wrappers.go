package codegen

import (
	"slices"
	"strings"

	"github.com/ddelnano/terraform-provider-mikrotik/client/pkg/inspect"
)

type (
	argumentWrapper struct {
		arg          *inspect.Argument
		possibleType string
	}

	propertyWrapper struct {
		property *inspect.Property
	}
)

func (aw *argumentWrapper) Name() string {
	return aw.arg.Name
}

func (aw *argumentWrapper) Options() []string {
	return aw.arg.Options
}

func (aw *argumentWrapper) Type() string {
	if aw.possibleType == "" {
		aw.possibleType = aw.guessType()
	}

	return aw.possibleType
}

func (aw *argumentWrapper) guessType() string {
	if len(aw.arg.Options) == 2 && slices.Contains(aw.arg.Options, "no") && slices.Contains(aw.arg.Options, "yes") {
		return "bool"
	}

	if strings.ToLower(aw.arg.Name) == "ttl" {
		return "types.MikrotikDuration"
	}

	if len(aw.arg.Options) == 1 && strings.ToLower(aw.arg.Options[0]) == "none" {
		// return "types.MikrotikNoneAware"
		return "string"
	}

	return "string"
}

func (pw *propertyWrapper) Name() string {
	return pw.property.Name
}

func (pw *propertyWrapper) Type() string {
	return "string"
}

func (pw *propertyWrapper) Readonly() bool {
	return true
}
