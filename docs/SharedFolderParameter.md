# SharedFolderParameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostPath** | **string** | Path of the host shared folder |
**Flags** | **int32** | The flags property specifies how the folder will be accessed by the VM. There is only one flag supported which is \&quot;4\&quot; and means read/write access.  |

## Methods

### NewSharedFolderParameter

`func NewSharedFolderParameter(hostPath string, flags int32, ) *SharedFolderParameter`

NewSharedFolderParameter instantiates a new SharedFolderParameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharedFolderParameterWithDefaults

`func NewSharedFolderParameterWithDefaults() *SharedFolderParameter`

NewSharedFolderParameterWithDefaults instantiates a new SharedFolderParameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostPath

`func (o *SharedFolderParameter) GetHostPath() string`

GetHostPath returns the HostPath field if non-nil, zero value otherwise.

### GetHostPathOk

`func (o *SharedFolderParameter) GetHostPathOk() (*string, bool)`

GetHostPathOk returns a tuple with the HostPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostPath

`func (o *SharedFolderParameter) SetHostPath(v string)`

SetHostPath sets HostPath field to given value.

### GetFlags

`func (o *SharedFolderParameter) GetFlags() int32`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *SharedFolderParameter) GetFlagsOk() (*int32, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *SharedFolderParameter) SetFlags(v int32)`

SetFlags sets Flags field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
