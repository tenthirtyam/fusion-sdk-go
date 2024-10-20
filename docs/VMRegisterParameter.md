# VMRegisterParameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Register VM name | [optional]
**Path** | Pointer to **string** | Register VM path | [optional]

## Methods

### NewVMRegisterParameter

`func NewVMRegisterParameter() *VMRegisterParameter`

NewVMRegisterParameter instantiates a new VMRegisterParameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMRegisterParameterWithDefaults

`func NewVMRegisterParameterWithDefaults() *VMRegisterParameter`

NewVMRegisterParameterWithDefaults instantiates a new VMRegisterParameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VMRegisterParameter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VMRegisterParameter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VMRegisterParameter) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VMRegisterParameter) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPath

`func (o *VMRegisterParameter) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *VMRegisterParameter) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *VMRegisterParameter) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *VMRegisterParameter) HasPath() bool`

HasPath returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
