// © Ryan Johnson. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

package fusion

import (
	"encoding/json"
)

// checks if the VMRrgistrationInformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VMRrgistrationInformation{}

// VMRrgistrationInformation struct for VMRrgistrationInformation
type VMRrgistrationInformation struct {
	// Registered VM name id
	Id *string `json:"id,omitempty"`
	// Registered VM path
	Path *string `json:"path,omitempty"`
}

// NewVMRrgistrationInformation instantiates a new VMRrgistrationInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVMRrgistrationInformation() *VMRrgistrationInformation {
	this := VMRrgistrationInformation{}
	return &this
}

// NewVMRrgistrationInformationWithDefaults instantiates a new VMRrgistrationInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVMRrgistrationInformationWithDefaults() *VMRrgistrationInformation {
	this := VMRrgistrationInformation{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VMRrgistrationInformation) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMRrgistrationInformation) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VMRrgistrationInformation) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *VMRrgistrationInformation) SetId(v string) {
	o.Id = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *VMRrgistrationInformation) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMRrgistrationInformation) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *VMRrgistrationInformation) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *VMRrgistrationInformation) SetPath(v string) {
	o.Path = &v
}

func (o VMRrgistrationInformation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VMRrgistrationInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	return toSerialize, nil
}

type NullableVMRrgistrationInformation struct {
	value *VMRrgistrationInformation
	isSet bool
}

func (v NullableVMRrgistrationInformation) Get() *VMRrgistrationInformation {
	return v.value
}

func (v *NullableVMRrgistrationInformation) Set(val *VMRrgistrationInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableVMRrgistrationInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableVMRrgistrationInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVMRrgistrationInformation(val *VMRrgistrationInformation) *NullableVMRrgistrationInformation {
	return &NullableVMRrgistrationInformation{value: val, isSet: true}
}

func (v NullableVMRrgistrationInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVMRrgistrationInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
