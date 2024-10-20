// © Ryan Johnson. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

package fusion

import (
	"encoding/json"
)

// checks if the VMGuestIsolation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VMGuestIsolation{}

// VMGuestIsolation struct for VMGuestIsolation
type VMGuestIsolation struct {
	CopyDisabled  *string `json:"copyDisabled,omitempty"`
	DndDisabled   *string `json:"dndDisabled,omitempty"`
	HgfsDisabled  *string `json:"hgfsDisabled,omitempty"`
	PasteDisabled *string `json:"pasteDisabled,omitempty"`
}

// NewVMGuestIsolation instantiates a new VMGuestIsolation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVMGuestIsolation() *VMGuestIsolation {
	this := VMGuestIsolation{}
	return &this
}

// NewVMGuestIsolationWithDefaults instantiates a new VMGuestIsolation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVMGuestIsolationWithDefaults() *VMGuestIsolation {
	this := VMGuestIsolation{}
	return &this
}

// GetCopyDisabled returns the CopyDisabled field value if set, zero value otherwise.
func (o *VMGuestIsolation) GetCopyDisabled() string {
	if o == nil || IsNil(o.CopyDisabled) {
		var ret string
		return ret
	}
	return *o.CopyDisabled
}

// GetCopyDisabledOk returns a tuple with the CopyDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMGuestIsolation) GetCopyDisabledOk() (*string, bool) {
	if o == nil || IsNil(o.CopyDisabled) {
		return nil, false
	}
	return o.CopyDisabled, true
}

// HasCopyDisabled returns a boolean if a field has been set.
func (o *VMGuestIsolation) HasCopyDisabled() bool {
	if o != nil && !IsNil(o.CopyDisabled) {
		return true
	}

	return false
}

// SetCopyDisabled gets a reference to the given string and assigns it to the CopyDisabled field.
func (o *VMGuestIsolation) SetCopyDisabled(v string) {
	o.CopyDisabled = &v
}

// GetDndDisabled returns the DndDisabled field value if set, zero value otherwise.
func (o *VMGuestIsolation) GetDndDisabled() string {
	if o == nil || IsNil(o.DndDisabled) {
		var ret string
		return ret
	}
	return *o.DndDisabled
}

// GetDndDisabledOk returns a tuple with the DndDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMGuestIsolation) GetDndDisabledOk() (*string, bool) {
	if o == nil || IsNil(o.DndDisabled) {
		return nil, false
	}
	return o.DndDisabled, true
}

// HasDndDisabled returns a boolean if a field has been set.
func (o *VMGuestIsolation) HasDndDisabled() bool {
	if o != nil && !IsNil(o.DndDisabled) {
		return true
	}

	return false
}

// SetDndDisabled gets a reference to the given string and assigns it to the DndDisabled field.
func (o *VMGuestIsolation) SetDndDisabled(v string) {
	o.DndDisabled = &v
}

// GetHgfsDisabled returns the HgfsDisabled field value if set, zero value otherwise.
func (o *VMGuestIsolation) GetHgfsDisabled() string {
	if o == nil || IsNil(o.HgfsDisabled) {
		var ret string
		return ret
	}
	return *o.HgfsDisabled
}

// GetHgfsDisabledOk returns a tuple with the HgfsDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMGuestIsolation) GetHgfsDisabledOk() (*string, bool) {
	if o == nil || IsNil(o.HgfsDisabled) {
		return nil, false
	}
	return o.HgfsDisabled, true
}

// HasHgfsDisabled returns a boolean if a field has been set.
func (o *VMGuestIsolation) HasHgfsDisabled() bool {
	if o != nil && !IsNil(o.HgfsDisabled) {
		return true
	}

	return false
}

// SetHgfsDisabled gets a reference to the given string and assigns it to the HgfsDisabled field.
func (o *VMGuestIsolation) SetHgfsDisabled(v string) {
	o.HgfsDisabled = &v
}

// GetPasteDisabled returns the PasteDisabled field value if set, zero value otherwise.
func (o *VMGuestIsolation) GetPasteDisabled() string {
	if o == nil || IsNil(o.PasteDisabled) {
		var ret string
		return ret
	}
	return *o.PasteDisabled
}

// GetPasteDisabledOk returns a tuple with the PasteDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMGuestIsolation) GetPasteDisabledOk() (*string, bool) {
	if o == nil || IsNil(o.PasteDisabled) {
		return nil, false
	}
	return o.PasteDisabled, true
}

// HasPasteDisabled returns a boolean if a field has been set.
func (o *VMGuestIsolation) HasPasteDisabled() bool {
	if o != nil && !IsNil(o.PasteDisabled) {
		return true
	}

	return false
}

// SetPasteDisabled gets a reference to the given string and assigns it to the PasteDisabled field.
func (o *VMGuestIsolation) SetPasteDisabled(v string) {
	o.PasteDisabled = &v
}

func (o VMGuestIsolation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VMGuestIsolation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CopyDisabled) {
		toSerialize["copyDisabled"] = o.CopyDisabled
	}
	if !IsNil(o.DndDisabled) {
		toSerialize["dndDisabled"] = o.DndDisabled
	}
	if !IsNil(o.HgfsDisabled) {
		toSerialize["hgfsDisabled"] = o.HgfsDisabled
	}
	if !IsNil(o.PasteDisabled) {
		toSerialize["pasteDisabled"] = o.PasteDisabled
	}
	return toSerialize, nil
}

type NullableVMGuestIsolation struct {
	value *VMGuestIsolation
	isSet bool
}

func (v NullableVMGuestIsolation) Get() *VMGuestIsolation {
	return v.value
}

func (v *NullableVMGuestIsolation) Set(val *VMGuestIsolation) {
	v.value = val
	v.isSet = true
}

func (v NullableVMGuestIsolation) IsSet() bool {
	return v.isSet
}

func (v *NullableVMGuestIsolation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVMGuestIsolation(val *VMGuestIsolation) *NullableVMGuestIsolation {
	return &NullableVMGuestIsolation{value: val, isSet: true}
}

func (v NullableVMGuestIsolation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVMGuestIsolation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
