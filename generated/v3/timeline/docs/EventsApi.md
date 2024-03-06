# \EventsAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCrmV3TimelineEventsEventTemplateIdEventIdDetailGetDetailById**](EventsAPI.md#GetCrmV3TimelineEventsEventTemplateIdEventIdDetailGetDetailById) | **Get** /crm/v3/timeline/events/{eventTemplateId}/{eventId}/detail | Gets the detailTemplate as rendered
[**GetCrmV3TimelineEventsEventTemplateIdEventIdGetById**](EventsAPI.md#GetCrmV3TimelineEventsEventTemplateIdEventIdGetById) | **Get** /crm/v3/timeline/events/{eventTemplateId}/{eventId} | Gets the event
[**GetCrmV3TimelineEventsEventTemplateIdEventIdRenderGetRenderById**](EventsAPI.md#GetCrmV3TimelineEventsEventTemplateIdEventIdRenderGetRenderById) | **Get** /crm/v3/timeline/events/{eventTemplateId}/{eventId}/render | Renders the header or detail as HTML
[**PostCrmV3TimelineEventsBatchCreateCreateBatch**](EventsAPI.md#PostCrmV3TimelineEventsBatchCreateCreateBatch) | **Post** /crm/v3/timeline/events/batch/create | Creates multiple events
[**PostCrmV3TimelineEventsCreate**](EventsAPI.md#PostCrmV3TimelineEventsCreate) | **Post** /crm/v3/timeline/events | Create a single event



## GetCrmV3TimelineEventsEventTemplateIdEventIdDetailGetDetailById

> EventDetail GetCrmV3TimelineEventsEventTemplateIdEventIdDetailGetDetailById(ctx, eventTemplateId, eventId).Execute()

Gets the detailTemplate as rendered



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
	eventTemplateId := "eventTemplateId_example" // string | The event template ID.
	eventId := "eventId_example" // string | The event ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.GetCrmV3TimelineEventsEventTemplateIdEventIdDetailGetDetailById(context.Background(), eventTemplateId, eventId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetCrmV3TimelineEventsEventTemplateIdEventIdDetailGetDetailById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrmV3TimelineEventsEventTemplateIdEventIdDetailGetDetailById`: EventDetail
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.GetCrmV3TimelineEventsEventTemplateIdEventIdDetailGetDetailById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventTemplateId** | **string** | The event template ID. | 
**eventId** | **string** | The event ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3TimelineEventsEventTemplateIdEventIdDetailGetDetailByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EventDetail**](EventDetail.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [oauth2](../README.md#oauth2), [private_apps_legacy](../README.md#private_apps_legacy), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrmV3TimelineEventsEventTemplateIdEventIdGetById

> TimelineEventResponse GetCrmV3TimelineEventsEventTemplateIdEventIdGetById(ctx, eventTemplateId, eventId).Execute()

Gets the event



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
	eventTemplateId := "eventTemplateId_example" // string | The event template ID.
	eventId := "eventId_example" // string | The event ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.GetCrmV3TimelineEventsEventTemplateIdEventIdGetById(context.Background(), eventTemplateId, eventId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetCrmV3TimelineEventsEventTemplateIdEventIdGetById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrmV3TimelineEventsEventTemplateIdEventIdGetById`: TimelineEventResponse
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.GetCrmV3TimelineEventsEventTemplateIdEventIdGetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventTemplateId** | **string** | The event template ID. | 
**eventId** | **string** | The event ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3TimelineEventsEventTemplateIdEventIdGetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TimelineEventResponse**](TimelineEventResponse.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [oauth2](../README.md#oauth2), [private_apps_legacy](../README.md#private_apps_legacy), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrmV3TimelineEventsEventTemplateIdEventIdRenderGetRenderById

> string GetCrmV3TimelineEventsEventTemplateIdEventIdRenderGetRenderById(ctx, eventTemplateId, eventId).Detail(detail).Execute()

Renders the header or detail as HTML



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
	eventTemplateId := "eventTemplateId_example" // string | The event template ID.
	eventId := "eventId_example" // string | The event ID.
	detail := true // bool | Set to 'true', we want to render the `detailTemplate` instead of the `headerTemplate`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.GetCrmV3TimelineEventsEventTemplateIdEventIdRenderGetRenderById(context.Background(), eventTemplateId, eventId).Detail(detail).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetCrmV3TimelineEventsEventTemplateIdEventIdRenderGetRenderById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrmV3TimelineEventsEventTemplateIdEventIdRenderGetRenderById`: string
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.GetCrmV3TimelineEventsEventTemplateIdEventIdRenderGetRenderById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventTemplateId** | **string** | The event template ID. | 
**eventId** | **string** | The event ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3TimelineEventsEventTemplateIdEventIdRenderGetRenderByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **detail** | **bool** | Set to &#39;true&#39;, we want to render the &#x60;detailTemplate&#x60; instead of the &#x60;headerTemplate&#x60;. | 

### Return type

**string**

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [oauth2](../README.md#oauth2), [private_apps_legacy](../README.md#private_apps_legacy), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/html, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCrmV3TimelineEventsBatchCreateCreateBatch

> PostCrmV3TimelineEventsBatchCreateCreateBatch(ctx).BatchInputTimelineEvent(batchInputTimelineEvent).Execute()

Creates multiple events



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
	batchInputTimelineEvent := *openapiclient.NewBatchInputTimelineEvent([]openapiclient.TimelineEvent{*openapiclient.NewTimelineEvent("1001298", map[string]string{"key": "Inner_example"})}) // BatchInputTimelineEvent | The timeline event definition.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.PostCrmV3TimelineEventsBatchCreateCreateBatch(context.Background()).BatchInputTimelineEvent(batchInputTimelineEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.PostCrmV3TimelineEventsBatchCreateCreateBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3TimelineEventsBatchCreateCreateBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputTimelineEvent** | [**BatchInputTimelineEvent**](BatchInputTimelineEvent.md) | The timeline event definition. | 

### Return type

 (empty response body)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [oauth2](../README.md#oauth2), [private_apps_legacy](../README.md#private_apps_legacy), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCrmV3TimelineEventsCreate

> TimelineEventResponse PostCrmV3TimelineEventsCreate(ctx).TimelineEvent(timelineEvent).Execute()

Create a single event



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
	timelineEvent := *openapiclient.NewTimelineEvent("1001298", map[string]string{"key": "Inner_example"}) // TimelineEvent | The timeline event definition.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.PostCrmV3TimelineEventsCreate(context.Background()).TimelineEvent(timelineEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.PostCrmV3TimelineEventsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCrmV3TimelineEventsCreate`: TimelineEventResponse
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.PostCrmV3TimelineEventsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3TimelineEventsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timelineEvent** | [**TimelineEvent**](TimelineEvent.md) | The timeline event definition. | 

### Return type

[**TimelineEventResponse**](TimelineEventResponse.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [oauth2](../README.md#oauth2), [private_apps_legacy](../README.md#private_apps_legacy), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

