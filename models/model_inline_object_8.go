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

// InlineObject8 struct for InlineObject8
type InlineObject8 struct {
	Conids *[]int32 `json:"conids,omitempty"`
}

// NewInlineObject8 instantiates a new InlineObject8 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject8() *InlineObject8 {
	this := InlineObject8{}
	return &this
}

// NewInlineObject8WithDefaults instantiates a new InlineObject8 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject8WithDefaults() *InlineObject8 {
	this := InlineObject8{}
	return &this
}

// GetConids returns the Conids field value if set, zero value otherwise.
func (o *InlineObject8) GetConids() []int32 {
	if o == nil || o.Conids == nil {
		var ret []int32
		return ret
	}
	return *o.Conids
}

// GetConidsOk returns a tuple with the Conids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject8) GetConidsOk() (*[]int32, bool) {
	if o == nil || o.Conids == nil {
		return nil, false
	}
	return o.Conids, true
}

// HasConids returns a boolean if a field has been set.
func (o *InlineObject8) HasConids() bool {
	if o != nil && o.Conids != nil {
		return true
	}

	return false
}

// SetConids gets a reference to the given []int32 and assigns it to the Conids field.
func (o *InlineObject8) SetConids(v []int32) {
	o.Conids = &v
}

func (o InlineObject8) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Conids != nil {
		toSerialize["conids"] = o.Conids
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject8 struct {
	value *InlineObject8
	isSet bool
}

func (v NullableInlineObject8) Get() *InlineObject8 {
	return v.value
}

func (v *NullableInlineObject8) Set(val *InlineObject8) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject8) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject8) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject8(val *InlineObject8) *NullableInlineObject8 {
	return &NullableInlineObject8{value: val, isSet: true}
}

func (v NullableInlineObject8) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject8) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


