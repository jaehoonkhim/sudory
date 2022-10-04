// Code generated by go-enum DO NOT EDIT.
// Version:
// Revision:
// Build Date:
// Built By:

package v3

import (
	"fmt"
	"strings"
)

const (
	// FormatTypeDisable is a FormatType of type Disable.
	FormatTypeDisable FormatType = iota
	// FormatTypeFields is a FormatType of type Fields.
	FormatTypeFields
	// FormatTypeJq is a FormatType of type Jq.
	FormatTypeJq
)

const _FormatTypeName = "disablefieldsjq"

var _FormatTypeNames = []string{
	_FormatTypeName[0:7],
	_FormatTypeName[7:13],
	_FormatTypeName[13:15],
}

// FormatTypeNames returns a list of possible string values of FormatType.
func FormatTypeNames() []string {
	tmp := make([]string, len(_FormatTypeNames))
	copy(tmp, _FormatTypeNames)
	return tmp
}

var _FormatTypeMap = map[FormatType]string{
	FormatTypeDisable: _FormatTypeName[0:7],
	FormatTypeFields:  _FormatTypeName[7:13],
	FormatTypeJq:      _FormatTypeName[13:15],
}

// String implements the Stringer interface.
func (x FormatType) String() string {
	if str, ok := _FormatTypeMap[x]; ok {
		return str
	}
	return fmt.Sprintf("FormatType(%d)", x)
}

var _FormatTypeValue = map[string]FormatType{
	_FormatTypeName[0:7]:                    FormatTypeDisable,
	strings.ToLower(_FormatTypeName[0:7]):   FormatTypeDisable,
	_FormatTypeName[7:13]:                   FormatTypeFields,
	strings.ToLower(_FormatTypeName[7:13]):  FormatTypeFields,
	_FormatTypeName[13:15]:                  FormatTypeJq,
	strings.ToLower(_FormatTypeName[13:15]): FormatTypeJq,
}

// ParseFormatType attempts to convert a string to a FormatType.
func ParseFormatType(name string) (FormatType, error) {
	if x, ok := _FormatTypeValue[name]; ok {
		return x, nil
	}
	// Case insensitive parse, do a separate lookup to prevent unnecessary cost of lowercasing a string if we don't need to.
	if x, ok := _FormatTypeValue[strings.ToLower(name)]; ok {
		return x, nil
	}
	return FormatType(0), fmt.Errorf("%s is not a valid FormatType, try [%s]", name, strings.Join(_FormatTypeNames, ", "))
}
