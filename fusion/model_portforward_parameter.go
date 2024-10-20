// © Ryan Johnson. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

package fusion

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the PortforwardParameter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PortforwardParameter{}

// PortforwardParameter struct for PortforwardParameter
type PortforwardParameter struct {
	GuestIp string `json:"guestIp"`
	// port of communication
	GuestPort int32   `json:"guestPort"`
	Desc      *string `json:"desc,omitempty"`
}

type _PortforwardParameter PortforwardParameter

// NewPortforwardParameter instantiates a new PortforwardParameter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortforwardParameter(guestIp string, guestPort int32) *PortforwardParameter {
	this := PortforwardParameter{}
	this.GuestIp = guestIp
	this.GuestPort = guestPort
	return &this
}

// NewPortforwardParameterWithDefaults instantiates a new PortforwardParameter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortforwardParameterWithDefaults() *PortforwardParameter {
	this := PortforwardParameter{}
	var guestPort int32 = 0
	this.GuestPort = guestPort
	return &this
}

// GetGuestIp returns the GuestIp field value
func (o *PortforwardParameter) GetGuestIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GuestIp
}

// GetGuestIpOk returns a tuple with the GuestIp field value
// and a boolean to check if the value has been set.
func (o *PortforwardParameter) GetGuestIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GuestIp, true
}

// SetGuestIp sets field value
func (o *PortforwardParameter) SetGuestIp(v string) {
	o.GuestIp = v
}

// GetGuestPort returns the GuestPort field value
func (o *PortforwardParameter) GetGuestPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.GuestPort
}

// GetGuestPortOk returns a tuple with the GuestPort field value
// and a boolean to check if the value has been set.
func (o *PortforwardParameter) GetGuestPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GuestPort, true
}

// SetGuestPort sets field value
func (o *PortforwardParameter) SetGuestPort(v int32) {
	o.GuestPort = v
}

// GetDesc returns the Desc field value if set, zero value otherwise.
func (o *PortforwardParameter) GetDesc() string {
	if o == nil || IsNil(o.Desc) {
		var ret string
		return ret
	}
	return *o.Desc
}

// GetDescOk returns a tuple with the Desc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortforwardParameter) GetDescOk() (*string, bool) {
	if o == nil || IsNil(o.Desc) {
		return nil, false
	}
	return o.Desc, true
}

// HasDesc returns a boolean if a field has been set.
func (o *PortforwardParameter) HasDesc() bool {
	if o != nil && !IsNil(o.Desc) {
		return true
	}

	return false
}

// SetDesc gets a reference to the given string and assigns it to the Desc field.
func (o *PortforwardParameter) SetDesc(v string) {
	o.Desc = &v
}

func (o PortforwardParameter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PortforwardParameter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["guestIp"] = o.GuestIp
	toSerialize["guestPort"] = o.GuestPort
	if !IsNil(o.Desc) {
		toSerialize["desc"] = o.Desc
	}
	return toSerialize, nil
}

func (o *PortforwardParameter) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"guestIp",
		"guestPort",
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

	varPortforwardParameter := _PortforwardParameter{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPortforwardParameter)

	if err != nil {
		return err
	}

	*o = PortforwardParameter(varPortforwardParameter)

	return err
}

type NullablePortforwardParameter struct {
	value *PortforwardParameter
	isSet bool
}

func (v NullablePortforwardParameter) Get() *PortforwardParameter {
	return v.value
}

func (v *NullablePortforwardParameter) Set(val *PortforwardParameter) {
	v.value = val
	v.isSet = true
}

func (v NullablePortforwardParameter) IsSet() bool {
	return v.isSet
}

func (v *NullablePortforwardParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortforwardParameter(val *PortforwardParameter) *NullablePortforwardParameter {
	return &NullablePortforwardParameter{value: val, isSet: true}
}

func (v NullablePortforwardParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortforwardParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
