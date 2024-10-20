# \VMNetworkAdaptersManagementAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNICDevice**](VMNetworkAdaptersManagementAPI.md#CreateNICDevice) | **Post** /vms/{id}/nic | Creates a network adapter in the VM
[**DeleteNICDevice**](VMNetworkAdaptersManagementAPI.md#DeleteNICDevice) | **Delete** /vms/{id}/nic/{index} | Deletes a VM network adapter
[**GetAllNICDevices**](VMNetworkAdaptersManagementAPI.md#GetAllNICDevices) | **Get** /vms/{id}/nic | Returns all network adapters in the VM
[**GetIPAddress**](VMNetworkAdaptersManagementAPI.md#GetIPAddress) | **Get** /vms/{id}/ip | Returns the IP address of a VM
[**GetNicInfo**](VMNetworkAdaptersManagementAPI.md#GetNicInfo) | **Get** /vms/{id}/nicips | Returns the IP stack configuration of all NICs of a VM
[**UpdateNICDevice**](VMNetworkAdaptersManagementAPI.md#UpdateNICDevice) | **Put** /vms/{id}/nic/{index} | Updates a network adapter in the VM



## CreateNICDevice

> NICDevice CreateNICDevice(ctx, id).Parameters(parameters).VmPassword(vmPassword).Execute()

Creates a network adapter in the VM

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
	parameters := *fusionClient.NewNICDeviceParameter("Type_example", "Vmnet_example") // NICDeviceParameter | Parameters of network adapter to create
	vmPassword := "vmPassword_example" // string | VM password (optional)

	configuration := fusionClient.NewConfiguration()
	apiClient := fusionClient.NewAPIClient(configuration)
	resp, r, err := apiClient.VMNetworkAdaptersManagementAPI.CreateNICDevice(context.Background(), id).Parameters(parameters).VmPassword(vmPassword).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMNetworkAdaptersManagementAPI.CreateNICDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNICDevice`: NICDevice
	fmt.Fprintf(os.Stdout, "Response from `VMNetworkAdaptersManagementAPI.CreateNICDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of VM |

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNICDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **parameters** | [**NICDeviceParameter**](NICDeviceParameter.md) | Parameters of network adapter to create |
 **vmPassword** | **string** | VM password |

### Return type

[**NICDevice**](NICDevice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.vmware.vmw.rest-v1+json
- **Accept**: application/vnd.vmware.vmw.rest-v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNICDevice

> DeleteNICDevice(ctx, id, index).VmPassword(vmPassword).Execute()

Deletes a VM network adapter

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
	index := "index_example" // string | Index of VM network adapter
	vmPassword := "vmPassword_example" // string | VM password (optional)

	configuration := fusionClient.NewConfiguration()
	apiClient := fusionClient.NewAPIClient(configuration)
	r, err := apiClient.VMNetworkAdaptersManagementAPI.DeleteNICDevice(context.Background(), id, index).VmPassword(vmPassword).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMNetworkAdaptersManagementAPI.DeleteNICDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of VM |
**index** | **string** | Index of VM network adapter |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNICDeviceRequest struct via the builder pattern


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


## GetAllNICDevices

> NICDevices GetAllNICDevices(ctx, id).VmPassword(vmPassword).Execute()

Returns all network adapters in the VM

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
	resp, r, err := apiClient.VMNetworkAdaptersManagementAPI.GetAllNICDevices(context.Background(), id).VmPassword(vmPassword).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMNetworkAdaptersManagementAPI.GetAllNICDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllNICDevices`: NICDevices
	fmt.Fprintf(os.Stdout, "Response from `VMNetworkAdaptersManagementAPI.GetAllNICDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of VM |

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllNICDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vmPassword** | **string** | VM password |

### Return type

[**NICDevices**](NICDevices.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.vmware.vmw.rest-v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIPAddress

> GetIPAddress200Response GetIPAddress(ctx, id).VmPassword(vmPassword).Execute()

Returns the IP address of a VM

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
	resp, r, err := apiClient.VMNetworkAdaptersManagementAPI.GetIPAddress(context.Background(), id).VmPassword(vmPassword).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMNetworkAdaptersManagementAPI.GetIPAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIPAddress`: GetIPAddress200Response
	fmt.Fprintf(os.Stdout, "Response from `VMNetworkAdaptersManagementAPI.GetIPAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of VM |

### Other Parameters

Other parameters are passed through a pointer to a apiGetIPAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vmPassword** | **string** | VM password |

### Return type

[**GetIPAddress200Response**](GetIPAddress200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.vmware.vmw.rest-v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNicInfo

> NicIpStackAll GetNicInfo(ctx, id).VmPassword(vmPassword).Execute()

Returns the IP stack configuration of all NICs of a VM

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
	resp, r, err := apiClient.VMNetworkAdaptersManagementAPI.GetNicInfo(context.Background(), id).VmPassword(vmPassword).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMNetworkAdaptersManagementAPI.GetNicInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNicInfo`: NicIpStackAll
	fmt.Fprintf(os.Stdout, "Response from `VMNetworkAdaptersManagementAPI.GetNicInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of VM |

### Other Parameters

Other parameters are passed through a pointer to a apiGetNicInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vmPassword** | **string** | VM password |

### Return type

[**NicIpStackAll**](NicIpStackAll.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.vmware.vmw.rest-v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNICDevice

> NICDevice UpdateNICDevice(ctx, id, index).Parameters(parameters).VmPassword(vmPassword).Execute()

Updates a network adapter in the VM

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
	index := "index_example" // string | Index of VM network adapter
	parameters := *fusionClient.NewNICDeviceParameter("Type_example", "Vmnet_example") // NICDeviceParameter | Parameters of network adapter to update to
	vmPassword := "vmPassword_example" // string | VM password (optional)

	configuration := fusionClient.NewConfiguration()
	apiClient := fusionClient.NewAPIClient(configuration)
	resp, r, err := apiClient.VMNetworkAdaptersManagementAPI.UpdateNICDevice(context.Background(), id, index).Parameters(parameters).VmPassword(vmPassword).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VMNetworkAdaptersManagementAPI.UpdateNICDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNICDevice`: NICDevice
	fmt.Fprintf(os.Stdout, "Response from `VMNetworkAdaptersManagementAPI.UpdateNICDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of VM |
**index** | **string** | Index of VM network adapter |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNICDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **parameters** | [**NICDeviceParameter**](NICDeviceParameter.md) | Parameters of network adapter to update to |
 **vmPassword** | **string** | VM password |

### Return type

[**NICDevice**](NICDevice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.vmware.vmw.rest-v1+json
- **Accept**: application/vnd.vmware.vmw.rest-v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

