// © Ryan Johnson. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

package fusion

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the MACToIP type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MACToIP{}

// MACToIP The MAC to IP setting
type MACToIP struct {
	Vmnet string `json:"vmnet"`
	Mac   string `json:"mac"`
	Ip    string `json:"ip"`
}

type _MACToIP MACToIP

// NewMACToIP instantiates a new MACToIP object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMACToIP(vmnet string, mac string, ip string) *MACToIP {
	this := MACToIP{}
	this.Vmnet = vmnet
	this.Mac = mac
	this.Ip = ip
	return &this
}

// NewMACToIPWithDefaults instantiates a new MACToIP object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMACToIPWithDefaults() *MACToIP {
	this := MACToIP{}
	return &this
}

// GetVmnet returns the Vmnet field value
func (o *MACToIP) GetVmnet() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Vmnet
}

// GetVmnetOk returns a tuple with the Vmnet field value
// and a boolean to check if the value has been set.
func (o *MACToIP) GetVmnetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vmnet, true
}

// SetVmnet sets field value
func (o *MACToIP) SetVmnet(v string) {
	o.Vmnet = v
}

// GetMac returns the Mac field value
func (o *MACToIP) GetMac() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mac
}

// GetMacOk returns a tuple with the Mac field value
// and a boolean to check if the value has been set.
func (o *MACToIP) GetMacOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mac, true
}

// SetMac sets field value
func (o *MACToIP) SetMac(v string) {
	o.Mac = v
}

// GetIp returns the Ip field value
func (o *MACToIP) GetIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ip
}

// GetIpOk returns a tuple with the Ip field value
// and a boolean to check if the value has been set.
func (o *MACToIP) GetIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ip, true
}

// SetIp sets field value
func (o *MACToIP) SetIp(v string) {
	o.Ip = v
}

func (o MACToIP) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MACToIP) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["vmnet"] = o.Vmnet
	toSerialize["mac"] = o.Mac
	toSerialize["ip"] = o.Ip
	return toSerialize, nil
}

func (o *MACToIP) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"vmnet",
		"mac",
		"ip",
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

	varMACToIP := _MACToIP{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMACToIP)

	if err != nil {
		return err
	}

	*o = MACToIP(varMACToIP)

	return err
}

type NullableMACToIP struct {
	value *MACToIP
	isSet bool
}

func (v NullableMACToIP) Get() *MACToIP {
	return v.value
}

func (v *NullableMACToIP) Set(val *MACToIP) {
	v.value = val
	v.isSet = true
}

func (v NullableMACToIP) IsSet() bool {
	return v.isSet
}

func (v *NullableMACToIP) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMACToIP(val *MACToIP) *NullableMACToIP {
	return &NullableMACToIP{value: val, isSet: true}
}

func (v NullableMACToIP) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMACToIP) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
