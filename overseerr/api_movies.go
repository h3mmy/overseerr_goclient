/*
Overseerr API

This is the documentation for the Overseerr API backend.  Two primary authentication methods are supported:  - **Cookie Authentication**: A valid sign-in to the `/auth/plex` or `/auth/local` will generate a valid authentication cookie. - **API Key Authentication**: Sign-in is also possible by passing an `X-Api-Key` header along with a valid API Key generated by Overseerr.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package overseerr_go

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// MoviesAPIService MoviesAPI service
type MoviesAPIService service

type MoviesAPIMovieMovieIdGetRequest struct {
	ctx        context.Context
	ApiService *MoviesAPIService
	movieId    float32
	language   *string
}

func (r MoviesAPIMovieMovieIdGetRequest) Language(language string) MoviesAPIMovieMovieIdGetRequest {
	r.language = &language
	return r
}

func (r MoviesAPIMovieMovieIdGetRequest) Execute() (*MovieDetails, *http.Response, error) {
	return r.ApiService.MovieMovieIdGetExecute(r)
}

/*
MovieMovieIdGet Get movie details

Returns full movie details in a JSON object.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param movieId
	@return MoviesAPIMovieMovieIdGetRequest
*/
func (a *MoviesAPIService) MovieMovieIdGet(ctx context.Context, movieId float32) MoviesAPIMovieMovieIdGetRequest {
	return MoviesAPIMovieMovieIdGetRequest{
		ApiService: a,
		ctx:        ctx,
		movieId:    movieId,
	}
}

// Execute executes the request
//
//	@return MovieDetails
func (a *MoviesAPIService) MovieMovieIdGetExecute(r MoviesAPIMovieMovieIdGetRequest) (*MovieDetails, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *MovieDetails
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MoviesAPIService.MovieMovieIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/movie/{movieId}"
	localVarPath = strings.Replace(localVarPath, "{"+"movieId"+"}", url.PathEscape(parameterValueToString(r.movieId, "movieId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.language != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "language", r.language, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Api-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MoviesAPIMovieMovieIdRatingsGetRequest struct {
	ctx        context.Context
	ApiService *MoviesAPIService
	movieId    float32
}

func (r MoviesAPIMovieMovieIdRatingsGetRequest) Execute() (*MovieMovieIdRatingsGet200Response, *http.Response, error) {
	return r.ApiService.MovieMovieIdRatingsGetExecute(r)
}

/*
MovieMovieIdRatingsGet Get movie ratings

Returns ratings based on the provided movieId in a JSON object.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param movieId
	@return MoviesAPIMovieMovieIdRatingsGetRequest
*/
func (a *MoviesAPIService) MovieMovieIdRatingsGet(ctx context.Context, movieId float32) MoviesAPIMovieMovieIdRatingsGetRequest {
	return MoviesAPIMovieMovieIdRatingsGetRequest{
		ApiService: a,
		ctx:        ctx,
		movieId:    movieId,
	}
}

// Execute executes the request
//
//	@return MovieMovieIdRatingsGet200Response
func (a *MoviesAPIService) MovieMovieIdRatingsGetExecute(r MoviesAPIMovieMovieIdRatingsGetRequest) (*MovieMovieIdRatingsGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *MovieMovieIdRatingsGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MoviesAPIService.MovieMovieIdRatingsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/movie/{movieId}/ratings"
	localVarPath = strings.Replace(localVarPath, "{"+"movieId"+"}", url.PathEscape(parameterValueToString(r.movieId, "movieId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Api-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MoviesAPIMovieMovieIdRatingscombinedGetRequest struct {
	ctx        context.Context
	ApiService *MoviesAPIService
	movieId    float32
}

func (r MoviesAPIMovieMovieIdRatingscombinedGetRequest) Execute() (*MovieMovieIdRatingscombinedGet200Response, *http.Response, error) {
	return r.ApiService.MovieMovieIdRatingscombinedGetExecute(r)
}

/*
MovieMovieIdRatingscombinedGet Get RT and IMDB movie ratings combined

Returns ratings from RottenTomatoes and IMDB based on the provided movieId in a JSON object.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param movieId
	@return MoviesAPIMovieMovieIdRatingscombinedGetRequest
*/
func (a *MoviesAPIService) MovieMovieIdRatingscombinedGet(ctx context.Context, movieId float32) MoviesAPIMovieMovieIdRatingscombinedGetRequest {
	return MoviesAPIMovieMovieIdRatingscombinedGetRequest{
		ApiService: a,
		ctx:        ctx,
		movieId:    movieId,
	}
}

// Execute executes the request
//
//	@return MovieMovieIdRatingscombinedGet200Response
func (a *MoviesAPIService) MovieMovieIdRatingscombinedGetExecute(r MoviesAPIMovieMovieIdRatingscombinedGetRequest) (*MovieMovieIdRatingscombinedGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *MovieMovieIdRatingscombinedGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MoviesAPIService.MovieMovieIdRatingscombinedGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/movie/{movieId}/ratingscombined"
	localVarPath = strings.Replace(localVarPath, "{"+"movieId"+"}", url.PathEscape(parameterValueToString(r.movieId, "movieId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Api-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MoviesAPIMovieMovieIdRecommendationsGetRequest struct {
	ctx        context.Context
	ApiService *MoviesAPIService
	movieId    float32
	page       *float32
	language   *string
}

func (r MoviesAPIMovieMovieIdRecommendationsGetRequest) Page(page float32) MoviesAPIMovieMovieIdRecommendationsGetRequest {
	r.page = &page
	return r
}

func (r MoviesAPIMovieMovieIdRecommendationsGetRequest) Language(language string) MoviesAPIMovieMovieIdRecommendationsGetRequest {
	r.language = &language
	return r
}

func (r MoviesAPIMovieMovieIdRecommendationsGetRequest) Execute() (*DiscoverMoviesGet200Response, *http.Response, error) {
	return r.ApiService.MovieMovieIdRecommendationsGetExecute(r)
}

/*
MovieMovieIdRecommendationsGet Get recommended movies

Returns list of recommended movies based on provided movie ID in a JSON object.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param movieId
	@return MoviesAPIMovieMovieIdRecommendationsGetRequest
*/
func (a *MoviesAPIService) MovieMovieIdRecommendationsGet(ctx context.Context, movieId float32) MoviesAPIMovieMovieIdRecommendationsGetRequest {
	return MoviesAPIMovieMovieIdRecommendationsGetRequest{
		ApiService: a,
		ctx:        ctx,
		movieId:    movieId,
	}
}

// Execute executes the request
//
//	@return DiscoverMoviesGet200Response
func (a *MoviesAPIService) MovieMovieIdRecommendationsGetExecute(r MoviesAPIMovieMovieIdRecommendationsGetRequest) (*DiscoverMoviesGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DiscoverMoviesGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MoviesAPIService.MovieMovieIdRecommendationsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/movie/{movieId}/recommendations"
	localVarPath = strings.Replace(localVarPath, "{"+"movieId"+"}", url.PathEscape(parameterValueToString(r.movieId, "movieId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	} else {
		var defaultValue float32 = 1
		r.page = &defaultValue
	}
	if r.language != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "language", r.language, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Api-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type MoviesAPIMovieMovieIdSimilarGetRequest struct {
	ctx        context.Context
	ApiService *MoviesAPIService
	movieId    float32
	page       *float32
	language   *string
}

func (r MoviesAPIMovieMovieIdSimilarGetRequest) Page(page float32) MoviesAPIMovieMovieIdSimilarGetRequest {
	r.page = &page
	return r
}

func (r MoviesAPIMovieMovieIdSimilarGetRequest) Language(language string) MoviesAPIMovieMovieIdSimilarGetRequest {
	r.language = &language
	return r
}

func (r MoviesAPIMovieMovieIdSimilarGetRequest) Execute() (*DiscoverMoviesGet200Response, *http.Response, error) {
	return r.ApiService.MovieMovieIdSimilarGetExecute(r)
}

/*
MovieMovieIdSimilarGet Get similar movies

Returns list of similar movies based on the provided movieId in a JSON object.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param movieId
	@return MoviesAPIMovieMovieIdSimilarGetRequest
*/
func (a *MoviesAPIService) MovieMovieIdSimilarGet(ctx context.Context, movieId float32) MoviesAPIMovieMovieIdSimilarGetRequest {
	return MoviesAPIMovieMovieIdSimilarGetRequest{
		ApiService: a,
		ctx:        ctx,
		movieId:    movieId,
	}
}

// Execute executes the request
//
//	@return DiscoverMoviesGet200Response
func (a *MoviesAPIService) MovieMovieIdSimilarGetExecute(r MoviesAPIMovieMovieIdSimilarGetRequest) (*DiscoverMoviesGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DiscoverMoviesGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MoviesAPIService.MovieMovieIdSimilarGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/movie/{movieId}/similar"
	localVarPath = strings.Replace(localVarPath, "{"+"movieId"+"}", url.PathEscape(parameterValueToString(r.movieId, "movieId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	} else {
		var defaultValue float32 = 1
		r.page = &defaultValue
	}
	if r.language != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "language", r.language, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Api-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
