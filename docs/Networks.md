# Networks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Num** | **int32** | Number of items | [default to 0]
**Vmnets** | [**[]Network**](Network.md) | The list of virtual networks |

## Methods

### NewNetworks

`func NewNetworks(num int32, vmnets []Network, ) *Networks`

NewNetworks instantiates a new Networks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksWithDefaults

`func NewNetworksWithDefaults() *Networks`

NewNetworksWithDefaults instantiates a new Networks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNum

`func (o *Networks) GetNum() int32`

GetNum returns the Num field if non-nil, zero value otherwise.

### GetNumOk

`func (o *Networks) GetNumOk() (*int32, bool)`

GetNumOk returns a tuple with the Num field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNum

`func (o *Networks) SetNum(v int32)`

SetNum sets Num field to given value.

### GetVmnets

`func (o *Networks) GetVmnets() []Network`

GetVmnets returns the Vmnets field if non-nil, zero value otherwise.

### GetVmnetsOk

`func (o *Networks) GetVmnetsOk() (*[]Network, bool)`

GetVmnetsOk returns a tuple with the Vmnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmnets

`func (o *Networks) SetVmnets(v []Network)`

SetVmnets sets Vmnets field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
