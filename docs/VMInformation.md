# VMInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  |
**Cpu** | Pointer to [**VMCPU**](VMCPU.md) |  | [optional]
**Memory** | Pointer to **int32** | Memory size in mega bytes | [optional] [default to 512]

## Methods

### NewVMInformation

`func NewVMInformation(id string, ) *VMInformation`

NewVMInformation instantiates a new VMInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMInformationWithDefaults

`func NewVMInformationWithDefaults() *VMInformation`

NewVMInformationWithDefaults instantiates a new VMInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VMInformation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VMInformation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VMInformation) SetId(v string)`

SetId sets Id field to given value.

### GetCpu

`func (o *VMInformation) GetCpu() VMCPU`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *VMInformation) GetCpuOk() (*VMCPU, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *VMInformation) SetCpu(v VMCPU)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *VMInformation) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetMemory

`func (o *VMInformation) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *VMInformation) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *VMInformation) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *VMInformation) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
