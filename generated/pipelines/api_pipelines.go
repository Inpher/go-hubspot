/*
CRM Pipelines

Pipelines represent distinct stages in a workflow, like closing a deal or servicing a support ticket. These endpoints provide access to read and modify pipelines in HubSpot. Pipelines support `deals` and `tickets` object types.  ## Pipeline ID validation  When calling endpoints that take pipelineId as a parameter, that ID must correspond to an existing, un-archived pipeline. Otherwise the request will fail with a `404 Not Found` response.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pipelines

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// PipelinesApiService PipelinesApi service
type PipelinesApiService service

type ApiDeleteCrmV3PipelinesObjectTypePipelineIdArchiveRequest struct {
	ctx                            context.Context
	ApiService                     *PipelinesApiService
	objectType                     string
	pipelineId                     string
	validateReferencesBeforeDelete *bool
}

func (r ApiDeleteCrmV3PipelinesObjectTypePipelineIdArchiveRequest) ValidateReferencesBeforeDelete(validateReferencesBeforeDelete bool) ApiDeleteCrmV3PipelinesObjectTypePipelineIdArchiveRequest {
	r.validateReferencesBeforeDelete = &validateReferencesBeforeDelete
	return r
}

func (r ApiDeleteCrmV3PipelinesObjectTypePipelineIdArchiveRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteCrmV3PipelinesObjectTypePipelineIdArchiveExecute(r)
}

/*
DeleteCrmV3PipelinesObjectTypePipelineIdArchive Delete a pipeline

Delete the pipeline identified by `{pipelineId}`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param objectType
 @param pipelineId
 @return ApiDeleteCrmV3PipelinesObjectTypePipelineIdArchiveRequest
*/
func (a *PipelinesApiService) DeleteCrmV3PipelinesObjectTypePipelineIdArchive(ctx context.Context, objectType string, pipelineId string) ApiDeleteCrmV3PipelinesObjectTypePipelineIdArchiveRequest {
	return ApiDeleteCrmV3PipelinesObjectTypePipelineIdArchiveRequest{
		ApiService: a,
		ctx:        ctx,
		objectType: objectType,
		pipelineId: pipelineId,
	}
}

// Execute executes the request
func (a *PipelinesApiService) DeleteCrmV3PipelinesObjectTypePipelineIdArchiveExecute(r ApiDeleteCrmV3PipelinesObjectTypePipelineIdArchiveRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PipelinesApiService.DeleteCrmV3PipelinesObjectTypePipelineIdArchive")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/pipelines/{objectType}/{pipelineId}"
	localVarPath = strings.Replace(localVarPath, "{"+"objectType"+"}", url.PathEscape(parameterToString(r.objectType, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pipelineId"+"}", url.PathEscape(parameterToString(r.pipelineId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.validateReferencesBeforeDelete != nil {
		localVarQueryParams.Add("validateReferencesBeforeDelete", parameterToString(*r.validateReferencesBeforeDelete, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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

type ApiGetCrmV3PipelinesObjectTypeGetAllRequest struct {
	ctx        context.Context
	ApiService *PipelinesApiService
	objectType string
}

func (r ApiGetCrmV3PipelinesObjectTypeGetAllRequest) Execute() (*CollectionResponsePipelineNoPaging, *http.Response, error) {
	return r.ApiService.GetCrmV3PipelinesObjectTypeGetAllExecute(r)
}

/*
GetCrmV3PipelinesObjectTypeGetAll Retrieve all pipelines

Return all pipelines for the object type specified by `{objectType}`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param objectType
 @return ApiGetCrmV3PipelinesObjectTypeGetAllRequest
*/
func (a *PipelinesApiService) GetCrmV3PipelinesObjectTypeGetAll(ctx context.Context, objectType string) ApiGetCrmV3PipelinesObjectTypeGetAllRequest {
	return ApiGetCrmV3PipelinesObjectTypeGetAllRequest{
		ApiService: a,
		ctx:        ctx,
		objectType: objectType,
	}
}

// Execute executes the request
//  @return CollectionResponsePipelineNoPaging
func (a *PipelinesApiService) GetCrmV3PipelinesObjectTypeGetAllExecute(r ApiGetCrmV3PipelinesObjectTypeGetAllRequest) (*CollectionResponsePipelineNoPaging, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CollectionResponsePipelineNoPaging
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PipelinesApiService.GetCrmV3PipelinesObjectTypeGetAll")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/pipelines/{objectType}"
	localVarPath = strings.Replace(localVarPath, "{"+"objectType"+"}", url.PathEscape(parameterToString(r.objectType, "")), -1)

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

type ApiGetCrmV3PipelinesObjectTypePipelineIdGetByIdRequest struct {
	ctx        context.Context
	ApiService *PipelinesApiService
	objectType string
	pipelineId string
}

func (r ApiGetCrmV3PipelinesObjectTypePipelineIdGetByIdRequest) Execute() (*Pipeline, *http.Response, error) {
	return r.ApiService.GetCrmV3PipelinesObjectTypePipelineIdGetByIdExecute(r)
}

/*
GetCrmV3PipelinesObjectTypePipelineIdGetById Return a pipeline by ID

Return a single pipeline object identified by its unique `{pipelineId}`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param objectType
 @param pipelineId
 @return ApiGetCrmV3PipelinesObjectTypePipelineIdGetByIdRequest
*/
func (a *PipelinesApiService) GetCrmV3PipelinesObjectTypePipelineIdGetById(ctx context.Context, objectType string, pipelineId string) ApiGetCrmV3PipelinesObjectTypePipelineIdGetByIdRequest {
	return ApiGetCrmV3PipelinesObjectTypePipelineIdGetByIdRequest{
		ApiService: a,
		ctx:        ctx,
		objectType: objectType,
		pipelineId: pipelineId,
	}
}

// Execute executes the request
//  @return Pipeline
func (a *PipelinesApiService) GetCrmV3PipelinesObjectTypePipelineIdGetByIdExecute(r ApiGetCrmV3PipelinesObjectTypePipelineIdGetByIdRequest) (*Pipeline, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Pipeline
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PipelinesApiService.GetCrmV3PipelinesObjectTypePipelineIdGetById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/pipelines/{objectType}/{pipelineId}"
	localVarPath = strings.Replace(localVarPath, "{"+"objectType"+"}", url.PathEscape(parameterToString(r.objectType, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pipelineId"+"}", url.PathEscape(parameterToString(r.pipelineId, "")), -1)

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

type ApiPatchCrmV3PipelinesObjectTypePipelineIdUpdateRequest struct {
	ctx                            context.Context
	ApiService                     *PipelinesApiService
	objectType                     string
	pipelineId                     string
	pipelinePatchInput             *PipelinePatchInput
	validateReferencesBeforeDelete *bool
}

func (r ApiPatchCrmV3PipelinesObjectTypePipelineIdUpdateRequest) PipelinePatchInput(pipelinePatchInput PipelinePatchInput) ApiPatchCrmV3PipelinesObjectTypePipelineIdUpdateRequest {
	r.pipelinePatchInput = &pipelinePatchInput
	return r
}

func (r ApiPatchCrmV3PipelinesObjectTypePipelineIdUpdateRequest) ValidateReferencesBeforeDelete(validateReferencesBeforeDelete bool) ApiPatchCrmV3PipelinesObjectTypePipelineIdUpdateRequest {
	r.validateReferencesBeforeDelete = &validateReferencesBeforeDelete
	return r
}

func (r ApiPatchCrmV3PipelinesObjectTypePipelineIdUpdateRequest) Execute() (*Pipeline, *http.Response, error) {
	return r.ApiService.PatchCrmV3PipelinesObjectTypePipelineIdUpdateExecute(r)
}

/*
PatchCrmV3PipelinesObjectTypePipelineIdUpdate Update a pipeline

Perform a partial update of the pipeline identified by `{pipelineId}`. The updated pipeline will be returned in the response.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param objectType
 @param pipelineId
 @return ApiPatchCrmV3PipelinesObjectTypePipelineIdUpdateRequest
*/
func (a *PipelinesApiService) PatchCrmV3PipelinesObjectTypePipelineIdUpdate(ctx context.Context, objectType string, pipelineId string) ApiPatchCrmV3PipelinesObjectTypePipelineIdUpdateRequest {
	return ApiPatchCrmV3PipelinesObjectTypePipelineIdUpdateRequest{
		ApiService: a,
		ctx:        ctx,
		objectType: objectType,
		pipelineId: pipelineId,
	}
}

// Execute executes the request
//  @return Pipeline
func (a *PipelinesApiService) PatchCrmV3PipelinesObjectTypePipelineIdUpdateExecute(r ApiPatchCrmV3PipelinesObjectTypePipelineIdUpdateRequest) (*Pipeline, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Pipeline
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PipelinesApiService.PatchCrmV3PipelinesObjectTypePipelineIdUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/pipelines/{objectType}/{pipelineId}"
	localVarPath = strings.Replace(localVarPath, "{"+"objectType"+"}", url.PathEscape(parameterToString(r.objectType, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pipelineId"+"}", url.PathEscape(parameterToString(r.pipelineId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pipelinePatchInput == nil {
		return localVarReturnValue, nil, reportError("pipelinePatchInput is required and must be specified")
	}

	if r.validateReferencesBeforeDelete != nil {
		localVarQueryParams.Add("validateReferencesBeforeDelete", parameterToString(*r.validateReferencesBeforeDelete, ""))
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
	localVarPostBody = r.pipelinePatchInput
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

type ApiPostCrmV3PipelinesObjectTypeCreateRequest struct {
	ctx           context.Context
	ApiService    *PipelinesApiService
	objectType    string
	pipelineInput *PipelineInput
}

func (r ApiPostCrmV3PipelinesObjectTypeCreateRequest) PipelineInput(pipelineInput PipelineInput) ApiPostCrmV3PipelinesObjectTypeCreateRequest {
	r.pipelineInput = &pipelineInput
	return r
}

func (r ApiPostCrmV3PipelinesObjectTypeCreateRequest) Execute() (*Pipeline, *http.Response, error) {
	return r.ApiService.PostCrmV3PipelinesObjectTypeCreateExecute(r)
}

/*
PostCrmV3PipelinesObjectTypeCreate Create a pipeline

Create a new pipeline with the provided property values. The entire pipeline object, including its unique ID, will be returned in the response.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param objectType
 @return ApiPostCrmV3PipelinesObjectTypeCreateRequest
*/
func (a *PipelinesApiService) PostCrmV3PipelinesObjectTypeCreate(ctx context.Context, objectType string) ApiPostCrmV3PipelinesObjectTypeCreateRequest {
	return ApiPostCrmV3PipelinesObjectTypeCreateRequest{
		ApiService: a,
		ctx:        ctx,
		objectType: objectType,
	}
}

// Execute executes the request
//  @return Pipeline
func (a *PipelinesApiService) PostCrmV3PipelinesObjectTypeCreateExecute(r ApiPostCrmV3PipelinesObjectTypeCreateRequest) (*Pipeline, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Pipeline
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PipelinesApiService.PostCrmV3PipelinesObjectTypeCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/pipelines/{objectType}"
	localVarPath = strings.Replace(localVarPath, "{"+"objectType"+"}", url.PathEscape(parameterToString(r.objectType, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pipelineInput == nil {
		return localVarReturnValue, nil, reportError("pipelineInput is required and must be specified")
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
	localVarPostBody = r.pipelineInput
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

type ApiPutCrmV3PipelinesObjectTypePipelineIdReplaceRequest struct {
	ctx                            context.Context
	ApiService                     *PipelinesApiService
	objectType                     string
	pipelineId                     string
	pipelineInput                  *PipelineInput
	validateReferencesBeforeDelete *bool
}

func (r ApiPutCrmV3PipelinesObjectTypePipelineIdReplaceRequest) PipelineInput(pipelineInput PipelineInput) ApiPutCrmV3PipelinesObjectTypePipelineIdReplaceRequest {
	r.pipelineInput = &pipelineInput
	return r
}

func (r ApiPutCrmV3PipelinesObjectTypePipelineIdReplaceRequest) ValidateReferencesBeforeDelete(validateReferencesBeforeDelete bool) ApiPutCrmV3PipelinesObjectTypePipelineIdReplaceRequest {
	r.validateReferencesBeforeDelete = &validateReferencesBeforeDelete
	return r
}

func (r ApiPutCrmV3PipelinesObjectTypePipelineIdReplaceRequest) Execute() (*Pipeline, *http.Response, error) {
	return r.ApiService.PutCrmV3PipelinesObjectTypePipelineIdReplaceExecute(r)
}

/*
PutCrmV3PipelinesObjectTypePipelineIdReplace Replace a pipeline

Replace all the properties of an existing pipeline with the values provided. This will overwrite any existing pipeline stages. The updated pipeline will be returned in the response.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param objectType
 @param pipelineId
 @return ApiPutCrmV3PipelinesObjectTypePipelineIdReplaceRequest
*/
func (a *PipelinesApiService) PutCrmV3PipelinesObjectTypePipelineIdReplace(ctx context.Context, objectType string, pipelineId string) ApiPutCrmV3PipelinesObjectTypePipelineIdReplaceRequest {
	return ApiPutCrmV3PipelinesObjectTypePipelineIdReplaceRequest{
		ApiService: a,
		ctx:        ctx,
		objectType: objectType,
		pipelineId: pipelineId,
	}
}

// Execute executes the request
//  @return Pipeline
func (a *PipelinesApiService) PutCrmV3PipelinesObjectTypePipelineIdReplaceExecute(r ApiPutCrmV3PipelinesObjectTypePipelineIdReplaceRequest) (*Pipeline, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Pipeline
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PipelinesApiService.PutCrmV3PipelinesObjectTypePipelineIdReplace")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/crm/v3/pipelines/{objectType}/{pipelineId}"
	localVarPath = strings.Replace(localVarPath, "{"+"objectType"+"}", url.PathEscape(parameterToString(r.objectType, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pipelineId"+"}", url.PathEscape(parameterToString(r.pipelineId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pipelineInput == nil {
		return localVarReturnValue, nil, reportError("pipelineInput is required and must be specified")
	}

	if r.validateReferencesBeforeDelete != nil {
		localVarQueryParams.Add("validateReferencesBeforeDelete", parameterToString(*r.validateReferencesBeforeDelete, ""))
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
	localVarPostBody = r.pipelineInput
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
