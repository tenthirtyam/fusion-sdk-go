// © Ryan Johnson. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

package fusion

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the VMInformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VMInformation{}

// VMInformation struct for VMInformation
type VMInformation struct {
	Id  string `json:"id"`
	Cpu *VMCPU `json:"cpu,omitempty"`
	// Memory size in mega bytes
	Memory *int32 `json:"memory,omitempty"`
}

type _VMInformation VMInformation

// NewVMInformation instantiates a new VMInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVMInformation(id string) *VMInformation {
	this := VMInformation{}
	this.Id = id
	var memory int32 = 512
	this.Memory = &memory
	return &this
}

// NewVMInformationWithDefaults instantiates a new VMInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVMInformationWithDefaults() *VMInformation {
	this := VMInformation{}
	var memory int32 = 512
	this.Memory = &memory
	return &this
}

// GetId returns the Id field value
func (o *VMInformation) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *VMInformation) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *VMInformation) SetId(v string) {
	o.Id = v
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *VMInformation) GetCpu() VMCPU {
	if o == nil || IsNil(o.Cpu) {
		var ret VMCPU
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMInformation) GetCpuOk() (*VMCPU, bool) {
	if o == nil || IsNil(o.Cpu) {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *VMInformation) HasCpu() bool {
	if o != nil && !IsNil(o.Cpu) {
		return true
	}

	return false
}

// SetCpu gets a reference to the given VMCPU and assigns it to the Cpu field.
func (o *VMInformation) SetCpu(v VMCPU) {
	o.Cpu = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *VMInformation) GetMemory() int32 {
	if o == nil || IsNil(o.Memory) {
		var ret int32
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMInformation) GetMemoryOk() (*int32, bool) {
	if o == nil || IsNil(o.Memory) {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *VMInformation) HasMemory() bool {
	if o != nil && !IsNil(o.Memory) {
		return true
	}

	return false
}

// SetMemory gets a reference to the given int32 and assigns it to the Memory field.
func (o *VMInformation) SetMemory(v int32) {
	o.Memory = &v
}

func (o VMInformation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VMInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Cpu) {
		toSerialize["cpu"] = o.Cpu
	}
	if !IsNil(o.Memory) {
		toSerialize["memory"] = o.Memory
	}
	return toSerialize, nil
}

func (o *VMInformation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varVMInformation := _VMInformation{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVMInformation)

	if err != nil {
		return err
	}

	*o = VMInformation(varVMInformation)

	return err
}

type NullableVMInformation struct {
	value *VMInformation
	isSet bool
}

func (v NullableVMInformation) Get() *VMInformation {
	return v.value
}

func (v *NullableVMInformation) Set(val *VMInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableVMInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableVMInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVMInformation(val *VMInformation) *NullableVMInformation {
	return &NullableVMInformation{value: val, isSet: true}
}

func (v NullableVMInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVMInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
