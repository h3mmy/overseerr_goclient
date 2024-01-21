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

// checks if the RegionsGet200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegionsGet200ResponseInner{}

// RegionsGet200ResponseInner struct for RegionsGet200ResponseInner
type RegionsGet200ResponseInner struct {
	Iso31661    *string `json:"iso_3166_1,omitempty"`
	EnglishName *string `json:"english_name,omitempty"`
}

// NewRegionsGet200ResponseInner instantiates a new RegionsGet200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegionsGet200ResponseInner() *RegionsGet200ResponseInner {
	this := RegionsGet200ResponseInner{}
	return &this
}

// NewRegionsGet200ResponseInnerWithDefaults instantiates a new RegionsGet200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegionsGet200ResponseInnerWithDefaults() *RegionsGet200ResponseInner {
	this := RegionsGet200ResponseInner{}
	return &this
}

// GetIso31661 returns the Iso31661 field value if set, zero value otherwise.
func (o *RegionsGet200ResponseInner) GetIso31661() string {
	if o == nil || IsNil(o.Iso31661) {
		var ret string
		return ret
	}
	return *o.Iso31661
}

// GetIso31661Ok returns a tuple with the Iso31661 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionsGet200ResponseInner) GetIso31661Ok() (*string, bool) {
	if o == nil || IsNil(o.Iso31661) {
		return nil, false
	}
	return o.Iso31661, true
}

// HasIso31661 returns a boolean if a field has been set.
func (o *RegionsGet200ResponseInner) HasIso31661() bool {
	if o != nil && !IsNil(o.Iso31661) {
		return true
	}

	return false
}

// SetIso31661 gets a reference to the given string and assigns it to the Iso31661 field.
func (o *RegionsGet200ResponseInner) SetIso31661(v string) {
	o.Iso31661 = &v
}

// GetEnglishName returns the EnglishName field value if set, zero value otherwise.
func (o *RegionsGet200ResponseInner) GetEnglishName() string {
	if o == nil || IsNil(o.EnglishName) {
		var ret string
		return ret
	}
	return *o.EnglishName
}

// GetEnglishNameOk returns a tuple with the EnglishName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionsGet200ResponseInner) GetEnglishNameOk() (*string, bool) {
	if o == nil || IsNil(o.EnglishName) {
		return nil, false
	}
	return o.EnglishName, true
}

// HasEnglishName returns a boolean if a field has been set.
func (o *RegionsGet200ResponseInner) HasEnglishName() bool {
	if o != nil && !IsNil(o.EnglishName) {
		return true
	}

	return false
}

// SetEnglishName gets a reference to the given string and assigns it to the EnglishName field.
func (o *RegionsGet200ResponseInner) SetEnglishName(v string) {
	o.EnglishName = &v
}

func (o RegionsGet200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegionsGet200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Iso31661) {
		toSerialize["iso_3166_1"] = o.Iso31661
	}
	if !IsNil(o.EnglishName) {
		toSerialize["english_name"] = o.EnglishName
	}
	return toSerialize, nil
}

type NullableRegionsGet200ResponseInner struct {
	value *RegionsGet200ResponseInner
	isSet bool
}

func (v NullableRegionsGet200ResponseInner) Get() *RegionsGet200ResponseInner {
	return v.value
}

func (v *NullableRegionsGet200ResponseInner) Set(val *RegionsGet200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableRegionsGet200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableRegionsGet200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegionsGet200ResponseInner(val *RegionsGet200ResponseInner) *NullableRegionsGet200ResponseInner {
	return &NullableRegionsGet200ResponseInner{value: val, isSet: true}
}

func (v NullableRegionsGet200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegionsGet200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
