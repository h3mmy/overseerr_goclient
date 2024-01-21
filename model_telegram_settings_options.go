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

// checks if the TelegramSettingsOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TelegramSettingsOptions{}

// TelegramSettingsOptions struct for TelegramSettingsOptions
type TelegramSettingsOptions struct {
	BotUsername *string `json:"botUsername,omitempty"`
	BotAPI *string `json:"botAPI,omitempty"`
	ChatId *string `json:"chatId,omitempty"`
	SendSilently *bool `json:"sendSilently,omitempty"`
}

// NewTelegramSettingsOptions instantiates a new TelegramSettingsOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelegramSettingsOptions() *TelegramSettingsOptions {
	this := TelegramSettingsOptions{}
	return &this
}

// NewTelegramSettingsOptionsWithDefaults instantiates a new TelegramSettingsOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelegramSettingsOptionsWithDefaults() *TelegramSettingsOptions {
	this := TelegramSettingsOptions{}
	return &this
}

// GetBotUsername returns the BotUsername field value if set, zero value otherwise.
func (o *TelegramSettingsOptions) GetBotUsername() string {
	if o == nil || IsNil(o.BotUsername) {
		var ret string
		return ret
	}
	return *o.BotUsername
}

// GetBotUsernameOk returns a tuple with the BotUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelegramSettingsOptions) GetBotUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.BotUsername) {
		return nil, false
	}
	return o.BotUsername, true
}

// HasBotUsername returns a boolean if a field has been set.
func (o *TelegramSettingsOptions) HasBotUsername() bool {
	if o != nil && !IsNil(o.BotUsername) {
		return true
	}

	return false
}

// SetBotUsername gets a reference to the given string and assigns it to the BotUsername field.
func (o *TelegramSettingsOptions) SetBotUsername(v string) {
	o.BotUsername = &v
}

// GetBotAPI returns the BotAPI field value if set, zero value otherwise.
func (o *TelegramSettingsOptions) GetBotAPI() string {
	if o == nil || IsNil(o.BotAPI) {
		var ret string
		return ret
	}
	return *o.BotAPI
}

// GetBotAPIOk returns a tuple with the BotAPI field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelegramSettingsOptions) GetBotAPIOk() (*string, bool) {
	if o == nil || IsNil(o.BotAPI) {
		return nil, false
	}
	return o.BotAPI, true
}

// HasBotAPI returns a boolean if a field has been set.
func (o *TelegramSettingsOptions) HasBotAPI() bool {
	if o != nil && !IsNil(o.BotAPI) {
		return true
	}

	return false
}

// SetBotAPI gets a reference to the given string and assigns it to the BotAPI field.
func (o *TelegramSettingsOptions) SetBotAPI(v string) {
	o.BotAPI = &v
}

// GetChatId returns the ChatId field value if set, zero value otherwise.
func (o *TelegramSettingsOptions) GetChatId() string {
	if o == nil || IsNil(o.ChatId) {
		var ret string
		return ret
	}
	return *o.ChatId
}

// GetChatIdOk returns a tuple with the ChatId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelegramSettingsOptions) GetChatIdOk() (*string, bool) {
	if o == nil || IsNil(o.ChatId) {
		return nil, false
	}
	return o.ChatId, true
}

// HasChatId returns a boolean if a field has been set.
func (o *TelegramSettingsOptions) HasChatId() bool {
	if o != nil && !IsNil(o.ChatId) {
		return true
	}

	return false
}

// SetChatId gets a reference to the given string and assigns it to the ChatId field.
func (o *TelegramSettingsOptions) SetChatId(v string) {
	o.ChatId = &v
}

// GetSendSilently returns the SendSilently field value if set, zero value otherwise.
func (o *TelegramSettingsOptions) GetSendSilently() bool {
	if o == nil || IsNil(o.SendSilently) {
		var ret bool
		return ret
	}
	return *o.SendSilently
}

// GetSendSilentlyOk returns a tuple with the SendSilently field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelegramSettingsOptions) GetSendSilentlyOk() (*bool, bool) {
	if o == nil || IsNil(o.SendSilently) {
		return nil, false
	}
	return o.SendSilently, true
}

// HasSendSilently returns a boolean if a field has been set.
func (o *TelegramSettingsOptions) HasSendSilently() bool {
	if o != nil && !IsNil(o.SendSilently) {
		return true
	}

	return false
}

// SetSendSilently gets a reference to the given bool and assigns it to the SendSilently field.
func (o *TelegramSettingsOptions) SetSendSilently(v bool) {
	o.SendSilently = &v
}

func (o TelegramSettingsOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TelegramSettingsOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BotUsername) {
		toSerialize["botUsername"] = o.BotUsername
	}
	if !IsNil(o.BotAPI) {
		toSerialize["botAPI"] = o.BotAPI
	}
	if !IsNil(o.ChatId) {
		toSerialize["chatId"] = o.ChatId
	}
	if !IsNil(o.SendSilently) {
		toSerialize["sendSilently"] = o.SendSilently
	}
	return toSerialize, nil
}

type NullableTelegramSettingsOptions struct {
	value *TelegramSettingsOptions
	isSet bool
}

func (v NullableTelegramSettingsOptions) Get() *TelegramSettingsOptions {
	return v.value
}

func (v *NullableTelegramSettingsOptions) Set(val *TelegramSettingsOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableTelegramSettingsOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableTelegramSettingsOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelegramSettingsOptions(val *TelegramSettingsOptions) *NullableTelegramSettingsOptions {
	return &NullableTelegramSettingsOptions{value: val, isSet: true}
}

func (v NullableTelegramSettingsOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelegramSettingsOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

