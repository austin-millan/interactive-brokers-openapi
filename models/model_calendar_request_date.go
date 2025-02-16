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

// CalendarRequestDate struct for CalendarRequestDate
type CalendarRequestDate struct {
	// start date of a period. for example 20180808-0400
	Start *string `json:"start,omitempty"`
	// end date of a period. for example 20180808-0400
	End *string `json:"end,omitempty"`
}

// NewCalendarRequestDate instantiates a new CalendarRequestDate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCalendarRequestDate() *CalendarRequestDate {
	this := CalendarRequestDate{}
	return &this
}

// NewCalendarRequestDateWithDefaults instantiates a new CalendarRequestDate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCalendarRequestDateWithDefaults() *CalendarRequestDate {
	this := CalendarRequestDate{}
	return &this
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *CalendarRequestDate) GetStart() string {
	if o == nil || o.Start == nil {
		var ret string
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendarRequestDate) GetStartOk() (*string, bool) {
	if o == nil || o.Start == nil {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *CalendarRequestDate) HasStart() bool {
	if o != nil && o.Start != nil {
		return true
	}

	return false
}

// SetStart gets a reference to the given string and assigns it to the Start field.
func (o *CalendarRequestDate) SetStart(v string) {
	o.Start = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *CalendarRequestDate) GetEnd() string {
	if o == nil || o.End == nil {
		var ret string
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendarRequestDate) GetEndOk() (*string, bool) {
	if o == nil || o.End == nil {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *CalendarRequestDate) HasEnd() bool {
	if o != nil && o.End != nil {
		return true
	}

	return false
}

// SetEnd gets a reference to the given string and assigns it to the End field.
func (o *CalendarRequestDate) SetEnd(v string) {
	o.End = &v
}

func (o CalendarRequestDate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Start != nil {
		toSerialize["start"] = o.Start
	}
	if o.End != nil {
		toSerialize["end"] = o.End
	}
	return json.Marshal(toSerialize)
}

type NullableCalendarRequestDate struct {
	value *CalendarRequestDate
	isSet bool
}

func (v NullableCalendarRequestDate) Get() *CalendarRequestDate {
	return v.value
}

func (v *NullableCalendarRequestDate) Set(val *CalendarRequestDate) {
	v.value = val
	v.isSet = true
}

func (v NullableCalendarRequestDate) IsSet() bool {
	return v.isSet
}

func (v *NullableCalendarRequestDate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCalendarRequestDate(val *CalendarRequestDate) *NullableCalendarRequestDate {
	return &NullableCalendarRequestDate{value: val, isSet: true}
}

func (v NullableCalendarRequestDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCalendarRequestDate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


