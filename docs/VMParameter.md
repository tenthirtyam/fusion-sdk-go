# VMParameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Processors** | Pointer to **int32** | Number of processor cores | [optional] [default to 1]
**Memory** | Pointer to **int32** | Memory size in mega bytes | [optional] [default to 512]

## Methods

### NewVMParameter

`func NewVMParameter() *VMParameter`

NewVMParameter instantiates a new VMParameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMParameterWithDefaults

`func NewVMParameterWithDefaults() *VMParameter`

NewVMParameterWithDefaults instantiates a new VMParameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProcessors

`func (o *VMParameter) GetProcessors() int32`

GetProcessors returns the Processors field if non-nil, zero value otherwise.

### GetProcessorsOk

`func (o *VMParameter) GetProcessorsOk() (*int32, bool)`

GetProcessorsOk returns a tuple with the Processors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessors

`func (o *VMParameter) SetProcessors(v int32)`

SetProcessors sets Processors field to given value.

### HasProcessors

`func (o *VMParameter) HasProcessors() bool`

HasProcessors returns a boolean if a field has been set.

### GetMemory

`func (o *VMParameter) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *VMParameter) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *VMParameter) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *VMParameter) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
