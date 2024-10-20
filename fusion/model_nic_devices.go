// © Ryan Johnson. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

package fusion

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the NICDevices type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NICDevices{}

// NICDevices The list of network adapters
type NICDevices struct {
	// Number of NIC devices
	Num int32 `json:"num"`
	// The network adapter added to this VM
	Nics []NICDevice `json:"nics"`
}

type _NICDevices NICDevices

// NewNICDevices instantiates a new NICDevices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNICDevices(num int32, nics []NICDevice) *NICDevices {
	this := NICDevices{}
	this.Num = num
	this.Nics = nics
	return &this
}

// NewNICDevicesWithDefaults instantiates a new NICDevices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNICDevicesWithDefaults() *NICDevices {
	this := NICDevices{}
	var num int32 = 1
	this.Num = num
	return &this
}

// GetNum returns the Num field value
func (o *NICDevices) GetNum() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Num
}

// GetNumOk returns a tuple with the Num field value
// and a boolean to check if the value has been set.
func (o *NICDevices) GetNumOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Num, true
}

// SetNum sets field value
func (o *NICDevices) SetNum(v int32) {
	o.Num = v
}

// GetNics returns the Nics field value
func (o *NICDevices) GetNics() []NICDevice {
	if o == nil {
		var ret []NICDevice
		return ret
	}

	return o.Nics
}

// GetNicsOk returns a tuple with the Nics field value
// and a boolean to check if the value has been set.
func (o *NICDevices) GetNicsOk() ([]NICDevice, bool) {
	if o == nil {
		return nil, false
	}
	return o.Nics, true
}

// SetNics sets field value
func (o *NICDevices) SetNics(v []NICDevice) {
	o.Nics = v
}

func (o NICDevices) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NICDevices) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["num"] = o.Num
	toSerialize["nics"] = o.Nics
	return toSerialize, nil
}

func (o *NICDevices) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"num",
		"nics",
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

	varNICDevices := _NICDevices{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNICDevices)

	if err != nil {
		return err
	}

	*o = NICDevices(varNICDevices)

	return err
}

type NullableNICDevices struct {
	value *NICDevices
	isSet bool
}

func (v NullableNICDevices) Get() *NICDevices {
	return v.value
}

func (v *NullableNICDevices) Set(val *NICDevices) {
	v.value = val
	v.isSet = true
}

func (v NullableNICDevices) IsSet() bool {
	return v.isSet
}

func (v *NullableNICDevices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNICDevices(val *NICDevices) *NullableNICDevices {
	return &NullableNICDevices{value: val, isSet: true}
}

func (v NullableNICDevices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNICDevices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
