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

// checks if the TvResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TvResult{}

// TvResult struct for TvResult
type TvResult struct {
	Id *float32 `json:"id,omitempty"`
	MediaType *string `json:"mediaType,omitempty"`
	Popularity *float32 `json:"popularity,omitempty"`
	PosterPath *string `json:"posterPath,omitempty"`
	BackdropPath *string `json:"backdropPath,omitempty"`
	VoteCount *float32 `json:"voteCount,omitempty"`
	VoteAverage *float32 `json:"voteAverage,omitempty"`
	GenreIds []float32 `json:"genreIds,omitempty"`
	Overview *string `json:"overview,omitempty"`
	OriginalLanguage *string `json:"originalLanguage,omitempty"`
	Name *string `json:"name,omitempty"`
	OriginalName *string `json:"originalName,omitempty"`
	OriginCountry []string `json:"originCountry,omitempty"`
	FirstAirDate *string `json:"firstAirDate,omitempty"`
	MediaInfo *MediaInfo `json:"mediaInfo,omitempty"`
}

// NewTvResult instantiates a new TvResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTvResult() *TvResult {
	this := TvResult{}
	return &this
}

// NewTvResultWithDefaults instantiates a new TvResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTvResultWithDefaults() *TvResult {
	this := TvResult{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TvResult) GetId() float32 {
	if o == nil || IsNil(o.Id) {
		var ret float32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvResult) GetIdOk() (*float32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TvResult) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given float32 and assigns it to the Id field.
func (o *TvResult) SetId(v float32) {
	o.Id = &v
}

// GetMediaType returns the MediaType field value if set, zero value otherwise.
func (o *TvResult) GetMediaType() string {
	if o == nil || IsNil(o.MediaType) {
		var ret string
		return ret
	}
	return *o.MediaType
}

// GetMediaTypeOk returns a tuple with the MediaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvResult) GetMediaTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MediaType) {
		return nil, false
	}
	return o.MediaType, true
}

// HasMediaType returns a boolean if a field has been set.
func (o *TvResult) HasMediaType() bool {
	if o != nil && !IsNil(o.MediaType) {
		return true
	}

	return false
}

// SetMediaType gets a reference to the given string and assigns it to the MediaType field.
func (o *TvResult) SetMediaType(v string) {
	o.MediaType = &v
}

// GetPopularity returns the Popularity field value if set, zero value otherwise.
func (o *TvResult) GetPopularity() float32 {
	if o == nil || IsNil(o.Popularity) {
		var ret float32
		return ret
	}
	return *o.Popularity
}

// GetPopularityOk returns a tuple with the Popularity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvResult) GetPopularityOk() (*float32, bool) {
	if o == nil || IsNil(o.Popularity) {
		return nil, false
	}
	return o.Popularity, true
}

// HasPopularity returns a boolean if a field has been set.
func (o *TvResult) HasPopularity() bool {
	if o != nil && !IsNil(o.Popularity) {
		return true
	}

	return false
}

// SetPopularity gets a reference to the given float32 and assigns it to the Popularity field.
func (o *TvResult) SetPopularity(v float32) {
	o.Popularity = &v
}

// GetPosterPath returns the PosterPath field value if set, zero value otherwise.
func (o *TvResult) GetPosterPath() string {
	if o == nil || IsNil(o.PosterPath) {
		var ret string
		return ret
	}
	return *o.PosterPath
}

// GetPosterPathOk returns a tuple with the PosterPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvResult) GetPosterPathOk() (*string, bool) {
	if o == nil || IsNil(o.PosterPath) {
		return nil, false
	}
	return o.PosterPath, true
}

// HasPosterPath returns a boolean if a field has been set.
func (o *TvResult) HasPosterPath() bool {
	if o != nil && !IsNil(o.PosterPath) {
		return true
	}

	return false
}

// SetPosterPath gets a reference to the given string and assigns it to the PosterPath field.
func (o *TvResult) SetPosterPath(v string) {
	o.PosterPath = &v
}

// GetBackdropPath returns the BackdropPath field value if set, zero value otherwise.
func (o *TvResult) GetBackdropPath() string {
	if o == nil || IsNil(o.BackdropPath) {
		var ret string
		return ret
	}
	return *o.BackdropPath
}

// GetBackdropPathOk returns a tuple with the BackdropPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvResult) GetBackdropPathOk() (*string, bool) {
	if o == nil || IsNil(o.BackdropPath) {
		return nil, false
	}
	return o.BackdropPath, true
}

// HasBackdropPath returns a boolean if a field has been set.
func (o *TvResult) HasBackdropPath() bool {
	if o != nil && !IsNil(o.BackdropPath) {
		return true
	}

	return false
}

// SetBackdropPath gets a reference to the given string and assigns it to the BackdropPath field.
func (o *TvResult) SetBackdropPath(v string) {
	o.BackdropPath = &v
}

// GetVoteCount returns the VoteCount field value if set, zero value otherwise.
func (o *TvResult) GetVoteCount() float32 {
	if o == nil || IsNil(o.VoteCount) {
		var ret float32
		return ret
	}
	return *o.VoteCount
}

// GetVoteCountOk returns a tuple with the VoteCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvResult) GetVoteCountOk() (*float32, bool) {
	if o == nil || IsNil(o.VoteCount) {
		return nil, false
	}
	return o.VoteCount, true
}

// HasVoteCount returns a boolean if a field has been set.
func (o *TvResult) HasVoteCount() bool {
	if o != nil && !IsNil(o.VoteCount) {
		return true
	}

	return false
}

// SetVoteCount gets a reference to the given float32 and assigns it to the VoteCount field.
func (o *TvResult) SetVoteCount(v float32) {
	o.VoteCount = &v
}

// GetVoteAverage returns the VoteAverage field value if set, zero value otherwise.
func (o *TvResult) GetVoteAverage() float32 {
	if o == nil || IsNil(o.VoteAverage) {
		var ret float32
		return ret
	}
	return *o.VoteAverage
}

// GetVoteAverageOk returns a tuple with the VoteAverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvResult) GetVoteAverageOk() (*float32, bool) {
	if o == nil || IsNil(o.VoteAverage) {
		return nil, false
	}
	return o.VoteAverage, true
}

// HasVoteAverage returns a boolean if a field has been set.
func (o *TvResult) HasVoteAverage() bool {
	if o != nil && !IsNil(o.VoteAverage) {
		return true
	}

	return false
}

// SetVoteAverage gets a reference to the given float32 and assigns it to the VoteAverage field.
func (o *TvResult) SetVoteAverage(v float32) {
	o.VoteAverage = &v
}

// GetGenreIds returns the GenreIds field value if set, zero value otherwise.
func (o *TvResult) GetGenreIds() []float32 {
	if o == nil || IsNil(o.GenreIds) {
		var ret []float32
		return ret
	}
	return o.GenreIds
}

// GetGenreIdsOk returns a tuple with the GenreIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvResult) GetGenreIdsOk() ([]float32, bool) {
	if o == nil || IsNil(o.GenreIds) {
		return nil, false
	}
	return o.GenreIds, true
}

// HasGenreIds returns a boolean if a field has been set.
func (o *TvResult) HasGenreIds() bool {
	if o != nil && !IsNil(o.GenreIds) {
		return true
	}

	return false
}

// SetGenreIds gets a reference to the given []float32 and assigns it to the GenreIds field.
func (o *TvResult) SetGenreIds(v []float32) {
	o.GenreIds = v
}

// GetOverview returns the Overview field value if set, zero value otherwise.
func (o *TvResult) GetOverview() string {
	if o == nil || IsNil(o.Overview) {
		var ret string
		return ret
	}
	return *o.Overview
}

// GetOverviewOk returns a tuple with the Overview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvResult) GetOverviewOk() (*string, bool) {
	if o == nil || IsNil(o.Overview) {
		return nil, false
	}
	return o.Overview, true
}

// HasOverview returns a boolean if a field has been set.
func (o *TvResult) HasOverview() bool {
	if o != nil && !IsNil(o.Overview) {
		return true
	}

	return false
}

// SetOverview gets a reference to the given string and assigns it to the Overview field.
func (o *TvResult) SetOverview(v string) {
	o.Overview = &v
}

// GetOriginalLanguage returns the OriginalLanguage field value if set, zero value otherwise.
func (o *TvResult) GetOriginalLanguage() string {
	if o == nil || IsNil(o.OriginalLanguage) {
		var ret string
		return ret
	}
	return *o.OriginalLanguage
}

// GetOriginalLanguageOk returns a tuple with the OriginalLanguage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvResult) GetOriginalLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.OriginalLanguage) {
		return nil, false
	}
	return o.OriginalLanguage, true
}

// HasOriginalLanguage returns a boolean if a field has been set.
func (o *TvResult) HasOriginalLanguage() bool {
	if o != nil && !IsNil(o.OriginalLanguage) {
		return true
	}

	return false
}

// SetOriginalLanguage gets a reference to the given string and assigns it to the OriginalLanguage field.
func (o *TvResult) SetOriginalLanguage(v string) {
	o.OriginalLanguage = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TvResult) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvResult) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TvResult) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TvResult) SetName(v string) {
	o.Name = &v
}

// GetOriginalName returns the OriginalName field value if set, zero value otherwise.
func (o *TvResult) GetOriginalName() string {
	if o == nil || IsNil(o.OriginalName) {
		var ret string
		return ret
	}
	return *o.OriginalName
}

// GetOriginalNameOk returns a tuple with the OriginalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvResult) GetOriginalNameOk() (*string, bool) {
	if o == nil || IsNil(o.OriginalName) {
		return nil, false
	}
	return o.OriginalName, true
}

// HasOriginalName returns a boolean if a field has been set.
func (o *TvResult) HasOriginalName() bool {
	if o != nil && !IsNil(o.OriginalName) {
		return true
	}

	return false
}

// SetOriginalName gets a reference to the given string and assigns it to the OriginalName field.
func (o *TvResult) SetOriginalName(v string) {
	o.OriginalName = &v
}

// GetOriginCountry returns the OriginCountry field value if set, zero value otherwise.
func (o *TvResult) GetOriginCountry() []string {
	if o == nil || IsNil(o.OriginCountry) {
		var ret []string
		return ret
	}
	return o.OriginCountry
}

// GetOriginCountryOk returns a tuple with the OriginCountry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvResult) GetOriginCountryOk() ([]string, bool) {
	if o == nil || IsNil(o.OriginCountry) {
		return nil, false
	}
	return o.OriginCountry, true
}

// HasOriginCountry returns a boolean if a field has been set.
func (o *TvResult) HasOriginCountry() bool {
	if o != nil && !IsNil(o.OriginCountry) {
		return true
	}

	return false
}

// SetOriginCountry gets a reference to the given []string and assigns it to the OriginCountry field.
func (o *TvResult) SetOriginCountry(v []string) {
	o.OriginCountry = v
}

// GetFirstAirDate returns the FirstAirDate field value if set, zero value otherwise.
func (o *TvResult) GetFirstAirDate() string {
	if o == nil || IsNil(o.FirstAirDate) {
		var ret string
		return ret
	}
	return *o.FirstAirDate
}

// GetFirstAirDateOk returns a tuple with the FirstAirDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvResult) GetFirstAirDateOk() (*string, bool) {
	if o == nil || IsNil(o.FirstAirDate) {
		return nil, false
	}
	return o.FirstAirDate, true
}

// HasFirstAirDate returns a boolean if a field has been set.
func (o *TvResult) HasFirstAirDate() bool {
	if o != nil && !IsNil(o.FirstAirDate) {
		return true
	}

	return false
}

// SetFirstAirDate gets a reference to the given string and assigns it to the FirstAirDate field.
func (o *TvResult) SetFirstAirDate(v string) {
	o.FirstAirDate = &v
}

// GetMediaInfo returns the MediaInfo field value if set, zero value otherwise.
func (o *TvResult) GetMediaInfo() MediaInfo {
	if o == nil || IsNil(o.MediaInfo) {
		var ret MediaInfo
		return ret
	}
	return *o.MediaInfo
}

// GetMediaInfoOk returns a tuple with the MediaInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvResult) GetMediaInfoOk() (*MediaInfo, bool) {
	if o == nil || IsNil(o.MediaInfo) {
		return nil, false
	}
	return o.MediaInfo, true
}

// HasMediaInfo returns a boolean if a field has been set.
func (o *TvResult) HasMediaInfo() bool {
	if o != nil && !IsNil(o.MediaInfo) {
		return true
	}

	return false
}

// SetMediaInfo gets a reference to the given MediaInfo and assigns it to the MediaInfo field.
func (o *TvResult) SetMediaInfo(v MediaInfo) {
	o.MediaInfo = &v
}

func (o TvResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TvResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.MediaType) {
		toSerialize["mediaType"] = o.MediaType
	}
	if !IsNil(o.Popularity) {
		toSerialize["popularity"] = o.Popularity
	}
	if !IsNil(o.PosterPath) {
		toSerialize["posterPath"] = o.PosterPath
	}
	if !IsNil(o.BackdropPath) {
		toSerialize["backdropPath"] = o.BackdropPath
	}
	if !IsNil(o.VoteCount) {
		toSerialize["voteCount"] = o.VoteCount
	}
	if !IsNil(o.VoteAverage) {
		toSerialize["voteAverage"] = o.VoteAverage
	}
	if !IsNil(o.GenreIds) {
		toSerialize["genreIds"] = o.GenreIds
	}
	if !IsNil(o.Overview) {
		toSerialize["overview"] = o.Overview
	}
	if !IsNil(o.OriginalLanguage) {
		toSerialize["originalLanguage"] = o.OriginalLanguage
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.OriginalName) {
		toSerialize["originalName"] = o.OriginalName
	}
	if !IsNil(o.OriginCountry) {
		toSerialize["originCountry"] = o.OriginCountry
	}
	if !IsNil(o.FirstAirDate) {
		toSerialize["firstAirDate"] = o.FirstAirDate
	}
	if !IsNil(o.MediaInfo) {
		toSerialize["mediaInfo"] = o.MediaInfo
	}
	return toSerialize, nil
}

type NullableTvResult struct {
	value *TvResult
	isSet bool
}

func (v NullableTvResult) Get() *TvResult {
	return v.value
}

func (v *NullableTvResult) Set(val *TvResult) {
	v.value = val
	v.isSet = true
}

func (v NullableTvResult) IsSet() bool {
	return v.isSet
}

func (v *NullableTvResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTvResult(val *TvResult) *NullableTvResult {
	return &NullableTvResult{value: val, isSet: true}
}

func (v NullableTvResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTvResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


