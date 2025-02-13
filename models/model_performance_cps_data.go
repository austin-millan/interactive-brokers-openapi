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

// PerformanceCpsData struct for PerformanceCpsData
type PerformanceCpsData struct {
	Id *string `json:"id,omitempty"`
	// for example-- acctid
	IdType *string `json:"idType,omitempty"`
	// start date-- yyyyMMdd
	Start *string `json:"start,omitempty"`
	BaseCurrency *string `json:"baseCurrency,omitempty"`
	// each value stands for price change percent of corresponding date in dates array
	Returns *[]float32 `json:"returns,omitempty"`
	// end date-- yyyyMMdd
	End *string `json:"end,omitempty"`
}

// NewPerformanceCpsData instantiates a new PerformanceCpsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPerformanceCpsData() *PerformanceCpsData {
	this := PerformanceCpsData{}
	return &this
}

// NewPerformanceCpsDataWithDefaults instantiates a new PerformanceCpsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPerformanceCpsDataWithDefaults() *PerformanceCpsData {
	this := PerformanceCpsData{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PerformanceCpsData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformanceCpsData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PerformanceCpsData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PerformanceCpsData) SetId(v string) {
	o.Id = &v
}

// GetIdType returns the IdType field value if set, zero value otherwise.
func (o *PerformanceCpsData) GetIdType() string {
	if o == nil || o.IdType == nil {
		var ret string
		return ret
	}
	return *o.IdType
}

// GetIdTypeOk returns a tuple with the IdType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformanceCpsData) GetIdTypeOk() (*string, bool) {
	if o == nil || o.IdType == nil {
		return nil, false
	}
	return o.IdType, true
}

// HasIdType returns a boolean if a field has been set.
func (o *PerformanceCpsData) HasIdType() bool {
	if o != nil && o.IdType != nil {
		return true
	}

	return false
}

// SetIdType gets a reference to the given string and assigns it to the IdType field.
func (o *PerformanceCpsData) SetIdType(v string) {
	o.IdType = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *PerformanceCpsData) GetStart() string {
	if o == nil || o.Start == nil {
		var ret string
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformanceCpsData) GetStartOk() (*string, bool) {
	if o == nil || o.Start == nil {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *PerformanceCpsData) HasStart() bool {
	if o != nil && o.Start != nil {
		return true
	}

	return false
}

// SetStart gets a reference to the given string and assigns it to the Start field.
func (o *PerformanceCpsData) SetStart(v string) {
	o.Start = &v
}

// GetBaseCurrency returns the BaseCurrency field value if set, zero value otherwise.
func (o *PerformanceCpsData) GetBaseCurrency() string {
	if o == nil || o.BaseCurrency == nil {
		var ret string
		return ret
	}
	return *o.BaseCurrency
}

// GetBaseCurrencyOk returns a tuple with the BaseCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformanceCpsData) GetBaseCurrencyOk() (*string, bool) {
	if o == nil || o.BaseCurrency == nil {
		return nil, false
	}
	return o.BaseCurrency, true
}

// HasBaseCurrency returns a boolean if a field has been set.
func (o *PerformanceCpsData) HasBaseCurrency() bool {
	if o != nil && o.BaseCurrency != nil {
		return true
	}

	return false
}

// SetBaseCurrency gets a reference to the given string and assigns it to the BaseCurrency field.
func (o *PerformanceCpsData) SetBaseCurrency(v string) {
	o.BaseCurrency = &v
}

// GetReturns returns the Returns field value if set, zero value otherwise.
func (o *PerformanceCpsData) GetReturns() []float32 {
	if o == nil || o.Returns == nil {
		var ret []float32
		return ret
	}
	return *o.Returns
}

// GetReturnsOk returns a tuple with the Returns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformanceCpsData) GetReturnsOk() (*[]float32, bool) {
	if o == nil || o.Returns == nil {
		return nil, false
	}
	return o.Returns, true
}

// HasReturns returns a boolean if a field has been set.
func (o *PerformanceCpsData) HasReturns() bool {
	if o != nil && o.Returns != nil {
		return true
	}

	return false
}

// SetReturns gets a reference to the given []float32 and assigns it to the Returns field.
func (o *PerformanceCpsData) SetReturns(v []float32) {
	o.Returns = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *PerformanceCpsData) GetEnd() string {
	if o == nil || o.End == nil {
		var ret string
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformanceCpsData) GetEndOk() (*string, bool) {
	if o == nil || o.End == nil {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *PerformanceCpsData) HasEnd() bool {
	if o != nil && o.End != nil {
		return true
	}

	return false
}

// SetEnd gets a reference to the given string and assigns it to the End field.
func (o *PerformanceCpsData) SetEnd(v string) {
	o.End = &v
}

func (o PerformanceCpsData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IdType != nil {
		toSerialize["idType"] = o.IdType
	}
	if o.Start != nil {
		toSerialize["start"] = o.Start
	}
	if o.BaseCurrency != nil {
		toSerialize["baseCurrency"] = o.BaseCurrency
	}
	if o.Returns != nil {
		toSerialize["returns"] = o.Returns
	}
	if o.End != nil {
		toSerialize["end"] = o.End
	}
	return json.Marshal(toSerialize)
}

type NullablePerformanceCpsData struct {
	value *PerformanceCpsData
	isSet bool
}

func (v NullablePerformanceCpsData) Get() *PerformanceCpsData {
	return v.value
}

func (v *NullablePerformanceCpsData) Set(val *PerformanceCpsData) {
	v.value = val
	v.isSet = true
}

func (v NullablePerformanceCpsData) IsSet() bool {
	return v.isSet
}

func (v *NullablePerformanceCpsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePerformanceCpsData(val *PerformanceCpsData) *NullablePerformanceCpsData {
	return &NullablePerformanceCpsData{value: val, isSet: true}
}

func (v NullablePerformanceCpsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePerformanceCpsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


