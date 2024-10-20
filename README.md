# Go SDK for VMware Fusion REST API

## Overview

- API version: 1.2.1

## Documentation for API Endpoints

Class | Method | HTTP Request | Description
------------ | ------------- | ------------- | -------------
*`HostNetworksManagementAPI`* | [**`CreateNetwork`**](docs/HostNetworksManagementAPI.md#createnetwork) | **Post** `/vmnets` | Creates a virtual network
*`HostNetworksManagementAPI`* | [**`DeletePortforward`**](docs/HostNetworksManagementAPI.md#deleteportforward) | **Delete** `/vmnet/{vmnet}/portforward/{protocol}/{port}` | Deletes port forwarding
*`HostNetworksManagementAPI`* | [**`GetAllNetworks`**](docs/HostNetworksManagementAPI.md#getallnetworks) | **Get** `/vmnet` | Returns all virtual networks
*`HostNetworksManagementAPI`* | [**`GetMACToIPs`**](docs/HostNetworksManagementAPI.md#getmactoips) | **Get** `/vmnet/{vmnet}/mactoip` | Returns all MAC-to-IP settings for DHCP service
*`HostNetworksManagementAPI`* | [**`GetPortforwards`**](docs/HostNetworksManagementAPI.md#getportforwards) | **Get** `/vmnet/{vmnet}/portforward` | Returns all port forwardings
*`HostNetworksManagementAPI`* | [**`UpdateMacToIP`**](docs/HostNetworksManagementAPI.md#updatemactoip) | **Put** `/vmnet/{vmnet}/mactoip/{mac}` | Updates the MAC-to-IP binding
*`HostNetworksManagementAPI`* | [**`UpdatePortforward`**](docs/HostNetworksManagementAPI.md#updateportforward) | **Put** `/vmnet/{vmnet}/portforward/{protocol}/{port}` | Updates port forwarding
*`VMManagementAPI`* | [**`ConfigVMParams`**](docs/VMManagementAPI.md#configvmparams) | **Put** `/vms/{id}/params` | update the vm config params
*`VMManagementAPI`* | [**`CreateVM`**](docs/VMManagementAPI.md#createvm) | **Post** `/vms` | Creates a copy of the VM
*`VMManagementAPI`* | [**`DeleteVM`**](docs/VMManagementAPI.md#deletevm) | **Delete** `/vms/{id}` | Deletes a VM
*`VMManagementAPI`* | [**`GetAllVMs`**](docs/VMManagementAPI.md#getallvms) | **Get** `/vms` | Returns a list of VM IDs and paths for all VMs
*`VMManagementAPI`* | [**`GetVM`**](docs/VMManagementAPI.md#getvm) | **Get** `/vms/{id}` | Returns the VM setting information of a VM
*`VMManagementAPI`* | [**`GetVMParams`**](docs/VMManagementAPI.md#getvmparams) | **Get** `/vms/{id}/params/{name}` | Get the VM config params
*`VMManagementAPI`* | [**`GetVMRestrictions`**](docs/VMManagementAPI.md#getvmrestrictions) | **Get** `/vms/{id}/restrictions` | Returns the restrictions information of the VM
*`VMManagementAPI`* | [**`RegisterVM`**](docs/VMManagementAPI.md#registervm) | **Post** `/vms/registration` | Register VM to VM Library
*`VMManagementAPI`* | [**`UpdateVM`**](docs/VMManagementAPI.md#updatevm) | **Put** `/vms/{id}` | Updates the VM settings
*`VMNetworkAdaptersManagementAPI`* | [**`CreateNICDevice`**](docs/VMNetworkAdaptersManagementAPI.md#createnicdevice) | **Post** `/vms/{id}/nic` | Creates a network adapter in the VM
*`VMNetworkAdaptersManagementAPI`* | [**`DeleteNICDevice`**](docs/VMNetworkAdaptersManagementAPI.md#deletenicdevice) | **Delete** `/vms/{id}/nic/{index}` | Deletes a VM network adapter
*`VMNetworkAdaptersManagementAPI`* | [**`GetAllNICDevices`**](docs/VMNetworkAdaptersManagementAPI.md#getallnicdevices) | **Get** `/vms/{id}/nic` | Returns all network adapters in the VM
*`VMNetworkAdaptersManagementAPI`* | [**`GetIPAddress`**](docs/VMNetworkAdaptersManagementAPI.md#getipaddress) | **Get** `/vms/{id}/ip` | Returns the IP address of a VM
*`VMNetworkAdaptersManagementAPI`* | [**`GetNicInfo`**](docs/VMNetworkAdaptersManagementAPI.md#getnicinfo) | **Get** `/vms/{id}/nicips` | Returns the IP stack configuration of all NICs of a VM
*`VMNetworkAdaptersManagementAPI`* | [**`UpdateNICDevice`**](docs/VMNetworkAdaptersManagementAPI.md#updatenicdevice) | **Put** `/vms/{id}/nic/{index}` | Updates a network adapter in the VM
*`VMPowerManagementAPI`* | [**`ChangePowerState`**](docs/VMPowerManagementAPI.md#changepowerstate) | **Put** `/vms/{id}/power` | Changes the VM power state
*`VMPowerManagementAPI`* | [**`GetPowerState`**](docs/VMPowerManagementAPI.md#getpowerstate) | **Get** `/vms/{id}/power` | Returns the power state of the VM
*`VMSharedFoldersManagementAPI`* | [**`CreateSharedFolder`**](docs/VMSharedFoldersManagementAPI.md#createsharedfolder) | **Post** `/vms/{id}/sharedfolders` | Mounts a new shared folder in the VM
*`VMSharedFoldersManagementAPI`* | [**`DeleteSharedFolder`**](docs/VMSharedFoldersManagementAPI.md#deletesharedfolder) | **Delete** `/vms/{id}/sharedfolders/{folder id}` | Deletes a shared folder
*`VMSharedFoldersManagementAPI`* | [**`GetAllSharedFolders`**](docs/VMSharedFoldersManagementAPI.md#getallsharedfolders) | **Get** `/vms/{id}/sharedfolders` | Returns all shared folders mounted in the VM
*`VMSharedFoldersManagementAPI`* | [**`UpdataSharedFolder`**](docs/VMSharedFoldersManagementAPI.md#updatasharedfolder) | **Put** `/vms/{id}/sharedfolders/{folder id}` | Updates a shared folder mounted in the VM

## Documentation For Models

- [`ConfigVMParamsParameter`](docs/ConfigVMParamsParameter.md)
- [`CreateVmnetParameter`](docs/CreateVmnetParameter.md)
- [`DaemonState`](docs/DaemonState.md)
- [`DhcpConfig`](docs/DhcpConfig.md)
- [`DnsConfig`](docs/DnsConfig.md)
- [`ErrorModel`](docs/ErrorModel.md)
- [`GetIPAddress200Response`](docs/GetIPAddress200Response.md)
- [`MACToIP`](docs/MACToIP.md)
- [`MACToIPs`](docs/MACToIPs.md)
- [`MacToIPParameter`](docs/MacToIPParameter.md)
- [`NICDevice`](docs/NICDevice.md)
- [`NICDeviceParameter`](docs/NICDeviceParameter.md)
- [`NICDevices`](docs/NICDevices.md)
- [`Network`](docs/Network.md)
- [`Networks`](docs/Networks.md)
- [`NicIpStack`](docs/NicIpStack.md)
- [`NicIpStackAll`](docs/NicIpStackAll.md)
- [`Portforward`](docs/Portforward.md)
- [`PortforwardGuest`](docs/PortforwardGuest.md)
- [`PortforwardParameter`](docs/PortforwardParameter.md)
- [`Portforwards`](docs/Portforwards.md)
- [`RouteEntry`](docs/RouteEntry.md)
- [`SharedFolder`](docs/SharedFolder.md)
- [`SharedFolderParameter`](docs/SharedFolderParameter.md)
- [`VMApplianceView`](docs/VMApplianceView.md)
- [`VMCPU`](docs/VMCPU.md)
- [`VMCloneParameter`](docs/VMCloneParameter.md)
- [`VMConnectedDevice`](docs/VMConnectedDevice.md)
- [`VMConnectedDeviceList`](docs/VMConnectedDeviceList.md)
- [`VMGuestIsolation`](docs/VMGuestIsolation.md)
- [`VMID`](docs/VMID.md)
- [`VMInformation`](docs/VMInformation.md)
- [`VMParameter`](docs/VMParameter.md)
- [`VMPowerOperation`](docs/VMPowerOperation.md)
- [`VMPowerState`](docs/VMPowerState.md)
- [`VMRegisterParameter`](docs/VMRegisterParameter.md)
- [`VMRemoteVNC`](docs/VMRemoteVNC.md)
- [`VMRestrictionsInformation`](docs/VMRestrictionsInformation.md)
- [`VMRrgistrationInformation`](docs/VMRrgistrationInformation.md)
- [`VMUsbDevice`](docs/VMUsbDevice.md)
- [`VMUsbList`](docs/VMUsbList.md)
- [`WinsConfig`](docs/WinsConfig.md)
