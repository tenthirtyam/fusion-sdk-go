# SharedFolder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FolderId** | **string** | ID of folder which be mounted to the host |
**HostPath** | **string** | Path of the host shared folder |
**Flags** | **int32** | The flags property specifies how the folder will be accessed by the VM. There is only one flag supported which is \&quot;4\&quot; and means read/write access.  |

## Methods

### NewSharedFolder

`func NewSharedFolder(folderId string, hostPath string, flags int32, ) *SharedFolder`

NewSharedFolder instantiates a new SharedFolder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharedFolderWithDefaults

`func NewSharedFolderWithDefaults() *SharedFolder`

NewSharedFolderWithDefaults instantiates a new SharedFolder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFolderId

`func (o *SharedFolder) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *SharedFolder) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *SharedFolder) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### GetHostPath

`func (o *SharedFolder) GetHostPath() string`

GetHostPath returns the HostPath field if non-nil, zero value otherwise.

### GetHostPathOk

`func (o *SharedFolder) GetHostPathOk() (*string, bool)`

GetHostPathOk returns a tuple with the HostPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostPath

`func (o *SharedFolder) SetHostPath(v string)`

SetHostPath sets HostPath field to given value.

### GetFlags

`func (o *SharedFolder) GetFlags() int32`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *SharedFolder) GetFlagsOk() (*int32, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *SharedFolder) SetFlags(v int32)`

SetFlags sets Flags field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
