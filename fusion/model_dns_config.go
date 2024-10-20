// © Ryan Johnson. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

package fusion

import (
	"encoding/json"
)

// checks if the DnsConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DnsConfig{}

// DnsConfig DNS configuration
type DnsConfig struct {
	Hostname   *string  `json:"hostname,omitempty"`
	Domainname *string  `json:"domainname,omitempty"`
	Server     []string `json:"server,omitempty"`
	Search     []string `json:"search,omitempty"`
}

// NewDnsConfig instantiates a new DnsConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnsConfig() *DnsConfig {
	this := DnsConfig{}
	return &this
}

// NewDnsConfigWithDefaults instantiates a new DnsConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnsConfigWithDefaults() *DnsConfig {
	this := DnsConfig{}
	return &this
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *DnsConfig) GetHostname() string {
	if o == nil || IsNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsConfig) GetHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.Hostname) {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *DnsConfig) HasHostname() bool {
	if o != nil && !IsNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *DnsConfig) SetHostname(v string) {
	o.Hostname = &v
}

// GetDomainname returns the Domainname field value if set, zero value otherwise.
func (o *DnsConfig) GetDomainname() string {
	if o == nil || IsNil(o.Domainname) {
		var ret string
		return ret
	}
	return *o.Domainname
}

// GetDomainnameOk returns a tuple with the Domainname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsConfig) GetDomainnameOk() (*string, bool) {
	if o == nil || IsNil(o.Domainname) {
		return nil, false
	}
	return o.Domainname, true
}

// HasDomainname returns a boolean if a field has been set.
func (o *DnsConfig) HasDomainname() bool {
	if o != nil && !IsNil(o.Domainname) {
		return true
	}

	return false
}

// SetDomainname gets a reference to the given string and assigns it to the Domainname field.
func (o *DnsConfig) SetDomainname(v string) {
	o.Domainname = &v
}

// GetServer returns the Server field value if set, zero value otherwise.
func (o *DnsConfig) GetServer() []string {
	if o == nil || IsNil(o.Server) {
		var ret []string
		return ret
	}
	return o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsConfig) GetServerOk() ([]string, bool) {
	if o == nil || IsNil(o.Server) {
		return nil, false
	}
	return o.Server, true
}

// HasServer returns a boolean if a field has been set.
func (o *DnsConfig) HasServer() bool {
	if o != nil && !IsNil(o.Server) {
		return true
	}

	return false
}

// SetServer gets a reference to the given []string and assigns it to the Server field.
func (o *DnsConfig) SetServer(v []string) {
	o.Server = v
}

// GetSearch returns the Search field value if set, zero value otherwise.
func (o *DnsConfig) GetSearch() []string {
	if o == nil || IsNil(o.Search) {
		var ret []string
		return ret
	}
	return o.Search
}

// GetSearchOk returns a tuple with the Search field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsConfig) GetSearchOk() ([]string, bool) {
	if o == nil || IsNil(o.Search) {
		return nil, false
	}
	return o.Search, true
}

// HasSearch returns a boolean if a field has been set.
func (o *DnsConfig) HasSearch() bool {
	if o != nil && !IsNil(o.Search) {
		return true
	}

	return false
}

// SetSearch gets a reference to the given []string and assigns it to the Search field.
func (o *DnsConfig) SetSearch(v []string) {
	o.Search = v
}

func (o DnsConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DnsConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Hostname) {
		toSerialize["hostname"] = o.Hostname
	}
	if !IsNil(o.Domainname) {
		toSerialize["domainname"] = o.Domainname
	}
	if !IsNil(o.Server) {
		toSerialize["server"] = o.Server
	}
	if !IsNil(o.Search) {
		toSerialize["search"] = o.Search
	}
	return toSerialize, nil
}

type NullableDnsConfig struct {
	value *DnsConfig
	isSet bool
}

func (v NullableDnsConfig) Get() *DnsConfig {
	return v.value
}

func (v *NullableDnsConfig) Set(val *DnsConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableDnsConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableDnsConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnsConfig(val *DnsConfig) *NullableDnsConfig {
	return &NullableDnsConfig{value: val, isSet: true}
}

func (v NullableDnsConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnsConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
