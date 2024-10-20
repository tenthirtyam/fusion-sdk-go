# PortforwardParameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GuestIp** | **string** |  |
**GuestPort** | **int32** | port of communication | [default to 0]
**Desc** | Pointer to **string** |  | [optional]

## Methods

### NewPortforwardParameter

`func NewPortforwardParameter(guestIp string, guestPort int32, ) *PortforwardParameter`

NewPortforwardParameter instantiates a new PortforwardParameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortforwardParameterWithDefaults

`func NewPortforwardParameterWithDefaults() *PortforwardParameter`

NewPortforwardParameterWithDefaults instantiates a new PortforwardParameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGuestIp

`func (o *PortforwardParameter) GetGuestIp() string`

GetGuestIp returns the GuestIp field if non-nil, zero value otherwise.

### GetGuestIpOk

`func (o *PortforwardParameter) GetGuestIpOk() (*string, bool)`

GetGuestIpOk returns a tuple with the GuestIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestIp

`func (o *PortforwardParameter) SetGuestIp(v string)`

SetGuestIp sets GuestIp field to given value.

### GetGuestPort

`func (o *PortforwardParameter) GetGuestPort() int32`

GetGuestPort returns the GuestPort field if non-nil, zero value otherwise.

### GetGuestPortOk

`func (o *PortforwardParameter) GetGuestPortOk() (*int32, bool)`

GetGuestPortOk returns a tuple with the GuestPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestPort

`func (o *PortforwardParameter) SetGuestPort(v int32)`

SetGuestPort sets GuestPort field to given value.

### GetDesc

`func (o *PortforwardParameter) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *PortforwardParameter) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *PortforwardParameter) SetDesc(v string)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *PortforwardParameter) HasDesc() bool`

HasDesc returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
