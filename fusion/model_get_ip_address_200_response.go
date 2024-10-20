// © Ryan Johnson. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

package fusion

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetIPAddress200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetIPAddress200Response{}

// GetIPAddress200Response struct for GetIPAddress200Response
type GetIPAddress200Response struct {
	// Guest OS IP address
	Ip string `json:"ip"`
}

type _GetIPAddress200Response GetIPAddress200Response

// NewGetIPAddress200Response instantiates a new GetIPAddress200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetIPAddress200Response(ip string) *GetIPAddress200Response {
	this := GetIPAddress200Response{}
	this.Ip = ip
	return &this
}

// NewGetIPAddress200ResponseWithDefaults instantiates a new GetIPAddress200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetIPAddress200ResponseWithDefaults() *GetIPAddress200Response {
	this := GetIPAddress200Response{}
	return &this
}

// GetIp returns the Ip field value
func (o *GetIPAddress200Response) GetIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ip
}

// GetIpOk returns a tuple with the Ip field value
// and a boolean to check if the value has been set.
func (o *GetIPAddress200Response) GetIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ip, true
}

// SetIp sets field value
func (o *GetIPAddress200Response) SetIp(v string) {
	o.Ip = v
}

func (o GetIPAddress200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetIPAddress200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ip"] = o.Ip
	return toSerialize, nil
}

func (o *GetIPAddress200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varGetIPAddress200Response := _GetIPAddress200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetIPAddress200Response)

	if err != nil {
		return err
	}

	*o = GetIPAddress200Response(varGetIPAddress200Response)

	return err
}

type NullableGetIPAddress200Response struct {
	value *GetIPAddress200Response
	isSet bool
}

func (v NullableGetIPAddress200Response) Get() *GetIPAddress200Response {
	return v.value
}

func (v *NullableGetIPAddress200Response) Set(val *GetIPAddress200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetIPAddress200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetIPAddress200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetIPAddress200Response(val *GetIPAddress200Response) *NullableGetIPAddress200Response {
	return &NullableGetIPAddress200Response{value: val, isSet: true}
}

func (v NullableGetIPAddress200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetIPAddress200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
