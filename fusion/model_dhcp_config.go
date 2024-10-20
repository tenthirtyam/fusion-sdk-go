// © Ryan Johnson. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

package fusion

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the DhcpConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DhcpConfig{}

// DhcpConfig DHCP configuration
type DhcpConfig struct {
	Enabled bool   `json:"enabled"`
	Setting string `json:"setting"`
}

type _DhcpConfig DhcpConfig

// NewDhcpConfig instantiates a new DhcpConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDhcpConfig(enabled bool, setting string) *DhcpConfig {
	this := DhcpConfig{}
	this.Enabled = enabled
	this.Setting = setting
	return &this
}

// NewDhcpConfigWithDefaults instantiates a new DhcpConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDhcpConfigWithDefaults() *DhcpConfig {
	this := DhcpConfig{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *DhcpConfig) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *DhcpConfig) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *DhcpConfig) SetEnabled(v bool) {
	o.Enabled = v
}

// GetSetting returns the Setting field value
func (o *DhcpConfig) GetSetting() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Setting
}

// GetSettingOk returns a tuple with the Setting field value
// and a boolean to check if the value has been set.
func (o *DhcpConfig) GetSettingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Setting, true
}

// SetSetting sets field value
func (o *DhcpConfig) SetSetting(v string) {
	o.Setting = v
}

func (o DhcpConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DhcpConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enabled"] = o.Enabled
	toSerialize["setting"] = o.Setting
	return toSerialize, nil
}

func (o *DhcpConfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"enabled",
		"setting",
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

	varDhcpConfig := _DhcpConfig{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDhcpConfig)

	if err != nil {
		return err
	}

	*o = DhcpConfig(varDhcpConfig)

	return err
}

type NullableDhcpConfig struct {
	value *DhcpConfig
	isSet bool
}

func (v NullableDhcpConfig) Get() *DhcpConfig {
	return v.value
}

func (v *NullableDhcpConfig) Set(val *DhcpConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableDhcpConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableDhcpConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDhcpConfig(val *DhcpConfig) *NullableDhcpConfig {
	return &NullableDhcpConfig{value: val, isSet: true}
}

func (v NullableDhcpConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDhcpConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
