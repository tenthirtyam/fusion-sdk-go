// © Ryan Johnson. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

package fusion

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the Networks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Networks{}

// Networks The list of virtual networks
type Networks struct {
	// Number of items
	Num int32 `json:"num"`
	// The list of virtual networks
	Vmnets []Network `json:"vmnets"`
}

type _Networks Networks

// NewNetworks instantiates a new Networks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworks(num int32, vmnets []Network) *Networks {
	this := Networks{}
	this.Num = num
	this.Vmnets = vmnets
	return &this
}

// NewNetworksWithDefaults instantiates a new Networks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksWithDefaults() *Networks {
	this := Networks{}
	var num int32 = 0
	this.Num = num
	return &this
}

// GetNum returns the Num field value
func (o *Networks) GetNum() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Num
}

// GetNumOk returns a tuple with the Num field value
// and a boolean to check if the value has been set.
func (o *Networks) GetNumOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Num, true
}

// SetNum sets field value
func (o *Networks) SetNum(v int32) {
	o.Num = v
}

// GetVmnets returns the Vmnets field value
func (o *Networks) GetVmnets() []Network {
	if o == nil {
		var ret []Network
		return ret
	}

	return o.Vmnets
}

// GetVmnetsOk returns a tuple with the Vmnets field value
// and a boolean to check if the value has been set.
func (o *Networks) GetVmnetsOk() ([]Network, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vmnets, true
}

// SetVmnets sets field value
func (o *Networks) SetVmnets(v []Network) {
	o.Vmnets = v
}

func (o Networks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Networks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["num"] = o.Num
	toSerialize["vmnets"] = o.Vmnets
	return toSerialize, nil
}

func (o *Networks) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"num",
		"vmnets",
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

	varNetworks := _Networks{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNetworks)

	if err != nil {
		return err
	}

	*o = Networks(varNetworks)

	return err
}

type NullableNetworks struct {
	value *Networks
	isSet bool
}

func (v NullableNetworks) Get() *Networks {
	return v.value
}

func (v *NullableNetworks) Set(val *Networks) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworks) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworks(val *Networks) *NullableNetworks {
	return &NullableNetworks{value: val, isSet: true}
}

func (v NullableNetworks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
