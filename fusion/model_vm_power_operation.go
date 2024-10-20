// © Ryan Johnson. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

package fusion

import (
	"encoding/json"
	"fmt"
)

// VMPowerOperation the model 'VMPowerOperation'
type VMPowerOperation string

// List of VMPowerOperation
const (
	ON       VMPowerOperation = "on"
	OFF      VMPowerOperation = "off"
	SHUTDOWN VMPowerOperation = "shutdown"
	SUSPEND  VMPowerOperation = "suspend"
	PAUSE    VMPowerOperation = "pause"
	UNPAUSE  VMPowerOperation = "unpause"
)

// All allowed values of VMPowerOperation enum
var AllowedVMPowerOperationEnumValues = []VMPowerOperation{
	"on",
	"off",
	"shutdown",
	"suspend",
	"pause",
	"unpause",
}

func (v *VMPowerOperation) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VMPowerOperation(value)
	for _, existing := range AllowedVMPowerOperationEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VMPowerOperation", value)
}

// NewVMPowerOperationFromValue returns a pointer to a valid VMPowerOperation
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVMPowerOperationFromValue(v string) (*VMPowerOperation, error) {
	ev := VMPowerOperation(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VMPowerOperation: valid values are %v", v, AllowedVMPowerOperationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VMPowerOperation) IsValid() bool {
	for _, existing := range AllowedVMPowerOperationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to VMPowerOperation value
func (v VMPowerOperation) Ptr() *VMPowerOperation {
	return &v
}

type NullableVMPowerOperation struct {
	value *VMPowerOperation
	isSet bool
}

func (v NullableVMPowerOperation) Get() *VMPowerOperation {
	return v.value
}

func (v *NullableVMPowerOperation) Set(val *VMPowerOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableVMPowerOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableVMPowerOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVMPowerOperation(val *VMPowerOperation) *NullableVMPowerOperation {
	return &NullableVMPowerOperation{value: val, isSet: true}
}

func (v NullableVMPowerOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVMPowerOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
