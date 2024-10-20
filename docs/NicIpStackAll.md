# NicIpStackAll

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nics** | Pointer to [**NicIpStack**](NicIpStack.md) |  | [optional]
**Routes** | Pointer to [**[]RouteEntry**](RouteEntry.md) |  | [optional]
**Dns** | Pointer to [**DnsConfig**](DnsConfig.md) |  | [optional]
**Wins** | Pointer to [**WinsConfig**](WinsConfig.md) |  | [optional]
**Dhcpv4** | Pointer to [**DhcpConfig**](DhcpConfig.md) |  | [optional]
**Dhcpv6** | Pointer to [**DhcpConfig**](DhcpConfig.md) |  | [optional]

## Methods

### NewNicIpStackAll

`func NewNicIpStackAll() *NicIpStackAll`

NewNicIpStackAll instantiates a new NicIpStackAll object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNicIpStackAllWithDefaults

`func NewNicIpStackAllWithDefaults() *NicIpStackAll`

NewNicIpStackAllWithDefaults instantiates a new NicIpStackAll object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNics

`func (o *NicIpStackAll) GetNics() NicIpStack`

GetNics returns the Nics field if non-nil, zero value otherwise.

### GetNicsOk

`func (o *NicIpStackAll) GetNicsOk() (*NicIpStack, bool)`

GetNicsOk returns a tuple with the Nics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNics

`func (o *NicIpStackAll) SetNics(v NicIpStack)`

SetNics sets Nics field to given value.

### HasNics

`func (o *NicIpStackAll) HasNics() bool`

HasNics returns a boolean if a field has been set.

### GetRoutes

`func (o *NicIpStackAll) GetRoutes() []RouteEntry`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *NicIpStackAll) GetRoutesOk() (*[]RouteEntry, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *NicIpStackAll) SetRoutes(v []RouteEntry)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *NicIpStackAll) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.

### GetDns

`func (o *NicIpStackAll) GetDns() DnsConfig`

GetDns returns the Dns field if non-nil, zero value otherwise.

### GetDnsOk

`func (o *NicIpStackAll) GetDnsOk() (*DnsConfig, bool)`

GetDnsOk returns a tuple with the Dns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns

`func (o *NicIpStackAll) SetDns(v DnsConfig)`

SetDns sets Dns field to given value.

### HasDns

`func (o *NicIpStackAll) HasDns() bool`

HasDns returns a boolean if a field has been set.

### GetWins

`func (o *NicIpStackAll) GetWins() WinsConfig`

GetWins returns the Wins field if non-nil, zero value otherwise.

### GetWinsOk

`func (o *NicIpStackAll) GetWinsOk() (*WinsConfig, bool)`

GetWinsOk returns a tuple with the Wins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWins

`func (o *NicIpStackAll) SetWins(v WinsConfig)`

SetWins sets Wins field to given value.

### HasWins

`func (o *NicIpStackAll) HasWins() bool`

HasWins returns a boolean if a field has been set.

### GetDhcpv4

`func (o *NicIpStackAll) GetDhcpv4() DhcpConfig`

GetDhcpv4 returns the Dhcpv4 field if non-nil, zero value otherwise.

### GetDhcpv4Ok

`func (o *NicIpStackAll) GetDhcpv4Ok() (*DhcpConfig, bool)`

GetDhcpv4Ok returns a tuple with the Dhcpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpv4

`func (o *NicIpStackAll) SetDhcpv4(v DhcpConfig)`

SetDhcpv4 sets Dhcpv4 field to given value.

### HasDhcpv4

`func (o *NicIpStackAll) HasDhcpv4() bool`

HasDhcpv4 returns a boolean if a field has been set.

### GetDhcpv6

`func (o *NicIpStackAll) GetDhcpv6() DhcpConfig`

GetDhcpv6 returns the Dhcpv6 field if non-nil, zero value otherwise.

### GetDhcpv6Ok

`func (o *NicIpStackAll) GetDhcpv6Ok() (*DhcpConfig, bool)`

GetDhcpv6Ok returns a tuple with the Dhcpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpv6

`func (o *NicIpStackAll) SetDhcpv6(v DhcpConfig)`

SetDhcpv6 sets Dhcpv6 field to given value.

### HasDhcpv6

`func (o *NicIpStackAll) HasDhcpv6() bool`

HasDhcpv6 returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
