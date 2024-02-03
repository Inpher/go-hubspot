/*
Marketing Events

These APIs allow you to interact with HubSpot's Marketing Events Extension. It allows you to: * Create, Read or update Marketing Event information in HubSpot * Specify whether a HubSpot contact has registered, attended or cancelled a registration to a Marketing Event. * Specify a URL that can be called to get the details of a Marketing Event.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package marketing_events_beta

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"

	"github.com/clarkmcc/go-hubspot"
	"net/url"
	"strings"
)

// MarketingEventsExternalApiService MarketingEventsExternalApi service
type MarketingEventsExternalApiService service

type ApiExternalCompleteCompleteRequest struct {
	ctx                                 context.Context
	ApiService                          *MarketingEventsExternalApiService
	externalEventId                     string
	externalAccountId                   *string
	marketingEventCompleteRequestParams *MarketingEventCompleteRequestParams
}

func (r ApiExternalCompleteCompleteRequest) ExternalAccountId(externalAccountId string) ApiExternalCompleteCompleteRequest {
	r.externalAccountId = &externalAccountId
	return r
}

func (r ApiExternalCompleteCompleteRequest) MarketingEventCompleteRequestParams(marketingEventCompleteRequestParams MarketingEventCompleteRequestParams) ApiExternalCompleteCompleteRequest {
	r.marketingEventCompleteRequestParams = &marketingEventCompleteRequestParams
	return r
}

func (r ApiExternalCompleteCompleteRequest) Execute() (*MarketingEventDefaultResponse, *http.Response, error) {
	return r.ApiService.ExternalCompleteCompleteExecute(r)
}

/*
ExternalCompleteComplete Method for ExternalCompleteComplete

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param externalEventId
 @return ApiExternalCompleteCompleteRequest
*/
func (a *MarketingEventsExternalApiService) ExternalCompleteComplete(ctx context.Context, externalEventId string) ApiExternalCompleteCompleteRequest {
	return ApiExternalCompleteCompleteRequest{
		ApiService:      a,
		ctx:             ctx,
		externalEventId: externalEventId,
	}
}

// Execute executes the request
//  @return MarketingEventDefaultResponse
func (a *MarketingEventsExternalApiService) ExternalCompleteCompleteExecute(r ApiExternalCompleteCompleteRequest) (*MarketingEventDefaultResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *MarketingEventDefaultResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MarketingEventsExternalApiService.ExternalCompleteComplete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/marketing/v3/marketing-events/events/{externalEventId}/complete"
	localVarPath = strings.Replace(localVarPath, "{"+"externalEventId"+"}", url.PathEscape(parameterToString(r.externalEventId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.externalAccountId == nil {
		return localVarReturnValue, nil, reportError("externalAccountId is required and must be specified")
	}
	if r.marketingEventCompleteRequestParams == nil {
		return localVarReturnValue, nil, reportError("marketingEventCompleteRequestParams is required and must be specified")
	}

	localVarQueryParams.Add("externalAccountId", parameterToString(*r.externalAccountId, ""))
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
	localVarPostBody = r.marketingEventCompleteRequestParams
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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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