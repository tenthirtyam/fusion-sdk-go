// © Ryan Johnson. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

package fusion

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the WinsConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WinsConfig{}

// WinsConfig WINS configuration
type WinsConfig struct {
	Primary   string `json:"primary"`
	Secondary string `json:"secondary"`
}

type _WinsConfig WinsConfig

// NewWinsConfig instantiates a new WinsConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWinsConfig(primary string, secondary string) *WinsConfig {
	this := WinsConfig{}
	this.Primary = primary
	this.Secondary = secondary
	return &this
}

// NewWinsConfigWithDefaults instantiates a new WinsConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWinsConfigWithDefaults() *WinsConfig {
	this := WinsConfig{}
	return &this
}

// GetPrimary returns the Primary field value
func (o *WinsConfig) GetPrimary() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Primary
}

// GetPrimaryOk returns a tuple with the Primary field value
// and a boolean to check if the value has been set.
func (o *WinsConfig) GetPrimaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Primary, true
}

// SetPrimary sets field value
func (o *WinsConfig) SetPrimary(v string) {
	o.Primary = v
}

// GetSecondary returns the Secondary field value
func (o *WinsConfig) GetSecondary() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Secondary
}

// GetSecondaryOk returns a tuple with the Secondary field value
// and a boolean to check if the value has been set.
func (o *WinsConfig) GetSecondaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Secondary, true
}

// SetSecondary sets field value
func (o *WinsConfig) SetSecondary(v string) {
	o.Secondary = v
}

func (o WinsConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WinsConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["primary"] = o.Primary
	toSerialize["secondary"] = o.Secondary
	return toSerialize, nil
}

func (o *WinsConfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"primary",
		"secondary",
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

	varWinsConfig := _WinsConfig{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWinsConfig)

	if err != nil {
		return err
	}

	*o = WinsConfig(varWinsConfig)

	return err
}

type NullableWinsConfig struct {
	value *WinsConfig
	isSet bool
}

func (v NullableWinsConfig) Get() *WinsConfig {
	return v.value
}

func (v *NullableWinsConfig) Set(val *WinsConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableWinsConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableWinsConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWinsConfig(val *WinsConfig) *NullableWinsConfig {
	return &NullableWinsConfig{value: val, isSet: true}
}

func (v NullableWinsConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWinsConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
