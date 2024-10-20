// © Ryan Johnson. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

package fusion

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the NicIpStack type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NicIpStack{}

// NicIpStack Information about a NIC
type NicIpStack struct {
	// Mac address, E.g., de:ad:be:ef:12:34
	Mac string `json:"mac"`
	// IP address(es) of the interface (CIDR)
	Ip    []string    `json:"ip,omitempty"`
	Dns   *DnsConfig  `json:"dns,omitempty"`
	Wins  *WinsConfig `json:"wins,omitempty"`
	Dhcp4 *DhcpConfig `json:"dhcp4,omitempty"`
	Dhcp6 *DhcpConfig `json:"dhcp6,omitempty"`
}

type _NicIpStack NicIpStack

// NewNicIpStack instantiates a new NicIpStack object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNicIpStack(mac string) *NicIpStack {
	this := NicIpStack{}
	this.Mac = mac
	return &this
}

// NewNicIpStackWithDefaults instantiates a new NicIpStack object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNicIpStackWithDefaults() *NicIpStack {
	this := NicIpStack{}
	return &this
}

// GetMac returns the Mac field value
func (o *NicIpStack) GetMac() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mac
}

// GetMacOk returns a tuple with the Mac field value
// and a boolean to check if the value has been set.
func (o *NicIpStack) GetMacOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mac, true
}

// SetMac sets field value
func (o *NicIpStack) SetMac(v string) {
	o.Mac = v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *NicIpStack) GetIp() []string {
	if o == nil || IsNil(o.Ip) {
		var ret []string
		return ret
	}
	return o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NicIpStack) GetIpOk() ([]string, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *NicIpStack) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given []string and assigns it to the Ip field.
func (o *NicIpStack) SetIp(v []string) {
	o.Ip = v
}

// GetDns returns the Dns field value if set, zero value otherwise.
func (o *NicIpStack) GetDns() DnsConfig {
	if o == nil || IsNil(o.Dns) {
		var ret DnsConfig
		return ret
	}
	return *o.Dns
}

// GetDnsOk returns a tuple with the Dns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NicIpStack) GetDnsOk() (*DnsConfig, bool) {
	if o == nil || IsNil(o.Dns) {
		return nil, false
	}
	return o.Dns, true
}

// HasDns returns a boolean if a field has been set.
func (o *NicIpStack) HasDns() bool {
	if o != nil && !IsNil(o.Dns) {
		return true
	}

	return false
}

// SetDns gets a reference to the given DnsConfig and assigns it to the Dns field.
func (o *NicIpStack) SetDns(v DnsConfig) {
	o.Dns = &v
}

// GetWins returns the Wins field value if set, zero value otherwise.
func (o *NicIpStack) GetWins() WinsConfig {
	if o == nil || IsNil(o.Wins) {
		var ret WinsConfig
		return ret
	}
	return *o.Wins
}

// GetWinsOk returns a tuple with the Wins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NicIpStack) GetWinsOk() (*WinsConfig, bool) {
	if o == nil || IsNil(o.Wins) {
		return nil, false
	}
	return o.Wins, true
}

// HasWins returns a boolean if a field has been set.
func (o *NicIpStack) HasWins() bool {
	if o != nil && !IsNil(o.Wins) {
		return true
	}

	return false
}

// SetWins gets a reference to the given WinsConfig and assigns it to the Wins field.
func (o *NicIpStack) SetWins(v WinsConfig) {
	o.Wins = &v
}

// GetDhcp4 returns the Dhcp4 field value if set, zero value otherwise.
func (o *NicIpStack) GetDhcp4() DhcpConfig {
	if o == nil || IsNil(o.Dhcp4) {
		var ret DhcpConfig
		return ret
	}
	return *o.Dhcp4
}

// GetDhcp4Ok returns a tuple with the Dhcp4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NicIpStack) GetDhcp4Ok() (*DhcpConfig, bool) {
	if o == nil || IsNil(o.Dhcp4) {
		return nil, false
	}
	return o.Dhcp4, true
}

// HasDhcp4 returns a boolean if a field has been set.
func (o *NicIpStack) HasDhcp4() bool {
	if o != nil && !IsNil(o.Dhcp4) {
		return true
	}

	return false
}

// SetDhcp4 gets a reference to the given DhcpConfig and assigns it to the Dhcp4 field.
func (o *NicIpStack) SetDhcp4(v DhcpConfig) {
	o.Dhcp4 = &v
}

// GetDhcp6 returns the Dhcp6 field value if set, zero value otherwise.
func (o *NicIpStack) GetDhcp6() DhcpConfig {
	if o == nil || IsNil(o.Dhcp6) {
		var ret DhcpConfig
		return ret
	}
	return *o.Dhcp6
}

// GetDhcp6Ok returns a tuple with the Dhcp6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NicIpStack) GetDhcp6Ok() (*DhcpConfig, bool) {
	if o == nil || IsNil(o.Dhcp6) {
		return nil, false
	}
	return o.Dhcp6, true
}

// HasDhcp6 returns a boolean if a field has been set.
func (o *NicIpStack) HasDhcp6() bool {
	if o != nil && !IsNil(o.Dhcp6) {
		return true
	}

	return false
}

// SetDhcp6 gets a reference to the given DhcpConfig and assigns it to the Dhcp6 field.
func (o *NicIpStack) SetDhcp6(v DhcpConfig) {
	o.Dhcp6 = &v
}

func (o NicIpStack) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NicIpStack) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mac"] = o.Mac
	if !IsNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !IsNil(o.Dns) {
		toSerialize["dns"] = o.Dns
	}
	if !IsNil(o.Wins) {
		toSerialize["wins"] = o.Wins
	}
	if !IsNil(o.Dhcp4) {
		toSerialize["dhcp4"] = o.Dhcp4
	}
	if !IsNil(o.Dhcp6) {
		toSerialize["dhcp6"] = o.Dhcp6
	}
	return toSerialize, nil
}

func (o *NicIpStack) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"mac",
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

	varNicIpStack := _NicIpStack{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNicIpStack)

	if err != nil {
		return err
	}

	*o = NicIpStack(varNicIpStack)

	return err
}

type NullableNicIpStack struct {
	value *NicIpStack
	isSet bool
}

func (v NullableNicIpStack) Get() *NicIpStack {
	return v.value
}

func (v *NullableNicIpStack) Set(val *NicIpStack) {
	v.value = val
	v.isSet = true
}

func (v NullableNicIpStack) IsSet() bool {
	return v.isSet
}

func (v *NullableNicIpStack) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNicIpStack(val *NicIpStack) *NullableNicIpStack {
	return &NullableNicIpStack{value: val, isSet: true}
}

func (v NullableNicIpStack) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNicIpStack) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
