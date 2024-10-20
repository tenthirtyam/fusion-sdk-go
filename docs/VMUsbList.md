# VMUsbList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Num** | Pointer to **int32** | Number of items | [optional] [default to 0]
**UsbDevices** | Pointer to [**[]VMUsbDevice**](VMUsbDevice.md) |  | [optional]

## Methods

### NewVMUsbList

`func NewVMUsbList() *VMUsbList`

NewVMUsbList instantiates a new VMUsbList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMUsbListWithDefaults

`func NewVMUsbListWithDefaults() *VMUsbList`

NewVMUsbListWithDefaults instantiates a new VMUsbList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNum

`func (o *VMUsbList) GetNum() int32`

GetNum returns the Num field if non-nil, zero value otherwise.

### GetNumOk

`func (o *VMUsbList) GetNumOk() (*int32, bool)`

GetNumOk returns a tuple with the Num field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNum

`func (o *VMUsbList) SetNum(v int32)`

SetNum sets Num field to given value.

### HasNum

`func (o *VMUsbList) HasNum() bool`

HasNum returns a boolean if a field has been set.

### GetUsbDevices

`func (o *VMUsbList) GetUsbDevices() []VMUsbDevice`

GetUsbDevices returns the UsbDevices field if non-nil, zero value otherwise.

### GetUsbDevicesOk

`func (o *VMUsbList) GetUsbDevicesOk() (*[]VMUsbDevice, bool)`

GetUsbDevicesOk returns a tuple with the UsbDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsbDevices

`func (o *VMUsbList) SetUsbDevices(v []VMUsbDevice)`

SetUsbDevices sets UsbDevices field to given value.

### HasUsbDevices

`func (o *VMUsbList) HasUsbDevices() bool`

HasUsbDevices returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
