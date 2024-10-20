# PortforwardGuest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | **string** |  |
**Port** | **int32** | port of communication | [default to 0]

## Methods

### NewPortforwardGuest

`func NewPortforwardGuest(ip string, port int32, ) *PortforwardGuest`

NewPortforwardGuest instantiates a new PortforwardGuest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortforwardGuestWithDefaults

`func NewPortforwardGuestWithDefaults() *PortforwardGuest`

NewPortforwardGuestWithDefaults instantiates a new PortforwardGuest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *PortforwardGuest) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *PortforwardGuest) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *PortforwardGuest) SetIp(v string)`

SetIp sets Ip field to given value.

### GetPort

`func (o *PortforwardGuest) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *PortforwardGuest) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *PortforwardGuest) SetPort(v int32)`

SetPort sets Port field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
