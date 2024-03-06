# \BlogPostsAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Archive**](BlogPostsAPI.md#Archive) | **Delete** /cms/v3/blogs/posts/{objectId} | Delete a Blog Post
[**AttachToLanguageGroup**](BlogPostsAPI.md#AttachToLanguageGroup) | **Post** /cms/v3/blogs/posts/multi-language/attach-to-lang-group | Attach a Blog Post to a multi-language group
[**BatchArchive**](BlogPostsAPI.md#BatchArchive) | **Post** /cms/v3/blogs/posts/batch/archive | Delete a batch of Blog Posts
[**BatchCreate**](BlogPostsAPI.md#BatchCreate) | **Post** /cms/v3/blogs/posts/batch/create | Create a batch of Blog Posts
[**BatchRead**](BlogPostsAPI.md#BatchRead) | **Post** /cms/v3/blogs/posts/batch/read | Retrieve a batch of Blog Posts
[**BatchUpdate**](BlogPostsAPI.md#BatchUpdate) | **Post** /cms/v3/blogs/posts/batch/update | Update a batch of Blog Posts
[**Clone**](BlogPostsAPI.md#Clone) | **Post** /cms/v3/blogs/posts/clone | Clone a Blog Post
[**Create**](BlogPostsAPI.md#Create) | **Post** /cms/v3/blogs/posts | Create a new Blog Post
[**CreateLanguageVariation**](BlogPostsAPI.md#CreateLanguageVariation) | **Post** /cms/v3/blogs/posts/multi-language/create-language-variation | Create a new language variation
[**DetachFromLanguageGroup**](BlogPostsAPI.md#DetachFromLanguageGroup) | **Post** /cms/v3/blogs/posts/multi-language/detach-from-lang-group | Detach a Blog Post from a multi-language group
[**GetByID**](BlogPostsAPI.md#GetByID) | **Get** /cms/v3/blogs/posts/{objectId} | Retrieve a Blog Post
[**GetDraftByID**](BlogPostsAPI.md#GetDraftByID) | **Get** /cms/v3/blogs/posts/{objectId}/draft | Retrieve the full draft version of the Blog Post
[**GetPage**](BlogPostsAPI.md#GetPage) | **Get** /cms/v3/blogs/posts | Get all Blog Posts
[**GetPreviousVersion**](BlogPostsAPI.md#GetPreviousVersion) | **Get** /cms/v3/blogs/posts/{objectId}/revisions/{revisionId} | Retrieves a previous version of a blog post
[**GetPreviousVersions**](BlogPostsAPI.md#GetPreviousVersions) | **Get** /cms/v3/blogs/posts/{objectId}/revisions | Retrieves all the previous versions of a blog post
[**PushLive**](BlogPostsAPI.md#PushLive) | **Post** /cms/v3/blogs/posts/{objectId}/draft/push-live | Push Blog Post draft edits live
[**ResetDraft**](BlogPostsAPI.md#ResetDraft) | **Post** /cms/v3/blogs/posts/{objectId}/draft/reset | Reset the Blog Post draft to the live version
[**RestorePreviousVersion**](BlogPostsAPI.md#RestorePreviousVersion) | **Post** /cms/v3/blogs/posts/{objectId}/revisions/{revisionId}/restore | Restore a previous version of a blog post
[**RestorePreviousVersionToDraft**](BlogPostsAPI.md#RestorePreviousVersionToDraft) | **Post** /cms/v3/blogs/posts/{objectId}/revisions/{revisionId}/restore-to-draft | Restore a previous version of a blog post, to the draft version of the blog post
[**Schedule**](BlogPostsAPI.md#Schedule) | **Post** /cms/v3/blogs/posts/schedule | Schedule a Blog Post to be Published
[**SetLanguagePrimary**](BlogPostsAPI.md#SetLanguagePrimary) | **Put** /cms/v3/blogs/posts/multi-language/set-new-lang-primary | Set a new primary language
[**Update**](BlogPostsAPI.md#Update) | **Patch** /cms/v3/blogs/posts/{objectId} | Update a Blog Post
[**UpdateDraft**](BlogPostsAPI.md#UpdateDraft) | **Patch** /cms/v3/blogs/posts/{objectId}/draft | Update a Blog Post draft
[**UpdateLanguages**](BlogPostsAPI.md#UpdateLanguages) | **Post** /cms/v3/blogs/posts/multi-language/update-languages | Update languages of multi-language group



## Archive

> Archive(ctx, objectId).Archived(archived).Execute()

Delete a Blog Post



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
	objectId := "objectId_example" // string | The Blog Post id.
	archived := true // bool | Whether to return only results that have been archived. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BlogPostsAPI.Archive(context.Background(), objectId).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsAPI.Archive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The Blog Post id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **archived** | **bool** | Whether to return only results that have been archived. | 

### Return type

 (empty response body)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AttachToLanguageGroup

> AttachToLanguageGroup(ctx).AttachToLangPrimaryRequestVNext(attachToLangPrimaryRequestVNext).Execute()

Attach a Blog Post to a multi-language group



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
	attachToLangPrimaryRequestVNext := *openapiclient.NewAttachToLangPrimaryRequestVNext("Language_example", "Id_example", "PrimaryId_example") // AttachToLangPrimaryRequestVNext | The JSON representation of the AttachToLangPrimaryRequest object.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BlogPostsAPI.AttachToLanguageGroup(context.Background()).AttachToLangPrimaryRequestVNext(attachToLangPrimaryRequestVNext).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsAPI.AttachToLanguageGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAttachToLanguageGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **attachToLangPrimaryRequestVNext** | [**AttachToLangPrimaryRequestVNext**](AttachToLangPrimaryRequestVNext.md) | The JSON representation of the AttachToLangPrimaryRequest object. | 

### Return type

 (empty response body)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchArchive

> BatchArchive(ctx).BatchInputString(batchInputString).Execute()

Delete a batch of Blog Posts



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
	batchInputString := *openapiclient.NewBatchInputString([]string{"Inputs_example"}) // BatchInputString | The JSON array of Blog Post ids.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BlogPostsAPI.BatchArchive(context.Background()).BatchInputString(batchInputString).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsAPI.BatchArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputString** | [**BatchInputString**](BatchInputString.md) | The JSON array of Blog Post ids. | 

### Return type

 (empty response body)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchCreate

> BatchResponseBlogPost BatchCreate(ctx).BatchInputBlogPost(batchInputBlogPost).Execute()

Create a batch of Blog Posts



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	batchInputBlogPost := *openapiclient.NewBatchInputBlogPost([]openapiclient.BlogPost{*openapiclient.NewBlogPost(time.Now(), "Language_example", false, "MetaDescription_example", []map[string]map[string]interface{}{map[string]map[string]interface{}{"key": map[string]interface{}(123)}}, "Password_example", "HtmlTitle_example", false, map[string]ContentLanguageVariation{"key": *openapiclient.NewContentLanguageVariation(false, time.Now(), time.Now(), []map[string]interface{}{map[string]interface{}(123)}, "Password_example", "AuthorName_example", false, "Name_example", "Campaign_example", int64(123), "State_example", time.Now(), "Slug_example")}, "Id_example", "State_example", "Slug_example", "CreatedById_example", "RssBody_example", false, false, time.Now(), "ContentTypeCategory_example", "MabExperimentId_example", "UpdatedById_example", "TranslatedFromId_example", "FolderId_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, int64(123), int32(123), "FeaturedImage_example", "AuthorName_example", "Domain_example", "Name_example", "DynamicPageHubDbTableId_example", "Campaign_example", "DynamicPageDataSourceId_example", false, false, map[string]LayoutSection{"key": *openapiclient.NewLayoutSection("CssStyle_example", "Label_example", "Type_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, []map[string]LayoutSection{map[string]LayoutSection{"key": *openapiclient.NewLayoutSection("CssStyle_example", "Label_example", "Type_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, []map[string]LayoutSection{map[string]LayoutSection{"key": }}, []openapiclient.RowMetaData{*openapiclient.NewRowMetaData("CssClass_example", *openapiclient.NewStyles(*openapiclient.NewRGBAColor(float32(123), int32(123), int32(123), int32(123)), "FlexboxPositioning_example", *openapiclient.NewBackgroundImage("ImageUrl_example", "BackgroundSize_example", "BackgroundPosition_example"), false, "VerticalAlignment_example", int32(123), *openapiclient.NewGradient(*openapiclient.NewAngle("Units_example", float32(123)), *openapiclient.NewSideOrCorner("HorizontalSide_example", "VerticalSide_example"), []openapiclient.ColorStop{*openapiclient.NewColorStop(*openapiclient.NewRGBAColor(float32(123), int32(123), int32(123), int32(123)))})))}, []openapiclient.LayoutSection{}, "CssClass_example", int32(123), "CssId_example", int32(123), "Name_example", *openapiclient.NewStyles(, "FlexboxPositioning_example", *openapiclient.NewBackgroundImage("ImageUrl_example", "BackgroundSize_example", "BackgroundPosition_example"), false, "VerticalAlignment_example", int32(123), *openapiclient.NewGradient(*openapiclient.NewAngle("Units_example", float32(123)), *openapiclient.NewSideOrCorner("HorizontalSide_example", "VerticalSide_example"), []openapiclient.ColorStop{*openapiclient.NewColorStop()})))}}, []openapiclient.RowMetaData{*openapiclient.NewRowMetaData("CssClass_example", )}, []openapiclient.LayoutSection{}, "CssClass_example", int32(123), "CssId_example", int32(123), "Name_example", )}, time.Now(), "FooterHtml_example", []int64{int64(123)}, map[string]map[string]interface{}{"key": map[string]interface{}(123)}, "PostSummary_example", "HeadHtml_example", "PageExpiryRedirectUrl_example", "AbStatus_example", false, "AbTestId_example", "FeaturedImageAltText_example", "BlogAuthorId_example", "ContentGroupId_example", "RssSummary_example", false, "Url_example", []map[string]interface{}{map[string]interface{}(123)}, false, int64(123), "PostBody_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, int64(123), false, "CurrentState_example", int32(123), "LinkRelCanonicalUrl_example")}) // BatchInputBlogPost | The JSON array of new Blog Posts to create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogPostsAPI.BatchCreate(context.Background()).BatchInputBlogPost(batchInputBlogPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsAPI.BatchCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchCreate`: BatchResponseBlogPost
	fmt.Fprintf(os.Stdout, "Response from `BlogPostsAPI.BatchCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputBlogPost** | [**BatchInputBlogPost**](BatchInputBlogPost.md) | The JSON array of new Blog Posts to create. | 

### Return type

[**BatchResponseBlogPost**](BatchResponseBlogPost.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchRead

> BatchResponseBlogPost BatchRead(ctx).BatchInputString(batchInputString).Archived(archived).Execute()

Retrieve a batch of Blog Posts



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
	batchInputString := *openapiclient.NewBatchInputString([]string{"Inputs_example"}) // BatchInputString | The JSON array of Blog Post ids.
	archived := true // bool | Specifies whether to return deleted Blog Posts. Defaults to `false`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogPostsAPI.BatchRead(context.Background()).BatchInputString(batchInputString).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsAPI.BatchRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchRead`: BatchResponseBlogPost
	fmt.Fprintf(os.Stdout, "Response from `BlogPostsAPI.BatchRead`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputString** | [**BatchInputString**](BatchInputString.md) | The JSON array of Blog Post ids. | 
 **archived** | **bool** | Specifies whether to return deleted Blog Posts. Defaults to &#x60;false&#x60;. | 

### Return type

[**BatchResponseBlogPost**](BatchResponseBlogPost.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchUpdate

> BatchResponseBlogPost BatchUpdate(ctx).BatchInputJsonNode(batchInputJsonNode).Archived(archived).Execute()

Update a batch of Blog Posts



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
	batchInputJsonNode := *openapiclient.NewBatchInputJsonNode([]map[string]interface{}{map[string]interface{}(123)}) // BatchInputJsonNode | A JSON array of the JSON representations of the updated Blog Posts.
	archived := true // bool | Specifies whether to update deleted Blog Posts. Defaults to `false`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogPostsAPI.BatchUpdate(context.Background()).BatchInputJsonNode(batchInputJsonNode).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsAPI.BatchUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchUpdate`: BatchResponseBlogPost
	fmt.Fprintf(os.Stdout, "Response from `BlogPostsAPI.BatchUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputJsonNode** | [**BatchInputJsonNode**](BatchInputJsonNode.md) | A JSON array of the JSON representations of the updated Blog Posts. | 
 **archived** | **bool** | Specifies whether to update deleted Blog Posts. Defaults to &#x60;false&#x60;. | 

### Return type

[**BatchResponseBlogPost**](BatchResponseBlogPost.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Clone

> BlogPost Clone(ctx).ContentCloneRequestVNext(contentCloneRequestVNext).Execute()

Clone a Blog Post



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
	contentCloneRequestVNext := *openapiclient.NewContentCloneRequestVNext("Id_example") // ContentCloneRequestVNext | The JSON representation of the ContentCloneRequest object.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogPostsAPI.Clone(context.Background()).ContentCloneRequestVNext(contentCloneRequestVNext).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsAPI.Clone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Clone`: BlogPost
	fmt.Fprintf(os.Stdout, "Response from `BlogPostsAPI.Clone`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentCloneRequestVNext** | [**ContentCloneRequestVNext**](ContentCloneRequestVNext.md) | The JSON representation of the ContentCloneRequest object. | 

### Return type

[**BlogPost**](BlogPost.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Create

> BlogPost Create(ctx).BlogPost(blogPost).Execute()

Create a new Blog Post



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	blogPost := *openapiclient.NewBlogPost(time.Now(), "Language_example", false, "MetaDescription_example", []map[string]map[string]interface{}{map[string]map[string]interface{}{"key": map[string]interface{}(123)}}, "Password_example", "HtmlTitle_example", false, map[string]ContentLanguageVariation{"key": *openapiclient.NewContentLanguageVariation(false, time.Now(), time.Now(), []map[string]interface{}{map[string]interface{}(123)}, "Password_example", "AuthorName_example", false, "Name_example", "Campaign_example", int64(123), "State_example", time.Now(), "Slug_example")}, "Id_example", "State_example", "Slug_example", "CreatedById_example", "RssBody_example", false, false, time.Now(), "ContentTypeCategory_example", "MabExperimentId_example", "UpdatedById_example", "TranslatedFromId_example", "FolderId_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, int64(123), int32(123), "FeaturedImage_example", "AuthorName_example", "Domain_example", "Name_example", "DynamicPageHubDbTableId_example", "Campaign_example", "DynamicPageDataSourceId_example", false, false, map[string]LayoutSection{"key": *openapiclient.NewLayoutSection("CssStyle_example", "Label_example", "Type_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, []map[string]LayoutSection{map[string]LayoutSection{"key": *openapiclient.NewLayoutSection("CssStyle_example", "Label_example", "Type_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, []map[string]LayoutSection{map[string]LayoutSection{"key": }}, []openapiclient.RowMetaData{*openapiclient.NewRowMetaData("CssClass_example", *openapiclient.NewStyles(*openapiclient.NewRGBAColor(float32(123), int32(123), int32(123), int32(123)), "FlexboxPositioning_example", *openapiclient.NewBackgroundImage("ImageUrl_example", "BackgroundSize_example", "BackgroundPosition_example"), false, "VerticalAlignment_example", int32(123), *openapiclient.NewGradient(*openapiclient.NewAngle("Units_example", float32(123)), *openapiclient.NewSideOrCorner("HorizontalSide_example", "VerticalSide_example"), []openapiclient.ColorStop{*openapiclient.NewColorStop(*openapiclient.NewRGBAColor(float32(123), int32(123), int32(123), int32(123)))})))}, []openapiclient.LayoutSection{}, "CssClass_example", int32(123), "CssId_example", int32(123), "Name_example", *openapiclient.NewStyles(, "FlexboxPositioning_example", *openapiclient.NewBackgroundImage("ImageUrl_example", "BackgroundSize_example", "BackgroundPosition_example"), false, "VerticalAlignment_example", int32(123), *openapiclient.NewGradient(*openapiclient.NewAngle("Units_example", float32(123)), *openapiclient.NewSideOrCorner("HorizontalSide_example", "VerticalSide_example"), []openapiclient.ColorStop{*openapiclient.NewColorStop()})))}}, []openapiclient.RowMetaData{*openapiclient.NewRowMetaData("CssClass_example", )}, []openapiclient.LayoutSection{}, "CssClass_example", int32(123), "CssId_example", int32(123), "Name_example", )}, time.Now(), "FooterHtml_example", []int64{int64(123)}, map[string]map[string]interface{}{"key": map[string]interface{}(123)}, "PostSummary_example", "HeadHtml_example", "PageExpiryRedirectUrl_example", "AbStatus_example", false, "AbTestId_example", "FeaturedImageAltText_example", "BlogAuthorId_example", "ContentGroupId_example", "RssSummary_example", false, "Url_example", []map[string]interface{}{map[string]interface{}(123)}, false, int64(123), "PostBody_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, int64(123), false, "CurrentState_example", int32(123), "LinkRelCanonicalUrl_example") // BlogPost | The JSON representation of a new Blog Post.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogPostsAPI.Create(context.Background()).BlogPost(blogPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsAPI.Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create`: BlogPost
	fmt.Fprintf(os.Stdout, "Response from `BlogPostsAPI.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blogPost** | [**BlogPost**](BlogPost.md) | The JSON representation of a new Blog Post. | 

### Return type

[**BlogPost**](BlogPost.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLanguageVariation

> BlogPost CreateLanguageVariation(ctx).BlogPostLanguageCloneRequestVNext(blogPostLanguageCloneRequestVNext).Execute()

Create a new language variation



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
	blogPostLanguageCloneRequestVNext := *openapiclient.NewBlogPostLanguageCloneRequestVNext("Id_example") // BlogPostLanguageCloneRequestVNext | The JSON representation of the BlogPostLanguageCloneRequestVNext object.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogPostsAPI.CreateLanguageVariation(context.Background()).BlogPostLanguageCloneRequestVNext(blogPostLanguageCloneRequestVNext).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsAPI.CreateLanguageVariation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLanguageVariation`: BlogPost
	fmt.Fprintf(os.Stdout, "Response from `BlogPostsAPI.CreateLanguageVariation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLanguageVariationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blogPostLanguageCloneRequestVNext** | [**BlogPostLanguageCloneRequestVNext**](BlogPostLanguageCloneRequestVNext.md) | The JSON representation of the BlogPostLanguageCloneRequestVNext object. | 

### Return type

[**BlogPost**](BlogPost.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DetachFromLanguageGroup

> DetachFromLanguageGroup(ctx).DetachFromLangGroupRequestVNext(detachFromLangGroupRequestVNext).Execute()

Detach a Blog Post from a multi-language group



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
	detachFromLangGroupRequestVNext := *openapiclient.NewDetachFromLangGroupRequestVNext("Id_example") // DetachFromLangGroupRequestVNext | The JSON representation of the DetachFromLangGroupRequest object.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BlogPostsAPI.DetachFromLanguageGroup(context.Background()).DetachFromLangGroupRequestVNext(detachFromLangGroupRequestVNext).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsAPI.DetachFromLanguageGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDetachFromLanguageGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **detachFromLangGroupRequestVNext** | [**DetachFromLangGroupRequestVNext**](DetachFromLangGroupRequestVNext.md) | The JSON representation of the DetachFromLangGroupRequest object. | 

### Return type

 (empty response body)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetByID

> BlogPost GetByID(ctx, objectId).Archived(archived).Property(property).Execute()

Retrieve a Blog Post



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
	objectId := "objectId_example" // string | The Blog Post id.
	archived := true // bool | Specifies whether to return deleted Blog Posts. Defaults to `false`. (optional)
	property := "property_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogPostsAPI.GetByID(context.Background(), objectId).Archived(archived).Property(property).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsAPI.GetByID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetByID`: BlogPost
	fmt.Fprintf(os.Stdout, "Response from `BlogPostsAPI.GetByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The Blog Post id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **archived** | **bool** | Specifies whether to return deleted Blog Posts. Defaults to &#x60;false&#x60;. | 
 **property** | **string** |  | 

### Return type

[**BlogPost**](BlogPost.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDraftByID

> BlogPost GetDraftByID(ctx, objectId).Execute()

Retrieve the full draft version of the Blog Post



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
	objectId := "objectId_example" // string | The Blog Post id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogPostsAPI.GetDraftByID(context.Background(), objectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsAPI.GetDraftByID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDraftByID`: BlogPost
	fmt.Fprintf(os.Stdout, "Response from `BlogPostsAPI.GetDraftByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The Blog Post id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDraftByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BlogPost**](BlogPost.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPage

> CollectionResponseWithTotalBlogPostForwardPaging GetPage(ctx).CreatedAt(createdAt).CreatedAfter(createdAfter).CreatedBefore(createdBefore).UpdatedAt(updatedAt).UpdatedAfter(updatedAfter).UpdatedBefore(updatedBefore).Sort(sort).After(after).Limit(limit).Archived(archived).Property(property).Execute()

Get all Blog Posts



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	createdAt := time.Now() // time.Time | Only return Blog Posts created at exactly the specified time. (optional)
	createdAfter := time.Now() // time.Time | Only return Blog Posts created after the specified time. (optional)
	createdBefore := time.Now() // time.Time | Only return Blog Posts created before the specified time. (optional)
	updatedAt := time.Now() // time.Time | Only return Blog Posts last updated at exactly the specified time. (optional)
	updatedAfter := time.Now() // time.Time | Only return Blog Posts last updated after the specified time. (optional)
	updatedBefore := time.Now() // time.Time | Only return Blog Posts last updated before the specified time. (optional)
	sort := []string{"Inner_example"} // []string | Specifies which fields to use for sorting results. Valid fields are `name`, `createdAt`, `updatedAt`, `createdBy`, `updatedBy`. `createdAt` will be used by default. (optional)
	after := "after_example" // string | The cursor token value to get the next set of results. You can get this from the `paging.next.after` JSON property of a paged response containing more results. (optional)
	limit := int32(56) // int32 | The maximum number of results to return. Default is 20. (optional)
	archived := true // bool | Specifies whether to return deleted Blog Posts. Defaults to `false`. (optional)
	property := "property_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogPostsAPI.GetPage(context.Background()).CreatedAt(createdAt).CreatedAfter(createdAfter).CreatedBefore(createdBefore).UpdatedAt(updatedAt).UpdatedAfter(updatedAfter).UpdatedBefore(updatedBefore).Sort(sort).After(after).Limit(limit).Archived(archived).Property(property).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsAPI.GetPage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPage`: CollectionResponseWithTotalBlogPostForwardPaging
	fmt.Fprintf(os.Stdout, "Response from `BlogPostsAPI.GetPage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createdAt** | **time.Time** | Only return Blog Posts created at exactly the specified time. | 
 **createdAfter** | **time.Time** | Only return Blog Posts created after the specified time. | 
 **createdBefore** | **time.Time** | Only return Blog Posts created before the specified time. | 
 **updatedAt** | **time.Time** | Only return Blog Posts last updated at exactly the specified time. | 
 **updatedAfter** | **time.Time** | Only return Blog Posts last updated after the specified time. | 
 **updatedBefore** | **time.Time** | Only return Blog Posts last updated before the specified time. | 
 **sort** | **[]string** | Specifies which fields to use for sorting results. Valid fields are &#x60;name&#x60;, &#x60;createdAt&#x60;, &#x60;updatedAt&#x60;, &#x60;createdBy&#x60;, &#x60;updatedBy&#x60;. &#x60;createdAt&#x60; will be used by default. | 
 **after** | **string** | The cursor token value to get the next set of results. You can get this from the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **limit** | **int32** | The maximum number of results to return. Default is 20. | 
 **archived** | **bool** | Specifies whether to return deleted Blog Posts. Defaults to &#x60;false&#x60;. | 
 **property** | **string** |  | 

### Return type

[**CollectionResponseWithTotalBlogPostForwardPaging**](CollectionResponseWithTotalBlogPostForwardPaging.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPreviousVersion

> VersionBlogPost GetPreviousVersion(ctx, objectId, revisionId).Execute()

Retrieves a previous version of a blog post



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
	objectId := "objectId_example" // string | The Blog Post id.
	revisionId := "revisionId_example" // string | The Blog Post version id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogPostsAPI.GetPreviousVersion(context.Background(), objectId, revisionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsAPI.GetPreviousVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPreviousVersion`: VersionBlogPost
	fmt.Fprintf(os.Stdout, "Response from `BlogPostsAPI.GetPreviousVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The Blog Post id. | 
**revisionId** | **string** | The Blog Post version id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPreviousVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**VersionBlogPost**](VersionBlogPost.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPreviousVersions

> CollectionResponseWithTotalVersionBlogPost GetPreviousVersions(ctx, objectId).After(after).Before(before).Limit(limit).Execute()

Retrieves all the previous versions of a blog post



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
	objectId := "objectId_example" // string | The Blog Post id.
	after := "after_example" // string | The cursor token value to get the next set of results. You can get this from the `paging.next.after` JSON property of a paged response containing more results. (optional)
	before := "before_example" // string |  (optional)
	limit := int32(56) // int32 | The maximum number of results to return. Default is 100. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogPostsAPI.GetPreviousVersions(context.Background(), objectId).After(after).Before(before).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsAPI.GetPreviousVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPreviousVersions`: CollectionResponseWithTotalVersionBlogPost
	fmt.Fprintf(os.Stdout, "Response from `BlogPostsAPI.GetPreviousVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The Blog Post id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPreviousVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **after** | **string** | The cursor token value to get the next set of results. You can get this from the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **before** | **string** |  | 
 **limit** | **int32** | The maximum number of results to return. Default is 100. | 

### Return type

[**CollectionResponseWithTotalVersionBlogPost**](CollectionResponseWithTotalVersionBlogPost.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PushLive

> PushLive(ctx, objectId).Execute()

Push Blog Post draft edits live



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
	objectId := "objectId_example" // string | The id of the Blog Post for which it's draft will be pushed live.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BlogPostsAPI.PushLive(context.Background(), objectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsAPI.PushLive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The id of the Blog Post for which it&#39;s draft will be pushed live. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPushLiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetDraft

> ResetDraft(ctx, objectId).Execute()

Reset the Blog Post draft to the live version



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
	objectId := "objectId_example" // string | The id of the Blog Post for which it's draft will be reset.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BlogPostsAPI.ResetDraft(context.Background(), objectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsAPI.ResetDraft``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The id of the Blog Post for which it&#39;s draft will be reset. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestorePreviousVersion

> BlogPost RestorePreviousVersion(ctx, objectId, revisionId).Execute()

Restore a previous version of a blog post



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
	objectId := "objectId_example" // string | The Blog Post id.
	revisionId := "revisionId_example" // string | The Blog Post version id to restore.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogPostsAPI.RestorePreviousVersion(context.Background(), objectId, revisionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsAPI.RestorePreviousVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RestorePreviousVersion`: BlogPost
	fmt.Fprintf(os.Stdout, "Response from `BlogPostsAPI.RestorePreviousVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The Blog Post id. | 
**revisionId** | **string** | The Blog Post version id to restore. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestorePreviousVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BlogPost**](BlogPost.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestorePreviousVersionToDraft

> BlogPost RestorePreviousVersionToDraft(ctx, objectId, revisionId).Execute()

Restore a previous version of a blog post, to the draft version of the blog post



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
	objectId := "objectId_example" // string | The Blog Post id.
	revisionId := int64(789) // int64 | The Blog Post version id to restore.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogPostsAPI.RestorePreviousVersionToDraft(context.Background(), objectId, revisionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsAPI.RestorePreviousVersionToDraft``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RestorePreviousVersionToDraft`: BlogPost
	fmt.Fprintf(os.Stdout, "Response from `BlogPostsAPI.RestorePreviousVersionToDraft`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The Blog Post id. | 
**revisionId** | **int64** | The Blog Post version id to restore. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestorePreviousVersionToDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BlogPost**](BlogPost.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Schedule

> Schedule(ctx).ContentScheduleRequestVNext(contentScheduleRequestVNext).Execute()

Schedule a Blog Post to be Published



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	contentScheduleRequestVNext := *openapiclient.NewContentScheduleRequestVNext(time.Now(), "Id_example") // ContentScheduleRequestVNext | The JSON representation of the ContentScheduleRequestVNext object.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BlogPostsAPI.Schedule(context.Background()).ContentScheduleRequestVNext(contentScheduleRequestVNext).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsAPI.Schedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentScheduleRequestVNext** | [**ContentScheduleRequestVNext**](ContentScheduleRequestVNext.md) | The JSON representation of the ContentScheduleRequestVNext object. | 

### Return type

 (empty response body)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetLanguagePrimary

> SetLanguagePrimary(ctx).SetNewLanguagePrimaryRequestVNext(setNewLanguagePrimaryRequestVNext).Execute()

Set a new primary language



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
	setNewLanguagePrimaryRequestVNext := *openapiclient.NewSetNewLanguagePrimaryRequestVNext("Id_example") // SetNewLanguagePrimaryRequestVNext | The JSON representation of the SetNewLanguagePrimaryRequest object.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BlogPostsAPI.SetLanguagePrimary(context.Background()).SetNewLanguagePrimaryRequestVNext(setNewLanguagePrimaryRequestVNext).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsAPI.SetLanguagePrimary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetLanguagePrimaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setNewLanguagePrimaryRequestVNext** | [**SetNewLanguagePrimaryRequestVNext**](SetNewLanguagePrimaryRequestVNext.md) | The JSON representation of the SetNewLanguagePrimaryRequest object. | 

### Return type

 (empty response body)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> BlogPost Update(ctx, objectId).BlogPost(blogPost).Archived(archived).Execute()

Update a Blog Post



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	objectId := "objectId_example" // string | The Blog Post id.
	blogPost := *openapiclient.NewBlogPost(time.Now(), "Language_example", false, "MetaDescription_example", []map[string]map[string]interface{}{map[string]map[string]interface{}{"key": map[string]interface{}(123)}}, "Password_example", "HtmlTitle_example", false, map[string]ContentLanguageVariation{"key": *openapiclient.NewContentLanguageVariation(false, time.Now(), time.Now(), []map[string]interface{}{map[string]interface{}(123)}, "Password_example", "AuthorName_example", false, "Name_example", "Campaign_example", int64(123), "State_example", time.Now(), "Slug_example")}, "Id_example", "State_example", "Slug_example", "CreatedById_example", "RssBody_example", false, false, time.Now(), "ContentTypeCategory_example", "MabExperimentId_example", "UpdatedById_example", "TranslatedFromId_example", "FolderId_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, int64(123), int32(123), "FeaturedImage_example", "AuthorName_example", "Domain_example", "Name_example", "DynamicPageHubDbTableId_example", "Campaign_example", "DynamicPageDataSourceId_example", false, false, map[string]LayoutSection{"key": *openapiclient.NewLayoutSection("CssStyle_example", "Label_example", "Type_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, []map[string]LayoutSection{map[string]LayoutSection{"key": *openapiclient.NewLayoutSection("CssStyle_example", "Label_example", "Type_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, []map[string]LayoutSection{map[string]LayoutSection{"key": }}, []openapiclient.RowMetaData{*openapiclient.NewRowMetaData("CssClass_example", *openapiclient.NewStyles(*openapiclient.NewRGBAColor(float32(123), int32(123), int32(123), int32(123)), "FlexboxPositioning_example", *openapiclient.NewBackgroundImage("ImageUrl_example", "BackgroundSize_example", "BackgroundPosition_example"), false, "VerticalAlignment_example", int32(123), *openapiclient.NewGradient(*openapiclient.NewAngle("Units_example", float32(123)), *openapiclient.NewSideOrCorner("HorizontalSide_example", "VerticalSide_example"), []openapiclient.ColorStop{*openapiclient.NewColorStop(*openapiclient.NewRGBAColor(float32(123), int32(123), int32(123), int32(123)))})))}, []openapiclient.LayoutSection{}, "CssClass_example", int32(123), "CssId_example", int32(123), "Name_example", *openapiclient.NewStyles(, "FlexboxPositioning_example", *openapiclient.NewBackgroundImage("ImageUrl_example", "BackgroundSize_example", "BackgroundPosition_example"), false, "VerticalAlignment_example", int32(123), *openapiclient.NewGradient(*openapiclient.NewAngle("Units_example", float32(123)), *openapiclient.NewSideOrCorner("HorizontalSide_example", "VerticalSide_example"), []openapiclient.ColorStop{*openapiclient.NewColorStop()})))}}, []openapiclient.RowMetaData{*openapiclient.NewRowMetaData("CssClass_example", )}, []openapiclient.LayoutSection{}, "CssClass_example", int32(123), "CssId_example", int32(123), "Name_example", )}, time.Now(), "FooterHtml_example", []int64{int64(123)}, map[string]map[string]interface{}{"key": map[string]interface{}(123)}, "PostSummary_example", "HeadHtml_example", "PageExpiryRedirectUrl_example", "AbStatus_example", false, "AbTestId_example", "FeaturedImageAltText_example", "BlogAuthorId_example", "ContentGroupId_example", "RssSummary_example", false, "Url_example", []map[string]interface{}{map[string]interface{}(123)}, false, int64(123), "PostBody_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, int64(123), false, "CurrentState_example", int32(123), "LinkRelCanonicalUrl_example") // BlogPost | The JSON representation of the updated Blog Post.
	archived := true // bool | Specifies whether to update deleted Blog Posts. Defaults to `false`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogPostsAPI.Update(context.Background(), objectId).BlogPost(blogPost).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsAPI.Update``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update`: BlogPost
	fmt.Fprintf(os.Stdout, "Response from `BlogPostsAPI.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The Blog Post id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **blogPost** | [**BlogPost**](BlogPost.md) | The JSON representation of the updated Blog Post. | 
 **archived** | **bool** | Specifies whether to update deleted Blog Posts. Defaults to &#x60;false&#x60;. | 

### Return type

[**BlogPost**](BlogPost.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDraft

> BlogPost UpdateDraft(ctx, objectId).BlogPost(blogPost).Execute()

Update a Blog Post draft



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	objectId := "objectId_example" // string | The Blog Post id.
	blogPost := *openapiclient.NewBlogPost(time.Now(), "Language_example", false, "MetaDescription_example", []map[string]map[string]interface{}{map[string]map[string]interface{}{"key": map[string]interface{}(123)}}, "Password_example", "HtmlTitle_example", false, map[string]ContentLanguageVariation{"key": *openapiclient.NewContentLanguageVariation(false, time.Now(), time.Now(), []map[string]interface{}{map[string]interface{}(123)}, "Password_example", "AuthorName_example", false, "Name_example", "Campaign_example", int64(123), "State_example", time.Now(), "Slug_example")}, "Id_example", "State_example", "Slug_example", "CreatedById_example", "RssBody_example", false, false, time.Now(), "ContentTypeCategory_example", "MabExperimentId_example", "UpdatedById_example", "TranslatedFromId_example", "FolderId_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, int64(123), int32(123), "FeaturedImage_example", "AuthorName_example", "Domain_example", "Name_example", "DynamicPageHubDbTableId_example", "Campaign_example", "DynamicPageDataSourceId_example", false, false, map[string]LayoutSection{"key": *openapiclient.NewLayoutSection("CssStyle_example", "Label_example", "Type_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, []map[string]LayoutSection{map[string]LayoutSection{"key": *openapiclient.NewLayoutSection("CssStyle_example", "Label_example", "Type_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, []map[string]LayoutSection{map[string]LayoutSection{"key": }}, []openapiclient.RowMetaData{*openapiclient.NewRowMetaData("CssClass_example", *openapiclient.NewStyles(*openapiclient.NewRGBAColor(float32(123), int32(123), int32(123), int32(123)), "FlexboxPositioning_example", *openapiclient.NewBackgroundImage("ImageUrl_example", "BackgroundSize_example", "BackgroundPosition_example"), false, "VerticalAlignment_example", int32(123), *openapiclient.NewGradient(*openapiclient.NewAngle("Units_example", float32(123)), *openapiclient.NewSideOrCorner("HorizontalSide_example", "VerticalSide_example"), []openapiclient.ColorStop{*openapiclient.NewColorStop(*openapiclient.NewRGBAColor(float32(123), int32(123), int32(123), int32(123)))})))}, []openapiclient.LayoutSection{}, "CssClass_example", int32(123), "CssId_example", int32(123), "Name_example", *openapiclient.NewStyles(, "FlexboxPositioning_example", *openapiclient.NewBackgroundImage("ImageUrl_example", "BackgroundSize_example", "BackgroundPosition_example"), false, "VerticalAlignment_example", int32(123), *openapiclient.NewGradient(*openapiclient.NewAngle("Units_example", float32(123)), *openapiclient.NewSideOrCorner("HorizontalSide_example", "VerticalSide_example"), []openapiclient.ColorStop{*openapiclient.NewColorStop()})))}}, []openapiclient.RowMetaData{*openapiclient.NewRowMetaData("CssClass_example", )}, []openapiclient.LayoutSection{}, "CssClass_example", int32(123), "CssId_example", int32(123), "Name_example", )}, time.Now(), "FooterHtml_example", []int64{int64(123)}, map[string]map[string]interface{}{"key": map[string]interface{}(123)}, "PostSummary_example", "HeadHtml_example", "PageExpiryRedirectUrl_example", "AbStatus_example", false, "AbTestId_example", "FeaturedImageAltText_example", "BlogAuthorId_example", "ContentGroupId_example", "RssSummary_example", false, "Url_example", []map[string]interface{}{map[string]interface{}(123)}, false, int64(123), "PostBody_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, int64(123), false, "CurrentState_example", int32(123), "LinkRelCanonicalUrl_example") // BlogPost | The JSON representation of the updated Blog Post to be applied to the draft.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogPostsAPI.UpdateDraft(context.Background(), objectId).BlogPost(blogPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsAPI.UpdateDraft``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDraft`: BlogPost
	fmt.Fprintf(os.Stdout, "Response from `BlogPostsAPI.UpdateDraft`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The Blog Post id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **blogPost** | [**BlogPost**](BlogPost.md) | The JSON representation of the updated Blog Post to be applied to the draft. | 

### Return type

[**BlogPost**](BlogPost.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLanguages

> UpdateLanguages(ctx).UpdateLanguagesRequestVNext(updateLanguagesRequestVNext).Execute()

Update languages of multi-language group



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
	updateLanguagesRequestVNext := *openapiclient.NewUpdateLanguagesRequestVNext(map[string]string{"key": "Inner_example"}, "PrimaryId_example") // UpdateLanguagesRequestVNext | The JSON representation of the SetNewLanguagePrimaryRequest object.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BlogPostsAPI.UpdateLanguages(context.Background()).UpdateLanguagesRequestVNext(updateLanguagesRequestVNext).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsAPI.UpdateLanguages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLanguagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateLanguagesRequestVNext** | [**UpdateLanguagesRequestVNext**](UpdateLanguagesRequestVNext.md) | The JSON representation of the SetNewLanguagePrimaryRequest object. | 

### Return type

 (empty response body)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

