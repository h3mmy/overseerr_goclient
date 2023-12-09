/*
Overseerr API

This is the documentation for the Overseerr API backend.  Two primary authentication methods are supported:  - **Cookie Authentication**: A valid sign-in to the `/auth/plex` or `/auth/local` will generate a valid authentication cookie. - **API Key Authentication**: Sign-in is also possible by passing an `X-Api-Key` header along with a valid API Key generated by Overseerr. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the TvTvIdRatingsGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TvTvIdRatingsGet200Response{}

// TvTvIdRatingsGet200Response struct for TvTvIdRatingsGet200Response
type TvTvIdRatingsGet200Response struct {
	Title *string `json:"title,omitempty"`
	Year *float32 `json:"year,omitempty"`
	Url *string `json:"url,omitempty"`
	CriticsScore *float32 `json:"criticsScore,omitempty"`
	CriticsRating *string `json:"criticsRating,omitempty"`
}

// NewTvTvIdRatingsGet200Response instantiates a new TvTvIdRatingsGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTvTvIdRatingsGet200Response() *TvTvIdRatingsGet200Response {
	this := TvTvIdRatingsGet200Response{}
	return &this
}

// NewTvTvIdRatingsGet200ResponseWithDefaults instantiates a new TvTvIdRatingsGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTvTvIdRatingsGet200ResponseWithDefaults() *TvTvIdRatingsGet200Response {
	this := TvTvIdRatingsGet200Response{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *TvTvIdRatingsGet200Response) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvTvIdRatingsGet200Response) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *TvTvIdRatingsGet200Response) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *TvTvIdRatingsGet200Response) SetTitle(v string) {
	o.Title = &v
}

// GetYear returns the Year field value if set, zero value otherwise.
func (o *TvTvIdRatingsGet200Response) GetYear() float32 {
	if o == nil || IsNil(o.Year) {
		var ret float32
		return ret
	}
	return *o.Year
}

// GetYearOk returns a tuple with the Year field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvTvIdRatingsGet200Response) GetYearOk() (*float32, bool) {
	if o == nil || IsNil(o.Year) {
		return nil, false
	}
	return o.Year, true
}

// HasYear returns a boolean if a field has been set.
func (o *TvTvIdRatingsGet200Response) HasYear() bool {
	if o != nil && !IsNil(o.Year) {
		return true
	}

	return false
}

// SetYear gets a reference to the given float32 and assigns it to the Year field.
func (o *TvTvIdRatingsGet200Response) SetYear(v float32) {
	o.Year = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *TvTvIdRatingsGet200Response) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvTvIdRatingsGet200Response) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *TvTvIdRatingsGet200Response) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *TvTvIdRatingsGet200Response) SetUrl(v string) {
	o.Url = &v
}

// GetCriticsScore returns the CriticsScore field value if set, zero value otherwise.
func (o *TvTvIdRatingsGet200Response) GetCriticsScore() float32 {
	if o == nil || IsNil(o.CriticsScore) {
		var ret float32
		return ret
	}
	return *o.CriticsScore
}

// GetCriticsScoreOk returns a tuple with the CriticsScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvTvIdRatingsGet200Response) GetCriticsScoreOk() (*float32, bool) {
	if o == nil || IsNil(o.CriticsScore) {
		return nil, false
	}
	return o.CriticsScore, true
}

// HasCriticsScore returns a boolean if a field has been set.
func (o *TvTvIdRatingsGet200Response) HasCriticsScore() bool {
	if o != nil && !IsNil(o.CriticsScore) {
		return true
	}

	return false
}

// SetCriticsScore gets a reference to the given float32 and assigns it to the CriticsScore field.
func (o *TvTvIdRatingsGet200Response) SetCriticsScore(v float32) {
	o.CriticsScore = &v
}

// GetCriticsRating returns the CriticsRating field value if set, zero value otherwise.
func (o *TvTvIdRatingsGet200Response) GetCriticsRating() string {
	if o == nil || IsNil(o.CriticsRating) {
		var ret string
		return ret
	}
	return *o.CriticsRating
}

// GetCriticsRatingOk returns a tuple with the CriticsRating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvTvIdRatingsGet200Response) GetCriticsRatingOk() (*string, bool) {
	if o == nil || IsNil(o.CriticsRating) {
		return nil, false
	}
	return o.CriticsRating, true
}

// HasCriticsRating returns a boolean if a field has been set.
func (o *TvTvIdRatingsGet200Response) HasCriticsRating() bool {
	if o != nil && !IsNil(o.CriticsRating) {
		return true
	}

	return false
}

// SetCriticsRating gets a reference to the given string and assigns it to the CriticsRating field.
func (o *TvTvIdRatingsGet200Response) SetCriticsRating(v string) {
	o.CriticsRating = &v
}

func (o TvTvIdRatingsGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TvTvIdRatingsGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Year) {
		toSerialize["year"] = o.Year
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.CriticsScore) {
		toSerialize["criticsScore"] = o.CriticsScore
	}
	if !IsNil(o.CriticsRating) {
		toSerialize["criticsRating"] = o.CriticsRating
	}
	return toSerialize, nil
}

type NullableTvTvIdRatingsGet200Response struct {
	value *TvTvIdRatingsGet200Response
	isSet bool
}

func (v NullableTvTvIdRatingsGet200Response) Get() *TvTvIdRatingsGet200Response {
	return v.value
}

func (v *NullableTvTvIdRatingsGet200Response) Set(val *TvTvIdRatingsGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableTvTvIdRatingsGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableTvTvIdRatingsGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTvTvIdRatingsGet200Response(val *TvTvIdRatingsGet200Response) *NullableTvTvIdRatingsGet200Response {
	return &NullableTvTvIdRatingsGet200Response{value: val, isSet: true}
}

func (v NullableTvTvIdRatingsGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTvTvIdRatingsGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


