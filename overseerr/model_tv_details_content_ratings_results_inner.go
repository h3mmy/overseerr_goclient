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

// checks if the TvDetailsContentRatingsResultsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TvDetailsContentRatingsResultsInner{}

// TvDetailsContentRatingsResultsInner struct for TvDetailsContentRatingsResultsInner
type TvDetailsContentRatingsResultsInner struct {
	Iso31661 *string `json:"iso_3166_1,omitempty"`
	Rating   *string `json:"rating,omitempty"`
}

// NewTvDetailsContentRatingsResultsInner instantiates a new TvDetailsContentRatingsResultsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTvDetailsContentRatingsResultsInner() *TvDetailsContentRatingsResultsInner {
	this := TvDetailsContentRatingsResultsInner{}
	return &this
}

// NewTvDetailsContentRatingsResultsInnerWithDefaults instantiates a new TvDetailsContentRatingsResultsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTvDetailsContentRatingsResultsInnerWithDefaults() *TvDetailsContentRatingsResultsInner {
	this := TvDetailsContentRatingsResultsInner{}
	return &this
}

// GetIso31661 returns the Iso31661 field value if set, zero value otherwise.
func (o *TvDetailsContentRatingsResultsInner) GetIso31661() string {
	if o == nil || IsNil(o.Iso31661) {
		var ret string
		return ret
	}
	return *o.Iso31661
}

// GetIso31661Ok returns a tuple with the Iso31661 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvDetailsContentRatingsResultsInner) GetIso31661Ok() (*string, bool) {
	if o == nil || IsNil(o.Iso31661) {
		return nil, false
	}
	return o.Iso31661, true
}

// HasIso31661 returns a boolean if a field has been set.
func (o *TvDetailsContentRatingsResultsInner) HasIso31661() bool {
	if o != nil && !IsNil(o.Iso31661) {
		return true
	}

	return false
}

// SetIso31661 gets a reference to the given string and assigns it to the Iso31661 field.
func (o *TvDetailsContentRatingsResultsInner) SetIso31661(v string) {
	o.Iso31661 = &v
}

// GetRating returns the Rating field value if set, zero value otherwise.
func (o *TvDetailsContentRatingsResultsInner) GetRating() string {
	if o == nil || IsNil(o.Rating) {
		var ret string
		return ret
	}
	return *o.Rating
}

// GetRatingOk returns a tuple with the Rating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvDetailsContentRatingsResultsInner) GetRatingOk() (*string, bool) {
	if o == nil || IsNil(o.Rating) {
		return nil, false
	}
	return o.Rating, true
}

// HasRating returns a boolean if a field has been set.
func (o *TvDetailsContentRatingsResultsInner) HasRating() bool {
	if o != nil && !IsNil(o.Rating) {
		return true
	}

	return false
}

// SetRating gets a reference to the given string and assigns it to the Rating field.
func (o *TvDetailsContentRatingsResultsInner) SetRating(v string) {
	o.Rating = &v
}

func (o TvDetailsContentRatingsResultsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TvDetailsContentRatingsResultsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Iso31661) {
		toSerialize["iso_3166_1"] = o.Iso31661
	}
	if !IsNil(o.Rating) {
		toSerialize["rating"] = o.Rating
	}
	return toSerialize, nil
}

type NullableTvDetailsContentRatingsResultsInner struct {
	value *TvDetailsContentRatingsResultsInner
	isSet bool
}

func (v NullableTvDetailsContentRatingsResultsInner) Get() *TvDetailsContentRatingsResultsInner {
	return v.value
}

func (v *NullableTvDetailsContentRatingsResultsInner) Set(val *TvDetailsContentRatingsResultsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableTvDetailsContentRatingsResultsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableTvDetailsContentRatingsResultsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTvDetailsContentRatingsResultsInner(val *TvDetailsContentRatingsResultsInner) *NullableTvDetailsContentRatingsResultsInner {
	return &NullableTvDetailsContentRatingsResultsInner{value: val, isSet: true}
}

func (v NullableTvDetailsContentRatingsResultsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTvDetailsContentRatingsResultsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
