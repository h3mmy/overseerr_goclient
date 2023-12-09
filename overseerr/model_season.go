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

// checks if the Season type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Season{}

// Season struct for Season
type Season struct {
	Id *float32 `json:"id,omitempty"`
	AirDate NullableString `json:"airDate,omitempty"`
	EpisodeCount *float32 `json:"episodeCount,omitempty"`
	Name *string `json:"name,omitempty"`
	Overview *string `json:"overview,omitempty"`
	PosterPath *string `json:"posterPath,omitempty"`
	SeasonNumber *float32 `json:"seasonNumber,omitempty"`
	Episodes []Episode `json:"episodes,omitempty"`
}

// NewSeason instantiates a new Season object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSeason() *Season {
	this := Season{}
	return &this
}

// NewSeasonWithDefaults instantiates a new Season object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSeasonWithDefaults() *Season {
	this := Season{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Season) GetId() float32 {
	if o == nil || IsNil(o.Id) {
		var ret float32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Season) GetIdOk() (*float32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Season) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given float32 and assigns it to the Id field.
func (o *Season) SetId(v float32) {
	o.Id = &v
}

// GetAirDate returns the AirDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Season) GetAirDate() string {
	if o == nil || IsNil(o.AirDate.Get()) {
		var ret string
		return ret
	}
	return *o.AirDate.Get()
}

// GetAirDateOk returns a tuple with the AirDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Season) GetAirDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AirDate.Get(), o.AirDate.IsSet()
}

// HasAirDate returns a boolean if a field has been set.
func (o *Season) HasAirDate() bool {
	if o != nil && o.AirDate.IsSet() {
		return true
	}

	return false
}

// SetAirDate gets a reference to the given NullableString and assigns it to the AirDate field.
func (o *Season) SetAirDate(v string) {
	o.AirDate.Set(&v)
}
// SetAirDateNil sets the value for AirDate to be an explicit nil
func (o *Season) SetAirDateNil() {
	o.AirDate.Set(nil)
}

// UnsetAirDate ensures that no value is present for AirDate, not even an explicit nil
func (o *Season) UnsetAirDate() {
	o.AirDate.Unset()
}

// GetEpisodeCount returns the EpisodeCount field value if set, zero value otherwise.
func (o *Season) GetEpisodeCount() float32 {
	if o == nil || IsNil(o.EpisodeCount) {
		var ret float32
		return ret
	}
	return *o.EpisodeCount
}

// GetEpisodeCountOk returns a tuple with the EpisodeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Season) GetEpisodeCountOk() (*float32, bool) {
	if o == nil || IsNil(o.EpisodeCount) {
		return nil, false
	}
	return o.EpisodeCount, true
}

// HasEpisodeCount returns a boolean if a field has been set.
func (o *Season) HasEpisodeCount() bool {
	if o != nil && !IsNil(o.EpisodeCount) {
		return true
	}

	return false
}

// SetEpisodeCount gets a reference to the given float32 and assigns it to the EpisodeCount field.
func (o *Season) SetEpisodeCount(v float32) {
	o.EpisodeCount = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Season) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Season) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Season) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Season) SetName(v string) {
	o.Name = &v
}

// GetOverview returns the Overview field value if set, zero value otherwise.
func (o *Season) GetOverview() string {
	if o == nil || IsNil(o.Overview) {
		var ret string
		return ret
	}
	return *o.Overview
}

// GetOverviewOk returns a tuple with the Overview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Season) GetOverviewOk() (*string, bool) {
	if o == nil || IsNil(o.Overview) {
		return nil, false
	}
	return o.Overview, true
}

// HasOverview returns a boolean if a field has been set.
func (o *Season) HasOverview() bool {
	if o != nil && !IsNil(o.Overview) {
		return true
	}

	return false
}

// SetOverview gets a reference to the given string and assigns it to the Overview field.
func (o *Season) SetOverview(v string) {
	o.Overview = &v
}

// GetPosterPath returns the PosterPath field value if set, zero value otherwise.
func (o *Season) GetPosterPath() string {
	if o == nil || IsNil(o.PosterPath) {
		var ret string
		return ret
	}
	return *o.PosterPath
}

// GetPosterPathOk returns a tuple with the PosterPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Season) GetPosterPathOk() (*string, bool) {
	if o == nil || IsNil(o.PosterPath) {
		return nil, false
	}
	return o.PosterPath, true
}

// HasPosterPath returns a boolean if a field has been set.
func (o *Season) HasPosterPath() bool {
	if o != nil && !IsNil(o.PosterPath) {
		return true
	}

	return false
}

// SetPosterPath gets a reference to the given string and assigns it to the PosterPath field.
func (o *Season) SetPosterPath(v string) {
	o.PosterPath = &v
}

// GetSeasonNumber returns the SeasonNumber field value if set, zero value otherwise.
func (o *Season) GetSeasonNumber() float32 {
	if o == nil || IsNil(o.SeasonNumber) {
		var ret float32
		return ret
	}
	return *o.SeasonNumber
}

// GetSeasonNumberOk returns a tuple with the SeasonNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Season) GetSeasonNumberOk() (*float32, bool) {
	if o == nil || IsNil(o.SeasonNumber) {
		return nil, false
	}
	return o.SeasonNumber, true
}

// HasSeasonNumber returns a boolean if a field has been set.
func (o *Season) HasSeasonNumber() bool {
	if o != nil && !IsNil(o.SeasonNumber) {
		return true
	}

	return false
}

// SetSeasonNumber gets a reference to the given float32 and assigns it to the SeasonNumber field.
func (o *Season) SetSeasonNumber(v float32) {
	o.SeasonNumber = &v
}

// GetEpisodes returns the Episodes field value if set, zero value otherwise.
func (o *Season) GetEpisodes() []Episode {
	if o == nil || IsNil(o.Episodes) {
		var ret []Episode
		return ret
	}
	return o.Episodes
}

// GetEpisodesOk returns a tuple with the Episodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Season) GetEpisodesOk() ([]Episode, bool) {
	if o == nil || IsNil(o.Episodes) {
		return nil, false
	}
	return o.Episodes, true
}

// HasEpisodes returns a boolean if a field has been set.
func (o *Season) HasEpisodes() bool {
	if o != nil && !IsNil(o.Episodes) {
		return true
	}

	return false
}

// SetEpisodes gets a reference to the given []Episode and assigns it to the Episodes field.
func (o *Season) SetEpisodes(v []Episode) {
	o.Episodes = v
}

func (o Season) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Season) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.AirDate.IsSet() {
		toSerialize["airDate"] = o.AirDate.Get()
	}
	if !IsNil(o.EpisodeCount) {
		toSerialize["episodeCount"] = o.EpisodeCount
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Overview) {
		toSerialize["overview"] = o.Overview
	}
	if !IsNil(o.PosterPath) {
		toSerialize["posterPath"] = o.PosterPath
	}
	if !IsNil(o.SeasonNumber) {
		toSerialize["seasonNumber"] = o.SeasonNumber
	}
	if !IsNil(o.Episodes) {
		toSerialize["episodes"] = o.Episodes
	}
	return toSerialize, nil
}

type NullableSeason struct {
	value *Season
	isSet bool
}

func (v NullableSeason) Get() *Season {
	return v.value
}

func (v *NullableSeason) Set(val *Season) {
	v.value = val
	v.isSet = true
}

func (v NullableSeason) IsSet() bool {
	return v.isSet
}

func (v *NullableSeason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSeason(val *Season) *NullableSeason {
	return &NullableSeason{value: val, isSet: true}
}

func (v NullableSeason) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSeason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


