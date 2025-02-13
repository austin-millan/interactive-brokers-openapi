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

// AccountMaster struct for AccountMaster
type AccountMaster struct {
	Title *string `json:"title,omitempty"`
	OfficialTitle *string `json:"officialTitle,omitempty"`
}

// NewAccountMaster instantiates a new AccountMaster object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountMaster() *AccountMaster {
	this := AccountMaster{}
	return &this
}

// NewAccountMasterWithDefaults instantiates a new AccountMaster object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountMasterWithDefaults() *AccountMaster {
	this := AccountMaster{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AccountMaster) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountMaster) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AccountMaster) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AccountMaster) SetTitle(v string) {
	o.Title = &v
}

// GetOfficialTitle returns the OfficialTitle field value if set, zero value otherwise.
func (o *AccountMaster) GetOfficialTitle() string {
	if o == nil || o.OfficialTitle == nil {
		var ret string
		return ret
	}
	return *o.OfficialTitle
}

// GetOfficialTitleOk returns a tuple with the OfficialTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountMaster) GetOfficialTitleOk() (*string, bool) {
	if o == nil || o.OfficialTitle == nil {
		return nil, false
	}
	return o.OfficialTitle, true
}

// HasOfficialTitle returns a boolean if a field has been set.
func (o *AccountMaster) HasOfficialTitle() bool {
	if o != nil && o.OfficialTitle != nil {
		return true
	}

	return false
}

// SetOfficialTitle gets a reference to the given string and assigns it to the OfficialTitle field.
func (o *AccountMaster) SetOfficialTitle(v string) {
	o.OfficialTitle = &v
}

func (o AccountMaster) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.OfficialTitle != nil {
		toSerialize["officialTitle"] = o.OfficialTitle
	}
	return json.Marshal(toSerialize)
}

type NullableAccountMaster struct {
	value *AccountMaster
	isSet bool
}

func (v NullableAccountMaster) Get() *AccountMaster {
	return v.value
}

func (v *NullableAccountMaster) Set(val *AccountMaster) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountMaster) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountMaster) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountMaster(val *AccountMaster) *NullableAccountMaster {
	return &NullableAccountMaster{value: val, isSet: true}
}

func (v NullableAccountMaster) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountMaster) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


