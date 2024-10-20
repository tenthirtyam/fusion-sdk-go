# DnsConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | Pointer to **string** |  | [optional]
**Domainname** | Pointer to **string** |  | [optional]
**Server** | Pointer to **[]string** |  | [optional]
**Search** | Pointer to **[]string** |  | [optional]

## Methods

### NewDnsConfig

`func NewDnsConfig() *DnsConfig`

NewDnsConfig instantiates a new DnsConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsConfigWithDefaults

`func NewDnsConfigWithDefaults() *DnsConfig`

NewDnsConfigWithDefaults instantiates a new DnsConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *DnsConfig) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *DnsConfig) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *DnsConfig) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *DnsConfig) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetDomainname

`func (o *DnsConfig) GetDomainname() string`

GetDomainname returns the Domainname field if non-nil, zero value otherwise.

### GetDomainnameOk

`func (o *DnsConfig) GetDomainnameOk() (*string, bool)`

GetDomainnameOk returns a tuple with the Domainname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainname

`func (o *DnsConfig) SetDomainname(v string)`

SetDomainname sets Domainname field to given value.

### HasDomainname

`func (o *DnsConfig) HasDomainname() bool`

HasDomainname returns a boolean if a field has been set.

### GetServer

`func (o *DnsConfig) GetServer() []string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *DnsConfig) GetServerOk() (*[]string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *DnsConfig) SetServer(v []string)`

SetServer sets Server field to given value.

### HasServer

`func (o *DnsConfig) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetSearch

`func (o *DnsConfig) GetSearch() []string`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *DnsConfig) GetSearchOk() (*[]string, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *DnsConfig) SetSearch(v []string)`

SetSearch sets Search field to given value.

### HasSearch

`func (o *DnsConfig) HasSearch() bool`

HasSearch returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
