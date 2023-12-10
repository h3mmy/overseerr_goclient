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

// checks if the ServiceProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceProfile{}

// ServiceProfile struct for ServiceProfile
type ServiceProfile struct {
	Id *float32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewServiceProfile instantiates a new ServiceProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceProfile() *ServiceProfile {
	this := ServiceProfile{}
	return &this
}

// NewServiceProfileWithDefaults instantiates a new ServiceProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceProfileWithDefaults() *ServiceProfile {
	this := ServiceProfile{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ServiceProfile) GetId() float32 {
	if o == nil || IsNil(o.Id) {
		var ret float32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProfile) GetIdOk() (*float32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ServiceProfile) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given float32 and assigns it to the Id field.
func (o *ServiceProfile) SetId(v float32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ServiceProfile) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProfile) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ServiceProfile) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ServiceProfile) SetName(v string) {
	o.Name = &v
}

func (o ServiceProfile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableServiceProfile struct {
	value *ServiceProfile
	isSet bool
}

func (v NullableServiceProfile) Get() *ServiceProfile {
	return v.value
}

func (v *NullableServiceProfile) Set(val *ServiceProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceProfile(val *ServiceProfile) *NullableServiceProfile {
	return &NullableServiceProfile{value: val, isSet: true}
}

func (v NullableServiceProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


