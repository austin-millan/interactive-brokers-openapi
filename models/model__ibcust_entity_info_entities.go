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

// IbcustEntityInfoEntities struct for IbcustEntityInfoEntities
type IbcustEntityInfoEntities struct {
	CanTrade *bool `json:"canTrade,omitempty"`
	CanSign *bool `json:"canSign,omitempty"`
	Type *string `json:"type,omitempty"`
	Name *IbcustEntityInfoName `json:"name,omitempty"`
	Address *IbcustEntityInfoAddress `json:"address,omitempty"`
	IdentDocs *[]map[string]interface{} `json:"identDocs,omitempty"`
}

// NewIbcustEntityInfoEntities instantiates a new IbcustEntityInfoEntities object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIbcustEntityInfoEntities() *IbcustEntityInfoEntities {
	this := IbcustEntityInfoEntities{}
	return &this
}

// NewIbcustEntityInfoEntitiesWithDefaults instantiates a new IbcustEntityInfoEntities object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIbcustEntityInfoEntitiesWithDefaults() *IbcustEntityInfoEntities {
	this := IbcustEntityInfoEntities{}
	return &this
}

// GetCanTrade returns the CanTrade field value if set, zero value otherwise.
func (o *IbcustEntityInfoEntities) GetCanTrade() bool {
	if o == nil || o.CanTrade == nil {
		var ret bool
		return ret
	}
	return *o.CanTrade
}

// GetCanTradeOk returns a tuple with the CanTrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IbcustEntityInfoEntities) GetCanTradeOk() (*bool, bool) {
	if o == nil || o.CanTrade == nil {
		return nil, false
	}
	return o.CanTrade, true
}

// HasCanTrade returns a boolean if a field has been set.
func (o *IbcustEntityInfoEntities) HasCanTrade() bool {
	if o != nil && o.CanTrade != nil {
		return true
	}

	return false
}

// SetCanTrade gets a reference to the given bool and assigns it to the CanTrade field.
func (o *IbcustEntityInfoEntities) SetCanTrade(v bool) {
	o.CanTrade = &v
}

// GetCanSign returns the CanSign field value if set, zero value otherwise.
func (o *IbcustEntityInfoEntities) GetCanSign() bool {
	if o == nil || o.CanSign == nil {
		var ret bool
		return ret
	}
	return *o.CanSign
}

// GetCanSignOk returns a tuple with the CanSign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IbcustEntityInfoEntities) GetCanSignOk() (*bool, bool) {
	if o == nil || o.CanSign == nil {
		return nil, false
	}
	return o.CanSign, true
}

// HasCanSign returns a boolean if a field has been set.
func (o *IbcustEntityInfoEntities) HasCanSign() bool {
	if o != nil && o.CanSign != nil {
		return true
	}

	return false
}

// SetCanSign gets a reference to the given bool and assigns it to the CanSign field.
func (o *IbcustEntityInfoEntities) SetCanSign(v bool) {
	o.CanSign = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IbcustEntityInfoEntities) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IbcustEntityInfoEntities) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IbcustEntityInfoEntities) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *IbcustEntityInfoEntities) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IbcustEntityInfoEntities) GetName() IbcustEntityInfoName {
	if o == nil || o.Name == nil {
		var ret IbcustEntityInfoName
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IbcustEntityInfoEntities) GetNameOk() (*IbcustEntityInfoName, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IbcustEntityInfoEntities) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given IbcustEntityInfoName and assigns it to the Name field.
func (o *IbcustEntityInfoEntities) SetName(v IbcustEntityInfoName) {
	o.Name = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *IbcustEntityInfoEntities) GetAddress() IbcustEntityInfoAddress {
	if o == nil || o.Address == nil {
		var ret IbcustEntityInfoAddress
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IbcustEntityInfoEntities) GetAddressOk() (*IbcustEntityInfoAddress, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *IbcustEntityInfoEntities) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given IbcustEntityInfoAddress and assigns it to the Address field.
func (o *IbcustEntityInfoEntities) SetAddress(v IbcustEntityInfoAddress) {
	o.Address = &v
}

// GetIdentDocs returns the IdentDocs field value if set, zero value otherwise.
func (o *IbcustEntityInfoEntities) GetIdentDocs() []map[string]interface{} {
	if o == nil || o.IdentDocs == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.IdentDocs
}

// GetIdentDocsOk returns a tuple with the IdentDocs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IbcustEntityInfoEntities) GetIdentDocsOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.IdentDocs == nil {
		return nil, false
	}
	return o.IdentDocs, true
}

// HasIdentDocs returns a boolean if a field has been set.
func (o *IbcustEntityInfoEntities) HasIdentDocs() bool {
	if o != nil && o.IdentDocs != nil {
		return true
	}

	return false
}

// SetIdentDocs gets a reference to the given []map[string]interface{} and assigns it to the IdentDocs field.
func (o *IbcustEntityInfoEntities) SetIdentDocs(v []map[string]interface{}) {
	o.IdentDocs = &v
}

func (o IbcustEntityInfoEntities) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CanTrade != nil {
		toSerialize["canTrade"] = o.CanTrade
	}
	if o.CanSign != nil {
		toSerialize["canSign"] = o.CanSign
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.IdentDocs != nil {
		toSerialize["identDocs"] = o.IdentDocs
	}
	return json.Marshal(toSerialize)
}

type NullableIbcustEntityInfoEntities struct {
	value *IbcustEntityInfoEntities
	isSet bool
}

func (v NullableIbcustEntityInfoEntities) Get() *IbcustEntityInfoEntities {
	return v.value
}

func (v *NullableIbcustEntityInfoEntities) Set(val *IbcustEntityInfoEntities) {
	v.value = val
	v.isSet = true
}

func (v NullableIbcustEntityInfoEntities) IsSet() bool {
	return v.isSet
}

func (v *NullableIbcustEntityInfoEntities) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIbcustEntityInfoEntities(val *IbcustEntityInfoEntities) *NullableIbcustEntityInfoEntities {
	return &NullableIbcustEntityInfoEntities{value: val, isSet: true}
}

func (v NullableIbcustEntityInfoEntities) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIbcustEntityInfoEntities) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


