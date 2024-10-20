// © Ryan Johnson. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

package fusion

import (
	"encoding/json"
)

// checks if the VMRemoteVNC type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VMRemoteVNC{}

// VMRemoteVNC struct for VMRemoteVNC
type VMRemoteVNC struct {
	VNCEnabled *string `json:"VNCEnabled,omitempty"`
	// port of communication
	VNCPort *int32 `json:"VNCPort,omitempty"`
}

// NewVMRemoteVNC instantiates a new VMRemoteVNC object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVMRemoteVNC() *VMRemoteVNC {
	this := VMRemoteVNC{}
	var vNCPort int32 = 0
	this.VNCPort = &vNCPort
	return &this
}

// NewVMRemoteVNCWithDefaults instantiates a new VMRemoteVNC object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVMRemoteVNCWithDefaults() *VMRemoteVNC {
	this := VMRemoteVNC{}
	var vNCPort int32 = 0
	this.VNCPort = &vNCPort
	return &this
}

// GetVNCEnabled returns the VNCEnabled field value if set, zero value otherwise.
func (o *VMRemoteVNC) GetVNCEnabled() string {
	if o == nil || IsNil(o.VNCEnabled) {
		var ret string
		return ret
	}
	return *o.VNCEnabled
}

// GetVNCEnabledOk returns a tuple with the VNCEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMRemoteVNC) GetVNCEnabledOk() (*string, bool) {
	if o == nil || IsNil(o.VNCEnabled) {
		return nil, false
	}
	return o.VNCEnabled, true
}

// HasVNCEnabled returns a boolean if a field has been set.
func (o *VMRemoteVNC) HasVNCEnabled() bool {
	if o != nil && !IsNil(o.VNCEnabled) {
		return true
	}

	return false
}

// SetVNCEnabled gets a reference to the given string and assigns it to the VNCEnabled field.
func (o *VMRemoteVNC) SetVNCEnabled(v string) {
	o.VNCEnabled = &v
}

// GetVNCPort returns the VNCPort field value if set, zero value otherwise.
func (o *VMRemoteVNC) GetVNCPort() int32 {
	if o == nil || IsNil(o.VNCPort) {
		var ret int32
		return ret
	}
	return *o.VNCPort
}

// GetVNCPortOk returns a tuple with the VNCPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMRemoteVNC) GetVNCPortOk() (*int32, bool) {
	if o == nil || IsNil(o.VNCPort) {
		return nil, false
	}
	return o.VNCPort, true
}

// HasVNCPort returns a boolean if a field has been set.
func (o *VMRemoteVNC) HasVNCPort() bool {
	if o != nil && !IsNil(o.VNCPort) {
		return true
	}

	return false
}

// SetVNCPort gets a reference to the given int32 and assigns it to the VNCPort field.
func (o *VMRemoteVNC) SetVNCPort(v int32) {
	o.VNCPort = &v
}

func (o VMRemoteVNC) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VMRemoteVNC) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VNCEnabled) {
		toSerialize["VNCEnabled"] = o.VNCEnabled
	}
	if !IsNil(o.VNCPort) {
		toSerialize["VNCPort"] = o.VNCPort
	}
	return toSerialize, nil
}

type NullableVMRemoteVNC struct {
	value *VMRemoteVNC
	isSet bool
}

func (v NullableVMRemoteVNC) Get() *VMRemoteVNC {
	return v.value
}

func (v *NullableVMRemoteVNC) Set(val *VMRemoteVNC) {
	v.value = val
	v.isSet = true
}

func (v NullableVMRemoteVNC) IsSet() bool {
	return v.isSet
}

func (v *NullableVMRemoteVNC) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVMRemoteVNC(val *VMRemoteVNC) *NullableVMRemoteVNC {
	return &NullableVMRemoteVNC{value: val, isSet: true}
}

func (v NullableVMRemoteVNC) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVMRemoteVNC) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
