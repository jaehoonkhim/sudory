// Code generated by go-enum DO NOT EDIT.
// Version:
// Revision:
// Build Date:
// Built By:

package v2

import (
	"fmt"
	"strings"
)

const (
	// OnCompletionNone is a OnCompletion of type None.
	OnCompletionNone OnCompletion = iota
	// OnCompletionRemove is a OnCompletion of type Remove.
	OnCompletionRemove
)

const _OnCompletionName = "noneremove"

var _OnCompletionNames = []string{
	_OnCompletionName[0:4],
	_OnCompletionName[4:10],
}

// OnCompletionNames returns a list of possible string values of OnCompletion.
func OnCompletionNames() []string {
	tmp := make([]string, len(_OnCompletionNames))
	copy(tmp, _OnCompletionNames)
	return tmp
}

var _OnCompletionMap = map[OnCompletion]string{
	OnCompletionNone:   _OnCompletionName[0:4],
	OnCompletionRemove: _OnCompletionName[4:10],
}

// String implements the Stringer interface.
func (x OnCompletion) String() string {
	if str, ok := _OnCompletionMap[x]; ok {
		return str
	}
	return fmt.Sprintf("OnCompletion(%d)", x)
}

var _OnCompletionValue = map[string]OnCompletion{
	_OnCompletionName[0:4]:                   OnCompletionNone,
	strings.ToLower(_OnCompletionName[0:4]):  OnCompletionNone,
	_OnCompletionName[4:10]:                  OnCompletionRemove,
	strings.ToLower(_OnCompletionName[4:10]): OnCompletionRemove,
}

// ParseOnCompletion attempts to convert a string to a OnCompletion.
func ParseOnCompletion(name string) (OnCompletion, error) {
	if x, ok := _OnCompletionValue[name]; ok {
		return x, nil
	}
	// Case insensitive parse, do a separate lookup to prevent unnecessary cost of lowercasing a string if we don't need to.
	if x, ok := _OnCompletionValue[strings.ToLower(name)]; ok {
		return x, nil
	}
	return OnCompletion(0), fmt.Errorf("%s is not a valid OnCompletion, try [%s]", name, strings.Join(_OnCompletionNames, ", "))
}

const (
	// ResultTypeNone is a ResultType of type None.
	ResultTypeNone ResultType = iota
	// ResultTypeDatabase is a ResultType of type Database.
	ResultTypeDatabase
	// ResultTypeDigitalOceanSpaces is a ResultType of type DigitalOcean:Spaces.
	ResultTypeDigitalOceanSpaces
)

const _ResultTypeName = "nonedatabaseDigitalOcean:Spaces"

var _ResultTypeNames = []string{
	_ResultTypeName[0:4],
	_ResultTypeName[4:12],
	_ResultTypeName[12:31],
}

// ResultTypeNames returns a list of possible string values of ResultType.
func ResultTypeNames() []string {
	tmp := make([]string, len(_ResultTypeNames))
	copy(tmp, _ResultTypeNames)
	return tmp
}

var _ResultTypeMap = map[ResultType]string{
	ResultTypeNone:               _ResultTypeName[0:4],
	ResultTypeDatabase:           _ResultTypeName[4:12],
	ResultTypeDigitalOceanSpaces: _ResultTypeName[12:31],
}

// String implements the Stringer interface.
func (x ResultType) String() string {
	if str, ok := _ResultTypeMap[x]; ok {
		return str
	}
	return fmt.Sprintf("ResultType(%d)", x)
}

var _ResultTypeValue = map[string]ResultType{
	_ResultTypeName[0:4]:                    ResultTypeNone,
	strings.ToLower(_ResultTypeName[0:4]):   ResultTypeNone,
	_ResultTypeName[4:12]:                   ResultTypeDatabase,
	strings.ToLower(_ResultTypeName[4:12]):  ResultTypeDatabase,
	_ResultTypeName[12:31]:                  ResultTypeDigitalOceanSpaces,
	strings.ToLower(_ResultTypeName[12:31]): ResultTypeDigitalOceanSpaces,
}

// ParseResultType attempts to convert a string to a ResultType.
func ParseResultType(name string) (ResultType, error) {
	if x, ok := _ResultTypeValue[name]; ok {
		return x, nil
	}
	// Case insensitive parse, do a separate lookup to prevent unnecessary cost of lowercasing a string if we don't need to.
	if x, ok := _ResultTypeValue[strings.ToLower(name)]; ok {
		return x, nil
	}
	return ResultType(0), fmt.Errorf("%s is not a valid ResultType, try [%s]", name, strings.Join(_ResultTypeNames, ", "))
}

const (
	// StepStatusRegist is a StepStatus of type Regist.
	StepStatusRegist StepStatus = iota
	// StepStatusSend is a StepStatus of type Send.
	StepStatusSend
	// StepStatusProcessing is a StepStatus of type Processing.
	StepStatusProcessing
	// StepStatusSuccess is a StepStatus of type Success.
	StepStatusSuccess StepStatus = iota + 1
	// StepStatusFail is a StepStatus of type Fail.
	StepStatusFail StepStatus = iota + 4
)

const _StepStatusName = "registsendprocessingsuccessfail"

var _StepStatusNames = []string{
	_StepStatusName[0:6],
	_StepStatusName[6:10],
	_StepStatusName[10:20],
	_StepStatusName[20:27],
	_StepStatusName[27:31],
}

// StepStatusNames returns a list of possible string values of StepStatus.
func StepStatusNames() []string {
	tmp := make([]string, len(_StepStatusNames))
	copy(tmp, _StepStatusNames)
	return tmp
}

var _StepStatusMap = map[StepStatus]string{
	StepStatusRegist:     _StepStatusName[0:6],
	StepStatusSend:       _StepStatusName[6:10],
	StepStatusProcessing: _StepStatusName[10:20],
	StepStatusSuccess:    _StepStatusName[20:27],
	StepStatusFail:       _StepStatusName[27:31],
}

// String implements the Stringer interface.
func (x StepStatus) String() string {
	if str, ok := _StepStatusMap[x]; ok {
		return str
	}
	return fmt.Sprintf("StepStatus(%d)", x)
}

var _StepStatusValue = map[string]StepStatus{
	_StepStatusName[0:6]:                    StepStatusRegist,
	strings.ToLower(_StepStatusName[0:6]):   StepStatusRegist,
	_StepStatusName[6:10]:                   StepStatusSend,
	strings.ToLower(_StepStatusName[6:10]):  StepStatusSend,
	_StepStatusName[10:20]:                  StepStatusProcessing,
	strings.ToLower(_StepStatusName[10:20]): StepStatusProcessing,
	_StepStatusName[20:27]:                  StepStatusSuccess,
	strings.ToLower(_StepStatusName[20:27]): StepStatusSuccess,
	_StepStatusName[27:31]:                  StepStatusFail,
	strings.ToLower(_StepStatusName[27:31]): StepStatusFail,
}

// ParseStepStatus attempts to convert a string to a StepStatus.
func ParseStepStatus(name string) (StepStatus, error) {
	if x, ok := _StepStatusValue[name]; ok {
		return x, nil
	}
	// Case insensitive parse, do a separate lookup to prevent unnecessary cost of lowercasing a string if we don't need to.
	if x, ok := _StepStatusValue[strings.ToLower(name)]; ok {
		return x, nil
	}
	return StepStatus(0), fmt.Errorf("%s is not a valid StepStatus, try [%s]", name, strings.Join(_StepStatusNames, ", "))
}
