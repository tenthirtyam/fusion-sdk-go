// © Ryan Johnson. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

package fusion

import (
	"encoding/json"
)

// checks if the ConfigVMParamsParameter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigVMParamsParameter{}

// ConfigVMParamsParameter struct for ConfigVMParamsParameter
type ConfigVMParamsParameter struct {
	// config params name
	Name *string `json:"name,omitempty"`
	// config params value
	Value *string `json:"value,omitempty"`
}

// NewConfigVMParamsParameter instantiates a new ConfigVMParamsParameter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigVMParamsParameter() *ConfigVMParamsParameter {
	this := ConfigVMParamsParameter{}
	return &this
}

// NewConfigVMParamsParameterWithDefaults instantiates a new ConfigVMParamsParameter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigVMParamsParameterWithDefaults() *ConfigVMParamsParameter {
	this := ConfigVMParamsParameter{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ConfigVMParamsParameter) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigVMParamsParameter) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ConfigVMParamsParameter) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ConfigVMParamsParameter) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ConfigVMParamsParameter) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigVMParamsParameter) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ConfigVMParamsParameter) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ConfigVMParamsParameter) SetValue(v string) {
	o.Value = &v
}

func (o ConfigVMParamsParameter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigVMParamsParameter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableConfigVMParamsParameter struct {
	value *ConfigVMParamsParameter
	isSet bool
}

func (v NullableConfigVMParamsParameter) Get() *ConfigVMParamsParameter {
	return v.value
}

func (v *NullableConfigVMParamsParameter) Set(val *ConfigVMParamsParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigVMParamsParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigVMParamsParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigVMParamsParameter(val *ConfigVMParamsParameter) *NullableConfigVMParamsParameter {
	return &NullableConfigVMParamsParameter{value: val, isSet: true}
}

func (v NullableConfigVMParamsParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigVMParamsParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
