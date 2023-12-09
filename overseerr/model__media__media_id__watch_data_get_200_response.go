/*
Overseerr API

This is the documentation for the Overseerr API backend.  Two primary authentication methods are supported:  - **Cookie Authentication**: A valid sign-in to the `/auth/plex` or `/auth/local` will generate a valid authentication cookie. - **API Key Authentication**: Sign-in is also possible by passing an `X-Api-Key` header along with a valid API Key generated by Overseerr. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the MediaMediaIdWatchDataGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MediaMediaIdWatchDataGet200Response{}

// MediaMediaIdWatchDataGet200Response struct for MediaMediaIdWatchDataGet200Response
type MediaMediaIdWatchDataGet200Response struct {
	Data *MediaMediaIdWatchDataGet200ResponseData `json:"data,omitempty"`
	Data4k *MediaMediaIdWatchDataGet200ResponseData `json:"data4k,omitempty"`
}

// NewMediaMediaIdWatchDataGet200Response instantiates a new MediaMediaIdWatchDataGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMediaMediaIdWatchDataGet200Response() *MediaMediaIdWatchDataGet200Response {
	this := MediaMediaIdWatchDataGet200Response{}
	return &this
}

// NewMediaMediaIdWatchDataGet200ResponseWithDefaults instantiates a new MediaMediaIdWatchDataGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMediaMediaIdWatchDataGet200ResponseWithDefaults() *MediaMediaIdWatchDataGet200Response {
	this := MediaMediaIdWatchDataGet200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *MediaMediaIdWatchDataGet200Response) GetData() MediaMediaIdWatchDataGet200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret MediaMediaIdWatchDataGet200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaMediaIdWatchDataGet200Response) GetDataOk() (*MediaMediaIdWatchDataGet200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *MediaMediaIdWatchDataGet200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given MediaMediaIdWatchDataGet200ResponseData and assigns it to the Data field.
func (o *MediaMediaIdWatchDataGet200Response) SetData(v MediaMediaIdWatchDataGet200ResponseData) {
	o.Data = &v
}

// GetData4k returns the Data4k field value if set, zero value otherwise.
func (o *MediaMediaIdWatchDataGet200Response) GetData4k() MediaMediaIdWatchDataGet200ResponseData {
	if o == nil || IsNil(o.Data4k) {
		var ret MediaMediaIdWatchDataGet200ResponseData
		return ret
	}
	return *o.Data4k
}

// GetData4kOk returns a tuple with the Data4k field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaMediaIdWatchDataGet200Response) GetData4kOk() (*MediaMediaIdWatchDataGet200ResponseData, bool) {
	if o == nil || IsNil(o.Data4k) {
		return nil, false
	}
	return o.Data4k, true
}

// HasData4k returns a boolean if a field has been set.
func (o *MediaMediaIdWatchDataGet200Response) HasData4k() bool {
	if o != nil && !IsNil(o.Data4k) {
		return true
	}

	return false
}

// SetData4k gets a reference to the given MediaMediaIdWatchDataGet200ResponseData and assigns it to the Data4k field.
func (o *MediaMediaIdWatchDataGet200Response) SetData4k(v MediaMediaIdWatchDataGet200ResponseData) {
	o.Data4k = &v
}

func (o MediaMediaIdWatchDataGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MediaMediaIdWatchDataGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Data4k) {
		toSerialize["data4k"] = o.Data4k
	}
	return toSerialize, nil
}

type NullableMediaMediaIdWatchDataGet200Response struct {
	value *MediaMediaIdWatchDataGet200Response
	isSet bool
}

func (v NullableMediaMediaIdWatchDataGet200Response) Get() *MediaMediaIdWatchDataGet200Response {
	return v.value
}

func (v *NullableMediaMediaIdWatchDataGet200Response) Set(val *MediaMediaIdWatchDataGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableMediaMediaIdWatchDataGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableMediaMediaIdWatchDataGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMediaMediaIdWatchDataGet200Response(val *MediaMediaIdWatchDataGet200Response) *NullableMediaMediaIdWatchDataGet200Response {
	return &NullableMediaMediaIdWatchDataGet200Response{value: val, isSet: true}
}

func (v NullableMediaMediaIdWatchDataGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMediaMediaIdWatchDataGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


