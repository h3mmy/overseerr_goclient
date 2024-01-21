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

// checks if the SettingsDiscoverAddPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SettingsDiscoverAddPostRequest{}

// SettingsDiscoverAddPostRequest struct for SettingsDiscoverAddPostRequest
type SettingsDiscoverAddPostRequest struct {
	Title *string  `json:"title,omitempty"`
	Type  *float32 `json:"type,omitempty"`
	Data  *string  `json:"data,omitempty"`
}

// NewSettingsDiscoverAddPostRequest instantiates a new SettingsDiscoverAddPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSettingsDiscoverAddPostRequest() *SettingsDiscoverAddPostRequest {
	this := SettingsDiscoverAddPostRequest{}
	return &this
}

// NewSettingsDiscoverAddPostRequestWithDefaults instantiates a new SettingsDiscoverAddPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSettingsDiscoverAddPostRequestWithDefaults() *SettingsDiscoverAddPostRequest {
	this := SettingsDiscoverAddPostRequest{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *SettingsDiscoverAddPostRequest) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsDiscoverAddPostRequest) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *SettingsDiscoverAddPostRequest) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *SettingsDiscoverAddPostRequest) SetTitle(v string) {
	o.Title = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SettingsDiscoverAddPostRequest) GetType() float32 {
	if o == nil || IsNil(o.Type) {
		var ret float32
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsDiscoverAddPostRequest) GetTypeOk() (*float32, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SettingsDiscoverAddPostRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given float32 and assigns it to the Type field.
func (o *SettingsDiscoverAddPostRequest) SetType(v float32) {
	o.Type = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SettingsDiscoverAddPostRequest) GetData() string {
	if o == nil || IsNil(o.Data) {
		var ret string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsDiscoverAddPostRequest) GetDataOk() (*string, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SettingsDiscoverAddPostRequest) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given string and assigns it to the Data field.
func (o *SettingsDiscoverAddPostRequest) SetData(v string) {
	o.Data = &v
}

func (o SettingsDiscoverAddPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SettingsDiscoverAddPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableSettingsDiscoverAddPostRequest struct {
	value *SettingsDiscoverAddPostRequest
	isSet bool
}

func (v NullableSettingsDiscoverAddPostRequest) Get() *SettingsDiscoverAddPostRequest {
	return v.value
}

func (v *NullableSettingsDiscoverAddPostRequest) Set(val *SettingsDiscoverAddPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSettingsDiscoverAddPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSettingsDiscoverAddPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSettingsDiscoverAddPostRequest(val *SettingsDiscoverAddPostRequest) *NullableSettingsDiscoverAddPostRequest {
	return &NullableSettingsDiscoverAddPostRequest{value: val, isSet: true}
}

func (v NullableSettingsDiscoverAddPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSettingsDiscoverAddPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
