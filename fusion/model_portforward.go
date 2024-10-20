// © Ryan Johnson. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

package fusion

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the Portforward type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Portforward{}

// Portforward The port forwarding
type Portforward struct {
	// port of communication
	Port     int32            `json:"port"`
	Protocol string           `json:"protocol"`
	Desc     string           `json:"desc"`
	Guest    PortforwardGuest `json:"guest"`
}

type _Portforward Portforward

// NewPortforward instantiates a new Portforward object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortforward(port int32, protocol string, desc string, guest PortforwardGuest) *Portforward {
	this := Portforward{}
	this.Port = port
	this.Protocol = protocol
	this.Desc = desc
	this.Guest = guest
	return &this
}

// NewPortforwardWithDefaults instantiates a new Portforward object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortforwardWithDefaults() *Portforward {
	this := Portforward{}
	var port int32 = 0
	this.Port = port
	return &this
}

// GetPort returns the Port field value
func (o *Portforward) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *Portforward) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *Portforward) SetPort(v int32) {
	o.Port = v
}

// GetProtocol returns the Protocol field value
func (o *Portforward) GetProtocol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
func (o *Portforward) GetProtocolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Protocol, true
}

// SetProtocol sets field value
func (o *Portforward) SetProtocol(v string) {
	o.Protocol = v
}

// GetDesc returns the Desc field value
func (o *Portforward) GetDesc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Desc
}

// GetDescOk returns a tuple with the Desc field value
// and a boolean to check if the value has been set.
func (o *Portforward) GetDescOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Desc, true
}

// SetDesc sets field value
func (o *Portforward) SetDesc(v string) {
	o.Desc = v
}

// GetGuest returns the Guest field value
func (o *Portforward) GetGuest() PortforwardGuest {
	if o == nil {
		var ret PortforwardGuest
		return ret
	}

	return o.Guest
}

// GetGuestOk returns a tuple with the Guest field value
// and a boolean to check if the value has been set.
func (o *Portforward) GetGuestOk() (*PortforwardGuest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Guest, true
}

// SetGuest sets field value
func (o *Portforward) SetGuest(v PortforwardGuest) {
	o.Guest = v
}

func (o Portforward) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Portforward) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["port"] = o.Port
	toSerialize["protocol"] = o.Protocol
	toSerialize["desc"] = o.Desc
	toSerialize["guest"] = o.Guest
	return toSerialize, nil
}

func (o *Portforward) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"port",
		"protocol",
		"desc",
		"guest",
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

	varPortforward := _Portforward{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPortforward)

	if err != nil {
		return err
	}

	*o = Portforward(varPortforward)

	return err
}

type NullablePortforward struct {
	value *Portforward
	isSet bool
}

func (v NullablePortforward) Get() *Portforward {
	return v.value
}

func (v *NullablePortforward) Set(val *Portforward) {
	v.value = val
	v.isSet = true
}

func (v NullablePortforward) IsSet() bool {
	return v.isSet
}

func (v *NullablePortforward) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortforward(val *Portforward) *NullablePortforward {
	return &NullablePortforward{value: val, isSet: true}
}

func (v NullablePortforward) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortforward) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
