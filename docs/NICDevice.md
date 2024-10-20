# NICDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | **int32** | Index of Network Adapters | [default to 1]
**Type** | **string** | The network type of network adapter |
**Vmnet** | **string** | The vmnet name |
**MacAddress** | **string** | Mac address |

## Methods

### NewNICDevice

`func NewNICDevice(index int32, type_ string, vmnet string, macAddress string, ) *NICDevice`

NewNICDevice instantiates a new NICDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNICDeviceWithDefaults

`func NewNICDeviceWithDefaults() *NICDevice`

NewNICDeviceWithDefaults instantiates a new NICDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *NICDevice) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *NICDevice) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *NICDevice) SetIndex(v int32)`

SetIndex sets Index field to given value.

### GetType

`func (o *NICDevice) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NICDevice) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NICDevice) SetType(v string)`

SetType sets Type field to given value.

### GetVmnet

`func (o *NICDevice) GetVmnet() string`

GetVmnet returns the Vmnet field if non-nil, zero value otherwise.

### GetVmnetOk

`func (o *NICDevice) GetVmnetOk() (*string, bool)`

GetVmnetOk returns a tuple with the Vmnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmnet

`func (o *NICDevice) SetVmnet(v string)`

SetVmnet sets Vmnet field to given value.

### GetMacAddress

`func (o *NICDevice) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *NICDevice) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *NICDevice) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
