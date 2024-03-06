# \FunctionsAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FunctionsArchive**](FunctionsAPI.md#FunctionsArchive) | **Delete** /automation/v4/actions/{appId}/{definitionId}/functions/{functionType}/{functionId} | Archive a function for a definition
[**FunctionsArchiveByType**](FunctionsAPI.md#FunctionsArchiveByType) | **Delete** /automation/v4/actions/{appId}/{definitionId}/functions/{functionType} | Delete a function for a definition
[**FunctionsCreateOrReplace**](FunctionsAPI.md#FunctionsCreateOrReplace) | **Put** /automation/v4/actions/{appId}/{definitionId}/functions/{functionType}/{functionId} | Insert a function for a definition
[**FunctionsCreateOrReplaceByType**](FunctionsAPI.md#FunctionsCreateOrReplaceByType) | **Put** /automation/v4/actions/{appId}/{definitionId}/functions/{functionType} | Insert a function for a definition
[**FunctionsGetByID**](FunctionsAPI.md#FunctionsGetByID) | **Get** /automation/v4/actions/{appId}/{definitionId}/functions/{functionType}/{functionId} | Get a function for a given definition
[**FunctionsGetByType**](FunctionsAPI.md#FunctionsGetByType) | **Get** /automation/v4/actions/{appId}/{definitionId}/functions/{functionType} | Get all functions by a type for a given definition
[**FunctionsGetPage**](FunctionsAPI.md#FunctionsGetPage) | **Get** /automation/v4/actions/{appId}/{definitionId}/functions | Get all functions for a given definition



## FunctionsArchive

> FunctionsArchive(ctx, definitionId, functionType, functionId, appId).Execute()

Archive a function for a definition

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	definitionId := "definitionId_example" // string | 
	functionType := "functionType_example" // string | 
	functionId := "functionId_example" // string | 
	appId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FunctionsAPI.FunctionsArchive(context.Background(), definitionId, functionType, functionId, appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.FunctionsArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**definitionId** | **string** |  | 
**functionType** | **string** |  | 
**functionId** | **string** |  | 
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFunctionsArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

 (empty response body)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FunctionsArchiveByType

> FunctionsArchiveByType(ctx, definitionId, functionType, appId).Execute()

Delete a function for a definition

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	definitionId := "definitionId_example" // string | 
	functionType := "functionType_example" // string | 
	appId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FunctionsAPI.FunctionsArchiveByType(context.Background(), definitionId, functionType, appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.FunctionsArchiveByType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**definitionId** | **string** |  | 
**functionType** | **string** |  | 
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFunctionsArchiveByTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FunctionsCreateOrReplace

> PublicActionFunctionIdentifier FunctionsCreateOrReplace(ctx, definitionId, functionType, functionId, appId).Body(body).Execute()

Insert a function for a definition

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	definitionId := "definitionId_example" // string | 
	functionType := "functionType_example" // string | 
	functionId := "functionId_example" // string | 
	appId := int32(56) // int32 | 
	body := "body_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsAPI.FunctionsCreateOrReplace(context.Background(), definitionId, functionType, functionId, appId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.FunctionsCreateOrReplace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FunctionsCreateOrReplace`: PublicActionFunctionIdentifier
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.FunctionsCreateOrReplace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**definitionId** | **string** |  | 
**functionType** | **string** |  | 
**functionId** | **string** |  | 
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFunctionsCreateOrReplaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | **string** |  | 

### Return type

[**PublicActionFunctionIdentifier**](PublicActionFunctionIdentifier.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FunctionsCreateOrReplaceByType

> PublicActionFunctionIdentifier FunctionsCreateOrReplaceByType(ctx, definitionId, functionType, appId).Body(body).Execute()

Insert a function for a definition

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	definitionId := "definitionId_example" // string | 
	functionType := "functionType_example" // string | 
	appId := int32(56) // int32 | 
	body := "body_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsAPI.FunctionsCreateOrReplaceByType(context.Background(), definitionId, functionType, appId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.FunctionsCreateOrReplaceByType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FunctionsCreateOrReplaceByType`: PublicActionFunctionIdentifier
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.FunctionsCreateOrReplaceByType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**definitionId** | **string** |  | 
**functionType** | **string** |  | 
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFunctionsCreateOrReplaceByTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | **string** |  | 

### Return type

[**PublicActionFunctionIdentifier**](PublicActionFunctionIdentifier.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FunctionsGetByID

> PublicActionFunction FunctionsGetByID(ctx, definitionId, functionType, functionId, appId).Execute()

Get a function for a given definition

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	definitionId := "definitionId_example" // string | 
	functionType := "functionType_example" // string | 
	functionId := "functionId_example" // string | 
	appId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsAPI.FunctionsGetByID(context.Background(), definitionId, functionType, functionId, appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.FunctionsGetByID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FunctionsGetByID`: PublicActionFunction
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.FunctionsGetByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**definitionId** | **string** |  | 
**functionType** | **string** |  | 
**functionId** | **string** |  | 
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFunctionsGetByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**PublicActionFunction**](PublicActionFunction.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FunctionsGetByType

> PublicActionFunction FunctionsGetByType(ctx, definitionId, functionType, appId).Execute()

Get all functions by a type for a given definition

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	definitionId := "definitionId_example" // string | 
	functionType := "functionType_example" // string | 
	appId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsAPI.FunctionsGetByType(context.Background(), definitionId, functionType, appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.FunctionsGetByType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FunctionsGetByType`: PublicActionFunction
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.FunctionsGetByType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**definitionId** | **string** |  | 
**functionType** | **string** |  | 
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFunctionsGetByTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**PublicActionFunction**](PublicActionFunction.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FunctionsGetPage

> CollectionResponsePublicActionFunctionIdentifierNoPaging FunctionsGetPage(ctx, definitionId, appId).Execute()

Get all functions for a given definition

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	definitionId := "definitionId_example" // string | 
	appId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsAPI.FunctionsGetPage(context.Background(), definitionId, appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.FunctionsGetPage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FunctionsGetPage`: CollectionResponsePublicActionFunctionIdentifierNoPaging
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.FunctionsGetPage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**definitionId** | **string** |  | 
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFunctionsGetPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CollectionResponsePublicActionFunctionIdentifierNoPaging**](CollectionResponsePublicActionFunctionIdentifierNoPaging.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

