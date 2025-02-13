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

// InlineResponse2004E device
type InlineResponse2004E struct {
	// device name
	NM *string `json:"NM,omitempty"`
	// device id
	I *string `json:"I,omitempty"`
	// unique device id
	UI *string `json:"UI,omitempty"`
	// device is enabled or not 0-true, 1-false.
	A *string `json:"A,omitempty"`
}

// NewInlineResponse2004E instantiates a new InlineResponse2004E object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2004E() *InlineResponse2004E {
	this := InlineResponse2004E{}
	return &this
}

// NewInlineResponse2004EWithDefaults instantiates a new InlineResponse2004E object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2004EWithDefaults() *InlineResponse2004E {
	this := InlineResponse2004E{}
	return &this
}

// GetNM returns the NM field value if set, zero value otherwise.
func (o *InlineResponse2004E) GetNM() string {
	if o == nil || o.NM == nil {
		var ret string
		return ret
	}
	return *o.NM
}

// GetNMOk returns a tuple with the NM field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2004E) GetNMOk() (*string, bool) {
	if o == nil || o.NM == nil {
		return nil, false
	}
	return o.NM, true
}

// HasNM returns a boolean if a field has been set.
func (o *InlineResponse2004E) HasNM() bool {
	if o != nil && o.NM != nil {
		return true
	}

	return false
}

// SetNM gets a reference to the given string and assigns it to the NM field.
func (o *InlineResponse2004E) SetNM(v string) {
	o.NM = &v
}

// GetI returns the I field value if set, zero value otherwise.
func (o *InlineResponse2004E) GetI() string {
	if o == nil || o.I == nil {
		var ret string
		return ret
	}
	return *o.I
}

// GetIOk returns a tuple with the I field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2004E) GetIOk() (*string, bool) {
	if o == nil || o.I == nil {
		return nil, false
	}
	return o.I, true
}

// HasI returns a boolean if a field has been set.
func (o *InlineResponse2004E) HasI() bool {
	if o != nil && o.I != nil {
		return true
	}

	return false
}

// SetI gets a reference to the given string and assigns it to the I field.
func (o *InlineResponse2004E) SetI(v string) {
	o.I = &v
}

// GetUI returns the UI field value if set, zero value otherwise.
func (o *InlineResponse2004E) GetUI() string {
	if o == nil || o.UI == nil {
		var ret string
		return ret
	}
	return *o.UI
}

// GetUIOk returns a tuple with the UI field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2004E) GetUIOk() (*string, bool) {
	if o == nil || o.UI == nil {
		return nil, false
	}
	return o.UI, true
}

// HasUI returns a boolean if a field has been set.
func (o *InlineResponse2004E) HasUI() bool {
	if o != nil && o.UI != nil {
		return true
	}

	return false
}

// SetUI gets a reference to the given string and assigns it to the UI field.
func (o *InlineResponse2004E) SetUI(v string) {
	o.UI = &v
}

// GetA returns the A field value if set, zero value otherwise.
func (o *InlineResponse2004E) GetA() string {
	if o == nil || o.A == nil {
		var ret string
		return ret
	}
	return *o.A
}

// GetAOk returns a tuple with the A field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2004E) GetAOk() (*string, bool) {
	if o == nil || o.A == nil {
		return nil, false
	}
	return o.A, true
}

// HasA returns a boolean if a field has been set.
func (o *InlineResponse2004E) HasA() bool {
	if o != nil && o.A != nil {
		return true
	}

	return false
}

// SetA gets a reference to the given string and assigns it to the A field.
func (o *InlineResponse2004E) SetA(v string) {
	o.A = &v
}

func (o InlineResponse2004E) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NM != nil {
		toSerialize["NM"] = o.NM
	}
	if o.I != nil {
		toSerialize["I"] = o.I
	}
	if o.UI != nil {
		toSerialize["UI"] = o.UI
	}
	if o.A != nil {
		toSerialize["A"] = o.A
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2004E struct {
	value *InlineResponse2004E
	isSet bool
}

func (v NullableInlineResponse2004E) Get() *InlineResponse2004E {
	return v.value
}

func (v *NullableInlineResponse2004E) Set(val *InlineResponse2004E) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2004E) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2004E) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2004E(val *InlineResponse2004E) *NullableInlineResponse2004E {
	return &NullableInlineResponse2004E{value: val, isSet: true}
}

func (v NullableInlineResponse2004E) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2004E) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


