# Portforward

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | **int32** | port of communication | [default to 0]
**Protocol** | **string** |  |
**Desc** | **string** |  |
**Guest** | [**PortforwardGuest**](PortforwardGuest.md) |  |

## Methods

### NewPortforward

`func NewPortforward(port int32, protocol string, desc string, guest PortforwardGuest, ) *Portforward`

NewPortforward instantiates a new Portforward object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortforwardWithDefaults

`func NewPortforwardWithDefaults() *Portforward`

NewPortforwardWithDefaults instantiates a new Portforward object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPort

`func (o *Portforward) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *Portforward) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *Portforward) SetPort(v int32)`

SetPort sets Port field to given value.

### GetProtocol

`func (o *Portforward) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *Portforward) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *Portforward) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### GetDesc

`func (o *Portforward) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *Portforward) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *Portforward) SetDesc(v string)`

SetDesc sets Desc field to given value.

### GetGuest

`func (o *Portforward) GetGuest() PortforwardGuest`

GetGuest returns the Guest field if non-nil, zero value otherwise.

### GetGuestOk

`func (o *Portforward) GetGuestOk() (*PortforwardGuest, bool)`

GetGuestOk returns a tuple with the Guest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuest

`func (o *Portforward) SetGuest(v PortforwardGuest)`

SetGuest sets Guest field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
