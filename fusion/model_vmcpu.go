// © Ryan Johnson. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

package fusion

import (
	"encoding/json"
)

// checks if the VMCPU type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VMCPU{}

// VMCPU The CPU information of VM
type VMCPU struct {
	// Number of processor cores
	Processors *int32 `json:"processors,omitempty"`
}

// NewVMCPU instantiates a new VMCPU object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVMCPU() *VMCPU {
	this := VMCPU{}
	var processors int32 = 1
	this.Processors = &processors
	return &this
}

// NewVMCPUWithDefaults instantiates a new VMCPU object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVMCPUWithDefaults() *VMCPU {
	this := VMCPU{}
	var processors int32 = 1
	this.Processors = &processors
	return &this
}

// GetProcessors returns the Processors field value if set, zero value otherwise.
func (o *VMCPU) GetProcessors() int32 {
	if o == nil || IsNil(o.Processors) {
		var ret int32
		return ret
	}
	return *o.Processors
}

// GetProcessorsOk returns a tuple with the Processors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMCPU) GetProcessorsOk() (*int32, bool) {
	if o == nil || IsNil(o.Processors) {
		return nil, false
	}
	return o.Processors, true
}

// HasProcessors returns a boolean if a field has been set.
func (o *VMCPU) HasProcessors() bool {
	if o != nil && !IsNil(o.Processors) {
		return true
	}

	return false
}

// SetProcessors gets a reference to the given int32 and assigns it to the Processors field.
func (o *VMCPU) SetProcessors(v int32) {
	o.Processors = &v
}

func (o VMCPU) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VMCPU) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Processors) {
		toSerialize["processors"] = o.Processors
	}
	return toSerialize, nil
}

type NullableVMCPU struct {
	value *VMCPU
	isSet bool
}

func (v NullableVMCPU) Get() *VMCPU {
	return v.value
}

func (v *NullableVMCPU) Set(val *VMCPU) {
	v.value = val
	v.isSet = true
}

func (v NullableVMCPU) IsSet() bool {
	return v.isSet
}

func (v *NullableVMCPU) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVMCPU(val *VMCPU) *NullableVMCPU {
	return &NullableVMCPU{value: val, isSet: true}
}

func (v NullableVMCPU) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVMCPU) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
