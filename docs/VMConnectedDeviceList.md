# VMConnectedDeviceList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Num** | Pointer to **int32** | Number of items | [optional] [default to 0]
**Devices** | Pointer to [**[]VMConnectedDevice**](VMConnectedDevice.md) |  | [optional]

## Methods

### NewVMConnectedDeviceList

`func NewVMConnectedDeviceList() *VMConnectedDeviceList`

NewVMConnectedDeviceList instantiates a new VMConnectedDeviceList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMConnectedDeviceListWithDefaults

`func NewVMConnectedDeviceListWithDefaults() *VMConnectedDeviceList`

NewVMConnectedDeviceListWithDefaults instantiates a new VMConnectedDeviceList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNum

`func (o *VMConnectedDeviceList) GetNum() int32`

GetNum returns the Num field if non-nil, zero value otherwise.

### GetNumOk

`func (o *VMConnectedDeviceList) GetNumOk() (*int32, bool)`

GetNumOk returns a tuple with the Num field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNum

`func (o *VMConnectedDeviceList) SetNum(v int32)`

SetNum sets Num field to given value.

### HasNum

`func (o *VMConnectedDeviceList) HasNum() bool`

HasNum returns a boolean if a field has been set.

### GetDevices

`func (o *VMConnectedDeviceList) GetDevices() []VMConnectedDevice`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *VMConnectedDeviceList) GetDevicesOk() (*[]VMConnectedDevice, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *VMConnectedDeviceList) SetDevices(v []VMConnectedDevice)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *VMConnectedDeviceList) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
