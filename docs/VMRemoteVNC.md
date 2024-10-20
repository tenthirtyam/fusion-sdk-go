# VMRemoteVNC

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VNCEnabled** | Pointer to **string** |  | [optional]
**VNCPort** | Pointer to **int32** | port of communication | [optional] [default to 0]

## Methods

### NewVMRemoteVNC

`func NewVMRemoteVNC() *VMRemoteVNC`

NewVMRemoteVNC instantiates a new VMRemoteVNC object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMRemoteVNCWithDefaults

`func NewVMRemoteVNCWithDefaults() *VMRemoteVNC`

NewVMRemoteVNCWithDefaults instantiates a new VMRemoteVNC object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVNCEnabled

`func (o *VMRemoteVNC) GetVNCEnabled() string`

GetVNCEnabled returns the VNCEnabled field if non-nil, zero value otherwise.

### GetVNCEnabledOk

`func (o *VMRemoteVNC) GetVNCEnabledOk() (*string, bool)`

GetVNCEnabledOk returns a tuple with the VNCEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVNCEnabled

`func (o *VMRemoteVNC) SetVNCEnabled(v string)`

SetVNCEnabled sets VNCEnabled field to given value.

### HasVNCEnabled

`func (o *VMRemoteVNC) HasVNCEnabled() bool`

HasVNCEnabled returns a boolean if a field has been set.

### GetVNCPort

`func (o *VMRemoteVNC) GetVNCPort() int32`

GetVNCPort returns the VNCPort field if non-nil, zero value otherwise.

### GetVNCPortOk

`func (o *VMRemoteVNC) GetVNCPortOk() (*int32, bool)`

GetVNCPortOk returns a tuple with the VNCPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVNCPort

`func (o *VMRemoteVNC) SetVNCPort(v int32)`

SetVNCPort sets VNCPort field to given value.

### HasVNCPort

`func (o *VMRemoteVNC) HasVNCPort() bool`

HasVNCPort returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
