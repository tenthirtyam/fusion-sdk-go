// © Ryan Johnson. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

package fusion

import (
	"encoding/json"
)

// checks if the NicIpStackAll type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NicIpStackAll{}

// NicIpStackAll Information about all NICs
type NicIpStackAll struct {
	Nics   *NicIpStack  `json:"nics,omitempty"`
	Routes []RouteEntry `json:"routes,omitempty"`
	Dns    *DnsConfig   `json:"dns,omitempty"`
	Wins   *WinsConfig  `json:"wins,omitempty"`
	Dhcpv4 *DhcpConfig  `json:"dhcpv4,omitempty"`
	Dhcpv6 *DhcpConfig  `json:"dhcpv6,omitempty"`
}

// NewNicIpStackAll instantiates a new NicIpStackAll object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNicIpStackAll() *NicIpStackAll {
	this := NicIpStackAll{}
	return &this
}

// NewNicIpStackAllWithDefaults instantiates a new NicIpStackAll object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNicIpStackAllWithDefaults() *NicIpStackAll {
	this := NicIpStackAll{}
	return &this
}

// GetNics returns the Nics field value if set, zero value otherwise.
func (o *NicIpStackAll) GetNics() NicIpStack {
	if o == nil || IsNil(o.Nics) {
		var ret NicIpStack
		return ret
	}
	return *o.Nics
}

// GetNicsOk returns a tuple with the Nics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NicIpStackAll) GetNicsOk() (*NicIpStack, bool) {
	if o == nil || IsNil(o.Nics) {
		return nil, false
	}
	return o.Nics, true
}

// HasNics returns a boolean if a field has been set.
func (o *NicIpStackAll) HasNics() bool {
	if o != nil && !IsNil(o.Nics) {
		return true
	}

	return false
}

// SetNics gets a reference to the given NicIpStack and assigns it to the Nics field.
func (o *NicIpStackAll) SetNics(v NicIpStack) {
	o.Nics = &v
}

// GetRoutes returns the Routes field value if set, zero value otherwise.
func (o *NicIpStackAll) GetRoutes() []RouteEntry {
	if o == nil || IsNil(o.Routes) {
		var ret []RouteEntry
		return ret
	}
	return o.Routes
}

// GetRoutesOk returns a tuple with the Routes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NicIpStackAll) GetRoutesOk() ([]RouteEntry, bool) {
	if o == nil || IsNil(o.Routes) {
		return nil, false
	}
	return o.Routes, true
}

// HasRoutes returns a boolean if a field has been set.
func (o *NicIpStackAll) HasRoutes() bool {
	if o != nil && !IsNil(o.Routes) {
		return true
	}

	return false
}

// SetRoutes gets a reference to the given []RouteEntry and assigns it to the Routes field.
func (o *NicIpStackAll) SetRoutes(v []RouteEntry) {
	o.Routes = v
}

// GetDns returns the Dns field value if set, zero value otherwise.
func (o *NicIpStackAll) GetDns() DnsConfig {
	if o == nil || IsNil(o.Dns) {
		var ret DnsConfig
		return ret
	}
	return *o.Dns
}

// GetDnsOk returns a tuple with the Dns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NicIpStackAll) GetDnsOk() (*DnsConfig, bool) {
	if o == nil || IsNil(o.Dns) {
		return nil, false
	}
	return o.Dns, true
}

// HasDns returns a boolean if a field has been set.
func (o *NicIpStackAll) HasDns() bool {
	if o != nil && !IsNil(o.Dns) {
		return true
	}

	return false
}

// SetDns gets a reference to the given DnsConfig and assigns it to the Dns field.
func (o *NicIpStackAll) SetDns(v DnsConfig) {
	o.Dns = &v
}

// GetWins returns the Wins field value if set, zero value otherwise.
func (o *NicIpStackAll) GetWins() WinsConfig {
	if o == nil || IsNil(o.Wins) {
		var ret WinsConfig
		return ret
	}
	return *o.Wins
}

// GetWinsOk returns a tuple with the Wins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NicIpStackAll) GetWinsOk() (*WinsConfig, bool) {
	if o == nil || IsNil(o.Wins) {
		return nil, false
	}
	return o.Wins, true
}

// HasWins returns a boolean if a field has been set.
func (o *NicIpStackAll) HasWins() bool {
	if o != nil && !IsNil(o.Wins) {
		return true
	}

	return false
}

// SetWins gets a reference to the given WinsConfig and assigns it to the Wins field.
func (o *NicIpStackAll) SetWins(v WinsConfig) {
	o.Wins = &v
}

// GetDhcpv4 returns the Dhcpv4 field value if set, zero value otherwise.
func (o *NicIpStackAll) GetDhcpv4() DhcpConfig {
	if o == nil || IsNil(o.Dhcpv4) {
		var ret DhcpConfig
		return ret
	}
	return *o.Dhcpv4
}

// GetDhcpv4Ok returns a tuple with the Dhcpv4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NicIpStackAll) GetDhcpv4Ok() (*DhcpConfig, bool) {
	if o == nil || IsNil(o.Dhcpv4) {
		return nil, false
	}
	return o.Dhcpv4, true
}

// HasDhcpv4 returns a boolean if a field has been set.
func (o *NicIpStackAll) HasDhcpv4() bool {
	if o != nil && !IsNil(o.Dhcpv4) {
		return true
	}

	return false
}

// SetDhcpv4 gets a reference to the given DhcpConfig and assigns it to the Dhcpv4 field.
func (o *NicIpStackAll) SetDhcpv4(v DhcpConfig) {
	o.Dhcpv4 = &v
}

// GetDhcpv6 returns the Dhcpv6 field value if set, zero value otherwise.
func (o *NicIpStackAll) GetDhcpv6() DhcpConfig {
	if o == nil || IsNil(o.Dhcpv6) {
		var ret DhcpConfig
		return ret
	}
	return *o.Dhcpv6
}

// GetDhcpv6Ok returns a tuple with the Dhcpv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NicIpStackAll) GetDhcpv6Ok() (*DhcpConfig, bool) {
	if o == nil || IsNil(o.Dhcpv6) {
		return nil, false
	}
	return o.Dhcpv6, true
}

// HasDhcpv6 returns a boolean if a field has been set.
func (o *NicIpStackAll) HasDhcpv6() bool {
	if o != nil && !IsNil(o.Dhcpv6) {
		return true
	}

	return false
}

// SetDhcpv6 gets a reference to the given DhcpConfig and assigns it to the Dhcpv6 field.
func (o *NicIpStackAll) SetDhcpv6(v DhcpConfig) {
	o.Dhcpv6 = &v
}

func (o NicIpStackAll) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NicIpStackAll) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Nics) {
		toSerialize["nics"] = o.Nics
	}
	if !IsNil(o.Routes) {
		toSerialize["routes"] = o.Routes
	}
	if !IsNil(o.Dns) {
		toSerialize["dns"] = o.Dns
	}
	if !IsNil(o.Wins) {
		toSerialize["wins"] = o.Wins
	}
	if !IsNil(o.Dhcpv4) {
		toSerialize["dhcpv4"] = o.Dhcpv4
	}
	if !IsNil(o.Dhcpv6) {
		toSerialize["dhcpv6"] = o.Dhcpv6
	}
	return toSerialize, nil
}

type NullableNicIpStackAll struct {
	value *NicIpStackAll
	isSet bool
}

func (v NullableNicIpStackAll) Get() *NicIpStackAll {
	return v.value
}

func (v *NullableNicIpStackAll) Set(val *NicIpStackAll) {
	v.value = val
	v.isSet = true
}

func (v NullableNicIpStackAll) IsSet() bool {
	return v.isSet
}

func (v *NullableNicIpStackAll) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNicIpStackAll(val *NicIpStackAll) *NullableNicIpStackAll {
	return &NullableNicIpStackAll{value: val, isSet: true}
}

func (v NullableNicIpStackAll) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNicIpStackAll) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
