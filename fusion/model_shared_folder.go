// © Ryan Johnson. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

package fusion

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the SharedFolder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SharedFolder{}

// SharedFolder struct for SharedFolder
type SharedFolder struct {
	// ID of folder which be mounted to the host
	FolderId string `json:"folder_id"`
	// Path of the host shared folder
	HostPath string `json:"host_path"`
	// The flags property specifies how the folder will be accessed by the VM. There is only one flag supported which is \"4\" and means read/write access.
	Flags int32 `json:"flags"`
}

type _SharedFolder SharedFolder

// NewSharedFolder instantiates a new SharedFolder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSharedFolder(folderId string, hostPath string, flags int32) *SharedFolder {
	this := SharedFolder{}
	this.FolderId = folderId
	this.HostPath = hostPath
	this.Flags = flags
	return &this
}

// NewSharedFolderWithDefaults instantiates a new SharedFolder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSharedFolderWithDefaults() *SharedFolder {
	this := SharedFolder{}
	return &this
}

// GetFolderId returns the FolderId field value
func (o *SharedFolder) GetFolderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FolderId
}

// GetFolderIdOk returns a tuple with the FolderId field value
// and a boolean to check if the value has been set.
func (o *SharedFolder) GetFolderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FolderId, true
}

// SetFolderId sets field value
func (o *SharedFolder) SetFolderId(v string) {
	o.FolderId = v
}

// GetHostPath returns the HostPath field value
func (o *SharedFolder) GetHostPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HostPath
}

// GetHostPathOk returns a tuple with the HostPath field value
// and a boolean to check if the value has been set.
func (o *SharedFolder) GetHostPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HostPath, true
}

// SetHostPath sets field value
func (o *SharedFolder) SetHostPath(v string) {
	o.HostPath = v
}

// GetFlags returns the Flags field value
func (o *SharedFolder) GetFlags() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value
// and a boolean to check if the value has been set.
func (o *SharedFolder) GetFlagsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Flags, true
}

// SetFlags sets field value
func (o *SharedFolder) SetFlags(v int32) {
	o.Flags = v
}

func (o SharedFolder) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SharedFolder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["folder_id"] = o.FolderId
	toSerialize["host_path"] = o.HostPath
	toSerialize["flags"] = o.Flags
	return toSerialize, nil
}

func (o *SharedFolder) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"folder_id",
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

	varSharedFolder := _SharedFolder{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSharedFolder)

	if err != nil {
		return err
	}

	*o = SharedFolder(varSharedFolder)

	return err
}

type NullableSharedFolder struct {
	value *SharedFolder
	isSet bool
}

func (v NullableSharedFolder) Get() *SharedFolder {
	return v.value
}

func (v *NullableSharedFolder) Set(val *SharedFolder) {
	v.value = val
	v.isSet = true
}

func (v NullableSharedFolder) IsSet() bool {
	return v.isSet
}

func (v *NullableSharedFolder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSharedFolder(val *SharedFolder) *NullableSharedFolder {
	return &NullableSharedFolder{value: val, isSet: true}
}

func (v NullableSharedFolder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSharedFolder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
