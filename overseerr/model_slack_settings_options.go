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

// checks if the SlackSettingsOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SlackSettingsOptions{}

// SlackSettingsOptions struct for SlackSettingsOptions
type SlackSettingsOptions struct {
	WebhookUrl *string `json:"webhookUrl,omitempty"`
}

// NewSlackSettingsOptions instantiates a new SlackSettingsOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSlackSettingsOptions() *SlackSettingsOptions {
	this := SlackSettingsOptions{}
	return &this
}

// NewSlackSettingsOptionsWithDefaults instantiates a new SlackSettingsOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSlackSettingsOptionsWithDefaults() *SlackSettingsOptions {
	this := SlackSettingsOptions{}
	return &this
}

// GetWebhookUrl returns the WebhookUrl field value if set, zero value otherwise.
func (o *SlackSettingsOptions) GetWebhookUrl() string {
	if o == nil || IsNil(o.WebhookUrl) {
		var ret string
		return ret
	}
	return *o.WebhookUrl
}

// GetWebhookUrlOk returns a tuple with the WebhookUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlackSettingsOptions) GetWebhookUrlOk() (*string, bool) {
	if o == nil || IsNil(o.WebhookUrl) {
		return nil, false
	}
	return o.WebhookUrl, true
}

// HasWebhookUrl returns a boolean if a field has been set.
func (o *SlackSettingsOptions) HasWebhookUrl() bool {
	if o != nil && !IsNil(o.WebhookUrl) {
		return true
	}

	return false
}

// SetWebhookUrl gets a reference to the given string and assigns it to the WebhookUrl field.
func (o *SlackSettingsOptions) SetWebhookUrl(v string) {
	o.WebhookUrl = &v
}

func (o SlackSettingsOptions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SlackSettingsOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.WebhookUrl) {
		toSerialize["webhookUrl"] = o.WebhookUrl
	}
	return toSerialize, nil
}

type NullableSlackSettingsOptions struct {
	value *SlackSettingsOptions
	isSet bool
}

func (v NullableSlackSettingsOptions) Get() *SlackSettingsOptions {
	return v.value
}

func (v *NullableSlackSettingsOptions) Set(val *SlackSettingsOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableSlackSettingsOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableSlackSettingsOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSlackSettingsOptions(val *SlackSettingsOptions) *NullableSlackSettingsOptions {
	return &NullableSlackSettingsOptions{value: val, isSet: true}
}

func (v NullableSlackSettingsOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSlackSettingsOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
