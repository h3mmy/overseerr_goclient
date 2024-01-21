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

// checks if the StatusGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StatusGet200Response{}

// StatusGet200Response struct for StatusGet200Response
type StatusGet200Response struct {
	Version         *string  `json:"version,omitempty"`
	CommitTag       *string  `json:"commitTag,omitempty"`
	UpdateAvailable *bool    `json:"updateAvailable,omitempty"`
	CommitsBehind   *float32 `json:"commitsBehind,omitempty"`
	RestartRequired *bool    `json:"restartRequired,omitempty"`
}

// NewStatusGet200Response instantiates a new StatusGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatusGet200Response() *StatusGet200Response {
	this := StatusGet200Response{}
	return &this
}

// NewStatusGet200ResponseWithDefaults instantiates a new StatusGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatusGet200ResponseWithDefaults() *StatusGet200Response {
	this := StatusGet200Response{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *StatusGet200Response) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusGet200Response) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *StatusGet200Response) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *StatusGet200Response) SetVersion(v string) {
	o.Version = &v
}

// GetCommitTag returns the CommitTag field value if set, zero value otherwise.
func (o *StatusGet200Response) GetCommitTag() string {
	if o == nil || IsNil(o.CommitTag) {
		var ret string
		return ret
	}
	return *o.CommitTag
}

// GetCommitTagOk returns a tuple with the CommitTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusGet200Response) GetCommitTagOk() (*string, bool) {
	if o == nil || IsNil(o.CommitTag) {
		return nil, false
	}
	return o.CommitTag, true
}

// HasCommitTag returns a boolean if a field has been set.
func (o *StatusGet200Response) HasCommitTag() bool {
	if o != nil && !IsNil(o.CommitTag) {
		return true
	}

	return false
}

// SetCommitTag gets a reference to the given string and assigns it to the CommitTag field.
func (o *StatusGet200Response) SetCommitTag(v string) {
	o.CommitTag = &v
}

// GetUpdateAvailable returns the UpdateAvailable field value if set, zero value otherwise.
func (o *StatusGet200Response) GetUpdateAvailable() bool {
	if o == nil || IsNil(o.UpdateAvailable) {
		var ret bool
		return ret
	}
	return *o.UpdateAvailable
}

// GetUpdateAvailableOk returns a tuple with the UpdateAvailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusGet200Response) GetUpdateAvailableOk() (*bool, bool) {
	if o == nil || IsNil(o.UpdateAvailable) {
		return nil, false
	}
	return o.UpdateAvailable, true
}

// HasUpdateAvailable returns a boolean if a field has been set.
func (o *StatusGet200Response) HasUpdateAvailable() bool {
	if o != nil && !IsNil(o.UpdateAvailable) {
		return true
	}

	return false
}

// SetUpdateAvailable gets a reference to the given bool and assigns it to the UpdateAvailable field.
func (o *StatusGet200Response) SetUpdateAvailable(v bool) {
	o.UpdateAvailable = &v
}

// GetCommitsBehind returns the CommitsBehind field value if set, zero value otherwise.
func (o *StatusGet200Response) GetCommitsBehind() float32 {
	if o == nil || IsNil(o.CommitsBehind) {
		var ret float32
		return ret
	}
	return *o.CommitsBehind
}

// GetCommitsBehindOk returns a tuple with the CommitsBehind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusGet200Response) GetCommitsBehindOk() (*float32, bool) {
	if o == nil || IsNil(o.CommitsBehind) {
		return nil, false
	}
	return o.CommitsBehind, true
}

// HasCommitsBehind returns a boolean if a field has been set.
func (o *StatusGet200Response) HasCommitsBehind() bool {
	if o != nil && !IsNil(o.CommitsBehind) {
		return true
	}

	return false
}

// SetCommitsBehind gets a reference to the given float32 and assigns it to the CommitsBehind field.
func (o *StatusGet200Response) SetCommitsBehind(v float32) {
	o.CommitsBehind = &v
}

// GetRestartRequired returns the RestartRequired field value if set, zero value otherwise.
func (o *StatusGet200Response) GetRestartRequired() bool {
	if o == nil || IsNil(o.RestartRequired) {
		var ret bool
		return ret
	}
	return *o.RestartRequired
}

// GetRestartRequiredOk returns a tuple with the RestartRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusGet200Response) GetRestartRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.RestartRequired) {
		return nil, false
	}
	return o.RestartRequired, true
}

// HasRestartRequired returns a boolean if a field has been set.
func (o *StatusGet200Response) HasRestartRequired() bool {
	if o != nil && !IsNil(o.RestartRequired) {
		return true
	}

	return false
}

// SetRestartRequired gets a reference to the given bool and assigns it to the RestartRequired field.
func (o *StatusGet200Response) SetRestartRequired(v bool) {
	o.RestartRequired = &v
}

func (o StatusGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StatusGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.CommitTag) {
		toSerialize["commitTag"] = o.CommitTag
	}
	if !IsNil(o.UpdateAvailable) {
		toSerialize["updateAvailable"] = o.UpdateAvailable
	}
	if !IsNil(o.CommitsBehind) {
		toSerialize["commitsBehind"] = o.CommitsBehind
	}
	if !IsNil(o.RestartRequired) {
		toSerialize["restartRequired"] = o.RestartRequired
	}
	return toSerialize, nil
}

type NullableStatusGet200Response struct {
	value *StatusGet200Response
	isSet bool
}

func (v NullableStatusGet200Response) Get() *StatusGet200Response {
	return v.value
}

func (v *NullableStatusGet200Response) Set(val *StatusGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableStatusGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableStatusGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatusGet200Response(val *StatusGet200Response) *NullableStatusGet200Response {
	return &NullableStatusGet200Response{value: val, isSet: true}
}

func (v NullableStatusGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatusGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
