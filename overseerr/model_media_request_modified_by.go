/*
Overseerr API

This is the documentation for the Overseerr API backend.  Two primary authentication methods are supported:  - **Cookie Authentication**: A valid sign-in to the `/auth/plex` or `/auth/local` will generate a valid authentication cookie. - **API Key Authentication**: Sign-in is also possible by passing an `X-Api-Key` header along with a valid API Key generated by Overseerr. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// MediaRequestModifiedBy struct for MediaRequestModifiedBy
type MediaRequestModifiedBy struct {
	User *User
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *MediaRequestModifiedBy) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into User
	err = json.Unmarshal(data, &dst.User);
	if err == nil {
		jsonUser, _ := json.Marshal(dst.User)
		if string(jsonUser) == "{}" { // empty struct
			dst.User = nil
		} else {
			return nil // data stored in dst.User, return on the first match
		}
	} else {
		dst.User = nil
	}

	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.string);
	if err == nil {
		jsonstring, _ := json.Marshal(dst.string)
		if string(jsonstring) == "{}" { // empty struct
			dst.string = nil
		} else {
			return nil // data stored in dst.string, return on the first match
		}
	} else {
		dst.string = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(MediaRequestModifiedBy)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *MediaRequestModifiedBy) MarshalJSON() ([]byte, error) {
	if src.User != nil {
		return json.Marshal(&src.User)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableMediaRequestModifiedBy struct {
	value *MediaRequestModifiedBy
	isSet bool
}

func (v NullableMediaRequestModifiedBy) Get() *MediaRequestModifiedBy {
	return v.value
}

func (v *NullableMediaRequestModifiedBy) Set(val *MediaRequestModifiedBy) {
	v.value = val
	v.isSet = true
}

func (v NullableMediaRequestModifiedBy) IsSet() bool {
	return v.isSet
}

func (v *NullableMediaRequestModifiedBy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMediaRequestModifiedBy(val *MediaRequestModifiedBy) *NullableMediaRequestModifiedBy {
	return &NullableMediaRequestModifiedBy{value: val, isSet: true}
}

func (v NullableMediaRequestModifiedBy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMediaRequestModifiedBy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


