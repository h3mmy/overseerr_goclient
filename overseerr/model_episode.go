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

// checks if the Episode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Episode{}

// Episode struct for Episode
type Episode struct {
	Id *float32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	AirDate NullableString `json:"airDate,omitempty"`
	EpisodeNumber *float32 `json:"episodeNumber,omitempty"`
	Overview *string `json:"overview,omitempty"`
	ProductionCode *string `json:"productionCode,omitempty"`
	SeasonNumber *float32 `json:"seasonNumber,omitempty"`
	ShowId *float32 `json:"showId,omitempty"`
	StillPath NullableString `json:"stillPath,omitempty"`
	VoteAverage *float32 `json:"voteAverage,omitempty"`
	VoteCount *float32 `json:"voteCount,omitempty"`
}

// NewEpisode instantiates a new Episode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEpisode() *Episode {
	this := Episode{}
	return &this
}

// NewEpisodeWithDefaults instantiates a new Episode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEpisodeWithDefaults() *Episode {
	this := Episode{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Episode) GetId() float32 {
	if o == nil || IsNil(o.Id) {
		var ret float32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Episode) GetIdOk() (*float32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Episode) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given float32 and assigns it to the Id field.
func (o *Episode) SetId(v float32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Episode) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Episode) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Episode) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Episode) SetName(v string) {
	o.Name = &v
}

// GetAirDate returns the AirDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Episode) GetAirDate() string {
	if o == nil || IsNil(o.AirDate.Get()) {
		var ret string
		return ret
	}
	return *o.AirDate.Get()
}

// GetAirDateOk returns a tuple with the AirDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Episode) GetAirDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AirDate.Get(), o.AirDate.IsSet()
}

// HasAirDate returns a boolean if a field has been set.
func (o *Episode) HasAirDate() bool {
	if o != nil && o.AirDate.IsSet() {
		return true
	}

	return false
}

// SetAirDate gets a reference to the given NullableString and assigns it to the AirDate field.
func (o *Episode) SetAirDate(v string) {
	o.AirDate.Set(&v)
}
// SetAirDateNil sets the value for AirDate to be an explicit nil
func (o *Episode) SetAirDateNil() {
	o.AirDate.Set(nil)
}

// UnsetAirDate ensures that no value is present for AirDate, not even an explicit nil
func (o *Episode) UnsetAirDate() {
	o.AirDate.Unset()
}

// GetEpisodeNumber returns the EpisodeNumber field value if set, zero value otherwise.
func (o *Episode) GetEpisodeNumber() float32 {
	if o == nil || IsNil(o.EpisodeNumber) {
		var ret float32
		return ret
	}
	return *o.EpisodeNumber
}

// GetEpisodeNumberOk returns a tuple with the EpisodeNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Episode) GetEpisodeNumberOk() (*float32, bool) {
	if o == nil || IsNil(o.EpisodeNumber) {
		return nil, false
	}
	return o.EpisodeNumber, true
}

// HasEpisodeNumber returns a boolean if a field has been set.
func (o *Episode) HasEpisodeNumber() bool {
	if o != nil && !IsNil(o.EpisodeNumber) {
		return true
	}

	return false
}

// SetEpisodeNumber gets a reference to the given float32 and assigns it to the EpisodeNumber field.
func (o *Episode) SetEpisodeNumber(v float32) {
	o.EpisodeNumber = &v
}

// GetOverview returns the Overview field value if set, zero value otherwise.
func (o *Episode) GetOverview() string {
	if o == nil || IsNil(o.Overview) {
		var ret string
		return ret
	}
	return *o.Overview
}

// GetOverviewOk returns a tuple with the Overview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Episode) GetOverviewOk() (*string, bool) {
	if o == nil || IsNil(o.Overview) {
		return nil, false
	}
	return o.Overview, true
}

// HasOverview returns a boolean if a field has been set.
func (o *Episode) HasOverview() bool {
	if o != nil && !IsNil(o.Overview) {
		return true
	}

	return false
}

// SetOverview gets a reference to the given string and assigns it to the Overview field.
func (o *Episode) SetOverview(v string) {
	o.Overview = &v
}

// GetProductionCode returns the ProductionCode field value if set, zero value otherwise.
func (o *Episode) GetProductionCode() string {
	if o == nil || IsNil(o.ProductionCode) {
		var ret string
		return ret
	}
	return *o.ProductionCode
}

// GetProductionCodeOk returns a tuple with the ProductionCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Episode) GetProductionCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ProductionCode) {
		return nil, false
	}
	return o.ProductionCode, true
}

// HasProductionCode returns a boolean if a field has been set.
func (o *Episode) HasProductionCode() bool {
	if o != nil && !IsNil(o.ProductionCode) {
		return true
	}

	return false
}

// SetProductionCode gets a reference to the given string and assigns it to the ProductionCode field.
func (o *Episode) SetProductionCode(v string) {
	o.ProductionCode = &v
}

// GetSeasonNumber returns the SeasonNumber field value if set, zero value otherwise.
func (o *Episode) GetSeasonNumber() float32 {
	if o == nil || IsNil(o.SeasonNumber) {
		var ret float32
		return ret
	}
	return *o.SeasonNumber
}

// GetSeasonNumberOk returns a tuple with the SeasonNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Episode) GetSeasonNumberOk() (*float32, bool) {
	if o == nil || IsNil(o.SeasonNumber) {
		return nil, false
	}
	return o.SeasonNumber, true
}

// HasSeasonNumber returns a boolean if a field has been set.
func (o *Episode) HasSeasonNumber() bool {
	if o != nil && !IsNil(o.SeasonNumber) {
		return true
	}

	return false
}

// SetSeasonNumber gets a reference to the given float32 and assigns it to the SeasonNumber field.
func (o *Episode) SetSeasonNumber(v float32) {
	o.SeasonNumber = &v
}

// GetShowId returns the ShowId field value if set, zero value otherwise.
func (o *Episode) GetShowId() float32 {
	if o == nil || IsNil(o.ShowId) {
		var ret float32
		return ret
	}
	return *o.ShowId
}

// GetShowIdOk returns a tuple with the ShowId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Episode) GetShowIdOk() (*float32, bool) {
	if o == nil || IsNil(o.ShowId) {
		return nil, false
	}
	return o.ShowId, true
}

// HasShowId returns a boolean if a field has been set.
func (o *Episode) HasShowId() bool {
	if o != nil && !IsNil(o.ShowId) {
		return true
	}

	return false
}

// SetShowId gets a reference to the given float32 and assigns it to the ShowId field.
func (o *Episode) SetShowId(v float32) {
	o.ShowId = &v
}

// GetStillPath returns the StillPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Episode) GetStillPath() string {
	if o == nil || IsNil(o.StillPath.Get()) {
		var ret string
		return ret
	}
	return *o.StillPath.Get()
}

// GetStillPathOk returns a tuple with the StillPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Episode) GetStillPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StillPath.Get(), o.StillPath.IsSet()
}

// HasStillPath returns a boolean if a field has been set.
func (o *Episode) HasStillPath() bool {
	if o != nil && o.StillPath.IsSet() {
		return true
	}

	return false
}

// SetStillPath gets a reference to the given NullableString and assigns it to the StillPath field.
func (o *Episode) SetStillPath(v string) {
	o.StillPath.Set(&v)
}
// SetStillPathNil sets the value for StillPath to be an explicit nil
func (o *Episode) SetStillPathNil() {
	o.StillPath.Set(nil)
}

// UnsetStillPath ensures that no value is present for StillPath, not even an explicit nil
func (o *Episode) UnsetStillPath() {
	o.StillPath.Unset()
}

// GetVoteAverage returns the VoteAverage field value if set, zero value otherwise.
func (o *Episode) GetVoteAverage() float32 {
	if o == nil || IsNil(o.VoteAverage) {
		var ret float32
		return ret
	}
	return *o.VoteAverage
}

// GetVoteAverageOk returns a tuple with the VoteAverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Episode) GetVoteAverageOk() (*float32, bool) {
	if o == nil || IsNil(o.VoteAverage) {
		return nil, false
	}
	return o.VoteAverage, true
}

// HasVoteAverage returns a boolean if a field has been set.
func (o *Episode) HasVoteAverage() bool {
	if o != nil && !IsNil(o.VoteAverage) {
		return true
	}

	return false
}

// SetVoteAverage gets a reference to the given float32 and assigns it to the VoteAverage field.
func (o *Episode) SetVoteAverage(v float32) {
	o.VoteAverage = &v
}

// GetVoteCount returns the VoteCount field value if set, zero value otherwise.
func (o *Episode) GetVoteCount() float32 {
	if o == nil || IsNil(o.VoteCount) {
		var ret float32
		return ret
	}
	return *o.VoteCount
}

// GetVoteCountOk returns a tuple with the VoteCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Episode) GetVoteCountOk() (*float32, bool) {
	if o == nil || IsNil(o.VoteCount) {
		return nil, false
	}
	return o.VoteCount, true
}

// HasVoteCount returns a boolean if a field has been set.
func (o *Episode) HasVoteCount() bool {
	if o != nil && !IsNil(o.VoteCount) {
		return true
	}

	return false
}

// SetVoteCount gets a reference to the given float32 and assigns it to the VoteCount field.
func (o *Episode) SetVoteCount(v float32) {
	o.VoteCount = &v
}

func (o Episode) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Episode) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.AirDate.IsSet() {
		toSerialize["airDate"] = o.AirDate.Get()
	}
	if !IsNil(o.EpisodeNumber) {
		toSerialize["episodeNumber"] = o.EpisodeNumber
	}
	if !IsNil(o.Overview) {
		toSerialize["overview"] = o.Overview
	}
	if !IsNil(o.ProductionCode) {
		toSerialize["productionCode"] = o.ProductionCode
	}
	if !IsNil(o.SeasonNumber) {
		toSerialize["seasonNumber"] = o.SeasonNumber
	}
	if !IsNil(o.ShowId) {
		toSerialize["showId"] = o.ShowId
	}
	if o.StillPath.IsSet() {
		toSerialize["stillPath"] = o.StillPath.Get()
	}
	if !IsNil(o.VoteAverage) {
		toSerialize["voteAverage"] = o.VoteAverage
	}
	if !IsNil(o.VoteCount) {
		toSerialize["voteCount"] = o.VoteCount
	}
	return toSerialize, nil
}

type NullableEpisode struct {
	value *Episode
	isSet bool
}

func (v NullableEpisode) Get() *Episode {
	return v.value
}

func (v *NullableEpisode) Set(val *Episode) {
	v.value = val
	v.isSet = true
}

func (v NullableEpisode) IsSet() bool {
	return v.isSet
}

func (v *NullableEpisode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEpisode(val *Episode) *NullableEpisode {
	return &NullableEpisode{value: val, isSet: true}
}

func (v NullableEpisode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEpisode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


