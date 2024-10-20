// © Ryan Johnson. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

package fusion

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the VMRestrictionsInformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VMRestrictionsInformation{}

// VMRestrictionsInformation struct for VMRestrictionsInformation
type VMRestrictionsInformation struct {
	Id string `json:"id"`
	// The organization manages the VM
	ManagedOrg          *string `json:"managedOrg,omitempty"`
	Integrityconstraint *string `json:"integrityconstraint,omitempty"`
	Cpu                 *VMCPU  `json:"cpu,omitempty"`
	// Memory size in mega bytes
	Memory        *int32                 `json:"memory,omitempty"`
	ApplianceView *VMApplianceView       `json:"applianceView,omitempty"`
	CddvdList     *VMConnectedDeviceList `json:"cddvdList,omitempty"`
	FloopyList    *VMConnectedDeviceList `json:"floopyList,omitempty"`
	// Number of items
	FirewareType     *int32                 `json:"firewareType,omitempty"`
	GuestIsolation   *VMGuestIsolation      `json:"guestIsolation,omitempty"`
	Niclist          *NICDevices            `json:"niclist,omitempty"`
	ParallelPortList *VMConnectedDeviceList `json:"parallelPortList,omitempty"`
	SerialPortList   *VMConnectedDeviceList `json:"serialPortList,omitempty"`
	UsbList          *VMUsbList             `json:"usbList,omitempty"`
	RemoteVNC        *VMRemoteVNC           `json:"remoteVNC,omitempty"`
}

type _VMRestrictionsInformation VMRestrictionsInformation

// NewVMRestrictionsInformation instantiates a new VMRestrictionsInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVMRestrictionsInformation(id string) *VMRestrictionsInformation {
	this := VMRestrictionsInformation{}
	this.Id = id
	var memory int32 = 512
	this.Memory = &memory
	var firewareType int32 = 0
	this.FirewareType = &firewareType
	return &this
}

// NewVMRestrictionsInformationWithDefaults instantiates a new VMRestrictionsInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVMRestrictionsInformationWithDefaults() *VMRestrictionsInformation {
	this := VMRestrictionsInformation{}
	var memory int32 = 512
	this.Memory = &memory
	var firewareType int32 = 0
	this.FirewareType = &firewareType
	return &this
}

// GetId returns the Id field value
func (o *VMRestrictionsInformation) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *VMRestrictionsInformation) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *VMRestrictionsInformation) SetId(v string) {
	o.Id = v
}

// GetManagedOrg returns the ManagedOrg field value if set, zero value otherwise.
func (o *VMRestrictionsInformation) GetManagedOrg() string {
	if o == nil || IsNil(o.ManagedOrg) {
		var ret string
		return ret
	}
	return *o.ManagedOrg
}

// GetManagedOrgOk returns a tuple with the ManagedOrg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMRestrictionsInformation) GetManagedOrgOk() (*string, bool) {
	if o == nil || IsNil(o.ManagedOrg) {
		return nil, false
	}
	return o.ManagedOrg, true
}

// HasManagedOrg returns a boolean if a field has been set.
func (o *VMRestrictionsInformation) HasManagedOrg() bool {
	if o != nil && !IsNil(o.ManagedOrg) {
		return true
	}

	return false
}

// SetManagedOrg gets a reference to the given string and assigns it to the ManagedOrg field.
func (o *VMRestrictionsInformation) SetManagedOrg(v string) {
	o.ManagedOrg = &v
}

// GetIntegrityconstraint returns the Integrityconstraint field value if set, zero value otherwise.
func (o *VMRestrictionsInformation) GetIntegrityconstraint() string {
	if o == nil || IsNil(o.Integrityconstraint) {
		var ret string
		return ret
	}
	return *o.Integrityconstraint
}

// GetIntegrityconstraintOk returns a tuple with the Integrityconstraint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMRestrictionsInformation) GetIntegrityconstraintOk() (*string, bool) {
	if o == nil || IsNil(o.Integrityconstraint) {
		return nil, false
	}
	return o.Integrityconstraint, true
}

// HasIntegrityconstraint returns a boolean if a field has been set.
func (o *VMRestrictionsInformation) HasIntegrityconstraint() bool {
	if o != nil && !IsNil(o.Integrityconstraint) {
		return true
	}

	return false
}

// SetIntegrityconstraint gets a reference to the given string and assigns it to the Integrityconstraint field.
func (o *VMRestrictionsInformation) SetIntegrityconstraint(v string) {
	o.Integrityconstraint = &v
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *VMRestrictionsInformation) GetCpu() VMCPU {
	if o == nil || IsNil(o.Cpu) {
		var ret VMCPU
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMRestrictionsInformation) GetCpuOk() (*VMCPU, bool) {
	if o == nil || IsNil(o.Cpu) {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *VMRestrictionsInformation) HasCpu() bool {
	if o != nil && !IsNil(o.Cpu) {
		return true
	}

	return false
}

// SetCpu gets a reference to the given VMCPU and assigns it to the Cpu field.
func (o *VMRestrictionsInformation) SetCpu(v VMCPU) {
	o.Cpu = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *VMRestrictionsInformation) GetMemory() int32 {
	if o == nil || IsNil(o.Memory) {
		var ret int32
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMRestrictionsInformation) GetMemoryOk() (*int32, bool) {
	if o == nil || IsNil(o.Memory) {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *VMRestrictionsInformation) HasMemory() bool {
	if o != nil && !IsNil(o.Memory) {
		return true
	}

	return false
}

// SetMemory gets a reference to the given int32 and assigns it to the Memory field.
func (o *VMRestrictionsInformation) SetMemory(v int32) {
	o.Memory = &v
}

// GetApplianceView returns the ApplianceView field value if set, zero value otherwise.
func (o *VMRestrictionsInformation) GetApplianceView() VMApplianceView {
	if o == nil || IsNil(o.ApplianceView) {
		var ret VMApplianceView
		return ret
	}
	return *o.ApplianceView
}

// GetApplianceViewOk returns a tuple with the ApplianceView field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMRestrictionsInformation) GetApplianceViewOk() (*VMApplianceView, bool) {
	if o == nil || IsNil(o.ApplianceView) {
		return nil, false
	}
	return o.ApplianceView, true
}

// HasApplianceView returns a boolean if a field has been set.
func (o *VMRestrictionsInformation) HasApplianceView() bool {
	if o != nil && !IsNil(o.ApplianceView) {
		return true
	}

	return false
}

// SetApplianceView gets a reference to the given VMApplianceView and assigns it to the ApplianceView field.
func (o *VMRestrictionsInformation) SetApplianceView(v VMApplianceView) {
	o.ApplianceView = &v
}

// GetCddvdList returns the CddvdList field value if set, zero value otherwise.
func (o *VMRestrictionsInformation) GetCddvdList() VMConnectedDeviceList {
	if o == nil || IsNil(o.CddvdList) {
		var ret VMConnectedDeviceList
		return ret
	}
	return *o.CddvdList
}

// GetCddvdListOk returns a tuple with the CddvdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMRestrictionsInformation) GetCddvdListOk() (*VMConnectedDeviceList, bool) {
	if o == nil || IsNil(o.CddvdList) {
		return nil, false
	}
	return o.CddvdList, true
}

// HasCddvdList returns a boolean if a field has been set.
func (o *VMRestrictionsInformation) HasCddvdList() bool {
	if o != nil && !IsNil(o.CddvdList) {
		return true
	}

	return false
}

// SetCddvdList gets a reference to the given VMConnectedDeviceList and assigns it to the CddvdList field.
func (o *VMRestrictionsInformation) SetCddvdList(v VMConnectedDeviceList) {
	o.CddvdList = &v
}

// GetFloopyList returns the FloopyList field value if set, zero value otherwise.
func (o *VMRestrictionsInformation) GetFloopyList() VMConnectedDeviceList {
	if o == nil || IsNil(o.FloopyList) {
		var ret VMConnectedDeviceList
		return ret
	}
	return *o.FloopyList
}

// GetFloopyListOk returns a tuple with the FloopyList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMRestrictionsInformation) GetFloopyListOk() (*VMConnectedDeviceList, bool) {
	if o == nil || IsNil(o.FloopyList) {
		return nil, false
	}
	return o.FloopyList, true
}

// HasFloopyList returns a boolean if a field has been set.
func (o *VMRestrictionsInformation) HasFloopyList() bool {
	if o != nil && !IsNil(o.FloopyList) {
		return true
	}

	return false
}

// SetFloopyList gets a reference to the given VMConnectedDeviceList and assigns it to the FloopyList field.
func (o *VMRestrictionsInformation) SetFloopyList(v VMConnectedDeviceList) {
	o.FloopyList = &v
}

// GetFirewareType returns the FirewareType field value if set, zero value otherwise.
func (o *VMRestrictionsInformation) GetFirewareType() int32 {
	if o == nil || IsNil(o.FirewareType) {
		var ret int32
		return ret
	}
	return *o.FirewareType
}

// GetFirewareTypeOk returns a tuple with the FirewareType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMRestrictionsInformation) GetFirewareTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.FirewareType) {
		return nil, false
	}
	return o.FirewareType, true
}

// HasFirewareType returns a boolean if a field has been set.
func (o *VMRestrictionsInformation) HasFirewareType() bool {
	if o != nil && !IsNil(o.FirewareType) {
		return true
	}

	return false
}

// SetFirewareType gets a reference to the given int32 and assigns it to the FirewareType field.
func (o *VMRestrictionsInformation) SetFirewareType(v int32) {
	o.FirewareType = &v
}

// GetGuestIsolation returns the GuestIsolation field value if set, zero value otherwise.
func (o *VMRestrictionsInformation) GetGuestIsolation() VMGuestIsolation {
	if o == nil || IsNil(o.GuestIsolation) {
		var ret VMGuestIsolation
		return ret
	}
	return *o.GuestIsolation
}

// GetGuestIsolationOk returns a tuple with the GuestIsolation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMRestrictionsInformation) GetGuestIsolationOk() (*VMGuestIsolation, bool) {
	if o == nil || IsNil(o.GuestIsolation) {
		return nil, false
	}
	return o.GuestIsolation, true
}

// HasGuestIsolation returns a boolean if a field has been set.
func (o *VMRestrictionsInformation) HasGuestIsolation() bool {
	if o != nil && !IsNil(o.GuestIsolation) {
		return true
	}

	return false
}

// SetGuestIsolation gets a reference to the given VMGuestIsolation and assigns it to the GuestIsolation field.
func (o *VMRestrictionsInformation) SetGuestIsolation(v VMGuestIsolation) {
	o.GuestIsolation = &v
}

// GetNiclist returns the Niclist field value if set, zero value otherwise.
func (o *VMRestrictionsInformation) GetNiclist() NICDevices {
	if o == nil || IsNil(o.Niclist) {
		var ret NICDevices
		return ret
	}
	return *o.Niclist
}

// GetNiclistOk returns a tuple with the Niclist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMRestrictionsInformation) GetNiclistOk() (*NICDevices, bool) {
	if o == nil || IsNil(o.Niclist) {
		return nil, false
	}
	return o.Niclist, true
}

// HasNiclist returns a boolean if a field has been set.
func (o *VMRestrictionsInformation) HasNiclist() bool {
	if o != nil && !IsNil(o.Niclist) {
		return true
	}

	return false
}

// SetNiclist gets a reference to the given NICDevices and assigns it to the Niclist field.
func (o *VMRestrictionsInformation) SetNiclist(v NICDevices) {
	o.Niclist = &v
}

// GetParallelPortList returns the ParallelPortList field value if set, zero value otherwise.
func (o *VMRestrictionsInformation) GetParallelPortList() VMConnectedDeviceList {
	if o == nil || IsNil(o.ParallelPortList) {
		var ret VMConnectedDeviceList
		return ret
	}
	return *o.ParallelPortList
}

// GetParallelPortListOk returns a tuple with the ParallelPortList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMRestrictionsInformation) GetParallelPortListOk() (*VMConnectedDeviceList, bool) {
	if o == nil || IsNil(o.ParallelPortList) {
		return nil, false
	}
	return o.ParallelPortList, true
}

// HasParallelPortList returns a boolean if a field has been set.
func (o *VMRestrictionsInformation) HasParallelPortList() bool {
	if o != nil && !IsNil(o.ParallelPortList) {
		return true
	}

	return false
}

// SetParallelPortList gets a reference to the given VMConnectedDeviceList and assigns it to the ParallelPortList field.
func (o *VMRestrictionsInformation) SetParallelPortList(v VMConnectedDeviceList) {
	o.ParallelPortList = &v
}

// GetSerialPortList returns the SerialPortList field value if set, zero value otherwise.
func (o *VMRestrictionsInformation) GetSerialPortList() VMConnectedDeviceList {
	if o == nil || IsNil(o.SerialPortList) {
		var ret VMConnectedDeviceList
		return ret
	}
	return *o.SerialPortList
}

// GetSerialPortListOk returns a tuple with the SerialPortList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMRestrictionsInformation) GetSerialPortListOk() (*VMConnectedDeviceList, bool) {
	if o == nil || IsNil(o.SerialPortList) {
		return nil, false
	}
	return o.SerialPortList, true
}

// HasSerialPortList returns a boolean if a field has been set.
func (o *VMRestrictionsInformation) HasSerialPortList() bool {
	if o != nil && !IsNil(o.SerialPortList) {
		return true
	}

	return false
}

// SetSerialPortList gets a reference to the given VMConnectedDeviceList and assigns it to the SerialPortList field.
func (o *VMRestrictionsInformation) SetSerialPortList(v VMConnectedDeviceList) {
	o.SerialPortList = &v
}

// GetUsbList returns the UsbList field value if set, zero value otherwise.
func (o *VMRestrictionsInformation) GetUsbList() VMUsbList {
	if o == nil || IsNil(o.UsbList) {
		var ret VMUsbList
		return ret
	}
	return *o.UsbList
}

// GetUsbListOk returns a tuple with the UsbList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMRestrictionsInformation) GetUsbListOk() (*VMUsbList, bool) {
	if o == nil || IsNil(o.UsbList) {
		return nil, false
	}
	return o.UsbList, true
}

// HasUsbList returns a boolean if a field has been set.
func (o *VMRestrictionsInformation) HasUsbList() bool {
	if o != nil && !IsNil(o.UsbList) {
		return true
	}

	return false
}

// SetUsbList gets a reference to the given VMUsbList and assigns it to the UsbList field.
func (o *VMRestrictionsInformation) SetUsbList(v VMUsbList) {
	o.UsbList = &v
}

// GetRemoteVNC returns the RemoteVNC field value if set, zero value otherwise.
func (o *VMRestrictionsInformation) GetRemoteVNC() VMRemoteVNC {
	if o == nil || IsNil(o.RemoteVNC) {
		var ret VMRemoteVNC
		return ret
	}
	return *o.RemoteVNC
}

// GetRemoteVNCOk returns a tuple with the RemoteVNC field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMRestrictionsInformation) GetRemoteVNCOk() (*VMRemoteVNC, bool) {
	if o == nil || IsNil(o.RemoteVNC) {
		return nil, false
	}
	return o.RemoteVNC, true
}

// HasRemoteVNC returns a boolean if a field has been set.
func (o *VMRestrictionsInformation) HasRemoteVNC() bool {
	if o != nil && !IsNil(o.RemoteVNC) {
		return true
	}

	return false
}

// SetRemoteVNC gets a reference to the given VMRemoteVNC and assigns it to the RemoteVNC field.
func (o *VMRestrictionsInformation) SetRemoteVNC(v VMRemoteVNC) {
	o.RemoteVNC = &v
}

func (o VMRestrictionsInformation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VMRestrictionsInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.ManagedOrg) {
		toSerialize["managedOrg"] = o.ManagedOrg
	}
	if !IsNil(o.Integrityconstraint) {
		toSerialize["integrityconstraint"] = o.Integrityconstraint
	}
	if !IsNil(o.Cpu) {
		toSerialize["cpu"] = o.Cpu
	}
	if !IsNil(o.Memory) {
		toSerialize["memory"] = o.Memory
	}
	if !IsNil(o.ApplianceView) {
		toSerialize["applianceView"] = o.ApplianceView
	}
	if !IsNil(o.CddvdList) {
		toSerialize["cddvdList"] = o.CddvdList
	}
	if !IsNil(o.FloopyList) {
		toSerialize["floopyList"] = o.FloopyList
	}
	if !IsNil(o.FirewareType) {
		toSerialize["firewareType"] = o.FirewareType
	}
	if !IsNil(o.GuestIsolation) {
		toSerialize["guestIsolation"] = o.GuestIsolation
	}
	if !IsNil(o.Niclist) {
		toSerialize["niclist"] = o.Niclist
	}
	if !IsNil(o.ParallelPortList) {
		toSerialize["parallelPortList"] = o.ParallelPortList
	}
	if !IsNil(o.SerialPortList) {
		toSerialize["serialPortList"] = o.SerialPortList
	}
	if !IsNil(o.UsbList) {
		toSerialize["usbList"] = o.UsbList
	}
	if !IsNil(o.RemoteVNC) {
		toSerialize["remoteVNC"] = o.RemoteVNC
	}
	return toSerialize, nil
}

func (o *VMRestrictionsInformation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varVMRestrictionsInformation := _VMRestrictionsInformation{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVMRestrictionsInformation)

	if err != nil {
		return err
	}

	*o = VMRestrictionsInformation(varVMRestrictionsInformation)

	return err
}

type NullableVMRestrictionsInformation struct {
	value *VMRestrictionsInformation
	isSet bool
}

func (v NullableVMRestrictionsInformation) Get() *VMRestrictionsInformation {
	return v.value
}

func (v *NullableVMRestrictionsInformation) Set(val *VMRestrictionsInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableVMRestrictionsInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableVMRestrictionsInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVMRestrictionsInformation(val *VMRestrictionsInformation) *NullableVMRestrictionsInformation {
	return &NullableVMRestrictionsInformation{value: val, isSet: true}
}

func (v NullableVMRestrictionsInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVMRestrictionsInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
