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

// checks if the AuthLocalPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthLocalPostRequest{}

// AuthLocalPostRequest struct for AuthLocalPostRequest
type AuthLocalPostRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type _AuthLocalPostRequest AuthLocalPostRequest

// NewAuthLocalPostRequest instantiates a new AuthLocalPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthLocalPostRequest(email string, password string) *AuthLocalPostRequest {
	this := AuthLocalPostRequest{}
	this.Email = email
	this.Password = password
	return &this
}

// NewAuthLocalPostRequestWithDefaults instantiates a new AuthLocalPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthLocalPostRequestWithDefaults() *AuthLocalPostRequest {
	this := AuthLocalPostRequest{}
	return &this
}

// GetEmail returns the Email field value
func (o *AuthLocalPostRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *AuthLocalPostRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *AuthLocalPostRequest) SetEmail(v string) {
	o.Email = v
}

// GetPassword returns the Password field value
func (o *AuthLocalPostRequest) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *AuthLocalPostRequest) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *AuthLocalPostRequest) SetPassword(v string) {
	o.Password = v
}

func (o AuthLocalPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthLocalPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	toSerialize["password"] = o.Password
	return toSerialize, nil
}

func (o *AuthLocalPostRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
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

	varAuthLocalPostRequest := _AuthLocalPostRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAuthLocalPostRequest)

	if err != nil {
		return err
	}

	*o = AuthLocalPostRequest(varAuthLocalPostRequest)

	return err
}

type NullableAuthLocalPostRequest struct {
	value *AuthLocalPostRequest
	isSet bool
}

func (v NullableAuthLocalPostRequest) Get() *AuthLocalPostRequest {
	return v.value
}

func (v *NullableAuthLocalPostRequest) Set(val *AuthLocalPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthLocalPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthLocalPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthLocalPostRequest(val *AuthLocalPostRequest) *NullableAuthLocalPostRequest {
	return &NullableAuthLocalPostRequest{value: val, isSet: true}
}

func (v NullableAuthLocalPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthLocalPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
