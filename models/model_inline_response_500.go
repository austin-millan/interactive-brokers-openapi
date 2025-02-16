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

// InlineResponse500 struct for InlineResponse500
type InlineResponse500 struct {
	Error *string `json:"error,omitempty"`
}

// NewInlineResponse500 instantiates a new InlineResponse500 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse500() *InlineResponse500 {
	this := InlineResponse500{}
	return &this
}

// NewInlineResponse500WithDefaults instantiates a new InlineResponse500 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse500WithDefaults() *InlineResponse500 {
	this := InlineResponse500{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *InlineResponse500) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse500) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *InlineResponse500) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *InlineResponse500) SetError(v string) {
	o.Error = &v
}

func (o InlineResponse500) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse500 struct {
	value *InlineResponse500
	isSet bool
}

func (v NullableInlineResponse500) Get() *InlineResponse500 {
	return v.value
}

func (v *NullableInlineResponse500) Set(val *InlineResponse500) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse500) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse500) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse500(val *InlineResponse500) *NullableInlineResponse500 {
	return &NullableInlineResponse500{value: val, isSet: true}
}

func (v NullableInlineResponse500) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse500) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


