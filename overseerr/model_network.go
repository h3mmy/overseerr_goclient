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

// checks if the Network type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Network{}

// Network struct for Network
type Network struct {
	Id *float32 `json:"id,omitempty"`
	LogoPath NullableString `json:"logoPath,omitempty"`
	OriginCountry *string `json:"originCountry,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewNetwork instantiates a new Network object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetwork() *Network {
	this := Network{}
	return &this
}

// NewNetworkWithDefaults instantiates a new Network object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkWithDefaults() *Network {
	this := Network{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Network) GetId() float32 {
	if o == nil || IsNil(o.Id) {
		var ret float32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetIdOk() (*float32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Network) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given float32 and assigns it to the Id field.
func (o *Network) SetId(v float32) {
	o.Id = &v
}

// GetLogoPath returns the LogoPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Network) GetLogoPath() string {
	if o == nil || IsNil(o.LogoPath.Get()) {
		var ret string
		return ret
	}
	return *o.LogoPath.Get()
}

// GetLogoPathOk returns a tuple with the LogoPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Network) GetLogoPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LogoPath.Get(), o.LogoPath.IsSet()
}

// HasLogoPath returns a boolean if a field has been set.
func (o *Network) HasLogoPath() bool {
	if o != nil && o.LogoPath.IsSet() {
		return true
	}

	return false
}

// SetLogoPath gets a reference to the given NullableString and assigns it to the LogoPath field.
func (o *Network) SetLogoPath(v string) {
	o.LogoPath.Set(&v)
}
// SetLogoPathNil sets the value for LogoPath to be an explicit nil
func (o *Network) SetLogoPathNil() {
	o.LogoPath.Set(nil)
}

// UnsetLogoPath ensures that no value is present for LogoPath, not even an explicit nil
func (o *Network) UnsetLogoPath() {
	o.LogoPath.Unset()
}

// GetOriginCountry returns the OriginCountry field value if set, zero value otherwise.
func (o *Network) GetOriginCountry() string {
	if o == nil || IsNil(o.OriginCountry) {
		var ret string
		return ret
	}
	return *o.OriginCountry
}

// GetOriginCountryOk returns a tuple with the OriginCountry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetOriginCountryOk() (*string, bool) {
	if o == nil || IsNil(o.OriginCountry) {
		return nil, false
	}
	return o.OriginCountry, true
}

// HasOriginCountry returns a boolean if a field has been set.
func (o *Network) HasOriginCountry() bool {
	if o != nil && !IsNil(o.OriginCountry) {
		return true
	}

	return false
}

// SetOriginCountry gets a reference to the given string and assigns it to the OriginCountry field.
func (o *Network) SetOriginCountry(v string) {
	o.OriginCountry = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Network) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Network) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Network) SetName(v string) {
	o.Name = &v
}

func (o Network) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Network) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.LogoPath.IsSet() {
		toSerialize["logoPath"] = o.LogoPath.Get()
	}
	if !IsNil(o.OriginCountry) {
		toSerialize["originCountry"] = o.OriginCountry
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableNetwork struct {
	value *Network
	isSet bool
}

func (v NullableNetwork) Get() *Network {
	return v.value
}

func (v *NullableNetwork) Set(val *Network) {
	v.value = val
	v.isSet = true
}

func (v NullableNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetwork(val *Network) *NullableNetwork {
	return &NullableNetwork{value: val, isSet: true}
}

func (v NullableNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


