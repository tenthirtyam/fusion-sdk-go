# CreateVmnetParameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The host network name |
**Type** | Pointer to **string** | The host network type | [optional]

## Methods

### NewCreateVmnetParameter

`func NewCreateVmnetParameter(name string, ) *CreateVmnetParameter`

NewCreateVmnetParameter instantiates a new CreateVmnetParameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVmnetParameterWithDefaults

`func NewCreateVmnetParameterWithDefaults() *CreateVmnetParameter`

NewCreateVmnetParameterWithDefaults instantiates a new CreateVmnetParameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateVmnetParameter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateVmnetParameter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateVmnetParameter) SetName(v string)`

SetName sets Name field to given value.

### GetType

`func (o *CreateVmnetParameter) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateVmnetParameter) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateVmnetParameter) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateVmnetParameter) HasType() bool`

HasType returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
