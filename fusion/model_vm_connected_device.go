// © Ryan Johnson. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

package fusion

import (
	"encoding/json"
)

// checks if the VMConnectedDevice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VMConnectedDevice{}

// VMConnectedDevice struct for VMConnectedDevice
type VMConnectedDevice struct {
	// Number of items
	Index          *int32  `json:"index,omitempty"`
	StartConnected *string `json:"startConnected,omitempty"`
	// Number of items
	ConnectionStatus *int32  `json:"connectionStatus,omitempty"`
	DevicePath       *string `json:"devicePath,omitempty"`
}

// NewVMConnectedDevice instantiates a new VMConnectedDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVMConnectedDevice() *VMConnectedDevice {
	this := VMConnectedDevice{}
	var index int32 = 0
	this.Index = &index
	var connectionStatus int32 = 0
	this.ConnectionStatus = &connectionStatus
	return &this
}

// NewVMConnectedDeviceWithDefaults instantiates a new VMConnectedDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVMConnectedDeviceWithDefaults() *VMConnectedDevice {
	this := VMConnectedDevice{}
	var index int32 = 0
	this.Index = &index
	var connectionStatus int32 = 0
	this.ConnectionStatus = &connectionStatus
	return &this
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *VMConnectedDevice) GetIndex() int32 {
	if o == nil || IsNil(o.Index) {
		var ret int32
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMConnectedDevice) GetIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.Index) {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *VMConnectedDevice) HasIndex() bool {
	if o != nil && !IsNil(o.Index) {
		return true
	}

	return false
}

// SetIndex gets a reference to the given int32 and assigns it to the Index field.
func (o *VMConnectedDevice) SetIndex(v int32) {
	o.Index = &v
}

// GetStartConnected returns the StartConnected field value if set, zero value otherwise.
func (o *VMConnectedDevice) GetStartConnected() string {
	if o == nil || IsNil(o.StartConnected) {
		var ret string
		return ret
	}
	return *o.StartConnected
}

// GetStartConnectedOk returns a tuple with the StartConnected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMConnectedDevice) GetStartConnectedOk() (*string, bool) {
	if o == nil || IsNil(o.StartConnected) {
		return nil, false
	}
	return o.StartConnected, true
}

// HasStartConnected returns a boolean if a field has been set.
func (o *VMConnectedDevice) HasStartConnected() bool {
	if o != nil && !IsNil(o.StartConnected) {
		return true
	}

	return false
}

// SetStartConnected gets a reference to the given string and assigns it to the StartConnected field.
func (o *VMConnectedDevice) SetStartConnected(v string) {
	o.StartConnected = &v
}

// GetConnectionStatus returns the ConnectionStatus field value if set, zero value otherwise.
func (o *VMConnectedDevice) GetConnectionStatus() int32 {
	if o == nil || IsNil(o.ConnectionStatus) {
		var ret int32
		return ret
	}
	return *o.ConnectionStatus
}

// GetConnectionStatusOk returns a tuple with the ConnectionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMConnectedDevice) GetConnectionStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.ConnectionStatus) {
		return nil, false
	}
	return o.ConnectionStatus, true
}

// HasConnectionStatus returns a boolean if a field has been set.
func (o *VMConnectedDevice) HasConnectionStatus() bool {
	if o != nil && !IsNil(o.ConnectionStatus) {
		return true
	}

	return false
}

// SetConnectionStatus gets a reference to the given int32 and assigns it to the ConnectionStatus field.
func (o *VMConnectedDevice) SetConnectionStatus(v int32) {
	o.ConnectionStatus = &v
}

// GetDevicePath returns the DevicePath field value if set, zero value otherwise.
func (o *VMConnectedDevice) GetDevicePath() string {
	if o == nil || IsNil(o.DevicePath) {
		var ret string
		return ret
	}
	return *o.DevicePath
}

// GetDevicePathOk returns a tuple with the DevicePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMConnectedDevice) GetDevicePathOk() (*string, bool) {
	if o == nil || IsNil(o.DevicePath) {
		return nil, false
	}
	return o.DevicePath, true
}

// HasDevicePath returns a boolean if a field has been set.
func (o *VMConnectedDevice) HasDevicePath() bool {
	if o != nil && !IsNil(o.DevicePath) {
		return true
	}

	return false
}

// SetDevicePath gets a reference to the given string and assigns it to the DevicePath field.
func (o *VMConnectedDevice) SetDevicePath(v string) {
	o.DevicePath = &v
}

func (o VMConnectedDevice) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VMConnectedDevice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Index) {
		toSerialize["index"] = o.Index
	}
	if !IsNil(o.StartConnected) {
		toSerialize["startConnected"] = o.StartConnected
	}
	if !IsNil(o.ConnectionStatus) {
		toSerialize["connectionStatus"] = o.ConnectionStatus
	}
	if !IsNil(o.DevicePath) {
		toSerialize["devicePath"] = o.DevicePath
	}
	return toSerialize, nil
}

type NullableVMConnectedDevice struct {
	value *VMConnectedDevice
	isSet bool
}

func (v NullableVMConnectedDevice) Get() *VMConnectedDevice {
	return v.value
}

func (v *NullableVMConnectedDevice) Set(val *VMConnectedDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableVMConnectedDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableVMConnectedDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVMConnectedDevice(val *VMConnectedDevice) *NullableVMConnectedDevice {
	return &NullableVMConnectedDevice{value: val, isSet: true}
}

func (v NullableVMConnectedDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVMConnectedDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
