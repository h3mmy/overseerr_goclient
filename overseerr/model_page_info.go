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

// checks if the PageInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PageInfo{}

// PageInfo struct for PageInfo
type PageInfo struct {
	Page *float32 `json:"page,omitempty"`
	Pages *float32 `json:"pages,omitempty"`
	Results *float32 `json:"results,omitempty"`
}

// NewPageInfo instantiates a new PageInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPageInfo() *PageInfo {
	this := PageInfo{}
	return &this
}

// NewPageInfoWithDefaults instantiates a new PageInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPageInfoWithDefaults() *PageInfo {
	this := PageInfo{}
	return &this
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *PageInfo) GetPage() float32 {
	if o == nil || IsNil(o.Page) {
		var ret float32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageInfo) GetPageOk() (*float32, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *PageInfo) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given float32 and assigns it to the Page field.
func (o *PageInfo) SetPage(v float32) {
	o.Page = &v
}

// GetPages returns the Pages field value if set, zero value otherwise.
func (o *PageInfo) GetPages() float32 {
	if o == nil || IsNil(o.Pages) {
		var ret float32
		return ret
	}
	return *o.Pages
}

// GetPagesOk returns a tuple with the Pages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageInfo) GetPagesOk() (*float32, bool) {
	if o == nil || IsNil(o.Pages) {
		return nil, false
	}
	return o.Pages, true
}

// HasPages returns a boolean if a field has been set.
func (o *PageInfo) HasPages() bool {
	if o != nil && !IsNil(o.Pages) {
		return true
	}

	return false
}

// SetPages gets a reference to the given float32 and assigns it to the Pages field.
func (o *PageInfo) SetPages(v float32) {
	o.Pages = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *PageInfo) GetResults() float32 {
	if o == nil || IsNil(o.Results) {
		var ret float32
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageInfo) GetResultsOk() (*float32, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *PageInfo) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given float32 and assigns it to the Results field.
func (o *PageInfo) SetResults(v float32) {
	o.Results = &v
}

func (o PageInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PageInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	if !IsNil(o.Pages) {
		toSerialize["pages"] = o.Pages
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullablePageInfo struct {
	value *PageInfo
	isSet bool
}

func (v NullablePageInfo) Get() *PageInfo {
	return v.value
}

func (v *NullablePageInfo) Set(val *PageInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePageInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePageInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePageInfo(val *PageInfo) *NullablePageInfo {
	return &NullablePageInfo{value: val, isSet: true}
}

func (v NullablePageInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePageInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


