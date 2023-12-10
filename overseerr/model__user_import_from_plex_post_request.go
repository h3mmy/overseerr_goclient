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

// checks if the UserImportFromPlexPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserImportFromPlexPostRequest{}

// UserImportFromPlexPostRequest struct for UserImportFromPlexPostRequest
type UserImportFromPlexPostRequest struct {
	PlexIds []string `json:"plexIds,omitempty"`
}

// NewUserImportFromPlexPostRequest instantiates a new UserImportFromPlexPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserImportFromPlexPostRequest() *UserImportFromPlexPostRequest {
	this := UserImportFromPlexPostRequest{}
	return &this
}

// NewUserImportFromPlexPostRequestWithDefaults instantiates a new UserImportFromPlexPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserImportFromPlexPostRequestWithDefaults() *UserImportFromPlexPostRequest {
	this := UserImportFromPlexPostRequest{}
	return &this
}

// GetPlexIds returns the PlexIds field value if set, zero value otherwise.
func (o *UserImportFromPlexPostRequest) GetPlexIds() []string {
	if o == nil || IsNil(o.PlexIds) {
		var ret []string
		return ret
	}
	return o.PlexIds
}

// GetPlexIdsOk returns a tuple with the PlexIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserImportFromPlexPostRequest) GetPlexIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.PlexIds) {
		return nil, false
	}
	return o.PlexIds, true
}

// HasPlexIds returns a boolean if a field has been set.
func (o *UserImportFromPlexPostRequest) HasPlexIds() bool {
	if o != nil && !IsNil(o.PlexIds) {
		return true
	}

	return false
}

// SetPlexIds gets a reference to the given []string and assigns it to the PlexIds field.
func (o *UserImportFromPlexPostRequest) SetPlexIds(v []string) {
	o.PlexIds = v
}

func (o UserImportFromPlexPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserImportFromPlexPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PlexIds) {
		toSerialize["plexIds"] = o.PlexIds
	}
	return toSerialize, nil
}

type NullableUserImportFromPlexPostRequest struct {
	value *UserImportFromPlexPostRequest
	isSet bool
}

func (v NullableUserImportFromPlexPostRequest) Get() *UserImportFromPlexPostRequest {
	return v.value
}

func (v *NullableUserImportFromPlexPostRequest) Set(val *UserImportFromPlexPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserImportFromPlexPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserImportFromPlexPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserImportFromPlexPostRequest(val *UserImportFromPlexPostRequest) *NullableUserImportFromPlexPostRequest {
	return &NullableUserImportFromPlexPostRequest{value: val, isSet: true}
}

func (v NullableUserImportFromPlexPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserImportFromPlexPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


