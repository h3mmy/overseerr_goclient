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

// checks if the TvDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TvDetails{}

// TvDetails struct for TvDetails
type TvDetails struct {
	Id *float32 `json:"id,omitempty"`
	BackdropPath *string `json:"backdropPath,omitempty"`
	PosterPath *string `json:"posterPath,omitempty"`
	ContentRatings *TvDetailsContentRatings `json:"contentRatings,omitempty"`
	CreatedBy []TvDetailsCreatedByInner `json:"createdBy,omitempty"`
	EpisodeRunTime []float32 `json:"episodeRunTime,omitempty"`
	FirstAirDate *string `json:"firstAirDate,omitempty"`
	Genres []Genre `json:"genres,omitempty"`
	Homepage *string `json:"homepage,omitempty"`
	InProduction *bool `json:"inProduction,omitempty"`
	Languages []string `json:"languages,omitempty"`
	LastAirDate *string `json:"lastAirDate,omitempty"`
	LastEpisodeToAir *Episode `json:"lastEpisodeToAir,omitempty"`
	Name *string `json:"name,omitempty"`
	NextEpisodeToAir *Episode `json:"nextEpisodeToAir,omitempty"`
	Networks []ProductionCompany `json:"networks,omitempty"`
	NumberOfEpisodes *float32 `json:"numberOfEpisodes,omitempty"`
	NumberOfSeason *float32 `json:"numberOfSeason,omitempty"`
	OriginCountry []string `json:"originCountry,omitempty"`
	OriginalLanguage *string `json:"originalLanguage,omitempty"`
	OriginalName *string `json:"originalName,omitempty"`
	Overview *string `json:"overview,omitempty"`
	Popularity *float32 `json:"popularity,omitempty"`
	ProductionCompanies []ProductionCompany `json:"productionCompanies,omitempty"`
	ProductionCountries []MovieDetailsProductionCountriesInner `json:"productionCountries,omitempty"`
	SpokenLanguages []SpokenLanguage `json:"spokenLanguages,omitempty"`
	Seasons []Season `json:"seasons,omitempty"`
	Status *string `json:"status,omitempty"`
	Tagline *string `json:"tagline,omitempty"`
	Type *string `json:"type,omitempty"`
	VoteAverage *float32 `json:"voteAverage,omitempty"`
	VoteCount *float32 `json:"voteCount,omitempty"`
	Credits *MovieDetailsCredits `json:"credits,omitempty"`
	ExternalIds *ExternalIds `json:"externalIds,omitempty"`
	Keywords []Keyword `json:"keywords,omitempty"`
	MediaInfo *MediaInfo `json:"mediaInfo,omitempty"`
	WatchProviders [][]WatchProvidersInner `json:"watchProviders,omitempty"`
}

// NewTvDetails instantiates a new TvDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTvDetails() *TvDetails {
	this := TvDetails{}
	return &this
}

// NewTvDetailsWithDefaults instantiates a new TvDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTvDetailsWithDefaults() *TvDetails {
	this := TvDetails{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TvDetails) GetId() float32 {
	if o == nil || IsNil(o.Id) {
		var ret float32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvDetails) GetIdOk() (*float32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TvDetails) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given float32 and assigns it to the Id field.
func (o *TvDetails) SetId(v float32) {
	o.Id = &v
}

// GetBackdropPath returns the BackdropPath field value if set, zero value otherwise.
func (o *TvDetails) GetBackdropPath() string {
	if o == nil || IsNil(o.BackdropPath) {
		var ret string
		return ret
	}
	return *o.BackdropPath
}

// GetBackdropPathOk returns a tuple with the BackdropPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvDetails) GetBackdropPathOk() (*string, bool) {
	if o == nil || IsNil(o.BackdropPath) {
		return nil, false
	}
	return o.BackdropPath, true
}

// HasBackdropPath returns a boolean if a field has been set.
func (o *TvDetails) HasBackdropPath() bool {
	if o != nil && !IsNil(o.BackdropPath) {
		return true
	}

	return false
}

// SetBackdropPath gets a reference to the given string and assigns it to the BackdropPath field.
func (o *TvDetails) SetBackdropPath(v string) {
	o.BackdropPath = &v
}

// GetPosterPath returns the PosterPath field value if set, zero value otherwise.
func (o *TvDetails) GetPosterPath() string {
	if o == nil || IsNil(o.PosterPath) {
		var ret string
		return ret
	}
	return *o.PosterPath
}

// GetPosterPathOk returns a tuple with the PosterPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvDetails) GetPosterPathOk() (*string, bool) {
	if o == nil || IsNil(o.PosterPath) {
		return nil, false
	}
	return o.PosterPath, true
}

// HasPosterPath returns a boolean if a field has been set.
func (o *TvDetails) HasPosterPath() bool {
	if o != nil && !IsNil(o.PosterPath) {
		return true
	}

	return false
}

// SetPosterPath gets a reference to the given string and assigns it to the PosterPath field.
func (o *TvDetails) SetPosterPath(v string) {
	o.PosterPath = &v
}

// GetContentRatings returns the ContentRatings field value if set, zero value otherwise.
func (o *TvDetails) GetContentRatings() TvDetailsContentRatings {
	if o == nil || IsNil(o.ContentRatings) {
		var ret TvDetailsContentRatings
		return ret
	}
	return *o.ContentRatings
}

// GetContentRatingsOk returns a tuple with the ContentRatings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvDetails) GetContentRatingsOk() (*TvDetailsContentRatings, bool) {
	if o == nil || IsNil(o.ContentRatings) {
		return nil, false
	}
	return o.ContentRatings, true
}

// HasContentRatings returns a boolean if a field has been set.
func (o *TvDetails) HasContentRatings() bool {
	if o != nil && !IsNil(o.ContentRatings) {
		return true
	}

	return false
}

// SetContentRatings gets a reference to the given TvDetailsContentRatings and assigns it to the ContentRatings field.
func (o *TvDetails) SetContentRatings(v TvDetailsContentRatings) {
	o.ContentRatings = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *TvDetails) GetCreatedBy() []TvDetailsCreatedByInner {
	if o == nil || IsNil(o.CreatedBy) {
		var ret []TvDetailsCreatedByInner
		return ret
	}
	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvDetails) GetCreatedByOk() ([]TvDetailsCreatedByInner, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *TvDetails) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given []TvDetailsCreatedByInner and assigns it to the CreatedBy field.
func (o *TvDetails) SetCreatedBy(v []TvDetailsCreatedByInner) {
	o.CreatedBy = v
}

// GetEpisodeRunTime returns the EpisodeRunTime field value if set, zero value otherwise.
func (o *TvDetails) GetEpisodeRunTime() []float32 {
	if o == nil || IsNil(o.EpisodeRunTime) {
		var ret []float32
		return ret
	}
	return o.EpisodeRunTime
}

// GetEpisodeRunTimeOk returns a tuple with the EpisodeRunTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvDetails) GetEpisodeRunTimeOk() ([]float32, bool) {
	if o == nil || IsNil(o.EpisodeRunTime) {
		return nil, false
	}
	return o.EpisodeRunTime, true
}

// HasEpisodeRunTime returns a boolean if a field has been set.
func (o *TvDetails) HasEpisodeRunTime() bool {
	if o != nil && !IsNil(o.EpisodeRunTime) {
		return true
	}

	return false
}

// SetEpisodeRunTime gets a reference to the given []float32 and assigns it to the EpisodeRunTime field.
func (o *TvDetails) SetEpisodeRunTime(v []float32) {
	o.EpisodeRunTime = v
}

// GetFirstAirDate returns the FirstAirDate field value if set, zero value otherwise.
func (o *TvDetails) GetFirstAirDate() string {
	if o == nil || IsNil(o.FirstAirDate) {
		var ret string
		return ret
	}
	return *o.FirstAirDate
}

// GetFirstAirDateOk returns a tuple with the FirstAirDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvDetails) GetFirstAirDateOk() (*string, bool) {
	if o == nil || IsNil(o.FirstAirDate) {
		return nil, false
	}
	return o.FirstAirDate, true
}

// HasFirstAirDate returns a boolean if a field has been set.
func (o *TvDetails) HasFirstAirDate() bool {
	if o != nil && !IsNil(o.FirstAirDate) {
		return true
	}

	return false
}

// SetFirstAirDate gets a reference to the given string and assigns it to the FirstAirDate field.
func (o *TvDetails) SetFirstAirDate(v string) {
	o.FirstAirDate = &v
}

// GetGenres returns the Genres field value if set, zero value otherwise.
func (o *TvDetails) GetGenres() []Genre {
	if o == nil || IsNil(o.Genres) {
		var ret []Genre
		return ret
	}
	return o.Genres
}

// GetGenresOk returns a tuple with the Genres field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvDetails) GetGenresOk() ([]Genre, bool) {
	if o == nil || IsNil(o.Genres) {
		return nil, false
	}
	return o.Genres, true
}

// HasGenres returns a boolean if a field has been set.
func (o *TvDetails) HasGenres() bool {
	if o != nil && !IsNil(o.Genres) {
		return true
	}

	return false
}

// SetGenres gets a reference to the given []Genre and assigns it to the Genres field.
func (o *TvDetails) SetGenres(v []Genre) {
	o.Genres = v
}

// GetHomepage returns the Homepage field value if set, zero value otherwise.
func (o *TvDetails) GetHomepage() string {
	if o == nil || IsNil(o.Homepage) {
		var ret string
		return ret
	}
	return *o.Homepage
}

// GetHomepageOk returns a tuple with the Homepage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvDetails) GetHomepageOk() (*string, bool) {
	if o == nil || IsNil(o.Homepage) {
		return nil, false
	}
	return o.Homepage, true
}

// HasHomepage returns a boolean if a field has been set.
func (o *TvDetails) HasHomepage() bool {
	if o != nil && !IsNil(o.Homepage) {
		return true
	}

	return false
}

// SetHomepage gets a reference to the given string and assigns it to the Homepage field.
func (o *TvDetails) SetHomepage(v string) {
	o.Homepage = &v
}

// GetInProduction returns the InProduction field value if set, zero value otherwise.
func (o *TvDetails) GetInProduction() bool {
	if o == nil || IsNil(o.InProduction) {
		var ret bool
		return ret
	}
	return *o.InProduction
}

// GetInProductionOk returns a tuple with the InProduction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvDetails) GetInProductionOk() (*bool, bool) {
	if o == nil || IsNil(o.InProduction) {
		return nil, false
	}
	return o.InProduction, true
}

// HasInProduction returns a boolean if a field has been set.
func (o *TvDetails) HasInProduction() bool {
	if o != nil && !IsNil(o.InProduction) {
		return true
	}

	return false
}

// SetInProduction gets a reference to the given bool and assigns it to the InProduction field.
func (o *TvDetails) SetInProduction(v bool) {
	o.InProduction = &v
}

// GetLanguages returns the Languages field value if set, zero value otherwise.
func (o *TvDetails) GetLanguages() []string {
	if o == nil || IsNil(o.Languages) {
		var ret []string
		return ret
	}
	return o.Languages
}

// GetLanguagesOk returns a tuple with the Languages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvDetails) GetLanguagesOk() ([]string, bool) {
	if o == nil || IsNil(o.Languages) {
		return nil, false
	}
	return o.Languages, true
}

// HasLanguages returns a boolean if a field has been set.
func (o *TvDetails) HasLanguages() bool {
	if o != nil && !IsNil(o.Languages) {
		return true
	}

	return false
}

// SetLanguages gets a reference to the given []string and assigns it to the Languages field.
func (o *TvDetails) SetLanguages(v []string) {
	o.Languages = v
}

// GetLastAirDate returns the LastAirDate field value if set, zero value otherwise.
func (o *TvDetails) GetLastAirDate() string {
	if o == nil || IsNil(o.LastAirDate) {
		var ret string
		return ret
	}
	return *o.LastAirDate
}

// GetLastAirDateOk returns a tuple with the LastAirDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvDetails) GetLastAirDateOk() (*string, bool) {
	if o == nil || IsNil(o.LastAirDate) {
		return nil, false
	}
	return o.LastAirDate, true
}

// HasLastAirDate returns a boolean if a field has been set.
func (o *TvDetails) HasLastAirDate() bool {
	if o != nil && !IsNil(o.LastAirDate) {
		return true
	}

	return false
}

// SetLastAirDate gets a reference to the given string and assigns it to the LastAirDate field.
func (o *TvDetails) SetLastAirDate(v string) {
	o.LastAirDate = &v
}

// GetLastEpisodeToAir returns the LastEpisodeToAir field value if set, zero value otherwise.
func (o *TvDetails) GetLastEpisodeToAir() Episode {
	if o == nil || IsNil(o.LastEpisodeToAir) {
		var ret Episode
		return ret
	}
	return *o.LastEpisodeToAir
}

// GetLastEpisodeToAirOk returns a tuple with the LastEpisodeToAir field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvDetails) GetLastEpisodeToAirOk() (*Episode, bool) {
	if o == nil || IsNil(o.LastEpisodeToAir) {
		return nil, false
	}
	return o.LastEpisodeToAir, true
}

// HasLastEpisodeToAir returns a boolean if a field has been set.
func (o *TvDetails) HasLastEpisodeToAir() bool {
	if o != nil && !IsNil(o.LastEpisodeToAir) {
		return true
	}

	return false
}

// SetLastEpisodeToAir gets a reference to the given Episode and assigns it to the LastEpisodeToAir field.
func (o *TvDetails) SetLastEpisodeToAir(v Episode) {
	o.LastEpisodeToAir = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TvDetails) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvDetails) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TvDetails) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TvDetails) SetName(v string) {
	o.Name = &v
}

// GetNextEpisodeToAir returns the NextEpisodeToAir field value if set, zero value otherwise.
func (o *TvDetails) GetNextEpisodeToAir() Episode {
	if o == nil || IsNil(o.NextEpisodeToAir) {
		var ret Episode
		return ret
	}
	return *o.NextEpisodeToAir
}

// GetNextEpisodeToAirOk returns a tuple with the NextEpisodeToAir field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvDetails) GetNextEpisodeToAirOk() (*Episode, bool) {
	if o == nil || IsNil(o.NextEpisodeToAir) {
		return nil, false
	}
	return o.NextEpisodeToAir, true
}

// HasNextEpisodeToAir returns a boolean if a field has been set.
func (o *TvDetails) HasNextEpisodeToAir() bool {
	if o != nil && !IsNil(o.NextEpisodeToAir) {
		return true
	}

	return false
}

// SetNextEpisodeToAir gets a reference to the given Episode and assigns it to the NextEpisodeToAir field.
func (o *TvDetails) SetNextEpisodeToAir(v Episode) {
	o.NextEpisodeToAir = &v
}

// GetNetworks returns the Networks field value if set, zero value otherwise.
func (o *TvDetails) GetNetworks() []ProductionCompany {
	if o == nil || IsNil(o.Networks) {
		var ret []ProductionCompany
		return ret
	}
	return o.Networks
}

// GetNetworksOk returns a tuple with the Networks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvDetails) GetNetworksOk() ([]ProductionCompany, bool) {
	if o == nil || IsNil(o.Networks) {
		return nil, false
	}
	return o.Networks, true
}

// HasNetworks returns a boolean if a field has been set.
func (o *TvDetails) HasNetworks() bool {
	if o != nil && !IsNil(o.Networks) {
		return true
	}

	return false
}

// SetNetworks gets a reference to the given []ProductionCompany and assigns it to the Networks field.
func (o *TvDetails) SetNetworks(v []ProductionCompany) {
	o.Networks = v
}

// GetNumberOfEpisodes returns the NumberOfEpisodes field value if set, zero value otherwise.
func (o *TvDetails) GetNumberOfEpisodes() float32 {
	if o == nil || IsNil(o.NumberOfEpisodes) {
		var ret float32
		return ret
	}
	return *o.NumberOfEpisodes
}

// GetNumberOfEpisodesOk returns a tuple with the NumberOfEpisodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvDetails) GetNumberOfEpisodesOk() (*float32, bool) {
	if o == nil || IsNil(o.NumberOfEpisodes) {
		return nil, false
	}
	return o.NumberOfEpisodes, true
}

// HasNumberOfEpisodes returns a boolean if a field has been set.
func (o *TvDetails) HasNumberOfEpisodes() bool {
	if o != nil && !IsNil(o.NumberOfEpisodes) {
		return true
	}

	return false
}

// SetNumberOfEpisodes gets a reference to the given float32 and assigns it to the NumberOfEpisodes field.
func (o *TvDetails) SetNumberOfEpisodes(v float32) {
	o.NumberOfEpisodes = &v
}

// GetNumberOfSeason returns the NumberOfSeason field value if set, zero value otherwise.
func (o *TvDetails) GetNumberOfSeason() float32 {
	if o == nil || IsNil(o.NumberOfSeason) {
		var ret float32
		return ret
	}
	return *o.NumberOfSeason
}

// GetNumberOfSeasonOk returns a tuple with the NumberOfSeason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvDetails) GetNumberOfSeasonOk() (*float32, bool) {
	if o == nil || IsNil(o.NumberOfSeason) {
		return nil, false
	}
	return o.NumberOfSeason, true
}

// HasNumberOfSeason returns a boolean if a field has been set.
func (o *TvDetails) HasNumberOfSeason() bool {
	if o != nil && !IsNil(o.NumberOfSeason) {
		return true
	}

	return false
}

// SetNumberOfSeason gets a reference to the given float32 and assigns it to the NumberOfSeason field.
func (o *TvDetails) SetNumberOfSeason(v float32) {
	o.NumberOfSeason = &v
}

// GetOriginCountry returns the OriginCountry field value if set, zero value otherwise.
func (o *TvDetails) GetOriginCountry() []string {
	if o == nil || IsNil(o.OriginCountry) {
		var ret []string
		return ret
	}
	return o.OriginCountry
}

// GetOriginCountryOk returns a tuple with the OriginCountry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvDetails) GetOriginCountryOk() ([]string, bool) {
	if o == nil || IsNil(o.OriginCountry) {
		return nil, false
	}
	return o.OriginCountry, true
}

// HasOriginCountry returns a boolean if a field has been set.
func (o *TvDetails) HasOriginCountry() bool {
	if o != nil && !IsNil(o.OriginCountry) {
		return true
	}

	return false
}

// SetOriginCountry gets a reference to the given []string and assigns it to the OriginCountry field.
func (o *TvDetails) SetOriginCountry(v []string) {
	o.OriginCountry = v
}

// GetOriginalLanguage returns the OriginalLanguage field value if set, zero value otherwise.
func (o *TvDetails) GetOriginalLanguage() string {
	if o == nil || IsNil(o.OriginalLanguage) {
		var ret string
		return ret
	}
	return *o.OriginalLanguage
}

// GetOriginalLanguageOk returns a tuple with the OriginalLanguage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvDetails) GetOriginalLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.OriginalLanguage) {
		return nil, false
	}
	return o.OriginalLanguage, true
}

// HasOriginalLanguage returns a boolean if a field has been set.
func (o *TvDetails) HasOriginalLanguage() bool {
	if o != nil && !IsNil(o.OriginalLanguage) {
		return true
	}

	return false
}

// SetOriginalLanguage gets a reference to the given string and assigns it to the OriginalLanguage field.
func (o *TvDetails) SetOriginalLanguage(v string) {
	o.OriginalLanguage = &v
}

// GetOriginalName returns the OriginalName field value if set, zero value otherwise.
func (o *TvDetails) GetOriginalName() string {
	if o == nil || IsNil(o.OriginalName) {
		var ret string
		return ret
	}
	return *o.OriginalName
}

// GetOriginalNameOk returns a tuple with the OriginalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvDetails) GetOriginalNameOk() (*string, bool) {
	if o == nil || IsNil(o.OriginalName) {
		return nil, false
	}
	return o.OriginalName, true
}

// HasOriginalName returns a boolean if a field has been set.
func (o *TvDetails) HasOriginalName() bool {
	if o != nil && !IsNil(o.OriginalName) {
		return true
	}

	return false
}

// SetOriginalName gets a reference to the given string and assigns it to the OriginalName field.
func (o *TvDetails) SetOriginalName(v string) {
	o.OriginalName = &v
}

// GetOverview returns the Overview field value if set, zero value otherwise.
func (o *TvDetails) GetOverview() string {
	if o == nil || IsNil(o.Overview) {
		var ret string
		return ret
	}
	return *o.Overview
}

// GetOverviewOk returns a tuple with the Overview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvDetails) GetOverviewOk() (*string, bool) {
	if o == nil || IsNil(o.Overview) {
		return nil, false
	}
	return o.Overview, true
}

// HasOverview returns a boolean if a field has been set.
func (o *TvDetails) HasOverview() bool {
	if o != nil && !IsNil(o.Overview) {
		return true
	}

	return false
}

// SetOverview gets a reference to the given string and assigns it to the Overview field.
func (o *TvDetails) SetOverview(v string) {
	o.Overview = &v
}

// GetPopularity returns the Popularity field value if set, zero value otherwise.
func (o *TvDetails) GetPopularity() float32 {
	if o == nil || IsNil(o.Popularity) {
		var ret float32
		return ret
	}
	return *o.Popularity
}

// GetPopularityOk returns a tuple with the Popularity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvDetails) GetPopularityOk() (*float32, bool) {
	if o == nil || IsNil(o.Popularity) {
		return nil, false
	}
	return o.Popularity, true
}

// HasPopularity returns a boolean if a field has been set.
func (o *TvDetails) HasPopularity() bool {
	if o != nil && !IsNil(o.Popularity) {
		return true
	}

	return false
}

// SetPopularity gets a reference to the given float32 and assigns it to the Popularity field.
func (o *TvDetails) SetPopularity(v float32) {
	o.Popularity = &v
}

// GetProductionCompanies returns the ProductionCompanies field value if set, zero value otherwise.
func (o *TvDetails) GetProductionCompanies() []ProductionCompany {
	if o == nil || IsNil(o.ProductionCompanies) {
		var ret []ProductionCompany
		return ret
	}
	return o.ProductionCompanies
}

// GetProductionCompaniesOk returns a tuple with the ProductionCompanies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvDetails) GetProductionCompaniesOk() ([]ProductionCompany, bool) {
	if o == nil || IsNil(o.ProductionCompanies) {
		return nil, false
	}
	return o.ProductionCompanies, true
}

// HasProductionCompanies returns a boolean if a field has been set.
func (o *TvDetails) HasProductionCompanies() bool {
	if o != nil && !IsNil(o.ProductionCompanies) {
		return true
	}

	return false
}

// SetProductionCompanies gets a reference to the given []ProductionCompany and assigns it to the ProductionCompanies field.
func (o *TvDetails) SetProductionCompanies(v []ProductionCompany) {
	o.ProductionCompanies = v
}

// GetProductionCountries returns the ProductionCountries field value if set, zero value otherwise.
func (o *TvDetails) GetProductionCountries() []MovieDetailsProductionCountriesInner {
	if o == nil || IsNil(o.ProductionCountries) {
		var ret []MovieDetailsProductionCountriesInner
		return ret
	}
	return o.ProductionCountries
}

// GetProductionCountriesOk returns a tuple with the ProductionCountries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvDetails) GetProductionCountriesOk() ([]MovieDetailsProductionCountriesInner, bool) {
	if o == nil || IsNil(o.ProductionCountries) {
		return nil, false
	}
	return o.ProductionCountries, true
}

// HasProductionCountries returns a boolean if a field has been set.
func (o *TvDetails) HasProductionCountries() bool {
	if o != nil && !IsNil(o.ProductionCountries) {
		return true
	}

	return false
}

// SetProductionCountries gets a reference to the given []MovieDetailsProductionCountriesInner and assigns it to the ProductionCountries field.
func (o *TvDetails) SetProductionCountries(v []MovieDetailsProductionCountriesInner) {
	o.ProductionCountries = v
}

// GetSpokenLanguages returns the SpokenLanguages field value if set, zero value otherwise.
func (o *TvDetails) GetSpokenLanguages() []SpokenLanguage {
	if o == nil || IsNil(o.SpokenLanguages) {
		var ret []SpokenLanguage
		return ret
	}
	return o.SpokenLanguages
}

// GetSpokenLanguagesOk returns a tuple with the SpokenLanguages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvDetails) GetSpokenLanguagesOk() ([]SpokenLanguage, bool) {
	if o == nil || IsNil(o.SpokenLanguages) {
		return nil, false
	}
	return o.SpokenLanguages, true
}

// HasSpokenLanguages returns a boolean if a field has been set.
func (o *TvDetails) HasSpokenLanguages() bool {
	if o != nil && !IsNil(o.SpokenLanguages) {
		return true
	}

	return false
}

// SetSpokenLanguages gets a reference to the given []SpokenLanguage and assigns it to the SpokenLanguages field.
func (o *TvDetails) SetSpokenLanguages(v []SpokenLanguage) {
	o.SpokenLanguages = v
}

// GetSeasons returns the Seasons field value if set, zero value otherwise.
func (o *TvDetails) GetSeasons() []Season {
	if o == nil || IsNil(o.Seasons) {
		var ret []Season
		return ret
	}
	return o.Seasons
}

// GetSeasonsOk returns a tuple with the Seasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvDetails) GetSeasonsOk() ([]Season, bool) {
	if o == nil || IsNil(o.Seasons) {
		return nil, false
	}
	return o.Seasons, true
}

// HasSeasons returns a boolean if a field has been set.
func (o *TvDetails) HasSeasons() bool {
	if o != nil && !IsNil(o.Seasons) {
		return true
	}

	return false
}

// SetSeasons gets a reference to the given []Season and assigns it to the Seasons field.
func (o *TvDetails) SetSeasons(v []Season) {
	o.Seasons = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *TvDetails) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvDetails) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *TvDetails) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *TvDetails) SetStatus(v string) {
	o.Status = &v
}

// GetTagline returns the Tagline field value if set, zero value otherwise.
func (o *TvDetails) GetTagline() string {
	if o == nil || IsNil(o.Tagline) {
		var ret string
		return ret
	}
	return *o.Tagline
}

// GetTaglineOk returns a tuple with the Tagline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvDetails) GetTaglineOk() (*string, bool) {
	if o == nil || IsNil(o.Tagline) {
		return nil, false
	}
	return o.Tagline, true
}

// HasTagline returns a boolean if a field has been set.
func (o *TvDetails) HasTagline() bool {
	if o != nil && !IsNil(o.Tagline) {
		return true
	}

	return false
}

// SetTagline gets a reference to the given string and assigns it to the Tagline field.
func (o *TvDetails) SetTagline(v string) {
	o.Tagline = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TvDetails) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvDetails) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TvDetails) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TvDetails) SetType(v string) {
	o.Type = &v
}

// GetVoteAverage returns the VoteAverage field value if set, zero value otherwise.
func (o *TvDetails) GetVoteAverage() float32 {
	if o == nil || IsNil(o.VoteAverage) {
		var ret float32
		return ret
	}
	return *o.VoteAverage
}

// GetVoteAverageOk returns a tuple with the VoteAverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvDetails) GetVoteAverageOk() (*float32, bool) {
	if o == nil || IsNil(o.VoteAverage) {
		return nil, false
	}
	return o.VoteAverage, true
}

// HasVoteAverage returns a boolean if a field has been set.
func (o *TvDetails) HasVoteAverage() bool {
	if o != nil && !IsNil(o.VoteAverage) {
		return true
	}

	return false
}

// SetVoteAverage gets a reference to the given float32 and assigns it to the VoteAverage field.
func (o *TvDetails) SetVoteAverage(v float32) {
	o.VoteAverage = &v
}

// GetVoteCount returns the VoteCount field value if set, zero value otherwise.
func (o *TvDetails) GetVoteCount() float32 {
	if o == nil || IsNil(o.VoteCount) {
		var ret float32
		return ret
	}
	return *o.VoteCount
}

// GetVoteCountOk returns a tuple with the VoteCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvDetails) GetVoteCountOk() (*float32, bool) {
	if o == nil || IsNil(o.VoteCount) {
		return nil, false
	}
	return o.VoteCount, true
}

// HasVoteCount returns a boolean if a field has been set.
func (o *TvDetails) HasVoteCount() bool {
	if o != nil && !IsNil(o.VoteCount) {
		return true
	}

	return false
}

// SetVoteCount gets a reference to the given float32 and assigns it to the VoteCount field.
func (o *TvDetails) SetVoteCount(v float32) {
	o.VoteCount = &v
}

// GetCredits returns the Credits field value if set, zero value otherwise.
func (o *TvDetails) GetCredits() MovieDetailsCredits {
	if o == nil || IsNil(o.Credits) {
		var ret MovieDetailsCredits
		return ret
	}
	return *o.Credits
}

// GetCreditsOk returns a tuple with the Credits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvDetails) GetCreditsOk() (*MovieDetailsCredits, bool) {
	if o == nil || IsNil(o.Credits) {
		return nil, false
	}
	return o.Credits, true
}

// HasCredits returns a boolean if a field has been set.
func (o *TvDetails) HasCredits() bool {
	if o != nil && !IsNil(o.Credits) {
		return true
	}

	return false
}

// SetCredits gets a reference to the given MovieDetailsCredits and assigns it to the Credits field.
func (o *TvDetails) SetCredits(v MovieDetailsCredits) {
	o.Credits = &v
}

// GetExternalIds returns the ExternalIds field value if set, zero value otherwise.
func (o *TvDetails) GetExternalIds() ExternalIds {
	if o == nil || IsNil(o.ExternalIds) {
		var ret ExternalIds
		return ret
	}
	return *o.ExternalIds
}

// GetExternalIdsOk returns a tuple with the ExternalIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvDetails) GetExternalIdsOk() (*ExternalIds, bool) {
	if o == nil || IsNil(o.ExternalIds) {
		return nil, false
	}
	return o.ExternalIds, true
}

// HasExternalIds returns a boolean if a field has been set.
func (o *TvDetails) HasExternalIds() bool {
	if o != nil && !IsNil(o.ExternalIds) {
		return true
	}

	return false
}

// SetExternalIds gets a reference to the given ExternalIds and assigns it to the ExternalIds field.
func (o *TvDetails) SetExternalIds(v ExternalIds) {
	o.ExternalIds = &v
}

// GetKeywords returns the Keywords field value if set, zero value otherwise.
func (o *TvDetails) GetKeywords() []Keyword {
	if o == nil || IsNil(o.Keywords) {
		var ret []Keyword
		return ret
	}
	return o.Keywords
}

// GetKeywordsOk returns a tuple with the Keywords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvDetails) GetKeywordsOk() ([]Keyword, bool) {
	if o == nil || IsNil(o.Keywords) {
		return nil, false
	}
	return o.Keywords, true
}

// HasKeywords returns a boolean if a field has been set.
func (o *TvDetails) HasKeywords() bool {
	if o != nil && !IsNil(o.Keywords) {
		return true
	}

	return false
}

// SetKeywords gets a reference to the given []Keyword and assigns it to the Keywords field.
func (o *TvDetails) SetKeywords(v []Keyword) {
	o.Keywords = v
}

// GetMediaInfo returns the MediaInfo field value if set, zero value otherwise.
func (o *TvDetails) GetMediaInfo() MediaInfo {
	if o == nil || IsNil(o.MediaInfo) {
		var ret MediaInfo
		return ret
	}
	return *o.MediaInfo
}

// GetMediaInfoOk returns a tuple with the MediaInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvDetails) GetMediaInfoOk() (*MediaInfo, bool) {
	if o == nil || IsNil(o.MediaInfo) {
		return nil, false
	}
	return o.MediaInfo, true
}

// HasMediaInfo returns a boolean if a field has been set.
func (o *TvDetails) HasMediaInfo() bool {
	if o != nil && !IsNil(o.MediaInfo) {
		return true
	}

	return false
}

// SetMediaInfo gets a reference to the given MediaInfo and assigns it to the MediaInfo field.
func (o *TvDetails) SetMediaInfo(v MediaInfo) {
	o.MediaInfo = &v
}

// GetWatchProviders returns the WatchProviders field value if set, zero value otherwise.
func (o *TvDetails) GetWatchProviders() [][]WatchProvidersInner {
	if o == nil || IsNil(o.WatchProviders) {
		var ret [][]WatchProvidersInner
		return ret
	}
	return o.WatchProviders
}

// GetWatchProvidersOk returns a tuple with the WatchProviders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvDetails) GetWatchProvidersOk() ([][]WatchProvidersInner, bool) {
	if o == nil || IsNil(o.WatchProviders) {
		return nil, false
	}
	return o.WatchProviders, true
}

// HasWatchProviders returns a boolean if a field has been set.
func (o *TvDetails) HasWatchProviders() bool {
	if o != nil && !IsNil(o.WatchProviders) {
		return true
	}

	return false
}

// SetWatchProviders gets a reference to the given [][]WatchProvidersInner and assigns it to the WatchProviders field.
func (o *TvDetails) SetWatchProviders(v [][]WatchProvidersInner) {
	o.WatchProviders = v
}

func (o TvDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TvDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.BackdropPath) {
		toSerialize["backdropPath"] = o.BackdropPath
	}
	if !IsNil(o.PosterPath) {
		toSerialize["posterPath"] = o.PosterPath
	}
	if !IsNil(o.ContentRatings) {
		toSerialize["contentRatings"] = o.ContentRatings
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if !IsNil(o.EpisodeRunTime) {
		toSerialize["episodeRunTime"] = o.EpisodeRunTime
	}
	if !IsNil(o.FirstAirDate) {
		toSerialize["firstAirDate"] = o.FirstAirDate
	}
	if !IsNil(o.Genres) {
		toSerialize["genres"] = o.Genres
	}
	if !IsNil(o.Homepage) {
		toSerialize["homepage"] = o.Homepage
	}
	if !IsNil(o.InProduction) {
		toSerialize["inProduction"] = o.InProduction
	}
	if !IsNil(o.Languages) {
		toSerialize["languages"] = o.Languages
	}
	if !IsNil(o.LastAirDate) {
		toSerialize["lastAirDate"] = o.LastAirDate
	}
	if !IsNil(o.LastEpisodeToAir) {
		toSerialize["lastEpisodeToAir"] = o.LastEpisodeToAir
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.NextEpisodeToAir) {
		toSerialize["nextEpisodeToAir"] = o.NextEpisodeToAir
	}
	if !IsNil(o.Networks) {
		toSerialize["networks"] = o.Networks
	}
	if !IsNil(o.NumberOfEpisodes) {
		toSerialize["numberOfEpisodes"] = o.NumberOfEpisodes
	}
	if !IsNil(o.NumberOfSeason) {
		toSerialize["numberOfSeason"] = o.NumberOfSeason
	}
	if !IsNil(o.OriginCountry) {
		toSerialize["originCountry"] = o.OriginCountry
	}
	if !IsNil(o.OriginalLanguage) {
		toSerialize["originalLanguage"] = o.OriginalLanguage
	}
	if !IsNil(o.OriginalName) {
		toSerialize["originalName"] = o.OriginalName
	}
	if !IsNil(o.Overview) {
		toSerialize["overview"] = o.Overview
	}
	if !IsNil(o.Popularity) {
		toSerialize["popularity"] = o.Popularity
	}
	if !IsNil(o.ProductionCompanies) {
		toSerialize["productionCompanies"] = o.ProductionCompanies
	}
	if !IsNil(o.ProductionCountries) {
		toSerialize["productionCountries"] = o.ProductionCountries
	}
	if !IsNil(o.SpokenLanguages) {
		toSerialize["spokenLanguages"] = o.SpokenLanguages
	}
	if !IsNil(o.Seasons) {
		toSerialize["seasons"] = o.Seasons
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Tagline) {
		toSerialize["tagline"] = o.Tagline
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.VoteAverage) {
		toSerialize["voteAverage"] = o.VoteAverage
	}
	if !IsNil(o.VoteCount) {
		toSerialize["voteCount"] = o.VoteCount
	}
	if !IsNil(o.Credits) {
		toSerialize["credits"] = o.Credits
	}
	if !IsNil(o.ExternalIds) {
		toSerialize["externalIds"] = o.ExternalIds
	}
	if !IsNil(o.Keywords) {
		toSerialize["keywords"] = o.Keywords
	}
	if !IsNil(o.MediaInfo) {
		toSerialize["mediaInfo"] = o.MediaInfo
	}
	if !IsNil(o.WatchProviders) {
		toSerialize["watchProviders"] = o.WatchProviders
	}
	return toSerialize, nil
}

type NullableTvDetails struct {
	value *TvDetails
	isSet bool
}

func (v NullableTvDetails) Get() *TvDetails {
	return v.value
}

func (v *NullableTvDetails) Set(val *TvDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableTvDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableTvDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTvDetails(val *TvDetails) *NullableTvDetails {
	return &NullableTvDetails{value: val, isSet: true}
}

func (v NullableTvDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTvDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


