// © Ryan Johnson. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

package fusion

import (
	"encoding/json"
)

// checks if the VMApplianceView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VMApplianceView{}

// VMApplianceView struct for VMApplianceView
type VMApplianceView struct {
	Author  *string `json:"author,omitempty"`
	Version *string `json:"version,omitempty"`
	// port of communication
	Port          *int32  `json:"port,omitempty"`
	ShowAtPowerOn *string `json:"showAtPowerOn,omitempty"`
}

// NewVMApplianceView instantiates a new VMApplianceView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVMApplianceView() *VMApplianceView {
	this := VMApplianceView{}
	var port int32 = 0
	this.Port = &port
	return &this
}

// NewVMApplianceViewWithDefaults instantiates a new VMApplianceView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVMApplianceViewWithDefaults() *VMApplianceView {
	this := VMApplianceView{}
	var port int32 = 0
	this.Port = &port
	return &this
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *VMApplianceView) GetAuthor() string {
	if o == nil || IsNil(o.Author) {
		var ret string
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMApplianceView) GetAuthorOk() (*string, bool) {
	if o == nil || IsNil(o.Author) {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *VMApplianceView) HasAuthor() bool {
	if o != nil && !IsNil(o.Author) {
		return true
	}

	return false
}

// SetAuthor gets a reference to the given string and assigns it to the Author field.
func (o *VMApplianceView) SetAuthor(v string) {
	o.Author = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *VMApplianceView) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMApplianceView) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *VMApplianceView) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *VMApplianceView) SetVersion(v string) {
	o.Version = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *VMApplianceView) GetPort() int32 {
	if o == nil || IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMApplianceView) GetPortOk() (*int32, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *VMApplianceView) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *VMApplianceView) SetPort(v int32) {
	o.Port = &v
}

// GetShowAtPowerOn returns the ShowAtPowerOn field value if set, zero value otherwise.
func (o *VMApplianceView) GetShowAtPowerOn() string {
	if o == nil || IsNil(o.ShowAtPowerOn) {
		var ret string
		return ret
	}
	return *o.ShowAtPowerOn
}

// GetShowAtPowerOnOk returns a tuple with the ShowAtPowerOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMApplianceView) GetShowAtPowerOnOk() (*string, bool) {
	if o == nil || IsNil(o.ShowAtPowerOn) {
		return nil, false
	}
	return o.ShowAtPowerOn, true
}

// HasShowAtPowerOn returns a boolean if a field has been set.
func (o *VMApplianceView) HasShowAtPowerOn() bool {
	if o != nil && !IsNil(o.ShowAtPowerOn) {
		return true
	}

	return false
}

// SetShowAtPowerOn gets a reference to the given string and assigns it to the ShowAtPowerOn field.
func (o *VMApplianceView) SetShowAtPowerOn(v string) {
	o.ShowAtPowerOn = &v
}

func (o VMApplianceView) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VMApplianceView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Author) {
		toSerialize["author"] = o.Author
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.ShowAtPowerOn) {
		toSerialize["showAtPowerOn"] = o.ShowAtPowerOn
	}
	return toSerialize, nil
}

type NullableVMApplianceView struct {
	value *VMApplianceView
	isSet bool
}

func (v NullableVMApplianceView) Get() *VMApplianceView {
	return v.value
}

func (v *NullableVMApplianceView) Set(val *VMApplianceView) {
	v.value = val
	v.isSet = true
}

func (v NullableVMApplianceView) IsSet() bool {
	return v.isSet
}

func (v *NullableVMApplianceView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVMApplianceView(val *VMApplianceView) *NullableVMApplianceView {
	return &NullableVMApplianceView{value: val, isSet: true}
}

func (v NullableVMApplianceView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVMApplianceView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
