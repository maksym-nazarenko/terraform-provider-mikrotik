package inspect

import (
	"strings"
)

func (us *UnquotedString) UnmarshalMikrotik(value string) error {
	*us = UnquotedString(strings.Trim(value, `"'`))

	return nil
}
