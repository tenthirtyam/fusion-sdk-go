# VMConnectedDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **int32** | Number of items | [optional] [default to 0]
**StartConnected** | Pointer to **string** |  | [optional]
**ConnectionStatus** | Pointer to **int32** | Number of items | [optional] [default to 0]
**DevicePath** | Pointer to **string** |  | [optional]

## Methods

### NewVMConnectedDevice

`func NewVMConnectedDevice() *VMConnectedDevice`

NewVMConnectedDevice instantiates a new VMConnectedDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMConnectedDeviceWithDefaults

`func NewVMConnectedDeviceWithDefaults() *VMConnectedDevice`

NewVMConnectedDeviceWithDefaults instantiates a new VMConnectedDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *VMConnectedDevice) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *VMConnectedDevice) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *VMConnectedDevice) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *VMConnectedDevice) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetStartConnected

`func (o *VMConnectedDevice) GetStartConnected() string`

GetStartConnected returns the StartConnected field if non-nil, zero value otherwise.

### GetStartConnectedOk

`func (o *VMConnectedDevice) GetStartConnectedOk() (*string, bool)`

GetStartConnectedOk returns a tuple with the StartConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartConnected

`func (o *VMConnectedDevice) SetStartConnected(v string)`

SetStartConnected sets StartConnected field to given value.

### HasStartConnected

`func (o *VMConnectedDevice) HasStartConnected() bool`

HasStartConnected returns a boolean if a field has been set.

### GetConnectionStatus

`func (o *VMConnectedDevice) GetConnectionStatus() int32`

GetConnectionStatus returns the ConnectionStatus field if non-nil, zero value otherwise.

### GetConnectionStatusOk

`func (o *VMConnectedDevice) GetConnectionStatusOk() (*int32, bool)`

GetConnectionStatusOk returns a tuple with the ConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatus

`func (o *VMConnectedDevice) SetConnectionStatus(v int32)`

SetConnectionStatus sets ConnectionStatus field to given value.

### HasConnectionStatus

`func (o *VMConnectedDevice) HasConnectionStatus() bool`

HasConnectionStatus returns a boolean if a field has been set.

### GetDevicePath

`func (o *VMConnectedDevice) GetDevicePath() string`

GetDevicePath returns the DevicePath field if non-nil, zero value otherwise.

### GetDevicePathOk

`func (o *VMConnectedDevice) GetDevicePathOk() (*string, bool)`

GetDevicePathOk returns a tuple with the DevicePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicePath

`func (o *VMConnectedDevice) SetDevicePath(v string)`

SetDevicePath sets DevicePath field to given value.

### HasDevicePath

`func (o *VMConnectedDevice) HasDevicePath() bool`

HasDevicePath returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
