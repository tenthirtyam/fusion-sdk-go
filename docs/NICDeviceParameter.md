# NICDeviceParameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The network type of network adapter |
**Vmnet** | **string** | The vmnet name, it should only be used while type is custom |

## Methods

### NewNICDeviceParameter

`func NewNICDeviceParameter(type_ string, vmnet string, ) *NICDeviceParameter`

NewNICDeviceParameter instantiates a new NICDeviceParameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNICDeviceParameterWithDefaults

`func NewNICDeviceParameterWithDefaults() *NICDeviceParameter`

NewNICDeviceParameterWithDefaults instantiates a new NICDeviceParameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NICDeviceParameter) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NICDeviceParameter) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NICDeviceParameter) SetType(v string)`

SetType sets Type field to given value.

### GetVmnet

`func (o *NICDeviceParameter) GetVmnet() string`

GetVmnet returns the Vmnet field if non-nil, zero value otherwise.

### GetVmnetOk

`func (o *NICDeviceParameter) GetVmnetOk() (*string, bool)`

GetVmnetOk returns a tuple with the Vmnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmnet

`func (o *NICDeviceParameter) SetVmnet(v string)`

SetVmnet sets Vmnet field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
