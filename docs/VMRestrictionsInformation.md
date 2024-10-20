# VMRestrictionsInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  |
**ManagedOrg** | Pointer to **string** | The organization manages the VM | [optional]
**Integrityconstraint** | Pointer to **string** |  | [optional]
**Cpu** | Pointer to [**VMCPU**](VMCPU.md) |  | [optional]
**Memory** | Pointer to **int32** | Memory size in mega bytes | [optional] [default to 512]
**ApplianceView** | Pointer to [**VMApplianceView**](VMApplianceView.md) |  | [optional]
**CddvdList** | Pointer to [**VMConnectedDeviceList**](VMConnectedDeviceList.md) |  | [optional]
**FloopyList** | Pointer to [**VMConnectedDeviceList**](VMConnectedDeviceList.md) |  | [optional]
**FirewareType** | Pointer to **int32** | Number of items | [optional] [default to 0]
**GuestIsolation** | Pointer to [**VMGuestIsolation**](VMGuestIsolation.md) |  | [optional]
**Niclist** | Pointer to [**NICDevices**](NICDevices.md) |  | [optional]
**ParallelPortList** | Pointer to [**VMConnectedDeviceList**](VMConnectedDeviceList.md) |  | [optional]
**SerialPortList** | Pointer to [**VMConnectedDeviceList**](VMConnectedDeviceList.md) |  | [optional]
**UsbList** | Pointer to [**VMUsbList**](VMUsbList.md) |  | [optional]
**RemoteVNC** | Pointer to [**VMRemoteVNC**](VMRemoteVNC.md) |  | [optional]

## Methods

### NewVMRestrictionsInformation

`func NewVMRestrictionsInformation(id string, ) *VMRestrictionsInformation`

NewVMRestrictionsInformation instantiates a new VMRestrictionsInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMRestrictionsInformationWithDefaults

`func NewVMRestrictionsInformationWithDefaults() *VMRestrictionsInformation`

NewVMRestrictionsInformationWithDefaults instantiates a new VMRestrictionsInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VMRestrictionsInformation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VMRestrictionsInformation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VMRestrictionsInformation) SetId(v string)`

SetId sets Id field to given value.

### GetManagedOrg

`func (o *VMRestrictionsInformation) GetManagedOrg() string`

GetManagedOrg returns the ManagedOrg field if non-nil, zero value otherwise.

### GetManagedOrgOk

`func (o *VMRestrictionsInformation) GetManagedOrgOk() (*string, bool)`

GetManagedOrgOk returns a tuple with the ManagedOrg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedOrg

`func (o *VMRestrictionsInformation) SetManagedOrg(v string)`

SetManagedOrg sets ManagedOrg field to given value.

### HasManagedOrg

`func (o *VMRestrictionsInformation) HasManagedOrg() bool`

HasManagedOrg returns a boolean if a field has been set.

### GetIntegrityconstraint

`func (o *VMRestrictionsInformation) GetIntegrityconstraint() string`

GetIntegrityconstraint returns the Integrityconstraint field if non-nil, zero value otherwise.

### GetIntegrityconstraintOk

`func (o *VMRestrictionsInformation) GetIntegrityconstraintOk() (*string, bool)`

GetIntegrityconstraintOk returns a tuple with the Integrityconstraint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrityconstraint

`func (o *VMRestrictionsInformation) SetIntegrityconstraint(v string)`

SetIntegrityconstraint sets Integrityconstraint field to given value.

### HasIntegrityconstraint

`func (o *VMRestrictionsInformation) HasIntegrityconstraint() bool`

HasIntegrityconstraint returns a boolean if a field has been set.

### GetCpu

`func (o *VMRestrictionsInformation) GetCpu() VMCPU`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *VMRestrictionsInformation) GetCpuOk() (*VMCPU, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *VMRestrictionsInformation) SetCpu(v VMCPU)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *VMRestrictionsInformation) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetMemory

`func (o *VMRestrictionsInformation) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *VMRestrictionsInformation) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *VMRestrictionsInformation) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *VMRestrictionsInformation) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetApplianceView

`func (o *VMRestrictionsInformation) GetApplianceView() VMApplianceView`

GetApplianceView returns the ApplianceView field if non-nil, zero value otherwise.

### GetApplianceViewOk

`func (o *VMRestrictionsInformation) GetApplianceViewOk() (*VMApplianceView, bool)`

GetApplianceViewOk returns a tuple with the ApplianceView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceView

`func (o *VMRestrictionsInformation) SetApplianceView(v VMApplianceView)`

SetApplianceView sets ApplianceView field to given value.

### HasApplianceView

`func (o *VMRestrictionsInformation) HasApplianceView() bool`

HasApplianceView returns a boolean if a field has been set.

### GetCddvdList

`func (o *VMRestrictionsInformation) GetCddvdList() VMConnectedDeviceList`

GetCddvdList returns the CddvdList field if non-nil, zero value otherwise.

### GetCddvdListOk

`func (o *VMRestrictionsInformation) GetCddvdListOk() (*VMConnectedDeviceList, bool)`

GetCddvdListOk returns a tuple with the CddvdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCddvdList

`func (o *VMRestrictionsInformation) SetCddvdList(v VMConnectedDeviceList)`

SetCddvdList sets CddvdList field to given value.

### HasCddvdList

`func (o *VMRestrictionsInformation) HasCddvdList() bool`

HasCddvdList returns a boolean if a field has been set.

### GetFloopyList

`func (o *VMRestrictionsInformation) GetFloopyList() VMConnectedDeviceList`

GetFloopyList returns the FloopyList field if non-nil, zero value otherwise.

### GetFloopyListOk

`func (o *VMRestrictionsInformation) GetFloopyListOk() (*VMConnectedDeviceList, bool)`

GetFloopyListOk returns a tuple with the FloopyList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloopyList

`func (o *VMRestrictionsInformation) SetFloopyList(v VMConnectedDeviceList)`

SetFloopyList sets FloopyList field to given value.

### HasFloopyList

`func (o *VMRestrictionsInformation) HasFloopyList() bool`

HasFloopyList returns a boolean if a field has been set.

### GetFirewareType

`func (o *VMRestrictionsInformation) GetFirewareType() int32`

GetFirewareType returns the FirewareType field if non-nil, zero value otherwise.

### GetFirewareTypeOk

`func (o *VMRestrictionsInformation) GetFirewareTypeOk() (*int32, bool)`

GetFirewareTypeOk returns a tuple with the FirewareType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewareType

`func (o *VMRestrictionsInformation) SetFirewareType(v int32)`

SetFirewareType sets FirewareType field to given value.

### HasFirewareType

`func (o *VMRestrictionsInformation) HasFirewareType() bool`

HasFirewareType returns a boolean if a field has been set.

### GetGuestIsolation

`func (o *VMRestrictionsInformation) GetGuestIsolation() VMGuestIsolation`

GetGuestIsolation returns the GuestIsolation field if non-nil, zero value otherwise.

### GetGuestIsolationOk

`func (o *VMRestrictionsInformation) GetGuestIsolationOk() (*VMGuestIsolation, bool)`

GetGuestIsolationOk returns a tuple with the GuestIsolation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestIsolation

`func (o *VMRestrictionsInformation) SetGuestIsolation(v VMGuestIsolation)`

SetGuestIsolation sets GuestIsolation field to given value.

### HasGuestIsolation

`func (o *VMRestrictionsInformation) HasGuestIsolation() bool`

HasGuestIsolation returns a boolean if a field has been set.

### GetNiclist

`func (o *VMRestrictionsInformation) GetNiclist() NICDevices`

GetNiclist returns the Niclist field if non-nil, zero value otherwise.

### GetNiclistOk

`func (o *VMRestrictionsInformation) GetNiclistOk() (*NICDevices, bool)`

GetNiclistOk returns a tuple with the Niclist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNiclist

`func (o *VMRestrictionsInformation) SetNiclist(v NICDevices)`

SetNiclist sets Niclist field to given value.

### HasNiclist

`func (o *VMRestrictionsInformation) HasNiclist() bool`

HasNiclist returns a boolean if a field has been set.

### GetParallelPortList

`func (o *VMRestrictionsInformation) GetParallelPortList() VMConnectedDeviceList`

GetParallelPortList returns the ParallelPortList field if non-nil, zero value otherwise.

### GetParallelPortListOk

`func (o *VMRestrictionsInformation) GetParallelPortListOk() (*VMConnectedDeviceList, bool)`

GetParallelPortListOk returns a tuple with the ParallelPortList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallelPortList

`func (o *VMRestrictionsInformation) SetParallelPortList(v VMConnectedDeviceList)`

SetParallelPortList sets ParallelPortList field to given value.

### HasParallelPortList

`func (o *VMRestrictionsInformation) HasParallelPortList() bool`

HasParallelPortList returns a boolean if a field has been set.

### GetSerialPortList

`func (o *VMRestrictionsInformation) GetSerialPortList() VMConnectedDeviceList`

GetSerialPortList returns the SerialPortList field if non-nil, zero value otherwise.

### GetSerialPortListOk

`func (o *VMRestrictionsInformation) GetSerialPortListOk() (*VMConnectedDeviceList, bool)`

GetSerialPortListOk returns a tuple with the SerialPortList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialPortList

`func (o *VMRestrictionsInformation) SetSerialPortList(v VMConnectedDeviceList)`

SetSerialPortList sets SerialPortList field to given value.

### HasSerialPortList

`func (o *VMRestrictionsInformation) HasSerialPortList() bool`

HasSerialPortList returns a boolean if a field has been set.

### GetUsbList

`func (o *VMRestrictionsInformation) GetUsbList() VMUsbList`

GetUsbList returns the UsbList field if non-nil, zero value otherwise.

### GetUsbListOk

`func (o *VMRestrictionsInformation) GetUsbListOk() (*VMUsbList, bool)`

GetUsbListOk returns a tuple with the UsbList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsbList

`func (o *VMRestrictionsInformation) SetUsbList(v VMUsbList)`

SetUsbList sets UsbList field to given value.

### HasUsbList

`func (o *VMRestrictionsInformation) HasUsbList() bool`

HasUsbList returns a boolean if a field has been set.

### GetRemoteVNC

`func (o *VMRestrictionsInformation) GetRemoteVNC() VMRemoteVNC`

GetRemoteVNC returns the RemoteVNC field if non-nil, zero value otherwise.

### GetRemoteVNCOk

`func (o *VMRestrictionsInformation) GetRemoteVNCOk() (*VMRemoteVNC, bool)`

GetRemoteVNCOk returns a tuple with the RemoteVNC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteVNC

`func (o *VMRestrictionsInformation) SetRemoteVNC(v VMRemoteVNC)`

SetRemoteVNC sets RemoteVNC field to given value.

### HasRemoteVNC

`func (o *VMRestrictionsInformation) HasRemoteVNC() bool`

HasRemoteVNC returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
