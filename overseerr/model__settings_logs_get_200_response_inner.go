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

// checks if the SettingsLogsGet200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SettingsLogsGet200ResponseInner{}

// SettingsLogsGet200ResponseInner struct for SettingsLogsGet200ResponseInner
type SettingsLogsGet200ResponseInner struct {
	Label     *string `json:"label,omitempty"`
	Level     *string `json:"level,omitempty"`
	Message   *string `json:"message,omitempty"`
	Timestamp *string `json:"timestamp,omitempty"`
}

// NewSettingsLogsGet200ResponseInner instantiates a new SettingsLogsGet200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSettingsLogsGet200ResponseInner() *SettingsLogsGet200ResponseInner {
	this := SettingsLogsGet200ResponseInner{}
	return &this
}

// NewSettingsLogsGet200ResponseInnerWithDefaults instantiates a new SettingsLogsGet200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSettingsLogsGet200ResponseInnerWithDefaults() *SettingsLogsGet200ResponseInner {
	this := SettingsLogsGet200ResponseInner{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *SettingsLogsGet200ResponseInner) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsLogsGet200ResponseInner) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *SettingsLogsGet200ResponseInner) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *SettingsLogsGet200ResponseInner) SetLabel(v string) {
	o.Label = &v
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *SettingsLogsGet200ResponseInner) GetLevel() string {
	if o == nil || IsNil(o.Level) {
		var ret string
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsLogsGet200ResponseInner) GetLevelOk() (*string, bool) {
	if o == nil || IsNil(o.Level) {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *SettingsLogsGet200ResponseInner) HasLevel() bool {
	if o != nil && !IsNil(o.Level) {
		return true
	}

	return false
}

// SetLevel gets a reference to the given string and assigns it to the Level field.
func (o *SettingsLogsGet200ResponseInner) SetLevel(v string) {
	o.Level = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *SettingsLogsGet200ResponseInner) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsLogsGet200ResponseInner) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *SettingsLogsGet200ResponseInner) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *SettingsLogsGet200ResponseInner) SetMessage(v string) {
	o.Message = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *SettingsLogsGet200ResponseInner) GetTimestamp() string {
	if o == nil || IsNil(o.Timestamp) {
		var ret string
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsLogsGet200ResponseInner) GetTimestampOk() (*string, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *SettingsLogsGet200ResponseInner) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given string and assigns it to the Timestamp field.
func (o *SettingsLogsGet200ResponseInner) SetTimestamp(v string) {
	o.Timestamp = &v
}

func (o SettingsLogsGet200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SettingsLogsGet200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Level) {
		toSerialize["level"] = o.Level
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	return toSerialize, nil
}

type NullableSettingsLogsGet200ResponseInner struct {
	value *SettingsLogsGet200ResponseInner
	isSet bool
}

func (v NullableSettingsLogsGet200ResponseInner) Get() *SettingsLogsGet200ResponseInner {
	return v.value
}

func (v *NullableSettingsLogsGet200ResponseInner) Set(val *SettingsLogsGet200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSettingsLogsGet200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSettingsLogsGet200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSettingsLogsGet200ResponseInner(val *SettingsLogsGet200ResponseInner) *NullableSettingsLogsGet200ResponseInner {
	return &NullableSettingsLogsGet200ResponseInner{value: val, isSet: true}
}

func (v NullableSettingsLogsGet200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSettingsLogsGet200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
