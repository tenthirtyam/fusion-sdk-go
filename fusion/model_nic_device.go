// © Ryan Johnson. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

package fusion

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the NICDevice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NICDevice{}

// NICDevice struct for NICDevice
type NICDevice struct {
	// Index of Network Adapters
	Index int32 `json:"index"`
	// The network type of network adapter
	Type string `json:"type"`
	// The vmnet name
	Vmnet string `json:"vmnet"`
	// Mac address
	MacAddress string `json:"macAddress"`
}

type _NICDevice NICDevice

// NewNICDevice instantiates a new NICDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNICDevice(index int32, type_ string, vmnet string, macAddress string) *NICDevice {
	this := NICDevice{}
	this.Index = index
	this.Type = type_
	this.Vmnet = vmnet
	this.MacAddress = macAddress
	return &this
}

// NewNICDeviceWithDefaults instantiates a new NICDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNICDeviceWithDefaults() *NICDevice {
	this := NICDevice{}
	var index int32 = 1
	this.Index = index
	return &this
}

// GetIndex returns the Index field value
func (o *NICDevice) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *NICDevice) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *NICDevice) SetIndex(v int32) {
	o.Index = v
}

// GetType returns the Type field value
func (o *NICDevice) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *NICDevice) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *NICDevice) SetType(v string) {
	o.Type = v
}

// GetVmnet returns the Vmnet field value
func (o *NICDevice) GetVmnet() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Vmnet
}

// GetVmnetOk returns a tuple with the Vmnet field value
// and a boolean to check if the value has been set.
func (o *NICDevice) GetVmnetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vmnet, true
}

// SetVmnet sets field value
func (o *NICDevice) SetVmnet(v string) {
	o.Vmnet = v
}

// GetMacAddress returns the MacAddress field value
func (o *NICDevice) GetMacAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value
// and a boolean to check if the value has been set.
func (o *NICDevice) GetMacAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MacAddress, true
}

// SetMacAddress sets field value
func (o *NICDevice) SetMacAddress(v string) {
	o.MacAddress = v
}

func (o NICDevice) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NICDevice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["index"] = o.Index
	toSerialize["type"] = o.Type
	toSerialize["vmnet"] = o.Vmnet
	toSerialize["macAddress"] = o.MacAddress
	return toSerialize, nil
}

func (o *NICDevice) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"index",
		"type",
		"vmnet",
		"macAddress",
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

	varNICDevice := _NICDevice{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNICDevice)

	if err != nil {
		return err
	}

	*o = NICDevice(varNICDevice)

	return err
}

type NullableNICDevice struct {
	value *NICDevice
	isSet bool
}

func (v NullableNICDevice) Get() *NICDevice {
	return v.value
}

func (v *NullableNICDevice) Set(val *NICDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableNICDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableNICDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNICDevice(val *NICDevice) *NullableNICDevice {
	return &NullableNICDevice{value: val, isSet: true}
}

func (v NullableNICDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNICDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
