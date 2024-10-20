// © Ryan Johnson. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

package fusion

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the VMPowerState type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VMPowerState{}

// VMPowerState struct for VMPowerState
type VMPowerState struct {
	PowerState string `json:"power_state"`
}

type _VMPowerState VMPowerState

// NewVMPowerState instantiates a new VMPowerState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVMPowerState(powerState string) *VMPowerState {
	this := VMPowerState{}
	this.PowerState = powerState
	return &this
}

// NewVMPowerStateWithDefaults instantiates a new VMPowerState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVMPowerStateWithDefaults() *VMPowerState {
	this := VMPowerState{}
	return &this
}

// GetPowerState returns the PowerState field value
func (o *VMPowerState) GetPowerState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PowerState
}

// GetPowerStateOk returns a tuple with the PowerState field value
// and a boolean to check if the value has been set.
func (o *VMPowerState) GetPowerStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PowerState, true
}

// SetPowerState sets field value
func (o *VMPowerState) SetPowerState(v string) {
	o.PowerState = v
}

func (o VMPowerState) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VMPowerState) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["power_state"] = o.PowerState
	return toSerialize, nil
}

func (o *VMPowerState) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"power_state",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varVMPowerState := _VMPowerState{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVMPowerState)

	if err != nil {
		return err
	}

	*o = VMPowerState(varVMPowerState)

	return err
}

type NullableVMPowerState struct {
	value *VMPowerState
	isSet bool
}

func (v NullableVMPowerState) Get() *VMPowerState {
	return v.value
}

func (v *NullableVMPowerState) Set(val *VMPowerState) {
	v.value = val
	v.isSet = true
}

func (v NullableVMPowerState) IsSet() bool {
	return v.isSet
}

func (v *NullableVMPowerState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVMPowerState(val *VMPowerState) *NullableVMPowerState {
	return &NullableVMPowerState{value: val, isSet: true}
}

func (v NullableVMPowerState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVMPowerState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
