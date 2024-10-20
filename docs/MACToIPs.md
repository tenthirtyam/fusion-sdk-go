# MACToIPs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Num** | **int32** | Number of items | [default to 0]
**Mactoips** | Pointer to [**[]MACToIP**](MACToIP.md) | The list of MAC to IP settings | [optional]

## Methods

### NewMACToIPs

`func NewMACToIPs(num int32, ) *MACToIPs`

NewMACToIPs instantiates a new MACToIPs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMACToIPsWithDefaults

`func NewMACToIPsWithDefaults() *MACToIPs`

NewMACToIPsWithDefaults instantiates a new MACToIPs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNum

`func (o *MACToIPs) GetNum() int32`

GetNum returns the Num field if non-nil, zero value otherwise.

### GetNumOk

`func (o *MACToIPs) GetNumOk() (*int32, bool)`

GetNumOk returns a tuple with the Num field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNum

`func (o *MACToIPs) SetNum(v int32)`

SetNum sets Num field to given value.

### GetMactoips

`func (o *MACToIPs) GetMactoips() []MACToIP`

GetMactoips returns the Mactoips field if non-nil, zero value otherwise.

### GetMactoipsOk

`func (o *MACToIPs) GetMactoipsOk() (*[]MACToIP, bool)`

GetMactoipsOk returns a tuple with the Mactoips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMactoips

`func (o *MACToIPs) SetMactoips(v []MACToIP)`

SetMactoips sets Mactoips field to given value.

### HasMactoips

`func (o *MACToIPs) HasMactoips() bool`

HasMactoips returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
