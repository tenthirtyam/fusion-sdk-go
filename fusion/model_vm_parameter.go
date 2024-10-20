// © Ryan Johnson. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

package fusion

import (
	"encoding/json"
)

// checks if the VMParameter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VMParameter{}

// VMParameter struct for VMParameter
type VMParameter struct {
	// Number of processor cores
	Processors *int32 `json:"processors,omitempty"`
	// Memory size in mega bytes
	Memory *int32 `json:"memory,omitempty"`
}

// NewVMParameter instantiates a new VMParameter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVMParameter() *VMParameter {
	this := VMParameter{}
	var processors int32 = 1
	this.Processors = &processors
	var memory int32 = 512
	this.Memory = &memory
	return &this
}

// NewVMParameterWithDefaults instantiates a new VMParameter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVMParameterWithDefaults() *VMParameter {
	this := VMParameter{}
	var processors int32 = 1
	this.Processors = &processors
	var memory int32 = 512
	this.Memory = &memory
	return &this
}

// GetProcessors returns the Processors field value if set, zero value otherwise.
func (o *VMParameter) GetProcessors() int32 {
	if o == nil || IsNil(o.Processors) {
		var ret int32
		return ret
	}
	return *o.Processors
}

// GetProcessorsOk returns a tuple with the Processors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMParameter) GetProcessorsOk() (*int32, bool) {
	if o == nil || IsNil(o.Processors) {
		return nil, false
	}
	return o.Processors, true
}

// HasProcessors returns a boolean if a field has been set.
func (o *VMParameter) HasProcessors() bool {
	if o != nil && !IsNil(o.Processors) {
		return true
	}

	return false
}

// SetProcessors gets a reference to the given int32 and assigns it to the Processors field.
func (o *VMParameter) SetProcessors(v int32) {
	o.Processors = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *VMParameter) GetMemory() int32 {
	if o == nil || IsNil(o.Memory) {
		var ret int32
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMParameter) GetMemoryOk() (*int32, bool) {
	if o == nil || IsNil(o.Memory) {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *VMParameter) HasMemory() bool {
	if o != nil && !IsNil(o.Memory) {
		return true
	}

	return false
}

// SetMemory gets a reference to the given int32 and assigns it to the Memory field.
func (o *VMParameter) SetMemory(v int32) {
	o.Memory = &v
}

func (o VMParameter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VMParameter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Processors) {
		toSerialize["processors"] = o.Processors
	}
	if !IsNil(o.Memory) {
		toSerialize["memory"] = o.Memory
	}
	return toSerialize, nil
}

type NullableVMParameter struct {
	value *VMParameter
	isSet bool
}

func (v NullableVMParameter) Get() *VMParameter {
	return v.value
}

func (v *NullableVMParameter) Set(val *VMParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableVMParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableVMParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVMParameter(val *VMParameter) *NullableVMParameter {
	return &NullableVMParameter{value: val, isSet: true}
}

func (v NullableVMParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVMParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
