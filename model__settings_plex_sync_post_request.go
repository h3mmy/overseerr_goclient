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

// checks if the SettingsPlexSyncPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SettingsPlexSyncPostRequest{}

// SettingsPlexSyncPostRequest struct for SettingsPlexSyncPostRequest
type SettingsPlexSyncPostRequest struct {
	Cancel *bool `json:"cancel,omitempty"`
	Start *bool `json:"start,omitempty"`
}

// NewSettingsPlexSyncPostRequest instantiates a new SettingsPlexSyncPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSettingsPlexSyncPostRequest() *SettingsPlexSyncPostRequest {
	this := SettingsPlexSyncPostRequest{}
	return &this
}

// NewSettingsPlexSyncPostRequestWithDefaults instantiates a new SettingsPlexSyncPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSettingsPlexSyncPostRequestWithDefaults() *SettingsPlexSyncPostRequest {
	this := SettingsPlexSyncPostRequest{}
	return &this
}

// GetCancel returns the Cancel field value if set, zero value otherwise.
func (o *SettingsPlexSyncPostRequest) GetCancel() bool {
	if o == nil || IsNil(o.Cancel) {
		var ret bool
		return ret
	}
	return *o.Cancel
}

// GetCancelOk returns a tuple with the Cancel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsPlexSyncPostRequest) GetCancelOk() (*bool, bool) {
	if o == nil || IsNil(o.Cancel) {
		return nil, false
	}
	return o.Cancel, true
}

// HasCancel returns a boolean if a field has been set.
func (o *SettingsPlexSyncPostRequest) HasCancel() bool {
	if o != nil && !IsNil(o.Cancel) {
		return true
	}

	return false
}

// SetCancel gets a reference to the given bool and assigns it to the Cancel field.
func (o *SettingsPlexSyncPostRequest) SetCancel(v bool) {
	o.Cancel = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *SettingsPlexSyncPostRequest) GetStart() bool {
	if o == nil || IsNil(o.Start) {
		var ret bool
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsPlexSyncPostRequest) GetStartOk() (*bool, bool) {
	if o == nil || IsNil(o.Start) {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *SettingsPlexSyncPostRequest) HasStart() bool {
	if o != nil && !IsNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given bool and assigns it to the Start field.
func (o *SettingsPlexSyncPostRequest) SetStart(v bool) {
	o.Start = &v
}

func (o SettingsPlexSyncPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SettingsPlexSyncPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cancel) {
		toSerialize["cancel"] = o.Cancel
	}
	if !IsNil(o.Start) {
		toSerialize["start"] = o.Start
	}
	return toSerialize, nil
}

type NullableSettingsPlexSyncPostRequest struct {
	value *SettingsPlexSyncPostRequest
	isSet bool
}

func (v NullableSettingsPlexSyncPostRequest) Get() *SettingsPlexSyncPostRequest {
	return v.value
}

func (v *NullableSettingsPlexSyncPostRequest) Set(val *SettingsPlexSyncPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSettingsPlexSyncPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSettingsPlexSyncPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSettingsPlexSyncPostRequest(val *SettingsPlexSyncPostRequest) *NullableSettingsPlexSyncPostRequest {
	return &NullableSettingsPlexSyncPostRequest{value: val, isSet: true}
}

func (v NullableSettingsPlexSyncPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSettingsPlexSyncPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

