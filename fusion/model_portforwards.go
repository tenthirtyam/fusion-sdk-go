// © Ryan Johnson. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

package fusion

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the Portforwards type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Portforwards{}

// Portforwards The list of port forwarding
type Portforwards struct {
	// Number of items
	Num int32 `json:"num"`
	// The list of port forwardings
	PortForwardings []Portforward `json:"port_forwardings"`
}

type _Portforwards Portforwards

// NewPortforwards instantiates a new Portforwards object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortforwards(num int32, portForwardings []Portforward) *Portforwards {
	this := Portforwards{}
	this.Num = num
	this.PortForwardings = portForwardings
	return &this
}

// NewPortforwardsWithDefaults instantiates a new Portforwards object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortforwardsWithDefaults() *Portforwards {
	this := Portforwards{}
	var num int32 = 0
	this.Num = num
	return &this
}

// GetNum returns the Num field value
func (o *Portforwards) GetNum() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Num
}

// GetNumOk returns a tuple with the Num field value
// and a boolean to check if the value has been set.
func (o *Portforwards) GetNumOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Num, true
}

// SetNum sets field value
func (o *Portforwards) SetNum(v int32) {
	o.Num = v
}

// GetPortForwardings returns the PortForwardings field value
func (o *Portforwards) GetPortForwardings() []Portforward {
	if o == nil {
		var ret []Portforward
		return ret
	}

	return o.PortForwardings
}

// GetPortForwardingsOk returns a tuple with the PortForwardings field value
// and a boolean to check if the value has been set.
func (o *Portforwards) GetPortForwardingsOk() ([]Portforward, bool) {
	if o == nil {
		return nil, false
	}
	return o.PortForwardings, true
}

// SetPortForwardings sets field value
func (o *Portforwards) SetPortForwardings(v []Portforward) {
	o.PortForwardings = v
}

func (o Portforwards) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Portforwards) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["num"] = o.Num
	toSerialize["port_forwardings"] = o.PortForwardings
	return toSerialize, nil
}

func (o *Portforwards) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"num",
		"port_forwardings",
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

	varPortforwards := _Portforwards{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPortforwards)

	if err != nil {
		return err
	}

	*o = Portforwards(varPortforwards)

	return err
}

type NullablePortforwards struct {
	value *Portforwards
	isSet bool
}

func (v NullablePortforwards) Get() *Portforwards {
	return v.value
}

func (v *NullablePortforwards) Set(val *Portforwards) {
	v.value = val
	v.isSet = true
}

func (v NullablePortforwards) IsSet() bool {
	return v.isSet
}

func (v *NullablePortforwards) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortforwards(val *Portforwards) *NullablePortforwards {
	return &NullablePortforwards{value: val, isSet: true}
}

func (v NullablePortforwards) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortforwards) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
