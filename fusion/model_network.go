// © Ryan Johnson. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

package fusion

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the Network type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Network{}

// Network The virtual network
type Network struct {
	// Name of virtual network
	Name   string `json:"name"`
	Type   string `json:"type"`
	Dhcp   string `json:"dhcp"`
	Subnet string `json:"subnet"`
	Mask   string `json:"mask"`
}

type _Network Network

// NewNetwork instantiates a new Network object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetwork(name string, type_ string, dhcp string, subnet string, mask string) *Network {
	this := Network{}
	this.Name = name
	this.Type = type_
	this.Dhcp = dhcp
	this.Subnet = subnet
	this.Mask = mask
	return &this
}

// NewNetworkWithDefaults instantiates a new Network object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkWithDefaults() *Network {
	this := Network{}
	return &this
}

// GetName returns the Name field value
func (o *Network) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Network) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Network) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *Network) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Network) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Network) SetType(v string) {
	o.Type = v
}

// GetDhcp returns the Dhcp field value
func (o *Network) GetDhcp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Dhcp
}

// GetDhcpOk returns a tuple with the Dhcp field value
// and a boolean to check if the value has been set.
func (o *Network) GetDhcpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dhcp, true
}

// SetDhcp sets field value
func (o *Network) SetDhcp(v string) {
	o.Dhcp = v
}

// GetSubnet returns the Subnet field value
func (o *Network) GetSubnet() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value
// and a boolean to check if the value has been set.
func (o *Network) GetSubnetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subnet, true
}

// SetSubnet sets field value
func (o *Network) SetSubnet(v string) {
	o.Subnet = v
}

// GetMask returns the Mask field value
func (o *Network) GetMask() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mask
}

// GetMaskOk returns a tuple with the Mask field value
// and a boolean to check if the value has been set.
func (o *Network) GetMaskOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mask, true
}

// SetMask sets field value
func (o *Network) SetMask(v string) {
	o.Mask = v
}

func (o Network) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Network) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	toSerialize["dhcp"] = o.Dhcp
	toSerialize["subnet"] = o.Subnet
	toSerialize["mask"] = o.Mask
	return toSerialize, nil
}

func (o *Network) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"type",
		"dhcp",
		"subnet",
		"mask",
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

	varNetwork := _Network{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNetwork)

	if err != nil {
		return err
	}

	*o = Network(varNetwork)

	return err
}

type NullableNetwork struct {
	value *Network
	isSet bool
}

func (v NullableNetwork) Get() *Network {
	return v.value
}

func (v *NullableNetwork) Set(val *Network) {
	v.value = val
	v.isSet = true
}

func (v NullableNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetwork(val *Network) *NullableNetwork {
	return &NullableNetwork{value: val, isSet: true}
}

func (v NullableNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
