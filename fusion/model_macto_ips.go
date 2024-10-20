// © Ryan Johnson. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

package fusion

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the MACToIPs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MACToIPs{}

// MACToIPs The list of MAC to IP settings
type MACToIPs struct {
	// Number of items
	Num int32 `json:"num"`
	// The list of MAC to IP settings
	Mactoips []MACToIP `json:"mactoips,omitempty"`
}

type _MACToIPs MACToIPs

// NewMACToIPs instantiates a new MACToIPs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMACToIPs(num int32) *MACToIPs {
	this := MACToIPs{}
	this.Num = num
	return &this
}

// NewMACToIPsWithDefaults instantiates a new MACToIPs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMACToIPsWithDefaults() *MACToIPs {
	this := MACToIPs{}
	var num int32 = 0
	this.Num = num
	return &this
}

// GetNum returns the Num field value
func (o *MACToIPs) GetNum() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Num
}

// GetNumOk returns a tuple with the Num field value
// and a boolean to check if the value has been set.
func (o *MACToIPs) GetNumOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Num, true
}

// SetNum sets field value
func (o *MACToIPs) SetNum(v int32) {
	o.Num = v
}

// GetMactoips returns the Mactoips field value if set, zero value otherwise.
func (o *MACToIPs) GetMactoips() []MACToIP {
	if o == nil || IsNil(o.Mactoips) {
		var ret []MACToIP
		return ret
	}
	return o.Mactoips
}

// GetMactoipsOk returns a tuple with the Mactoips field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MACToIPs) GetMactoipsOk() ([]MACToIP, bool) {
	if o == nil || IsNil(o.Mactoips) {
		return nil, false
	}
	return o.Mactoips, true
}

// HasMactoips returns a boolean if a field has been set.
func (o *MACToIPs) HasMactoips() bool {
	if o != nil && !IsNil(o.Mactoips) {
		return true
	}

	return false
}

// SetMactoips gets a reference to the given []MACToIP and assigns it to the Mactoips field.
func (o *MACToIPs) SetMactoips(v []MACToIP) {
	o.Mactoips = v
}

func (o MACToIPs) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MACToIPs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["num"] = o.Num
	if !IsNil(o.Mactoips) {
		toSerialize["mactoips"] = o.Mactoips
	}
	return toSerialize, nil
}

func (o *MACToIPs) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"num",
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

	varMACToIPs := _MACToIPs{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMACToIPs)

	if err != nil {
		return err
	}

	*o = MACToIPs(varMACToIPs)

	return err
}

type NullableMACToIPs struct {
	value *MACToIPs
	isSet bool
}

func (v NullableMACToIPs) Get() *MACToIPs {
	return v.value
}

func (v *NullableMACToIPs) Set(val *MACToIPs) {
	v.value = val
	v.isSet = true
}

func (v NullableMACToIPs) IsSet() bool {
	return v.isSet
}

func (v *NullableMACToIPs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMACToIPs(val *MACToIPs) *NullableMACToIPs {
	return &NullableMACToIPs{value: val, isSet: true}
}

func (v NullableMACToIPs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMACToIPs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
