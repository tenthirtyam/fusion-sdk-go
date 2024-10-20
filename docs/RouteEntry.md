# RouteEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dest** | **string** | IP address |
**Prefix** | **int32** | Number of items | [default to 0]
**Nexthop** | Pointer to **string** | IP address | [optional]
**Interface** | **int32** | Number of items | [default to 0]
**Type** | **int32** | Number of items | [default to 0]
**Metric** | **int32** | Number of items | [default to 0]

## Methods

### NewRouteEntry

`func NewRouteEntry(dest string, prefix int32, interface_ int32, type_ int32, metric int32, ) *RouteEntry`

NewRouteEntry instantiates a new RouteEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteEntryWithDefaults

`func NewRouteEntryWithDefaults() *RouteEntry`

NewRouteEntryWithDefaults instantiates a new RouteEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDest

`func (o *RouteEntry) GetDest() string`

GetDest returns the Dest field if non-nil, zero value otherwise.

### GetDestOk

`func (o *RouteEntry) GetDestOk() (*string, bool)`

GetDestOk returns a tuple with the Dest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDest

`func (o *RouteEntry) SetDest(v string)`

SetDest sets Dest field to given value.

### GetPrefix

`func (o *RouteEntry) GetPrefix() int32`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *RouteEntry) GetPrefixOk() (*int32, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *RouteEntry) SetPrefix(v int32)`

SetPrefix sets Prefix field to given value.

### GetNexthop

`func (o *RouteEntry) GetNexthop() string`

GetNexthop returns the Nexthop field if non-nil, zero value otherwise.

### GetNexthopOk

`func (o *RouteEntry) GetNexthopOk() (*string, bool)`

GetNexthopOk returns a tuple with the Nexthop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNexthop

`func (o *RouteEntry) SetNexthop(v string)`

SetNexthop sets Nexthop field to given value.

### HasNexthop

`func (o *RouteEntry) HasNexthop() bool`

HasNexthop returns a boolean if a field has been set.

### GetInterface

`func (o *RouteEntry) GetInterface() int32`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *RouteEntry) GetInterfaceOk() (*int32, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *RouteEntry) SetInterface(v int32)`

SetInterface sets Interface field to given value.

### GetType

`func (o *RouteEntry) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RouteEntry) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RouteEntry) SetType(v int32)`

SetType sets Type field to given value.

### GetMetric

`func (o *RouteEntry) GetMetric() int32`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *RouteEntry) GetMetricOk() (*int32, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *RouteEntry) SetMetric(v int32)`

SetMetric sets Metric field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
