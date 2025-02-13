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

// ScannerParamsFilter struct for ScannerParamsFilter
type ScannerParamsFilter struct {
	Code *string `json:"code,omitempty"`
	Value *float32 `json:"value,omitempty"`
}

// NewScannerParamsFilter instantiates a new ScannerParamsFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScannerParamsFilter() *ScannerParamsFilter {
	this := ScannerParamsFilter{}
	return &this
}

// NewScannerParamsFilterWithDefaults instantiates a new ScannerParamsFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScannerParamsFilterWithDefaults() *ScannerParamsFilter {
	this := ScannerParamsFilter{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ScannerParamsFilter) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScannerParamsFilter) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ScannerParamsFilter) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ScannerParamsFilter) SetCode(v string) {
	o.Code = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ScannerParamsFilter) GetValue() float32 {
	if o == nil || o.Value == nil {
		var ret float32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScannerParamsFilter) GetValueOk() (*float32, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ScannerParamsFilter) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given float32 and assigns it to the Value field.
func (o *ScannerParamsFilter) SetValue(v float32) {
	o.Value = &v
}

func (o ScannerParamsFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableScannerParamsFilter struct {
	value *ScannerParamsFilter
	isSet bool
}

func (v NullableScannerParamsFilter) Get() *ScannerParamsFilter {
	return v.value
}

func (v *NullableScannerParamsFilter) Set(val *ScannerParamsFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableScannerParamsFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableScannerParamsFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScannerParamsFilter(val *ScannerParamsFilter) *NullableScannerParamsFilter {
	return &NullableScannerParamsFilter{value: val, isSet: true}
}

func (v NullableScannerParamsFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScannerParamsFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


