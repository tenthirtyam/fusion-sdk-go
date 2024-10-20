// © Ryan Johnson. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

package fusion

import (
	"encoding/json"
)

// checks if the VMConnectedDeviceList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VMConnectedDeviceList{}

// VMConnectedDeviceList struct for VMConnectedDeviceList
type VMConnectedDeviceList struct {
	// Number of items
	Num     *int32              `json:"num,omitempty"`
	Devices []VMConnectedDevice `json:"devices,omitempty"`
}

// NewVMConnectedDeviceList instantiates a new VMConnectedDeviceList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVMConnectedDeviceList() *VMConnectedDeviceList {
	this := VMConnectedDeviceList{}
	var num int32 = 0
	this.Num = &num
	return &this
}

// NewVMConnectedDeviceListWithDefaults instantiates a new VMConnectedDeviceList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVMConnectedDeviceListWithDefaults() *VMConnectedDeviceList {
	this := VMConnectedDeviceList{}
	var num int32 = 0
	this.Num = &num
	return &this
}

// GetNum returns the Num field value if set, zero value otherwise.
func (o *VMConnectedDeviceList) GetNum() int32 {
	if o == nil || IsNil(o.Num) {
		var ret int32
		return ret
	}
	return *o.Num
}

// GetNumOk returns a tuple with the Num field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMConnectedDeviceList) GetNumOk() (*int32, bool) {
	if o == nil || IsNil(o.Num) {
		return nil, false
	}
	return o.Num, true
}

// HasNum returns a boolean if a field has been set.
func (o *VMConnectedDeviceList) HasNum() bool {
	if o != nil && !IsNil(o.Num) {
		return true
	}

	return false
}

// SetNum gets a reference to the given int32 and assigns it to the Num field.
func (o *VMConnectedDeviceList) SetNum(v int32) {
	o.Num = &v
}

// GetDevices returns the Devices field value if set, zero value otherwise.
func (o *VMConnectedDeviceList) GetDevices() []VMConnectedDevice {
	if o == nil || IsNil(o.Devices) {
		var ret []VMConnectedDevice
		return ret
	}
	return o.Devices
}

// GetDevicesOk returns a tuple with the Devices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMConnectedDeviceList) GetDevicesOk() ([]VMConnectedDevice, bool) {
	if o == nil || IsNil(o.Devices) {
		return nil, false
	}
	return o.Devices, true
}

// HasDevices returns a boolean if a field has been set.
func (o *VMConnectedDeviceList) HasDevices() bool {
	if o != nil && !IsNil(o.Devices) {
		return true
	}

	return false
}

// SetDevices gets a reference to the given []VMConnectedDevice and assigns it to the Devices field.
func (o *VMConnectedDeviceList) SetDevices(v []VMConnectedDevice) {
	o.Devices = v
}

func (o VMConnectedDeviceList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VMConnectedDeviceList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Num) {
		toSerialize["num"] = o.Num
	}
	if !IsNil(o.Devices) {
		toSerialize["devices"] = o.Devices
	}
	return toSerialize, nil
}

type NullableVMConnectedDeviceList struct {
	value *VMConnectedDeviceList
	isSet bool
}

func (v NullableVMConnectedDeviceList) Get() *VMConnectedDeviceList {
	return v.value
}

func (v *NullableVMConnectedDeviceList) Set(val *VMConnectedDeviceList) {
	v.value = val
	v.isSet = true
}

func (v NullableVMConnectedDeviceList) IsSet() bool {
	return v.isSet
}

func (v *NullableVMConnectedDeviceList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVMConnectedDeviceList(val *VMConnectedDeviceList) *NullableVMConnectedDeviceList {
	return &NullableVMConnectedDeviceList{value: val, isSet: true}
}

func (v NullableVMConnectedDeviceList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVMConnectedDeviceList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
