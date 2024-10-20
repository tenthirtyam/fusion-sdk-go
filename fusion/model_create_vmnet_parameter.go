// © Ryan Johnson. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

package fusion

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the CreateVmnetParameter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateVmnetParameter{}

// CreateVmnetParameter struct for CreateVmnetParameter
type CreateVmnetParameter struct {
	// The host network name
	Name string `json:"name"`
	// The host network type
	Type *string `json:"type,omitempty"`
}

type _CreateVmnetParameter CreateVmnetParameter

// NewCreateVmnetParameter instantiates a new CreateVmnetParameter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateVmnetParameter(name string) *CreateVmnetParameter {
	this := CreateVmnetParameter{}
	this.Name = name
	return &this
}

// NewCreateVmnetParameterWithDefaults instantiates a new CreateVmnetParameter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateVmnetParameterWithDefaults() *CreateVmnetParameter {
	this := CreateVmnetParameter{}
	return &this
}

// GetName returns the Name field value
func (o *CreateVmnetParameter) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateVmnetParameter) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateVmnetParameter) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CreateVmnetParameter) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVmnetParameter) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CreateVmnetParameter) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CreateVmnetParameter) SetType(v string) {
	o.Type = &v
}

func (o CreateVmnetParameter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateVmnetParameter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

func (o *CreateVmnetParameter) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varCreateVmnetParameter := _CreateVmnetParameter{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateVmnetParameter)

	if err != nil {
		return err
	}

	*o = CreateVmnetParameter(varCreateVmnetParameter)

	return err
}

type NullableCreateVmnetParameter struct {
	value *CreateVmnetParameter
	isSet bool
}

func (v NullableCreateVmnetParameter) Get() *CreateVmnetParameter {
	return v.value
}

func (v *NullableCreateVmnetParameter) Set(val *CreateVmnetParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateVmnetParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateVmnetParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateVmnetParameter(val *CreateVmnetParameter) *NullableCreateVmnetParameter {
	return &NullableCreateVmnetParameter{value: val, isSet: true}
}

func (v NullableCreateVmnetParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateVmnetParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
