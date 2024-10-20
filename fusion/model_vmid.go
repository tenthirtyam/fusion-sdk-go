// © Ryan Johnson. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

package fusion

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the VMID type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VMID{}

// VMID struct for VMID
type VMID struct {
	// ID of the VM
	Id string `json:"id"`
	// Path of the VM
	Path string `json:"path"`
}

type _VMID VMID

// NewVMID instantiates a new VMID object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVMID(id string, path string) *VMID {
	this := VMID{}
	this.Id = id
	this.Path = path
	return &this
}

// NewVMIDWithDefaults instantiates a new VMID object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVMIDWithDefaults() *VMID {
	this := VMID{}
	return &this
}

// GetId returns the Id field value
func (o *VMID) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *VMID) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *VMID) SetId(v string) {
	o.Id = v
}

// GetPath returns the Path field value
func (o *VMID) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *VMID) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *VMID) SetPath(v string) {
	o.Path = v
}

func (o VMID) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VMID) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["path"] = o.Path
	return toSerialize, nil
}

func (o *VMID) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"path",
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

	varVMID := _VMID{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVMID)

	if err != nil {
		return err
	}

	*o = VMID(varVMID)

	return err
}

type NullableVMID struct {
	value *VMID
	isSet bool
}

func (v NullableVMID) Get() *VMID {
	return v.value
}

func (v *NullableVMID) Set(val *VMID) {
	v.value = val
	v.isSet = true
}

func (v NullableVMID) IsSet() bool {
	return v.isSet
}

func (v *NullableVMID) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVMID(val *VMID) *NullableVMID {
	return &NullableVMID{value: val, isSet: true}
}

func (v NullableVMID) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVMID) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
