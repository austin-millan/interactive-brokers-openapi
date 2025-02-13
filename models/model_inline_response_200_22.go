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

// InlineResponse20022 struct for InlineResponse20022
type InlineResponse20022 struct {
	AcctId *map[string]interface{} `json:"acctId,omitempty"`
}

// NewInlineResponse20022 instantiates a new InlineResponse20022 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20022() *InlineResponse20022 {
	this := InlineResponse20022{}
	return &this
}

// NewInlineResponse20022WithDefaults instantiates a new InlineResponse20022 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20022WithDefaults() *InlineResponse20022 {
	this := InlineResponse20022{}
	return &this
}

// GetAcctId returns the AcctId field value if set, zero value otherwise.
func (o *InlineResponse20022) GetAcctId() map[string]interface{} {
	if o == nil || o.AcctId == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.AcctId
}

// GetAcctIdOk returns a tuple with the AcctId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20022) GetAcctIdOk() (*map[string]interface{}, bool) {
	if o == nil || o.AcctId == nil {
		return nil, false
	}
	return o.AcctId, true
}

// HasAcctId returns a boolean if a field has been set.
func (o *InlineResponse20022) HasAcctId() bool {
	if o != nil && o.AcctId != nil {
		return true
	}

	return false
}

// SetAcctId gets a reference to the given map[string]interface{} and assigns it to the AcctId field.
func (o *InlineResponse20022) SetAcctId(v map[string]interface{}) {
	o.AcctId = &v
}

func (o InlineResponse20022) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AcctId != nil {
		toSerialize["acctId"] = o.AcctId
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20022 struct {
	value *InlineResponse20022
	isSet bool
}

func (v NullableInlineResponse20022) Get() *InlineResponse20022 {
	return v.value
}

func (v *NullableInlineResponse20022) Set(val *InlineResponse20022) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20022) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20022) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20022(val *InlineResponse20022) *NullableInlineResponse20022 {
	return &NullableInlineResponse20022{value: val, isSet: true}
}

func (v NullableInlineResponse20022) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20022) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


