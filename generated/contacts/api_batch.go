/*
Contacts

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package contacts

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)

// BatchApiService BatchApi service
type BatchApiService service

type ApiPostCrmV3ObjectsContactsBatchArchiveArchiveRequest struct {
	ctx                            context.Context
	ApiService                     *BatchApiService
	batchInputSimplePublicObjectId *BatchInputSimplePublicObjectId
}

func (r ApiPostCrmV3ObjectsContactsBatchArchiveArchiveRequest) BatchInputSimplePublicObjectId(batchInputSimplePublicObjectId BatchInputSimplePublicObjectId) ApiPostCrmV3ObjectsContactsBatchArchiveArchiveRequest {
	r.batchInputSimplePublicObjectId = &batchInputSimplePublicObjectId
	return r
}

func (r ApiPostCrmV3ObjectsContactsBatchArchiveArchiveRequest) Execute() (*http.Response, error) {
	return r.ApiService.PostCrmV3ObjectsContactsBatchArchiveArchiveExecute(r)
}

/*
PostCrmV3ObjectsContactsBatchArchiveArchive Archive a batch of contacts by ID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostCrmV3ObjectsContactsBatchArchiveArchiveRequest
*/
func (a *BatchApiService) PostCrmV3ObjectsContactsBatchArchiveArchive(ctx context.Context) ApiPostCrmV3ObjectsContactsBatchArchiveArchiveRequest {
	return ApiPostCrmV3ObjectsContactsBatchArchiveArchiveRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *BatchApiService) PostCrmV3ObjectsContactsBatchArchiveArchiveExecute(r ApiPostCrmV3ObjectsContactsBatchArchiveArchiveRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BatchApiService.PostCrmV3ObjectsContactsBatchArchiveArchive")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/objects/contacts/batch/archive"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.batchInputSimplePublicObjectId == nil {
		return nil, reportError("batchInputSimplePublicObjectId is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.batchInputSimplePublicObjectId
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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
		newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiPostCrmV3ObjectsContactsBatchCreateCreateRequest struct {
	ctx                               context.Context
	ApiService                        *BatchApiService
	batchInputSimplePublicObjectInput *BatchInputSimplePublicObjectInput
}

func (r ApiPostCrmV3ObjectsContactsBatchCreateCreateRequest) BatchInputSimplePublicObjectInput(batchInputSimplePublicObjectInput BatchInputSimplePublicObjectInput) ApiPostCrmV3ObjectsContactsBatchCreateCreateRequest {
	r.batchInputSimplePublicObjectInput = &batchInputSimplePublicObjectInput
	return r
}

func (r ApiPostCrmV3ObjectsContactsBatchCreateCreateRequest) Execute() (*BatchResponseSimplePublicObject, *http.Response, error) {
	return r.ApiService.PostCrmV3ObjectsContactsBatchCreateCreateExecute(r)
}

/*
PostCrmV3ObjectsContactsBatchCreateCreate Create a batch of contacts

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostCrmV3ObjectsContactsBatchCreateCreateRequest
*/
func (a *BatchApiService) PostCrmV3ObjectsContactsBatchCreateCreate(ctx context.Context) ApiPostCrmV3ObjectsContactsBatchCreateCreateRequest {
	return ApiPostCrmV3ObjectsContactsBatchCreateCreateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return BatchResponseSimplePublicObject
func (a *BatchApiService) PostCrmV3ObjectsContactsBatchCreateCreateExecute(r ApiPostCrmV3ObjectsContactsBatchCreateCreateRequest) (*BatchResponseSimplePublicObject, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BatchResponseSimplePublicObject
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BatchApiService.PostCrmV3ObjectsContactsBatchCreateCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/objects/contacts/batch/create"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.batchInputSimplePublicObjectInput == nil {
		return localVarReturnValue, nil, reportError("batchInputSimplePublicObjectInput is required and must be specified")
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
	localVarPostBody = r.batchInputSimplePublicObjectInput
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

type ApiPostCrmV3ObjectsContactsBatchReadReadRequest struct {
	ctx                                context.Context
	ApiService                         *BatchApiService
	batchReadInputSimplePublicObjectId *BatchReadInputSimplePublicObjectId
	archived                           *bool
}

func (r ApiPostCrmV3ObjectsContactsBatchReadReadRequest) BatchReadInputSimplePublicObjectId(batchReadInputSimplePublicObjectId BatchReadInputSimplePublicObjectId) ApiPostCrmV3ObjectsContactsBatchReadReadRequest {
	r.batchReadInputSimplePublicObjectId = &batchReadInputSimplePublicObjectId
	return r
}

// Whether to return only results that have been archived.
func (r ApiPostCrmV3ObjectsContactsBatchReadReadRequest) Archived(archived bool) ApiPostCrmV3ObjectsContactsBatchReadReadRequest {
	r.archived = &archived
	return r
}

func (r ApiPostCrmV3ObjectsContactsBatchReadReadRequest) Execute() (*BatchResponseSimplePublicObject, *http.Response, error) {
	return r.ApiService.PostCrmV3ObjectsContactsBatchReadReadExecute(r)
}

/*
PostCrmV3ObjectsContactsBatchReadRead Read a batch of contacts by internal ID, or unique property values

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostCrmV3ObjectsContactsBatchReadReadRequest
*/
func (a *BatchApiService) PostCrmV3ObjectsContactsBatchReadRead(ctx context.Context) ApiPostCrmV3ObjectsContactsBatchReadReadRequest {
	return ApiPostCrmV3ObjectsContactsBatchReadReadRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return BatchResponseSimplePublicObject
func (a *BatchApiService) PostCrmV3ObjectsContactsBatchReadReadExecute(r ApiPostCrmV3ObjectsContactsBatchReadReadRequest) (*BatchResponseSimplePublicObject, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BatchResponseSimplePublicObject
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BatchApiService.PostCrmV3ObjectsContactsBatchReadRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/objects/contacts/batch/read"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.batchReadInputSimplePublicObjectId == nil {
		return localVarReturnValue, nil, reportError("batchReadInputSimplePublicObjectId is required and must be specified")
	}

	if r.archived != nil {
		localVarQueryParams.Add("archived", parameterToString(*r.archived, ""))
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
	localVarPostBody = r.batchReadInputSimplePublicObjectId
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

type ApiPostCrmV3ObjectsContactsBatchUpdateUpdateRequest struct {
	ctx                                    context.Context
	ApiService                             *BatchApiService
	batchInputSimplePublicObjectBatchInput *BatchInputSimplePublicObjectBatchInput
}

func (r ApiPostCrmV3ObjectsContactsBatchUpdateUpdateRequest) BatchInputSimplePublicObjectBatchInput(batchInputSimplePublicObjectBatchInput BatchInputSimplePublicObjectBatchInput) ApiPostCrmV3ObjectsContactsBatchUpdateUpdateRequest {
	r.batchInputSimplePublicObjectBatchInput = &batchInputSimplePublicObjectBatchInput
	return r
}

func (r ApiPostCrmV3ObjectsContactsBatchUpdateUpdateRequest) Execute() (*BatchResponseSimplePublicObject, *http.Response, error) {
	return r.ApiService.PostCrmV3ObjectsContactsBatchUpdateUpdateExecute(r)
}

/*
PostCrmV3ObjectsContactsBatchUpdateUpdate Update a batch of contacts

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostCrmV3ObjectsContactsBatchUpdateUpdateRequest
*/
func (a *BatchApiService) PostCrmV3ObjectsContactsBatchUpdateUpdate(ctx context.Context) ApiPostCrmV3ObjectsContactsBatchUpdateUpdateRequest {
	return ApiPostCrmV3ObjectsContactsBatchUpdateUpdateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return BatchResponseSimplePublicObject
func (a *BatchApiService) PostCrmV3ObjectsContactsBatchUpdateUpdateExecute(r ApiPostCrmV3ObjectsContactsBatchUpdateUpdateRequest) (*BatchResponseSimplePublicObject, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BatchResponseSimplePublicObject
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BatchApiService.PostCrmV3ObjectsContactsBatchUpdateUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/objects/contacts/batch/update"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.batchInputSimplePublicObjectBatchInput == nil {
		return localVarReturnValue, nil, reportError("batchInputSimplePublicObjectBatchInput is required and must be specified")
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
	localVarPostBody = r.batchInputSimplePublicObjectBatchInput
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
