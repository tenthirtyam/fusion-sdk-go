# \HostNetworksManagementAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetwork**](HostNetworksManagementAPI.md#CreateNetwork) | **Post** /vmnets | Creates a virtual network
[**DeletePortforward**](HostNetworksManagementAPI.md#DeletePortforward) | **Delete** /vmnet/{vmnet}/portforward/{protocol}/{port} | Deletes port forwarding
[**GetAllNetworks**](HostNetworksManagementAPI.md#GetAllNetworks) | **Get** /vmnet | Returns all virtual networks
[**GetMACToIPs**](HostNetworksManagementAPI.md#GetMACToIPs) | **Get** /vmnet/{vmnet}/mactoip | Returns all MAC-to-IP settings for DHCP service
[**GetPortforwards**](HostNetworksManagementAPI.md#GetPortforwards) | **Get** /vmnet/{vmnet}/portforward | Returns all port forwardings
[**UpdateMacToIP**](HostNetworksManagementAPI.md#UpdateMacToIP) | **Put** /vmnet/{vmnet}/mactoip/{mac} | Updates the MAC-to-IP binding
[**UpdatePortforward**](HostNetworksManagementAPI.md#UpdatePortforward) | **Put** /vmnet/{vmnet}/portforward/{protocol}/{port} | Updates port forwarding

## CreateNetwork

> Network CreateNetwork(ctx).Parameters(parameters).Execute()

Creates a virtual network

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
 parameters := *fusionClient.NewCreateVmnetParameter("Name_example") // CreateVmnetParameter | Host network to be created

 configuration := fusionClient.NewConfiguration()
 apiClient := fusionClient.NewAPIClient(configuration)
 resp, r, err := apiClient.HostNetworksManagementAPI.CreateNetwork(context.Background()).Parameters(parameters).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `HostNetworksManagementAPI.CreateNetwork``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `CreateNetwork`: Network
 fmt.Fprintf(os.Stdout, "Response from `HostNetworksManagementAPI.CreateNetwork`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **parameters** | [**CreateVmnetParameter**](CreateVmnetParameter.md) | Host network to be created |

### Return type

[**Network**](Network.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.vmware.vmw.rest-v1+json
- **Accept**: application/vnd.vmware.vmw.rest-v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## DeletePortforward

> DeletePortforward(ctx, vmnet, protocol, port).Execute()

Deletes port forwarding

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
 vmnet := "vmnet_example" // string | NAT type of virtual network
 protocol := "protocol_example" // string | Protocol type: tcp, udp
 port := int32(56) // int32 | Host port number

 configuration := fusionClient.NewConfiguration()
 apiClient := fusionClient.NewAPIClient(configuration)
 r, err := apiClient.HostNetworksManagementAPI.DeletePortforward(context.Background(), vmnet, protocol, port).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `HostNetworksManagementAPI.DeletePortforward``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmnet** | **string** | NAT type of virtual network |
**protocol** | **string** | Protocol type: tcp, udp |
**port** | **int32** | Host port number |

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePortforwardRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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

## GetAllNetworks

> Networks GetAllNetworks(ctx).Execute()

Returns all virtual networks

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

 configuration := fusionClient.NewConfiguration()
 apiClient := fusionClient.NewAPIClient(configuration)
 resp, r, err := apiClient.HostNetworksManagementAPI.GetAllNetworks(context.Background()).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `HostNetworksManagementAPI.GetAllNetworks``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetAllNetworks`: Networks
 fmt.Fprintf(os.Stdout, "Response from `HostNetworksManagementAPI.GetAllNetworks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllNetworksRequest struct via the builder pattern

### Return type

[**Networks**](Networks.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.vmware.vmw.rest-v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetMACToIPs

> MACToIPs GetMACToIPs(ctx, vmnet).Execute()

Returns all MAC-to-IP settings for DHCP service

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
 vmnet := "vmnet_example" // string | Virtual network that has DHCP enabled

 configuration := fusionClient.NewConfiguration()
 apiClient := fusionClient.NewAPIClient(configuration)
 resp, r, err := apiClient.HostNetworksManagementAPI.GetMACToIPs(context.Background(), vmnet).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `HostNetworksManagementAPI.GetMACToIPs``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetMACToIPs`: MACToIPs
 fmt.Fprintf(os.Stdout, "Response from `HostNetworksManagementAPI.GetMACToIPs`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmnet** | **string** | Virtual network that has DHCP enabled |

### Other Parameters

Other parameters are passed through a pointer to a apiGetMACToIPsRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**MACToIPs**](MACToIPs.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.vmware.vmw.rest-v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetPortforwards

> Portforwards GetPortforwards(ctx, vmnet).Execute()

Returns all port forwardings

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
 vmnet := "vmnet_example" // string | NAT type of virtual network

 configuration := fusionClient.NewConfiguration()
 apiClient := fusionClient.NewAPIClient(configuration)
 resp, r, err := apiClient.HostNetworksManagementAPI.GetPortforwards(context.Background(), vmnet).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `HostNetworksManagementAPI.GetPortforwards``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetPortforwards`: Portforwards
 fmt.Fprintf(os.Stdout, "Response from `HostNetworksManagementAPI.GetPortforwards`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmnet** | **string** | NAT type of virtual network |

### Other Parameters

Other parameters are passed through a pointer to a apiGetPortforwardsRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**Portforwards**](Portforwards.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.vmware.vmw.rest-v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateMacToIP

> ErrorModel UpdateMacToIP(ctx, vmnet, mac).Parameters(parameters).Execute()

Updates the MAC-to-IP binding

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
 vmnet := "vmnet_example" // string | Virtual network that enabled DHCP
 mac := "mac_example" // string | Mac address that want to be mapped with a given IP
 parameters := *fusionClient.NewMacToIPParameter("IP_example") // MacToIPParameter | IP that will be assigned to given Mac address. If empty IP, the original Mac to IP binding will be deleted

 configuration := fusionClient.NewConfiguration()
 apiClient := fusionClient.NewAPIClient(configuration)
 resp, r, err := apiClient.HostNetworksManagementAPI.UpdateMacToIP(context.Background(), vmnet, mac).Parameters(parameters).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `HostNetworksManagementAPI.UpdateMacToIP``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `UpdateMacToIP`: ErrorModel
 fmt.Fprintf(os.Stdout, "Response from `HostNetworksManagementAPI.UpdateMacToIP`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmnet** | **string** | Virtual network that enabled DHCP |
**mac** | **string** | Mac address that want to be mapped with a given IP |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMacToIPRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **parameters** | [**MacToIPParameter**](MacToIPParameter.md) | IP that will be assigned to given Mac address. If empty IP, the original Mac to IP binding will be deleted |

### Return type

[**ErrorModel**](ErrorModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.vmware.vmw.rest-v1+json
- **Accept**: application/vnd.vmware.vmw.rest-v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdatePortforward

> ErrorModel UpdatePortforward(ctx, vmnet, protocol, port).Parameters(parameters).Execute()

Updates port forwarding

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
 vmnet := "vmnet_example" // string | NAT type of virtual network
 protocol := "protocol_example" // string | Protocol type: tcp, udp
 port := int32(56) // int32 | Host port number
 parameters := *fusionClient.NewPortforwardParameter("GuestIp_example", int32(123)) // PortforwardParameter | Guest to forward to

 configuration := fusionClient.NewConfiguration()
 apiClient := fusionClient.NewAPIClient(configuration)
 resp, r, err := apiClient.HostNetworksManagementAPI.UpdatePortforward(context.Background(), vmnet, protocol, port).Parameters(parameters).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `HostNetworksManagementAPI.UpdatePortforward``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `UpdatePortforward`: ErrorModel
 fmt.Fprintf(os.Stdout, "Response from `HostNetworksManagementAPI.UpdatePortforward`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vmnet** | **string** | NAT type of virtual network |
**protocol** | **string** | Protocol type: tcp, udp |
**port** | **int32** | Host port number |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePortforwardRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **parameters** | [**PortforwardParameter**](PortforwardParameter.md) | Guest to forward to |

### Return type

[**ErrorModel**](ErrorModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.vmware.vmw.rest-v1+json
- **Accept**: application/vnd.vmware.vmw.rest-v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
