# \VMSharedFoldersManagementAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSharedFolder**](VMSharedFoldersManagementAPI.md#CreateSharedFolder) | **Post** /vms/{id}/sharedfolders | Mounts a new shared folder in the VM
[**DeleteSharedFolder**](VMSharedFoldersManagementAPI.md#DeleteSharedFolder) | **Delete** /vms/{id}/sharedfolders/{folder id} | Deletes a shared folder
[**GetAllSharedFolders**](VMSharedFoldersManagementAPI.md#GetAllSharedFolders) | **Get** /vms/{id}/sharedfolders | Returns all shared folders mounted in the VM
[**UpdataSharedFolder**](VMSharedFoldersManagementAPI.md#UpdataSharedFolder) | **Put** /vms/{id}/sharedfolders/{folder id} | Updates a shared folder mounted in the VM

## CreateSharedFolder

> []SharedFolder CreateSharedFolder(ctx, id).Parameters(parameters).VmPassword(vmPassword).Execute()

Mounts a new shared folder in the VM

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	fusionClient "github.com/tenthirtyam/fusion-sdk-go/fusion"
)

func main() {
	id := "id_example" // string | ID of VM
	parameters := *fusionClient.NewSharedFolder("FolderId_example", "HostPath_example", int32(123)) // SharedFolder | Parameters of the shared folder to mount
	vmPassword := "vmPassword_example" // string | VM password (optional)

	configuration := fusionClient.NewConfiguration()
	apiClient := fusionClient.NewAPIClient(configuration)
	resp, r, err := apiClient.VMSharedFoldersManagementAPI.CreateSharedFolder(context.Background(), id).Parameters(parameters).VmPassword(vmPassword).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMSharedFoldersManagementAPI.CreateSharedFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSharedFolder`: []SharedFolder
	fmt.Fprintf(os.Stdout, "Response from `VMSharedFoldersManagementAPI.CreateSharedFolder`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of VM |

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSharedFolderRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **parameters** | [**SharedFolder**](SharedFolder.md) | Parameters of the shared folder to mount |
 **vmPassword** | **string** | VM password |

### Return type

[**[]SharedFolder**](SharedFolder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.vmware.vmw.rest-v1+json
- **Accept**: application/vnd.vmware.vmw.rest-v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## DeleteSharedFolder

> DeleteSharedFolder(ctx, id, folderId).VmPassword(vmPassword).Execute()

Deletes a shared folder

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	fusionClient "github.com/tenthirtyam/fusion-sdk-go/fusion"
)

func main() {
	id := "id_example" // string | ID of VM
	folderId := "folderId_example" // string | ID of shared folder
	vmPassword := "vmPassword_example" // string | VM password (optional)

	configuration := fusionClient.NewConfiguration()
	apiClient := fusionClient.NewAPIClient(configuration)
	r, err := apiClient.VMSharedFoldersManagementAPI.DeleteSharedFolder(context.Background(), id, folderId).VmPassword(vmPassword).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMSharedFoldersManagementAPI.DeleteSharedFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of VM |
**folderId** | **string** | ID of shared folder |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSharedFolderRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vmPassword** | **string** | VM password |

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.vmware.vmw.rest-v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetAllSharedFolders

> []SharedFolder GetAllSharedFolders(ctx, id).VmPassword(vmPassword).Execute()

Returns all shared folders mounted in the VM

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	fusionClient "github.com/tenthirtyam/fusion-sdk-go/fusion"
)

func main() {
	id := "id_example" // string | ID of VM
	vmPassword := "vmPassword_example" // string | VM password (optional)

	configuration := fusionClient.NewConfiguration()
	apiClient := fusionClient.NewAPIClient(configuration)
	resp, r, err := apiClient.VMSharedFoldersManagementAPI.GetAllSharedFolders(context.Background(), id).VmPassword(vmPassword).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMSharedFoldersManagementAPI.GetAllSharedFolders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllSharedFolders`: []SharedFolder
	fmt.Fprintf(os.Stdout, "Response from `VMSharedFoldersManagementAPI.GetAllSharedFolders`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of VM |

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllSharedFoldersRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vmPassword** | **string** | VM password |

### Return type

[**[]SharedFolder**](SharedFolder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.vmware.vmw.rest-v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdataSharedFolder

> []SharedFolder UpdataSharedFolder(ctx, id, folderId).Parameters(parameters).VmPassword(vmPassword).Execute()

Updates a shared folder mounted in the VM

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	fusionClient "github.com/tenthirtyam/fusion-sdk-go/fusion"
)

func main() {
	id := "id_example" // string | ID of VM
	folderId := "folderId_example" // string | ID of VM shared folder
	parameters := *fusionClient.NewSharedFolderParameter("HostPath_example", int32(123)) // SharedFolderParameter | Parameters of the shared folder to update to
	vmPassword := "vmPassword_example" // string | VM password (optional)

	configuration := fusionClient.NewConfiguration()
	apiClient := fusionClient.NewAPIClient(configuration)
	resp, r, err := apiClient.VMSharedFoldersManagementAPI.UpdataSharedFolder(context.Background(), id, folderId).Parameters(parameters).VmPassword(vmPassword).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMSharedFoldersManagementAPI.UpdataSharedFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdataSharedFolder`: []SharedFolder
	fmt.Fprintf(os.Stdout, "Response from `VMSharedFoldersManagementAPI.UpdataSharedFolder`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of VM |
**folderId** | **string** | ID of VM shared folder |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdataSharedFolderRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **parameters** | [**SharedFolderParameter**](SharedFolderParameter.md) | Parameters of the shared folder to update to |
 **vmPassword** | **string** | VM password |

### Return type

[**[]SharedFolder**](SharedFolder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.vmware.vmw.rest-v1+json
- **Accept**: application/vnd.vmware.vmw.rest-v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
