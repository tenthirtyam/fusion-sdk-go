# NICDevices

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Num** | **int32** | Number of NIC devices | [default to 1]
**Nics** | [**[]NICDevice**](NICDevice.md) | The network adapter added to this VM |

## Methods

### NewNICDevices

`func NewNICDevices(num int32, nics []NICDevice, ) *NICDevices`

NewNICDevices instantiates a new NICDevices object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNICDevicesWithDefaults

`func NewNICDevicesWithDefaults() *NICDevices`

NewNICDevicesWithDefaults instantiates a new NICDevices object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNum

`func (o *NICDevices) GetNum() int32`

GetNum returns the Num field if non-nil, zero value otherwise.

### GetNumOk

`func (o *NICDevices) GetNumOk() (*int32, bool)`

GetNumOk returns a tuple with the Num field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNum

`func (o *NICDevices) SetNum(v int32)`

SetNum sets Num field to given value.

### GetNics

`func (o *NICDevices) GetNics() []NICDevice`

GetNics returns the Nics field if non-nil, zero value otherwise.

### GetNicsOk

`func (o *NICDevices) GetNicsOk() (*[]NICDevice, bool)`

GetNicsOk returns a tuple with the Nics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNics

`func (o *NICDevices) SetNics(v []NICDevice)`

SetNics sets Nics field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
