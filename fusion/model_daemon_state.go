// © Ryan Johnson. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

package fusion

import (
	"encoding/json"
	"fmt"
)

// DaemonState the model 'DaemonState'
type DaemonState string

// All allowed values of DaemonState enum
var AllowedDaemonStateEnumValues = []DaemonState{
	"on",
	"off",
}

func (v *DaemonState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DaemonState(value)
	for _, existing := range AllowedDaemonStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DaemonState", value)
}

// NewDaemonStateFromValue returns a pointer to a valid DaemonState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDaemonStateFromValue(v string) (*DaemonState, error) {
	ev := DaemonState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DaemonState: valid values are %v", v, AllowedDaemonStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DaemonState) IsValid() bool {
	for _, existing := range AllowedDaemonStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DaemonState value
func (v DaemonState) Ptr() *DaemonState {
	return &v
}

type NullableDaemonState struct {
	value *DaemonState
	isSet bool
}

func (v NullableDaemonState) Get() *DaemonState {
	return v.value
}

func (v *NullableDaemonState) Set(val *DaemonState) {
	v.value = val
	v.isSet = true
}

func (v NullableDaemonState) IsSet() bool {
	return v.isSet
}

func (v *NullableDaemonState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDaemonState(val *DaemonState) *NullableDaemonState {
	return &NullableDaemonState{value: val, isSet: true}
}

func (v NullableDaemonState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDaemonState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
