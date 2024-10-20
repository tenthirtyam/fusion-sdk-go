// © Ryan Johnson. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

package fusion

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the PortforwardGuest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PortforwardGuest{}

// PortforwardGuest struct for PortforwardGuest
type PortforwardGuest struct {
	Ip string `json:"ip"`
	// port of communication
	Port int32 `json:"port"`
}

type _PortforwardGuest PortforwardGuest

// NewPortforwardGuest instantiates a new PortforwardGuest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortforwardGuest(ip string, port int32) *PortforwardGuest {
	this := PortforwardGuest{}
	this.Ip = ip
	this.Port = port
	return &this
}

// NewPortforwardGuestWithDefaults instantiates a new PortforwardGuest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortforwardGuestWithDefaults() *PortforwardGuest {
	this := PortforwardGuest{}
	var port int32 = 0
	this.Port = port
	return &this
}

// GetIp returns the Ip field value
func (o *PortforwardGuest) GetIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ip
}

// GetIpOk returns a tuple with the Ip field value
// and a boolean to check if the value has been set.
func (o *PortforwardGuest) GetIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ip, true
}

// SetIp sets field value
func (o *PortforwardGuest) SetIp(v string) {
	o.Ip = v
}

// GetPort returns the Port field value
func (o *PortforwardGuest) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *PortforwardGuest) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *PortforwardGuest) SetPort(v int32) {
	o.Port = v
}

func (o PortforwardGuest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PortforwardGuest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ip"] = o.Ip
	toSerialize["port"] = o.Port
	return toSerialize, nil
}

func (o *PortforwardGuest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ip",
		"port",
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

	varPortforwardGuest := _PortforwardGuest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPortforwardGuest)

	if err != nil {
		return err
	}

	*o = PortforwardGuest(varPortforwardGuest)

	return err
}

type NullablePortforwardGuest struct {
	value *PortforwardGuest
	isSet bool
}

func (v NullablePortforwardGuest) Get() *PortforwardGuest {
	return v.value
}

func (v *NullablePortforwardGuest) Set(val *PortforwardGuest) {
	v.value = val
	v.isSet = true
}

func (v NullablePortforwardGuest) IsSet() bool {
	return v.isSet
}

func (v *NullablePortforwardGuest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortforwardGuest(val *PortforwardGuest) *NullablePortforwardGuest {
	return &NullablePortforwardGuest{value: val, isSet: true}
}

func (v NullablePortforwardGuest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortforwardGuest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
