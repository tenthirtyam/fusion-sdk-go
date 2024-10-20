# DhcpConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** |  |
**Setting** | **string** |  |

## Methods

### NewDhcpConfig

`func NewDhcpConfig(enabled bool, setting string, ) *DhcpConfig`

NewDhcpConfig instantiates a new DhcpConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpConfigWithDefaults

`func NewDhcpConfigWithDefaults() *DhcpConfig`

NewDhcpConfigWithDefaults instantiates a new DhcpConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *DhcpConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DhcpConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DhcpConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### GetSetting

`func (o *DhcpConfig) GetSetting() string`

GetSetting returns the Setting field if non-nil, zero value otherwise.

### GetSettingOk

`func (o *DhcpConfig) GetSettingOk() (*string, bool)`

GetSettingOk returns a tuple with the Setting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetting

`func (o *DhcpConfig) SetSetting(v string)`

SetSetting sets Setting field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
