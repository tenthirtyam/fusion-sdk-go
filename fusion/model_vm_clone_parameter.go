// © Ryan Johnson. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

package fusion

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the VMCloneParameter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VMCloneParameter{}

// VMCloneParameter struct for VMCloneParameter
type VMCloneParameter struct {
	// New VM name
	Name string `json:"name"`
	// Existing VM ID to clone.
	ParentId string `json:"parentId"`
}

type _VMCloneParameter VMCloneParameter

// NewVMCloneParameter instantiates a new VMCloneParameter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVMCloneParameter(name string, parentId string) *VMCloneParameter {
	this := VMCloneParameter{}
	this.Name = name
	this.ParentId = parentId
	return &this
}

// NewVMCloneParameterWithDefaults instantiates a new VMCloneParameter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVMCloneParameterWithDefaults() *VMCloneParameter {
	this := VMCloneParameter{}
	return &this
}

// GetName returns the Name field value
func (o *VMCloneParameter) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VMCloneParameter) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VMCloneParameter) SetName(v string) {
	o.Name = v
}

// GetParentId returns the ParentId field value
func (o *VMCloneParameter) GetParentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value
// and a boolean to check if the value has been set.
func (o *VMCloneParameter) GetParentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentId, true
}

// SetParentId sets field value
func (o *VMCloneParameter) SetParentId(v string) {
	o.ParentId = v
}

func (o VMCloneParameter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VMCloneParameter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["parentId"] = o.ParentId
	return toSerialize, nil
}

func (o *VMCloneParameter) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"parentId",
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

	varVMCloneParameter := _VMCloneParameter{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVMCloneParameter)

	if err != nil {
		return err
	}

	*o = VMCloneParameter(varVMCloneParameter)

	return err
}

type NullableVMCloneParameter struct {
	value *VMCloneParameter
	isSet bool
}

func (v NullableVMCloneParameter) Get() *VMCloneParameter {
	return v.value
}

func (v *NullableVMCloneParameter) Set(val *VMCloneParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableVMCloneParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableVMCloneParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVMCloneParameter(val *VMCloneParameter) *NullableVMCloneParameter {
	return &NullableVMCloneParameter{value: val, isSet: true}
}

func (v NullableVMCloneParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVMCloneParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
