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

// checks if the MovieMovieIdRatingscombinedGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MovieMovieIdRatingscombinedGet200Response{}

// MovieMovieIdRatingscombinedGet200Response struct for MovieMovieIdRatingscombinedGet200Response
type MovieMovieIdRatingscombinedGet200Response struct {
	Rt *MovieMovieIdRatingsGet200Response `json:"rt,omitempty"`
	Imdb *MovieMovieIdRatingscombinedGet200ResponseImdb `json:"imdb,omitempty"`
}

// NewMovieMovieIdRatingscombinedGet200Response instantiates a new MovieMovieIdRatingscombinedGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMovieMovieIdRatingscombinedGet200Response() *MovieMovieIdRatingscombinedGet200Response {
	this := MovieMovieIdRatingscombinedGet200Response{}
	return &this
}

// NewMovieMovieIdRatingscombinedGet200ResponseWithDefaults instantiates a new MovieMovieIdRatingscombinedGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMovieMovieIdRatingscombinedGet200ResponseWithDefaults() *MovieMovieIdRatingscombinedGet200Response {
	this := MovieMovieIdRatingscombinedGet200Response{}
	return &this
}

// GetRt returns the Rt field value if set, zero value otherwise.
func (o *MovieMovieIdRatingscombinedGet200Response) GetRt() MovieMovieIdRatingsGet200Response {
	if o == nil || IsNil(o.Rt) {
		var ret MovieMovieIdRatingsGet200Response
		return ret
	}
	return *o.Rt
}

// GetRtOk returns a tuple with the Rt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieMovieIdRatingscombinedGet200Response) GetRtOk() (*MovieMovieIdRatingsGet200Response, bool) {
	if o == nil || IsNil(o.Rt) {
		return nil, false
	}
	return o.Rt, true
}

// HasRt returns a boolean if a field has been set.
func (o *MovieMovieIdRatingscombinedGet200Response) HasRt() bool {
	if o != nil && !IsNil(o.Rt) {
		return true
	}

	return false
}

// SetRt gets a reference to the given MovieMovieIdRatingsGet200Response and assigns it to the Rt field.
func (o *MovieMovieIdRatingscombinedGet200Response) SetRt(v MovieMovieIdRatingsGet200Response) {
	o.Rt = &v
}

// GetImdb returns the Imdb field value if set, zero value otherwise.
func (o *MovieMovieIdRatingscombinedGet200Response) GetImdb() MovieMovieIdRatingscombinedGet200ResponseImdb {
	if o == nil || IsNil(o.Imdb) {
		var ret MovieMovieIdRatingscombinedGet200ResponseImdb
		return ret
	}
	return *o.Imdb
}

// GetImdbOk returns a tuple with the Imdb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovieMovieIdRatingscombinedGet200Response) GetImdbOk() (*MovieMovieIdRatingscombinedGet200ResponseImdb, bool) {
	if o == nil || IsNil(o.Imdb) {
		return nil, false
	}
	return o.Imdb, true
}

// HasImdb returns a boolean if a field has been set.
func (o *MovieMovieIdRatingscombinedGet200Response) HasImdb() bool {
	if o != nil && !IsNil(o.Imdb) {
		return true
	}

	return false
}

// SetImdb gets a reference to the given MovieMovieIdRatingscombinedGet200ResponseImdb and assigns it to the Imdb field.
func (o *MovieMovieIdRatingscombinedGet200Response) SetImdb(v MovieMovieIdRatingscombinedGet200ResponseImdb) {
	o.Imdb = &v
}

func (o MovieMovieIdRatingscombinedGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MovieMovieIdRatingscombinedGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Rt) {
		toSerialize["rt"] = o.Rt
	}
	if !IsNil(o.Imdb) {
		toSerialize["imdb"] = o.Imdb
	}
	return toSerialize, nil
}

type NullableMovieMovieIdRatingscombinedGet200Response struct {
	value *MovieMovieIdRatingscombinedGet200Response
	isSet bool
}

func (v NullableMovieMovieIdRatingscombinedGet200Response) Get() *MovieMovieIdRatingscombinedGet200Response {
	return v.value
}

func (v *NullableMovieMovieIdRatingscombinedGet200Response) Set(val *MovieMovieIdRatingscombinedGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableMovieMovieIdRatingscombinedGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableMovieMovieIdRatingscombinedGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMovieMovieIdRatingscombinedGet200Response(val *MovieMovieIdRatingscombinedGet200Response) *NullableMovieMovieIdRatingscombinedGet200Response {
	return &NullableMovieMovieIdRatingscombinedGet200Response{value: val, isSet: true}
}

func (v NullableMovieMovieIdRatingscombinedGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMovieMovieIdRatingscombinedGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


