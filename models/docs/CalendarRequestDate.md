# CalendarRequestDate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | Pointer to **string** | start date of a period. for example 20180808-0400 | [optional] 
**End** | Pointer to **string** | end date of a period. for example 20180808-0400 | [optional] 

## Methods

### NewCalendarRequestDate

`func NewCalendarRequestDate() *CalendarRequestDate`

NewCalendarRequestDate instantiates a new CalendarRequestDate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalendarRequestDateWithDefaults

`func NewCalendarRequestDateWithDefaults() *CalendarRequestDate`

NewCalendarRequestDateWithDefaults instantiates a new CalendarRequestDate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *CalendarRequestDate) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *CalendarRequestDate) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *CalendarRequestDate) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *CalendarRequestDate) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetEnd

`func (o *CalendarRequestDate) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *CalendarRequestDate) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *CalendarRequestDate) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *CalendarRequestDate) HasEnd() bool`

HasEnd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


