/*
Overseerr API

This is the documentation for the Overseerr API backend.  Two primary authentication methods are supported:  - **Cookie Authentication**: A valid sign-in to the `/auth/plex` or `/auth/local` will generate a valid authentication cookie. - **API Key Authentication**: Sign-in is also possible by passing an `X-Api-Key` header along with a valid API Key generated by Overseerr.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package overseerr_go

import (
	"encoding/json"
)

// checks if the MediaMediaIdWatchDataGet200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MediaMediaIdWatchDataGet200ResponseData{}

// MediaMediaIdWatchDataGet200ResponseData struct for MediaMediaIdWatchDataGet200ResponseData
type MediaMediaIdWatchDataGet200ResponseData struct {
	PlayCount7Days  *float32 `json:"playCount7Days,omitempty"`
	PlayCount30Days *float32 `json:"playCount30Days,omitempty"`
	PlayCount       *float32 `json:"playCount,omitempty"`
	Users           []User   `json:"users,omitempty"`
}

// NewMediaMediaIdWatchDataGet200ResponseData instantiates a new MediaMediaIdWatchDataGet200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMediaMediaIdWatchDataGet200ResponseData() *MediaMediaIdWatchDataGet200ResponseData {
	this := MediaMediaIdWatchDataGet200ResponseData{}
	return &this
}

// NewMediaMediaIdWatchDataGet200ResponseDataWithDefaults instantiates a new MediaMediaIdWatchDataGet200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMediaMediaIdWatchDataGet200ResponseDataWithDefaults() *MediaMediaIdWatchDataGet200ResponseData {
	this := MediaMediaIdWatchDataGet200ResponseData{}
	return &this
}

// GetPlayCount7Days returns the PlayCount7Days field value if set, zero value otherwise.
func (o *MediaMediaIdWatchDataGet200ResponseData) GetPlayCount7Days() float32 {
	if o == nil || IsNil(o.PlayCount7Days) {
		var ret float32
		return ret
	}
	return *o.PlayCount7Days
}

// GetPlayCount7DaysOk returns a tuple with the PlayCount7Days field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaMediaIdWatchDataGet200ResponseData) GetPlayCount7DaysOk() (*float32, bool) {
	if o == nil || IsNil(o.PlayCount7Days) {
		return nil, false
	}
	return o.PlayCount7Days, true
}

// HasPlayCount7Days returns a boolean if a field has been set.
func (o *MediaMediaIdWatchDataGet200ResponseData) HasPlayCount7Days() bool {
	if o != nil && !IsNil(o.PlayCount7Days) {
		return true
	}

	return false
}

// SetPlayCount7Days gets a reference to the given float32 and assigns it to the PlayCount7Days field.
func (o *MediaMediaIdWatchDataGet200ResponseData) SetPlayCount7Days(v float32) {
	o.PlayCount7Days = &v
}

// GetPlayCount30Days returns the PlayCount30Days field value if set, zero value otherwise.
func (o *MediaMediaIdWatchDataGet200ResponseData) GetPlayCount30Days() float32 {
	if o == nil || IsNil(o.PlayCount30Days) {
		var ret float32
		return ret
	}
	return *o.PlayCount30Days
}

// GetPlayCount30DaysOk returns a tuple with the PlayCount30Days field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaMediaIdWatchDataGet200ResponseData) GetPlayCount30DaysOk() (*float32, bool) {
	if o == nil || IsNil(o.PlayCount30Days) {
		return nil, false
	}
	return o.PlayCount30Days, true
}

// HasPlayCount30Days returns a boolean if a field has been set.
func (o *MediaMediaIdWatchDataGet200ResponseData) HasPlayCount30Days() bool {
	if o != nil && !IsNil(o.PlayCount30Days) {
		return true
	}

	return false
}

// SetPlayCount30Days gets a reference to the given float32 and assigns it to the PlayCount30Days field.
func (o *MediaMediaIdWatchDataGet200ResponseData) SetPlayCount30Days(v float32) {
	o.PlayCount30Days = &v
}

// GetPlayCount returns the PlayCount field value if set, zero value otherwise.
func (o *MediaMediaIdWatchDataGet200ResponseData) GetPlayCount() float32 {
	if o == nil || IsNil(o.PlayCount) {
		var ret float32
		return ret
	}
	return *o.PlayCount
}

// GetPlayCountOk returns a tuple with the PlayCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaMediaIdWatchDataGet200ResponseData) GetPlayCountOk() (*float32, bool) {
	if o == nil || IsNil(o.PlayCount) {
		return nil, false
	}
	return o.PlayCount, true
}

// HasPlayCount returns a boolean if a field has been set.
func (o *MediaMediaIdWatchDataGet200ResponseData) HasPlayCount() bool {
	if o != nil && !IsNil(o.PlayCount) {
		return true
	}

	return false
}

// SetPlayCount gets a reference to the given float32 and assigns it to the PlayCount field.
func (o *MediaMediaIdWatchDataGet200ResponseData) SetPlayCount(v float32) {
	o.PlayCount = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *MediaMediaIdWatchDataGet200ResponseData) GetUsers() []User {
	if o == nil || IsNil(o.Users) {
		var ret []User
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaMediaIdWatchDataGet200ResponseData) GetUsersOk() ([]User, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *MediaMediaIdWatchDataGet200ResponseData) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []User and assigns it to the Users field.
func (o *MediaMediaIdWatchDataGet200ResponseData) SetUsers(v []User) {
	o.Users = v
}

func (o MediaMediaIdWatchDataGet200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MediaMediaIdWatchDataGet200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PlayCount7Days) {
		toSerialize["playCount7Days"] = o.PlayCount7Days
	}
	if !IsNil(o.PlayCount30Days) {
		toSerialize["playCount30Days"] = o.PlayCount30Days
	}
	if !IsNil(o.PlayCount) {
		toSerialize["playCount"] = o.PlayCount
	}
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}
	return toSerialize, nil
}

type NullableMediaMediaIdWatchDataGet200ResponseData struct {
	value *MediaMediaIdWatchDataGet200ResponseData
	isSet bool
}

func (v NullableMediaMediaIdWatchDataGet200ResponseData) Get() *MediaMediaIdWatchDataGet200ResponseData {
	return v.value
}

func (v *NullableMediaMediaIdWatchDataGet200ResponseData) Set(val *MediaMediaIdWatchDataGet200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMediaMediaIdWatchDataGet200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMediaMediaIdWatchDataGet200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMediaMediaIdWatchDataGet200ResponseData(val *MediaMediaIdWatchDataGet200ResponseData) *NullableMediaMediaIdWatchDataGet200ResponseData {
	return &NullableMediaMediaIdWatchDataGet200ResponseData{value: val, isSet: true}
}

func (v NullableMediaMediaIdWatchDataGet200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMediaMediaIdWatchDataGet200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
