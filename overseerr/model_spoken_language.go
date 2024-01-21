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

// checks if the SpokenLanguage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpokenLanguage{}

// SpokenLanguage struct for SpokenLanguage
type SpokenLanguage struct {
	EnglishName NullableString `json:"englishName,omitempty"`
	Iso6391     *string        `json:"iso_639_1,omitempty"`
	Name        *string        `json:"name,omitempty"`
}

// NewSpokenLanguage instantiates a new SpokenLanguage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpokenLanguage() *SpokenLanguage {
	this := SpokenLanguage{}
	return &this
}

// NewSpokenLanguageWithDefaults instantiates a new SpokenLanguage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpokenLanguageWithDefaults() *SpokenLanguage {
	this := SpokenLanguage{}
	return &this
}

// GetEnglishName returns the EnglishName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SpokenLanguage) GetEnglishName() string {
	if o == nil || IsNil(o.EnglishName.Get()) {
		var ret string
		return ret
	}
	return *o.EnglishName.Get()
}

// GetEnglishNameOk returns a tuple with the EnglishName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SpokenLanguage) GetEnglishNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnglishName.Get(), o.EnglishName.IsSet()
}

// HasEnglishName returns a boolean if a field has been set.
func (o *SpokenLanguage) HasEnglishName() bool {
	if o != nil && o.EnglishName.IsSet() {
		return true
	}

	return false
}

// SetEnglishName gets a reference to the given NullableString and assigns it to the EnglishName field.
func (o *SpokenLanguage) SetEnglishName(v string) {
	o.EnglishName.Set(&v)
}

// SetEnglishNameNil sets the value for EnglishName to be an explicit nil
func (o *SpokenLanguage) SetEnglishNameNil() {
	o.EnglishName.Set(nil)
}

// UnsetEnglishName ensures that no value is present for EnglishName, not even an explicit nil
func (o *SpokenLanguage) UnsetEnglishName() {
	o.EnglishName.Unset()
}

// GetIso6391 returns the Iso6391 field value if set, zero value otherwise.
func (o *SpokenLanguage) GetIso6391() string {
	if o == nil || IsNil(o.Iso6391) {
		var ret string
		return ret
	}
	return *o.Iso6391
}

// GetIso6391Ok returns a tuple with the Iso6391 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpokenLanguage) GetIso6391Ok() (*string, bool) {
	if o == nil || IsNil(o.Iso6391) {
		return nil, false
	}
	return o.Iso6391, true
}

// HasIso6391 returns a boolean if a field has been set.
func (o *SpokenLanguage) HasIso6391() bool {
	if o != nil && !IsNil(o.Iso6391) {
		return true
	}

	return false
}

// SetIso6391 gets a reference to the given string and assigns it to the Iso6391 field.
func (o *SpokenLanguage) SetIso6391(v string) {
	o.Iso6391 = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SpokenLanguage) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpokenLanguage) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SpokenLanguage) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SpokenLanguage) SetName(v string) {
	o.Name = &v
}

func (o SpokenLanguage) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpokenLanguage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.EnglishName.IsSet() {
		toSerialize["englishName"] = o.EnglishName.Get()
	}
	if !IsNil(o.Iso6391) {
		toSerialize["iso_639_1"] = o.Iso6391
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableSpokenLanguage struct {
	value *SpokenLanguage
	isSet bool
}

func (v NullableSpokenLanguage) Get() *SpokenLanguage {
	return v.value
}

func (v *NullableSpokenLanguage) Set(val *SpokenLanguage) {
	v.value = val
	v.isSet = true
}

func (v NullableSpokenLanguage) IsSet() bool {
	return v.isSet
}

func (v *NullableSpokenLanguage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpokenLanguage(val *SpokenLanguage) *NullableSpokenLanguage {
	return &NullableSpokenLanguage{value: val, isSet: true}
}

func (v NullableSpokenLanguage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpokenLanguage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
