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

// checks if the MovieDetailsReleases type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MovieDetailsReleases{}

// MovieDetailsReleases struct for MovieDetailsReleases
type MovieDetailsReleases struct {
	Results []MovieDetailsReleasesResultsInner `json:"results,omitempty"`
}

// NewMovieDetailsReleases instantiates a new MovieDetailsReleases object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMovieDetailsReleases() *MovieDetailsReleases {
	this := MovieDetailsReleases{}
	return &this
}

// NewMovieDetailsReleasesWithDefaults instantiates a new MovieDetailsReleases object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMovieDetailsReleasesWithDefaults() *MovieDetailsReleases {
	this := MovieDetailsReleases{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *MovieDetailsReleases) GetResults() []MovieDetailsReleasesResultsInner {
	if o == nil || IsNil(o.Results) {
		var ret []MovieDetailsReleasesResultsInner
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieDetailsReleases) GetResultsOk() ([]MovieDetailsReleasesResultsInner, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *MovieDetailsReleases) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []MovieDetailsReleasesResultsInner and assigns it to the Results field.
func (o *MovieDetailsReleases) SetResults(v []MovieDetailsReleasesResultsInner) {
	o.Results = v
}

func (o MovieDetailsReleases) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MovieDetailsReleases) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullableMovieDetailsReleases struct {
	value *MovieDetailsReleases
	isSet bool
}

func (v NullableMovieDetailsReleases) Get() *MovieDetailsReleases {
	return v.value
}

func (v *NullableMovieDetailsReleases) Set(val *MovieDetailsReleases) {
	v.value = val
	v.isSet = true
}

func (v NullableMovieDetailsReleases) IsSet() bool {
	return v.isSet
}

func (v *NullableMovieDetailsReleases) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMovieDetailsReleases(val *MovieDetailsReleases) *NullableMovieDetailsReleases {
	return &NullableMovieDetailsReleases{value: val, isSet: true}
}

func (v NullableMovieDetailsReleases) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMovieDetailsReleases) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
