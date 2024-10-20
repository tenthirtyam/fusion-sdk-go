# VMApplianceView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Author** | Pointer to **string** |  | [optional]
**Version** | Pointer to **string** |  | [optional]
**Port** | Pointer to **int32** | port of communication | [optional] [default to 0]
**ShowAtPowerOn** | Pointer to **string** |  | [optional]

## Methods

### NewVMApplianceView

`func NewVMApplianceView() *VMApplianceView`

NewVMApplianceView instantiates a new VMApplianceView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMApplianceViewWithDefaults

`func NewVMApplianceViewWithDefaults() *VMApplianceView`

NewVMApplianceViewWithDefaults instantiates a new VMApplianceView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthor

`func (o *VMApplianceView) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *VMApplianceView) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *VMApplianceView) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *VMApplianceView) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetVersion

`func (o *VMApplianceView) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *VMApplianceView) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *VMApplianceView) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *VMApplianceView) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetPort

`func (o *VMApplianceView) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *VMApplianceView) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *VMApplianceView) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *VMApplianceView) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetShowAtPowerOn

`func (o *VMApplianceView) GetShowAtPowerOn() string`

GetShowAtPowerOn returns the ShowAtPowerOn field if non-nil, zero value otherwise.

### GetShowAtPowerOnOk

`func (o *VMApplianceView) GetShowAtPowerOnOk() (*string, bool)`

GetShowAtPowerOnOk returns a tuple with the ShowAtPowerOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowAtPowerOn

`func (o *VMApplianceView) SetShowAtPowerOn(v string)`

SetShowAtPowerOn sets ShowAtPowerOn field to given value.

### HasShowAtPowerOn

`func (o *VMApplianceView) HasShowAtPowerOn() bool`

HasShowAtPowerOn returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
