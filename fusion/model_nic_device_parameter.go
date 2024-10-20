// © Ryan Johnson. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

package fusion

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the NICDeviceParameter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NICDeviceParameter{}

// NICDeviceParameter struct for NICDeviceParameter
type NICDeviceParameter struct {
	// The network type of network adapter
	Type string `json:"type"`
	// The vmnet name, it should only be used while type is custom
	Vmnet string `json:"vmnet"`
}

type _NICDeviceParameter NICDeviceParameter

// NewNICDeviceParameter instantiates a new NICDeviceParameter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNICDeviceParameter(type_ string, vmnet string) *NICDeviceParameter {
	this := NICDeviceParameter{}
	this.Type = type_
	this.Vmnet = vmnet
	return &this
}

// NewNICDeviceParameterWithDefaults instantiates a new NICDeviceParameter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNICDeviceParameterWithDefaults() *NICDeviceParameter {
	this := NICDeviceParameter{}
	return &this
}

// GetType returns the Type field value
func (o *NICDeviceParameter) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *NICDeviceParameter) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *NICDeviceParameter) SetType(v string) {
	o.Type = v
}

// GetVmnet returns the Vmnet field value
func (o *NICDeviceParameter) GetVmnet() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Vmnet
}

// GetVmnetOk returns a tuple with the Vmnet field value
// and a boolean to check if the value has been set.
func (o *NICDeviceParameter) GetVmnetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vmnet, true
}

// SetVmnet sets field value
func (o *NICDeviceParameter) SetVmnet(v string) {
	o.Vmnet = v
}

func (o NICDeviceParameter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NICDeviceParameter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["vmnet"] = o.Vmnet
	return toSerialize, nil
}

func (o *NICDeviceParameter) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"vmnet",
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

	varNICDeviceParameter := _NICDeviceParameter{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNICDeviceParameter)

	if err != nil {
		return err
	}

	*o = NICDeviceParameter(varNICDeviceParameter)

	return err
}

type NullableNICDeviceParameter struct {
	value *NICDeviceParameter
	isSet bool
}

func (v NullableNICDeviceParameter) Get() *NICDeviceParameter {
	return v.value
}

func (v *NullableNICDeviceParameter) Set(val *NICDeviceParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableNICDeviceParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableNICDeviceParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNICDeviceParameter(val *NICDeviceParameter) *NullableNICDeviceParameter {
	return &NullableNICDeviceParameter{value: val, isSet: true}
}

func (v NullableNICDeviceParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNICDeviceParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
