// © Ryan Johnson. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

package fusion

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the RouteEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RouteEntry{}

// RouteEntry Routing entry
type RouteEntry struct {
	// IP address
	Dest string `json:"dest"`
	// Number of items
	Prefix int32 `json:"prefix"`
	// IP address
	Nexthop *string `json:"nexthop,omitempty"`
	// Number of items
	Interface int32 `json:"interface"`
	// Number of items
	Type int32 `json:"type"`
	// Number of items
	Metric int32 `json:"metric"`
}

type _RouteEntry RouteEntry

// NewRouteEntry instantiates a new RouteEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRouteEntry(dest string, prefix int32, interface_ int32, type_ int32, metric int32) *RouteEntry {
	this := RouteEntry{}
	this.Dest = dest
	this.Prefix = prefix
	this.Interface = interface_
	this.Type = type_
	this.Metric = metric
	return &this
}

// NewRouteEntryWithDefaults instantiates a new RouteEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRouteEntryWithDefaults() *RouteEntry {
	this := RouteEntry{}
	var prefix int32 = 0
	this.Prefix = prefix
	var interface_ int32 = 0
	this.Interface = interface_
	var type_ int32 = 0
	this.Type = type_
	var metric int32 = 0
	this.Metric = metric
	return &this
}

// GetDest returns the Dest field value
func (o *RouteEntry) GetDest() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Dest
}

// GetDestOk returns a tuple with the Dest field value
// and a boolean to check if the value has been set.
func (o *RouteEntry) GetDestOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dest, true
}

// SetDest sets field value
func (o *RouteEntry) SetDest(v string) {
	o.Dest = v
}

// GetPrefix returns the Prefix field value
func (o *RouteEntry) GetPrefix() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value
// and a boolean to check if the value has been set.
func (o *RouteEntry) GetPrefixOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Prefix, true
}

// SetPrefix sets field value
func (o *RouteEntry) SetPrefix(v int32) {
	o.Prefix = v
}

// GetNexthop returns the Nexthop field value if set, zero value otherwise.
func (o *RouteEntry) GetNexthop() string {
	if o == nil || IsNil(o.Nexthop) {
		var ret string
		return ret
	}
	return *o.Nexthop
}

// GetNexthopOk returns a tuple with the Nexthop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteEntry) GetNexthopOk() (*string, bool) {
	if o == nil || IsNil(o.Nexthop) {
		return nil, false
	}
	return o.Nexthop, true
}

// HasNexthop returns a boolean if a field has been set.
func (o *RouteEntry) HasNexthop() bool {
	if o != nil && !IsNil(o.Nexthop) {
		return true
	}

	return false
}

// SetNexthop gets a reference to the given string and assigns it to the Nexthop field.
func (o *RouteEntry) SetNexthop(v string) {
	o.Nexthop = &v
}

// GetInterface returns the Interface field value
func (o *RouteEntry) GetInterface() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Interface
}

// GetInterfaceOk returns a tuple with the Interface field value
// and a boolean to check if the value has been set.
func (o *RouteEntry) GetInterfaceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interface, true
}

// SetInterface sets field value
func (o *RouteEntry) SetInterface(v int32) {
	o.Interface = v
}

// GetType returns the Type field value
func (o *RouteEntry) GetType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RouteEntry) GetTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RouteEntry) SetType(v int32) {
	o.Type = v
}

// GetMetric returns the Metric field value
func (o *RouteEntry) GetMetric() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Metric
}

// GetMetricOk returns a tuple with the Metric field value
// and a boolean to check if the value has been set.
func (o *RouteEntry) GetMetricOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metric, true
}

// SetMetric sets field value
func (o *RouteEntry) SetMetric(v int32) {
	o.Metric = v
}

func (o RouteEntry) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RouteEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dest"] = o.Dest
	toSerialize["prefix"] = o.Prefix
	if !IsNil(o.Nexthop) {
		toSerialize["nexthop"] = o.Nexthop
	}
	toSerialize["interface"] = o.Interface
	toSerialize["type"] = o.Type
	toSerialize["metric"] = o.Metric
	return toSerialize, nil
}

func (o *RouteEntry) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dest",
		"prefix",
		"interface",
		"type",
		"metric",
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

	varRouteEntry := _RouteEntry{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRouteEntry)

	if err != nil {
		return err
	}

	*o = RouteEntry(varRouteEntry)

	return err
}

type NullableRouteEntry struct {
	value *RouteEntry
	isSet bool
}

func (v NullableRouteEntry) Get() *RouteEntry {
	return v.value
}

func (v *NullableRouteEntry) Set(val *RouteEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableRouteEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableRouteEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRouteEntry(val *RouteEntry) *NullableRouteEntry {
	return &NullableRouteEntry{value: val, isSet: true}
}

func (v NullableRouteEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRouteEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
