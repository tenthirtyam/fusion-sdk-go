// © Ryan Johnson. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

package fusion

import (
	"encoding/json"
)

// checks if the VMUsbDevice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VMUsbDevice{}

// VMUsbDevice struct for VMUsbDevice
type VMUsbDevice struct {
	// Number of items
	Index       *int32  `json:"index,omitempty"`
	Connected   *string `json:"connected,omitempty"`
	BackingInfo *string `json:"backingInfo,omitempty"`
	// Number of items
	BackingType *int32 `json:"BackingType,omitempty"`
}

// NewVMUsbDevice instantiates a new VMUsbDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVMUsbDevice() *VMUsbDevice {
	this := VMUsbDevice{}
	var index int32 = 0
	this.Index = &index
	var backingType int32 = 0
	this.BackingType = &backingType
	return &this
}

// NewVMUsbDeviceWithDefaults instantiates a new VMUsbDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVMUsbDeviceWithDefaults() *VMUsbDevice {
	this := VMUsbDevice{}
	var index int32 = 0
	this.Index = &index
	var backingType int32 = 0
	this.BackingType = &backingType
	return &this
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *VMUsbDevice) GetIndex() int32 {
	if o == nil || IsNil(o.Index) {
		var ret int32
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMUsbDevice) GetIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.Index) {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *VMUsbDevice) HasIndex() bool {
	if o != nil && !IsNil(o.Index) {
		return true
	}

	return false
}

// SetIndex gets a reference to the given int32 and assigns it to the Index field.
func (o *VMUsbDevice) SetIndex(v int32) {
	o.Index = &v
}

// GetConnected returns the Connected field value if set, zero value otherwise.
func (o *VMUsbDevice) GetConnected() string {
	if o == nil || IsNil(o.Connected) {
		var ret string
		return ret
	}
	return *o.Connected
}

// GetConnectedOk returns a tuple with the Connected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMUsbDevice) GetConnectedOk() (*string, bool) {
	if o == nil || IsNil(o.Connected) {
		return nil, false
	}
	return o.Connected, true
}

// HasConnected returns a boolean if a field has been set.
func (o *VMUsbDevice) HasConnected() bool {
	if o != nil && !IsNil(o.Connected) {
		return true
	}

	return false
}

// SetConnected gets a reference to the given string and assigns it to the Connected field.
func (o *VMUsbDevice) SetConnected(v string) {
	o.Connected = &v
}

// GetBackingInfo returns the BackingInfo field value if set, zero value otherwise.
func (o *VMUsbDevice) GetBackingInfo() string {
	if o == nil || IsNil(o.BackingInfo) {
		var ret string
		return ret
	}
	return *o.BackingInfo
}

// GetBackingInfoOk returns a tuple with the BackingInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMUsbDevice) GetBackingInfoOk() (*string, bool) {
	if o == nil || IsNil(o.BackingInfo) {
		return nil, false
	}
	return o.BackingInfo, true
}

// HasBackingInfo returns a boolean if a field has been set.
func (o *VMUsbDevice) HasBackingInfo() bool {
	if o != nil && !IsNil(o.BackingInfo) {
		return true
	}

	return false
}

// SetBackingInfo gets a reference to the given string and assigns it to the BackingInfo field.
func (o *VMUsbDevice) SetBackingInfo(v string) {
	o.BackingInfo = &v
}

// GetBackingType returns the BackingType field value if set, zero value otherwise.
func (o *VMUsbDevice) GetBackingType() int32 {
	if o == nil || IsNil(o.BackingType) {
		var ret int32
		return ret
	}
	return *o.BackingType
}

// GetBackingTypeOk returns a tuple with the BackingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMUsbDevice) GetBackingTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.BackingType) {
		return nil, false
	}
	return o.BackingType, true
}

// HasBackingType returns a boolean if a field has been set.
func (o *VMUsbDevice) HasBackingType() bool {
	if o != nil && !IsNil(o.BackingType) {
		return true
	}

	return false
}

// SetBackingType gets a reference to the given int32 and assigns it to the BackingType field.
func (o *VMUsbDevice) SetBackingType(v int32) {
	o.BackingType = &v
}

func (o VMUsbDevice) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VMUsbDevice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Index) {
		toSerialize["index"] = o.Index
	}
	if !IsNil(o.Connected) {
		toSerialize["connected"] = o.Connected
	}
	if !IsNil(o.BackingInfo) {
		toSerialize["backingInfo"] = o.BackingInfo
	}
	if !IsNil(o.BackingType) {
		toSerialize["BackingType"] = o.BackingType
	}
	return toSerialize, nil
}

type NullableVMUsbDevice struct {
	value *VMUsbDevice
	isSet bool
}

func (v NullableVMUsbDevice) Get() *VMUsbDevice {
	return v.value
}

func (v *NullableVMUsbDevice) Set(val *VMUsbDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableVMUsbDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableVMUsbDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVMUsbDevice(val *VMUsbDevice) *NullableVMUsbDevice {
	return &NullableVMUsbDevice{value: val, isSet: true}
}

func (v NullableVMUsbDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVMUsbDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
