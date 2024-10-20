// © Ryan Johnson. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

package fusion

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the SharedFolderParameter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SharedFolderParameter{}

// SharedFolderParameter struct for SharedFolderParameter
type SharedFolderParameter struct {
	// Path of the host shared folder
	HostPath string `json:"host_path"`
	// The flags property specifies how the folder will be accessed by the VM. There is only one flag supported which is \"4\" and means read/write access.
	Flags int32 `json:"flags"`
}

type _SharedFolderParameter SharedFolderParameter

// NewSharedFolderParameter instantiates a new SharedFolderParameter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSharedFolderParameter(hostPath string, flags int32) *SharedFolderParameter {
	this := SharedFolderParameter{}
	this.HostPath = hostPath
	this.Flags = flags
	return &this
}

// NewSharedFolderParameterWithDefaults instantiates a new SharedFolderParameter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSharedFolderParameterWithDefaults() *SharedFolderParameter {
	this := SharedFolderParameter{}
	return &this
}

// GetHostPath returns the HostPath field value
func (o *SharedFolderParameter) GetHostPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HostPath
}

// GetHostPathOk returns a tuple with the HostPath field value
// and a boolean to check if the value has been set.
func (o *SharedFolderParameter) GetHostPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HostPath, true
}

// SetHostPath sets field value
func (o *SharedFolderParameter) SetHostPath(v string) {
	o.HostPath = v
}

// GetFlags returns the Flags field value
func (o *SharedFolderParameter) GetFlags() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value
// and a boolean to check if the value has been set.
func (o *SharedFolderParameter) GetFlagsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Flags, true
}

// SetFlags sets field value
func (o *SharedFolderParameter) SetFlags(v int32) {
	o.Flags = v
}

func (o SharedFolderParameter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SharedFolderParameter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["host_path"] = o.HostPath
	toSerialize["flags"] = o.Flags
	return toSerialize, nil
}

func (o *SharedFolderParameter) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"host_path",
		"flags",
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

	varSharedFolderParameter := _SharedFolderParameter{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSharedFolderParameter)

	if err != nil {
		return err
	}

	*o = SharedFolderParameter(varSharedFolderParameter)

	return err
}

type NullableSharedFolderParameter struct {
	value *SharedFolderParameter
	isSet bool
}

func (v NullableSharedFolderParameter) Get() *SharedFolderParameter {
	return v.value
}

func (v *NullableSharedFolderParameter) Set(val *SharedFolderParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableSharedFolderParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableSharedFolderParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSharedFolderParameter(val *SharedFolderParameter) *NullableSharedFolderParameter {
	return &NullableSharedFolderParameter{value: val, isSet: true}
}

func (v NullableSharedFolderParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSharedFolderParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
