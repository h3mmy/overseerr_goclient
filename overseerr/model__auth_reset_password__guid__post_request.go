/*
Overseerr API

This is the documentation for the Overseerr API backend.  Two primary authentication methods are supported:  - **Cookie Authentication**: A valid sign-in to the `/auth/plex` or `/auth/local` will generate a valid authentication cookie. - **API Key Authentication**: Sign-in is also possible by passing an `X-Api-Key` header along with a valid API Key generated by Overseerr.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package overseerr_go

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the AuthResetPasswordGuidPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthResetPasswordGuidPostRequest{}

// AuthResetPasswordGuidPostRequest struct for AuthResetPasswordGuidPostRequest
type AuthResetPasswordGuidPostRequest struct {
	Password string `json:"password"`
}

type _AuthResetPasswordGuidPostRequest AuthResetPasswordGuidPostRequest

// NewAuthResetPasswordGuidPostRequest instantiates a new AuthResetPasswordGuidPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthResetPasswordGuidPostRequest(password string) *AuthResetPasswordGuidPostRequest {
	this := AuthResetPasswordGuidPostRequest{}
	this.Password = password
	return &this
}

// NewAuthResetPasswordGuidPostRequestWithDefaults instantiates a new AuthResetPasswordGuidPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthResetPasswordGuidPostRequestWithDefaults() *AuthResetPasswordGuidPostRequest {
	this := AuthResetPasswordGuidPostRequest{}
	return &this
}

// GetPassword returns the Password field value
func (o *AuthResetPasswordGuidPostRequest) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *AuthResetPasswordGuidPostRequest) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *AuthResetPasswordGuidPostRequest) SetPassword(v string) {
	o.Password = v
}

func (o AuthResetPasswordGuidPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthResetPasswordGuidPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["password"] = o.Password
	return toSerialize, nil
}

func (o *AuthResetPasswordGuidPostRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"password",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAuthResetPasswordGuidPostRequest := _AuthResetPasswordGuidPostRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAuthResetPasswordGuidPostRequest)

	if err != nil {
		return err
	}

	*o = AuthResetPasswordGuidPostRequest(varAuthResetPasswordGuidPostRequest)

	return err
}

type NullableAuthResetPasswordGuidPostRequest struct {
	value *AuthResetPasswordGuidPostRequest
	isSet bool
}

func (v NullableAuthResetPasswordGuidPostRequest) Get() *AuthResetPasswordGuidPostRequest {
	return v.value
}

func (v *NullableAuthResetPasswordGuidPostRequest) Set(val *AuthResetPasswordGuidPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthResetPasswordGuidPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthResetPasswordGuidPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthResetPasswordGuidPostRequest(val *AuthResetPasswordGuidPostRequest) *NullableAuthResetPasswordGuidPostRequest {
	return &NullableAuthResetPasswordGuidPostRequest{value: val, isSet: true}
}

func (v NullableAuthResetPasswordGuidPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthResetPasswordGuidPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
