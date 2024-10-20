# NicIpStack

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | **string** | Mac address, E.g., de:ad:be:ef:12:34 |
**Ip** | Pointer to **[]string** | IP address(es) of the interface (CIDR) | [optional]
**Dns** | Pointer to [**DnsConfig**](DnsConfig.md) |  | [optional]
**Wins** | Pointer to [**WinsConfig**](WinsConfig.md) |  | [optional]
**Dhcp4** | Pointer to [**DhcpConfig**](DhcpConfig.md) |  | [optional]
**Dhcp6** | Pointer to [**DhcpConfig**](DhcpConfig.md) |  | [optional]

## Methods

### NewNicIpStack

`func NewNicIpStack(mac string, ) *NicIpStack`

NewNicIpStack instantiates a new NicIpStack object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNicIpStackWithDefaults

`func NewNicIpStackWithDefaults() *NicIpStack`

NewNicIpStackWithDefaults instantiates a new NicIpStack object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *NicIpStack) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *NicIpStack) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *NicIpStack) SetMac(v string)`

SetMac sets Mac field to given value.

### GetIp

`func (o *NicIpStack) GetIp() []string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *NicIpStack) GetIpOk() (*[]string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *NicIpStack) SetIp(v []string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *NicIpStack) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetDns

`func (o *NicIpStack) GetDns() DnsConfig`

GetDns returns the Dns field if non-nil, zero value otherwise.

### GetDnsOk

`func (o *NicIpStack) GetDnsOk() (*DnsConfig, bool)`

GetDnsOk returns a tuple with the Dns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns

`func (o *NicIpStack) SetDns(v DnsConfig)`

SetDns sets Dns field to given value.

### HasDns

`func (o *NicIpStack) HasDns() bool`

HasDns returns a boolean if a field has been set.

### GetWins

`func (o *NicIpStack) GetWins() WinsConfig`

GetWins returns the Wins field if non-nil, zero value otherwise.

### GetWinsOk

`func (o *NicIpStack) GetWinsOk() (*WinsConfig, bool)`

GetWinsOk returns a tuple with the Wins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWins

`func (o *NicIpStack) SetWins(v WinsConfig)`

SetWins sets Wins field to given value.

### HasWins

`func (o *NicIpStack) HasWins() bool`

HasWins returns a boolean if a field has been set.

### GetDhcp4

`func (o *NicIpStack) GetDhcp4() DhcpConfig`

GetDhcp4 returns the Dhcp4 field if non-nil, zero value otherwise.

### GetDhcp4Ok

`func (o *NicIpStack) GetDhcp4Ok() (*DhcpConfig, bool)`

GetDhcp4Ok returns a tuple with the Dhcp4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcp4

`func (o *NicIpStack) SetDhcp4(v DhcpConfig)`

SetDhcp4 sets Dhcp4 field to given value.

### HasDhcp4

`func (o *NicIpStack) HasDhcp4() bool`

HasDhcp4 returns a boolean if a field has been set.

### GetDhcp6

`func (o *NicIpStack) GetDhcp6() DhcpConfig`

GetDhcp6 returns the Dhcp6 field if non-nil, zero value otherwise.

### GetDhcp6Ok

`func (o *NicIpStack) GetDhcp6Ok() (*DhcpConfig, bool)`

GetDhcp6Ok returns a tuple with the Dhcp6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcp6

`func (o *NicIpStack) SetDhcp6(v DhcpConfig)`

SetDhcp6 sets Dhcp6 field to given value.

### HasDhcp6

`func (o *NicIpStack) HasDhcp6() bool`

HasDhcp6 returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
