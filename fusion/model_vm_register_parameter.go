// © Ryan Johnson. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

package fusion

import (
	"encoding/json"
)

// checks if the VMRegisterParameter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VMRegisterParameter{}

// VMRegisterParameter struct for VMRegisterParameter
type VMRegisterParameter struct {
	// Register VM name
	Name *string `json:"name,omitempty"`
	// Register VM path
	Path *string `json:"path,omitempty"`
}

// NewVMRegisterParameter instantiates a new VMRegisterParameter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVMRegisterParameter() *VMRegisterParameter {
	this := VMRegisterParameter{}
	return &this
}

// NewVMRegisterParameterWithDefaults instantiates a new VMRegisterParameter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVMRegisterParameterWithDefaults() *VMRegisterParameter {
	this := VMRegisterParameter{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VMRegisterParameter) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMRegisterParameter) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VMRegisterParameter) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VMRegisterParameter) SetName(v string) {
	o.Name = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *VMRegisterParameter) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMRegisterParameter) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *VMRegisterParameter) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *VMRegisterParameter) SetPath(v string) {
	o.Path = &v
}

func (o VMRegisterParameter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VMRegisterParameter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	return toSerialize, nil
}

type NullableVMRegisterParameter struct {
	value *VMRegisterParameter
	isSet bool
}

func (v NullableVMRegisterParameter) Get() *VMRegisterParameter {
	return v.value
}

func (v *NullableVMRegisterParameter) Set(val *VMRegisterParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableVMRegisterParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableVMRegisterParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVMRegisterParameter(val *VMRegisterParameter) *NullableVMRegisterParameter {
	return &NullableVMRegisterParameter{value: val, isSet: true}
}

func (v NullableVMRegisterParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVMRegisterParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
