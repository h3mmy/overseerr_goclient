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

// checks if the IssueGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IssueGet200Response{}

// IssueGet200Response struct for IssueGet200Response
type IssueGet200Response struct {
	PageInfo *PageInfo `json:"pageInfo,omitempty"`
	Results []Issue `json:"results,omitempty"`
}

// NewIssueGet200Response instantiates a new IssueGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssueGet200Response() *IssueGet200Response {
	this := IssueGet200Response{}
	return &this
}

// NewIssueGet200ResponseWithDefaults instantiates a new IssueGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssueGet200ResponseWithDefaults() *IssueGet200Response {
	this := IssueGet200Response{}
	return &this
}

// GetPageInfo returns the PageInfo field value if set, zero value otherwise.
func (o *IssueGet200Response) GetPageInfo() PageInfo {
	if o == nil || IsNil(o.PageInfo) {
		var ret PageInfo
		return ret
	}
	return *o.PageInfo
}

// GetPageInfoOk returns a tuple with the PageInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueGet200Response) GetPageInfoOk() (*PageInfo, bool) {
	if o == nil || IsNil(o.PageInfo) {
		return nil, false
	}
	return o.PageInfo, true
}

// HasPageInfo returns a boolean if a field has been set.
func (o *IssueGet200Response) HasPageInfo() bool {
	if o != nil && !IsNil(o.PageInfo) {
		return true
	}

	return false
}

// SetPageInfo gets a reference to the given PageInfo and assigns it to the PageInfo field.
func (o *IssueGet200Response) SetPageInfo(v PageInfo) {
	o.PageInfo = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *IssueGet200Response) GetResults() []Issue {
	if o == nil || IsNil(o.Results) {
		var ret []Issue
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueGet200Response) GetResultsOk() ([]Issue, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *IssueGet200Response) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []Issue and assigns it to the Results field.
func (o *IssueGet200Response) SetResults(v []Issue) {
	o.Results = v
}

func (o IssueGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IssueGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PageInfo) {
		toSerialize["pageInfo"] = o.PageInfo
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullableIssueGet200Response struct {
	value *IssueGet200Response
	isSet bool
}

func (v NullableIssueGet200Response) Get() *IssueGet200Response {
	return v.value
}

func (v *NullableIssueGet200Response) Set(val *IssueGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableIssueGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableIssueGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssueGet200Response(val *IssueGet200Response) *NullableIssueGet200Response {
	return &NullableIssueGet200Response{value: val, isSet: true}
}

func (v NullableIssueGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssueGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


