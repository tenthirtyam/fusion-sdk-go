# Portforwards

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Num** | **int32** | Number of items | [default to 0]
**PortForwardings** | [**[]Portforward**](Portforward.md) | The list of port forwardings |

## Methods

### NewPortforwards

`func NewPortforwards(num int32, portForwardings []Portforward, ) *Portforwards`

NewPortforwards instantiates a new Portforwards object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortforwardsWithDefaults

`func NewPortforwardsWithDefaults() *Portforwards`

NewPortforwardsWithDefaults instantiates a new Portforwards object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNum

`func (o *Portforwards) GetNum() int32`

GetNum returns the Num field if non-nil, zero value otherwise.

### GetNumOk

`func (o *Portforwards) GetNumOk() (*int32, bool)`

GetNumOk returns a tuple with the Num field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNum

`func (o *Portforwards) SetNum(v int32)`

SetNum sets Num field to given value.

### GetPortForwardings

`func (o *Portforwards) GetPortForwardings() []Portforward`

GetPortForwardings returns the PortForwardings field if non-nil, zero value otherwise.

### GetPortForwardingsOk

`func (o *Portforwards) GetPortForwardingsOk() (*[]Portforward, bool)`

GetPortForwardingsOk returns a tuple with the PortForwardings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortForwardings

`func (o *Portforwards) SetPortForwardings(v []Portforward)`

SetPortForwardings sets PortForwardings field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
