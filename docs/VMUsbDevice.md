# VMUsbDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **int32** | Number of items | [optional] [default to 0]
**Connected** | Pointer to **string** |  | [optional]
**BackingInfo** | Pointer to **string** |  | [optional]
**BackingType** | Pointer to **int32** | Number of items | [optional] [default to 0]

## Methods

### NewVMUsbDevice

`func NewVMUsbDevice() *VMUsbDevice`

NewVMUsbDevice instantiates a new VMUsbDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMUsbDeviceWithDefaults

`func NewVMUsbDeviceWithDefaults() *VMUsbDevice`

NewVMUsbDeviceWithDefaults instantiates a new VMUsbDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *VMUsbDevice) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *VMUsbDevice) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *VMUsbDevice) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *VMUsbDevice) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetConnected

`func (o *VMUsbDevice) GetConnected() string`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *VMUsbDevice) GetConnectedOk() (*string, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *VMUsbDevice) SetConnected(v string)`

SetConnected sets Connected field to given value.

### HasConnected

`func (o *VMUsbDevice) HasConnected() bool`

HasConnected returns a boolean if a field has been set.

### GetBackingInfo

`func (o *VMUsbDevice) GetBackingInfo() string`

GetBackingInfo returns the BackingInfo field if non-nil, zero value otherwise.

### GetBackingInfoOk

`func (o *VMUsbDevice) GetBackingInfoOk() (*string, bool)`

GetBackingInfoOk returns a tuple with the BackingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackingInfo

`func (o *VMUsbDevice) SetBackingInfo(v string)`

SetBackingInfo sets BackingInfo field to given value.

### HasBackingInfo

`func (o *VMUsbDevice) HasBackingInfo() bool`

HasBackingInfo returns a boolean if a field has been set.

### GetBackingType

`func (o *VMUsbDevice) GetBackingType() int32`

GetBackingType returns the BackingType field if non-nil, zero value otherwise.

### GetBackingTypeOk

`func (o *VMUsbDevice) GetBackingTypeOk() (*int32, bool)`

GetBackingTypeOk returns a tuple with the BackingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackingType

`func (o *VMUsbDevice) SetBackingType(v int32)`

SetBackingType sets BackingType field to given value.

### HasBackingType

`func (o *VMUsbDevice) HasBackingType() bool`

HasBackingType returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
