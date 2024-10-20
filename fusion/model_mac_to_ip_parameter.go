// © Ryan Johnson. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

package fusion

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the MacToIPParameter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MacToIPParameter{}

// MacToIPParameter struct for MacToIPParameter
type MacToIPParameter struct {
	IP string `json:"IP"`
}

type _MacToIPParameter MacToIPParameter

// NewMacToIPParameter instantiates a new MacToIPParameter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMacToIPParameter(iP string) *MacToIPParameter {
	this := MacToIPParameter{}
	this.IP = iP
	return &this
}

// NewMacToIPParameterWithDefaults instantiates a new MacToIPParameter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMacToIPParameterWithDefaults() *MacToIPParameter {
	this := MacToIPParameter{}
	return &this
}

// GetIP returns the IP field value
func (o *MacToIPParameter) GetIP() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IP
}

// GetIPOk returns a tuple with the IP field value
// and a boolean to check if the value has been set.
func (o *MacToIPParameter) GetIPOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IP, true
}

// SetIP sets field value
func (o *MacToIPParameter) SetIP(v string) {
	o.IP = v
}

func (o MacToIPParameter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MacToIPParameter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["IP"] = o.IP
	return toSerialize, nil
}

func (o *MacToIPParameter) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"IP",
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

	varMacToIPParameter := _MacToIPParameter{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMacToIPParameter)

	if err != nil {
		return err
	}

	*o = MacToIPParameter(varMacToIPParameter)

	return err
}

type NullableMacToIPParameter struct {
	value *MacToIPParameter
	isSet bool
}

func (v NullableMacToIPParameter) Get() *MacToIPParameter {
	return v.value
}

func (v *NullableMacToIPParameter) Set(val *MacToIPParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableMacToIPParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableMacToIPParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMacToIPParameter(val *MacToIPParameter) *NullableMacToIPParameter {
	return &NullableMacToIPParameter{value: val, isSet: true}
}

func (v NullableMacToIPParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMacToIPParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
