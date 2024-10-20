// © Ryan Johnson. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

package fusion

import (
	"encoding/json"
)

// checks if the VMUsbList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VMUsbList{}

// VMUsbList struct for VMUsbList
type VMUsbList struct {
	// Number of items
	Num        *int32        `json:"num,omitempty"`
	UsbDevices []VMUsbDevice `json:"usbDevices,omitempty"`
}

// NewVMUsbList instantiates a new VMUsbList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVMUsbList() *VMUsbList {
	this := VMUsbList{}
	var num int32 = 0
	this.Num = &num
	return &this
}

// NewVMUsbListWithDefaults instantiates a new VMUsbList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVMUsbListWithDefaults() *VMUsbList {
	this := VMUsbList{}
	var num int32 = 0
	this.Num = &num
	return &this
}

// GetNum returns the Num field value if set, zero value otherwise.
func (o *VMUsbList) GetNum() int32 {
	if o == nil || IsNil(o.Num) {
		var ret int32
		return ret
	}
	return *o.Num
}

// GetNumOk returns a tuple with the Num field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMUsbList) GetNumOk() (*int32, bool) {
	if o == nil || IsNil(o.Num) {
		return nil, false
	}
	return o.Num, true
}

// HasNum returns a boolean if a field has been set.
func (o *VMUsbList) HasNum() bool {
	if o != nil && !IsNil(o.Num) {
		return true
	}

	return false
}

// SetNum gets a reference to the given int32 and assigns it to the Num field.
func (o *VMUsbList) SetNum(v int32) {
	o.Num = &v
}

// GetUsbDevices returns the UsbDevices field value if set, zero value otherwise.
func (o *VMUsbList) GetUsbDevices() []VMUsbDevice {
	if o == nil || IsNil(o.UsbDevices) {
		var ret []VMUsbDevice
		return ret
	}
	return o.UsbDevices
}

// GetUsbDevicesOk returns a tuple with the UsbDevices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMUsbList) GetUsbDevicesOk() ([]VMUsbDevice, bool) {
	if o == nil || IsNil(o.UsbDevices) {
		return nil, false
	}
	return o.UsbDevices, true
}

// HasUsbDevices returns a boolean if a field has been set.
func (o *VMUsbList) HasUsbDevices() bool {
	if o != nil && !IsNil(o.UsbDevices) {
		return true
	}

	return false
}

// SetUsbDevices gets a reference to the given []VMUsbDevice and assigns it to the UsbDevices field.
func (o *VMUsbList) SetUsbDevices(v []VMUsbDevice) {
	o.UsbDevices = v
}

func (o VMUsbList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VMUsbList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Num) {
		toSerialize["num"] = o.Num
	}
	if !IsNil(o.UsbDevices) {
		toSerialize["usbDevices"] = o.UsbDevices
	}
	return toSerialize, nil
}

type NullableVMUsbList struct {
	value *VMUsbList
	isSet bool
}

func (v NullableVMUsbList) Get() *VMUsbList {
	return v.value
}

func (v *NullableVMUsbList) Set(val *VMUsbList) {
	v.value = val
	v.isSet = true
}

func (v NullableVMUsbList) IsSet() bool {
	return v.isSet
}

func (v *NullableVMUsbList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVMUsbList(val *VMUsbList) *NullableVMUsbList {
	return &NullableVMUsbList{value: val, isSet: true}
}

func (v NullableVMUsbList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVMUsbList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
