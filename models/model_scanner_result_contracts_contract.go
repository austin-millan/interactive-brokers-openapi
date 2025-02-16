/*
 * Client Portal Web API
 *
 * Production version of the Client Portal Web API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ScannerResultContractsContract struct for ScannerResultContractsContract
type ScannerResultContractsContract struct {
	InScanTime *string `json:"inScanTime,omitempty"`
	Distance *int32 `json:"distance,omitempty"`
	ContractID *int32 `json:"contractID,omitempty"`
}

// NewScannerResultContractsContract instantiates a new ScannerResultContractsContract object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScannerResultContractsContract() *ScannerResultContractsContract {
	this := ScannerResultContractsContract{}
	return &this
}

// NewScannerResultContractsContractWithDefaults instantiates a new ScannerResultContractsContract object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScannerResultContractsContractWithDefaults() *ScannerResultContractsContract {
	this := ScannerResultContractsContract{}
	return &this
}

// GetInScanTime returns the InScanTime field value if set, zero value otherwise.
func (o *ScannerResultContractsContract) GetInScanTime() string {
	if o == nil || o.InScanTime == nil {
		var ret string
		return ret
	}
	return *o.InScanTime
}

// GetInScanTimeOk returns a tuple with the InScanTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScannerResultContractsContract) GetInScanTimeOk() (*string, bool) {
	if o == nil || o.InScanTime == nil {
		return nil, false
	}
	return o.InScanTime, true
}

// HasInScanTime returns a boolean if a field has been set.
func (o *ScannerResultContractsContract) HasInScanTime() bool {
	if o != nil && o.InScanTime != nil {
		return true
	}

	return false
}

// SetInScanTime gets a reference to the given string and assigns it to the InScanTime field.
func (o *ScannerResultContractsContract) SetInScanTime(v string) {
	o.InScanTime = &v
}

// GetDistance returns the Distance field value if set, zero value otherwise.
func (o *ScannerResultContractsContract) GetDistance() int32 {
	if o == nil || o.Distance == nil {
		var ret int32
		return ret
	}
	return *o.Distance
}

// GetDistanceOk returns a tuple with the Distance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScannerResultContractsContract) GetDistanceOk() (*int32, bool) {
	if o == nil || o.Distance == nil {
		return nil, false
	}
	return o.Distance, true
}

// HasDistance returns a boolean if a field has been set.
func (o *ScannerResultContractsContract) HasDistance() bool {
	if o != nil && o.Distance != nil {
		return true
	}

	return false
}

// SetDistance gets a reference to the given int32 and assigns it to the Distance field.
func (o *ScannerResultContractsContract) SetDistance(v int32) {
	o.Distance = &v
}

// GetContractID returns the ContractID field value if set, zero value otherwise.
func (o *ScannerResultContractsContract) GetContractID() int32 {
	if o == nil || o.ContractID == nil {
		var ret int32
		return ret
	}
	return *o.ContractID
}

// GetContractIDOk returns a tuple with the ContractID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScannerResultContractsContract) GetContractIDOk() (*int32, bool) {
	if o == nil || o.ContractID == nil {
		return nil, false
	}
	return o.ContractID, true
}

// HasContractID returns a boolean if a field has been set.
func (o *ScannerResultContractsContract) HasContractID() bool {
	if o != nil && o.ContractID != nil {
		return true
	}

	return false
}

// SetContractID gets a reference to the given int32 and assigns it to the ContractID field.
func (o *ScannerResultContractsContract) SetContractID(v int32) {
	o.ContractID = &v
}

func (o ScannerResultContractsContract) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.InScanTime != nil {
		toSerialize["inScanTime"] = o.InScanTime
	}
	if o.Distance != nil {
		toSerialize["distance"] = o.Distance
	}
	if o.ContractID != nil {
		toSerialize["contractID"] = o.ContractID
	}
	return json.Marshal(toSerialize)
}

type NullableScannerResultContractsContract struct {
	value *ScannerResultContractsContract
	isSet bool
}

func (v NullableScannerResultContractsContract) Get() *ScannerResultContractsContract {
	return v.value
}

func (v *NullableScannerResultContractsContract) Set(val *ScannerResultContractsContract) {
	v.value = val
	v.isSet = true
}

func (v NullableScannerResultContractsContract) IsSet() bool {
	return v.isSet
}

func (v *NullableScannerResultContractsContract) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScannerResultContractsContract(val *ScannerResultContractsContract) *NullableScannerResultContractsContract {
	return &NullableScannerResultContractsContract{value: val, isSet: true}
}

func (v NullableScannerResultContractsContract) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScannerResultContractsContract) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


