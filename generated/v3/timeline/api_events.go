/*
CRM Timeline

This feature allows an app to create and configure custom events that can show up in the timelines of certain CRM objects like contacts, companies, tickets, or deals. You'll find multiple use cases for this API in the sections below.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package timeline

import (
	"bytes"
	"context"
	"io"
	"net/http"

	"github.com/inpher/go-hubspot"
	"net/url"
	"strings"
)

// EventsAPIService EventsAPI service
type EventsAPIService service

type ApiGetCrmV3TimelineEventsEventTemplateIdEventIdDetailGetDetailByIdRequest struct {
	ctx             context.Context
	ApiService      *EventsAPIService
	eventTemplateId string
	eventId         string
}

func (r ApiGetCrmV3TimelineEventsEventTemplateIdEventIdDetailGetDetailByIdRequest) Execute() (*EventDetail, *http.Response, error) {
	return r.ApiService.GetCrmV3TimelineEventsEventTemplateIdEventIdDetailGetDetailByIdExecute(r)
}

/*
GetCrmV3TimelineEventsEventTemplateIdEventIdDetailGetDetailById Gets the detailTemplate as rendered

This will take the `detailTemplate` from the event template and return an object rendering the specified event. If the template references `extraData` that isn't found in the event, it will be ignored and we'll render without it.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param eventTemplateId The event template ID.
	@param eventId The event ID.
	@return ApiGetCrmV3TimelineEventsEventTemplateIdEventIdDetailGetDetailByIdRequest
*/
func (a *EventsAPIService) GetCrmV3TimelineEventsEventTemplateIdEventIdDetailGetDetailById(ctx context.Context, eventTemplateId string, eventId string) ApiGetCrmV3TimelineEventsEventTemplateIdEventIdDetailGetDetailByIdRequest {
	return ApiGetCrmV3TimelineEventsEventTemplateIdEventIdDetailGetDetailByIdRequest{
		ApiService:      a,
		ctx:             ctx,
		eventTemplateId: eventTemplateId,
		eventId:         eventId,
	}
}

// Execute executes the request
//
//	@return EventDetail
func (a *EventsAPIService) GetCrmV3TimelineEventsEventTemplateIdEventIdDetailGetDetailByIdExecute(r ApiGetCrmV3TimelineEventsEventTemplateIdEventIdDetailGetDetailByIdRequest) (*EventDetail, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *EventDetail
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventsAPIService.GetCrmV3TimelineEventsEventTemplateIdEventIdDetailGetDetailById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/timeline/events/{eventTemplateId}/{eventId}/detail"
	localVarPath = strings.Replace(localVarPath, "{"+"eventTemplateId"+"}", url.PathEscape(parameterValueToString(r.eventTemplateId, "eventTemplateId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"eventId"+"}", url.PathEscape(parameterValueToString(r.eventId, "eventId")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json", "*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["private_apps_legacy"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["private-app-legacy"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(hubspot.ContextKey).(hubspot.Authorizer); ok {
			auth.Apply(hubspot.AuthorizationRequest{
				QueryParams: localVarQueryParams,
				FormParams:  localVarFormParams,
				Headers:     localVarHeaderParams,
			})
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
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
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

type ApiGetCrmV3TimelineEventsEventTemplateIdEventIdGetByIdRequest struct {
	ctx             context.Context
	ApiService      *EventsAPIService
	eventTemplateId string
	eventId         string
}

func (r ApiGetCrmV3TimelineEventsEventTemplateIdEventIdGetByIdRequest) Execute() (*TimelineEventResponse, *http.Response, error) {
	return r.ApiService.GetCrmV3TimelineEventsEventTemplateIdEventIdGetByIdExecute(r)
}

/*
GetCrmV3TimelineEventsEventTemplateIdEventIdGetById Gets the event

This returns the previously created event. It contains all existing info for the event, but not necessarily the CRM object.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param eventTemplateId The event template ID.
	@param eventId The event ID.
	@return ApiGetCrmV3TimelineEventsEventTemplateIdEventIdGetByIdRequest
*/
func (a *EventsAPIService) GetCrmV3TimelineEventsEventTemplateIdEventIdGetById(ctx context.Context, eventTemplateId string, eventId string) ApiGetCrmV3TimelineEventsEventTemplateIdEventIdGetByIdRequest {
	return ApiGetCrmV3TimelineEventsEventTemplateIdEventIdGetByIdRequest{
		ApiService:      a,
		ctx:             ctx,
		eventTemplateId: eventTemplateId,
		eventId:         eventId,
	}
}

// Execute executes the request
//
//	@return TimelineEventResponse
func (a *EventsAPIService) GetCrmV3TimelineEventsEventTemplateIdEventIdGetByIdExecute(r ApiGetCrmV3TimelineEventsEventTemplateIdEventIdGetByIdRequest) (*TimelineEventResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TimelineEventResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventsAPIService.GetCrmV3TimelineEventsEventTemplateIdEventIdGetById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/timeline/events/{eventTemplateId}/{eventId}"
	localVarPath = strings.Replace(localVarPath, "{"+"eventTemplateId"+"}", url.PathEscape(parameterValueToString(r.eventTemplateId, "eventTemplateId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"eventId"+"}", url.PathEscape(parameterValueToString(r.eventId, "eventId")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json", "*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["private_apps_legacy"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["private-app-legacy"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(hubspot.ContextKey).(hubspot.Authorizer); ok {
			auth.Apply(hubspot.AuthorizationRequest{
				QueryParams: localVarQueryParams,
				FormParams:  localVarFormParams,
				Headers:     localVarHeaderParams,
			})
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
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
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

type ApiGetCrmV3TimelineEventsEventTemplateIdEventIdRenderGetRenderByIdRequest struct {
	ctx             context.Context
	ApiService      *EventsAPIService
	eventTemplateId string
	eventId         string
	detail          *bool
}

// Set to &#39;true&#39;, we want to render the &#x60;detailTemplate&#x60; instead of the &#x60;headerTemplate&#x60;.
func (r ApiGetCrmV3TimelineEventsEventTemplateIdEventIdRenderGetRenderByIdRequest) Detail(detail bool) ApiGetCrmV3TimelineEventsEventTemplateIdEventIdRenderGetRenderByIdRequest {
	r.detail = &detail
	return r
}

func (r ApiGetCrmV3TimelineEventsEventTemplateIdEventIdRenderGetRenderByIdRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.GetCrmV3TimelineEventsEventTemplateIdEventIdRenderGetRenderByIdExecute(r)
}

/*
GetCrmV3TimelineEventsEventTemplateIdEventIdRenderGetRenderById Renders the header or detail as HTML

This will take either the `headerTemplate` or `detailTemplate` from the event template and render for the specified event as HTML. If the template references `extraData` that isn't found in the event, it will be ignored and we'll render without it.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param eventTemplateId The event template ID.
	@param eventId The event ID.
	@return ApiGetCrmV3TimelineEventsEventTemplateIdEventIdRenderGetRenderByIdRequest
*/
func (a *EventsAPIService) GetCrmV3TimelineEventsEventTemplateIdEventIdRenderGetRenderById(ctx context.Context, eventTemplateId string, eventId string) ApiGetCrmV3TimelineEventsEventTemplateIdEventIdRenderGetRenderByIdRequest {
	return ApiGetCrmV3TimelineEventsEventTemplateIdEventIdRenderGetRenderByIdRequest{
		ApiService:      a,
		ctx:             ctx,
		eventTemplateId: eventTemplateId,
		eventId:         eventId,
	}
}

// Execute executes the request
//
//	@return string
func (a *EventsAPIService) GetCrmV3TimelineEventsEventTemplateIdEventIdRenderGetRenderByIdExecute(r ApiGetCrmV3TimelineEventsEventTemplateIdEventIdRenderGetRenderByIdRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventsAPIService.GetCrmV3TimelineEventsEventTemplateIdEventIdRenderGetRenderById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/timeline/events/{eventTemplateId}/{eventId}/render"
	localVarPath = strings.Replace(localVarPath, "{"+"eventTemplateId"+"}", url.PathEscape(parameterValueToString(r.eventTemplateId, "eventTemplateId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"eventId"+"}", url.PathEscape(parameterValueToString(r.eventId, "eventId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.detail != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "detail", r.detail, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/html", "*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["private_apps_legacy"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["private-app-legacy"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(hubspot.ContextKey).(hubspot.Authorizer); ok {
			auth.Apply(hubspot.AuthorizationRequest{
				QueryParams: localVarQueryParams,
				FormParams:  localVarFormParams,
				Headers:     localVarHeaderParams,
			})
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
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
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

type ApiPostCrmV3TimelineEventsBatchCreateCreateBatchRequest struct {
	ctx                     context.Context
	ApiService              *EventsAPIService
	batchInputTimelineEvent *BatchInputTimelineEvent
}

// The timeline event definition.
func (r ApiPostCrmV3TimelineEventsBatchCreateCreateBatchRequest) BatchInputTimelineEvent(batchInputTimelineEvent BatchInputTimelineEvent) ApiPostCrmV3TimelineEventsBatchCreateCreateBatchRequest {
	r.batchInputTimelineEvent = &batchInputTimelineEvent
	return r
}

func (r ApiPostCrmV3TimelineEventsBatchCreateCreateBatchRequest) Execute() (*http.Response, error) {
	return r.ApiService.PostCrmV3TimelineEventsBatchCreateCreateBatchExecute(r)
}

/*
PostCrmV3TimelineEventsBatchCreateCreateBatch Creates multiple events

Creates multiple instances of timeline events based on an event template. Once created, these event are immutable on the object timeline and cannot be modified. If the event template was configured to update object properties via `objectPropertyName`, this call will also attempt to updates those properties, or add them if they don't exist.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiPostCrmV3TimelineEventsBatchCreateCreateBatchRequest
*/
func (a *EventsAPIService) PostCrmV3TimelineEventsBatchCreateCreateBatch(ctx context.Context) ApiPostCrmV3TimelineEventsBatchCreateCreateBatchRequest {
	return ApiPostCrmV3TimelineEventsBatchCreateCreateBatchRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *EventsAPIService) PostCrmV3TimelineEventsBatchCreateCreateBatchExecute(r ApiPostCrmV3TimelineEventsBatchCreateCreateBatchRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventsAPIService.PostCrmV3TimelineEventsBatchCreateCreateBatch")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/timeline/events/batch/create"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.batchInputTimelineEvent == nil {
		return nil, reportError("batchInputTimelineEvent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.batchInputTimelineEvent
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["private_apps_legacy"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["private-app-legacy"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(hubspot.ContextKey).(hubspot.Authorizer); ok {
			auth.Apply(hubspot.AuthorizationRequest{
				QueryParams: localVarQueryParams,
				FormParams:  localVarFormParams,
				Headers:     localVarHeaderParams,
			})
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiPostCrmV3TimelineEventsCreateRequest struct {
	ctx           context.Context
	ApiService    *EventsAPIService
	timelineEvent *TimelineEvent
}

// The timeline event definition.
func (r ApiPostCrmV3TimelineEventsCreateRequest) TimelineEvent(timelineEvent TimelineEvent) ApiPostCrmV3TimelineEventsCreateRequest {
	r.timelineEvent = &timelineEvent
	return r
}

func (r ApiPostCrmV3TimelineEventsCreateRequest) Execute() (*TimelineEventResponse, *http.Response, error) {
	return r.ApiService.PostCrmV3TimelineEventsCreateExecute(r)
}

/*
PostCrmV3TimelineEventsCreate Create a single event

Creates an instance of a timeline event based on an event template. Once created, this event is immutable on the object timeline and cannot be modified. If the event template was configured to update object properties via `objectPropertyName`, this call will also attempt to updates those properties, or add them if they don't exist.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiPostCrmV3TimelineEventsCreateRequest
*/
func (a *EventsAPIService) PostCrmV3TimelineEventsCreate(ctx context.Context) ApiPostCrmV3TimelineEventsCreateRequest {
	return ApiPostCrmV3TimelineEventsCreateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return TimelineEventResponse
func (a *EventsAPIService) PostCrmV3TimelineEventsCreateExecute(r ApiPostCrmV3TimelineEventsCreateRequest) (*TimelineEventResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TimelineEventResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventsAPIService.PostCrmV3TimelineEventsCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/timeline/events"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.timelineEvent == nil {
		return localVarReturnValue, nil, reportError("timelineEvent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.timelineEvent
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["private_apps_legacy"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["private-app-legacy"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(hubspot.ContextKey).(hubspot.Authorizer); ok {
			auth.Apply(hubspot.AuthorizationRequest{
				QueryParams: localVarQueryParams,
				FormParams:  localVarFormParams,
				Headers:     localVarHeaderParams,
			})
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
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
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
