// © Ryan Johnson. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

package fusion

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the ErrorModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorModel{}

// ErrorModel struct for ErrorModel
type ErrorModel struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

type _ErrorModel ErrorModel

// NewErrorModel instantiates a new ErrorModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorModel(code int32, message string) *ErrorModel {
	this := ErrorModel{}
	this.Code = code
	this.Message = message
	return &this
}

// NewErrorModelWithDefaults instantiates a new ErrorModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorModelWithDefaults() *ErrorModel {
	this := ErrorModel{}
	return &this
}

// GetCode returns the Code field value
func (o *ErrorModel) GetCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *ErrorModel) GetCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *ErrorModel) SetCode(v int32) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *ErrorModel) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ErrorModel) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ErrorModel) SetMessage(v string) {
	o.Message = v
}

func (o ErrorModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

func (o *ErrorModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"code",
		"message",
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

	varErrorModel := _ErrorModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varErrorModel)

	if err != nil {
		return err
	}

	*o = ErrorModel(varErrorModel)

	return err
}

type NullableErrorModel struct {
	value *ErrorModel
	isSet bool
}

func (v NullableErrorModel) Get() *ErrorModel {
	return v.value
}

func (v *NullableErrorModel) Set(val *ErrorModel) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorModel) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorModel(val *ErrorModel) *NullableErrorModel {
	return &NullableErrorModel{value: val, isSet: true}
}

func (v NullableErrorModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
