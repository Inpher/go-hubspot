# BatchResponseMarketingEventPublicDefaultResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletedAt** | **time.Time** |  | 
**NumErrors** | Pointer to **int32** |  | [optional] 
**RequestedAt** | Pointer to **time.Time** |  | [optional] 
**StartedAt** | **time.Time** |  | 
**Links** | Pointer to **map[string]string** |  | [optional] 
**Results** | [**[]MarketingEventPublicDefaultResponse**](MarketingEventPublicDefaultResponse.md) |  | 
**Errors** | Pointer to [**[]StandardError**](StandardError.md) |  | [optional] 
**Status** | **string** |  | 

## Methods

### NewBatchResponseMarketingEventPublicDefaultResponse

`func NewBatchResponseMarketingEventPublicDefaultResponse(completedAt time.Time, startedAt time.Time, results []MarketingEventPublicDefaultResponse, status string, ) *BatchResponseMarketingEventPublicDefaultResponse`

NewBatchResponseMarketingEventPublicDefaultResponse instantiates a new BatchResponseMarketingEventPublicDefaultResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchResponseMarketingEventPublicDefaultResponseWithDefaults

`func NewBatchResponseMarketingEventPublicDefaultResponseWithDefaults() *BatchResponseMarketingEventPublicDefaultResponse`

NewBatchResponseMarketingEventPublicDefaultResponseWithDefaults instantiates a new BatchResponseMarketingEventPublicDefaultResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletedAt

`func (o *BatchResponseMarketingEventPublicDefaultResponse) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *BatchResponseMarketingEventPublicDefaultResponse) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *BatchResponseMarketingEventPublicDefaultResponse) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.


### GetNumErrors

`func (o *BatchResponseMarketingEventPublicDefaultResponse) GetNumErrors() int32`

GetNumErrors returns the NumErrors field if non-nil, zero value otherwise.

### GetNumErrorsOk

`func (o *BatchResponseMarketingEventPublicDefaultResponse) GetNumErrorsOk() (*int32, bool)`

GetNumErrorsOk returns a tuple with the NumErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumErrors

`func (o *BatchResponseMarketingEventPublicDefaultResponse) SetNumErrors(v int32)`

SetNumErrors sets NumErrors field to given value.

### HasNumErrors

`func (o *BatchResponseMarketingEventPublicDefaultResponse) HasNumErrors() bool`

HasNumErrors returns a boolean if a field has been set.

### GetRequestedAt

`func (o *BatchResponseMarketingEventPublicDefaultResponse) GetRequestedAt() time.Time`

GetRequestedAt returns the RequestedAt field if non-nil, zero value otherwise.

### GetRequestedAtOk

`func (o *BatchResponseMarketingEventPublicDefaultResponse) GetRequestedAtOk() (*time.Time, bool)`

GetRequestedAtOk returns a tuple with the RequestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedAt

`func (o *BatchResponseMarketingEventPublicDefaultResponse) SetRequestedAt(v time.Time)`

SetRequestedAt sets RequestedAt field to given value.

### HasRequestedAt

`func (o *BatchResponseMarketingEventPublicDefaultResponse) HasRequestedAt() bool`

HasRequestedAt returns a boolean if a field has been set.

### GetStartedAt

`func (o *BatchResponseMarketingEventPublicDefaultResponse) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *BatchResponseMarketingEventPublicDefaultResponse) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *BatchResponseMarketingEventPublicDefaultResponse) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetLinks

`func (o *BatchResponseMarketingEventPublicDefaultResponse) GetLinks() map[string]string`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BatchResponseMarketingEventPublicDefaultResponse) GetLinksOk() (*map[string]string, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BatchResponseMarketingEventPublicDefaultResponse) SetLinks(v map[string]string)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *BatchResponseMarketingEventPublicDefaultResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetResults

`func (o *BatchResponseMarketingEventPublicDefaultResponse) GetResults() []MarketingEventPublicDefaultResponse`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *BatchResponseMarketingEventPublicDefaultResponse) GetResultsOk() (*[]MarketingEventPublicDefaultResponse, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *BatchResponseMarketingEventPublicDefaultResponse) SetResults(v []MarketingEventPublicDefaultResponse)`

SetResults sets Results field to given value.


### GetErrors

`func (o *BatchResponseMarketingEventPublicDefaultResponse) GetErrors() []StandardError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *BatchResponseMarketingEventPublicDefaultResponse) GetErrorsOk() (*[]StandardError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *BatchResponseMarketingEventPublicDefaultResponse) SetErrors(v []StandardError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *BatchResponseMarketingEventPublicDefaultResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetStatus

`func (o *BatchResponseMarketingEventPublicDefaultResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BatchResponseMarketingEventPublicDefaultResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BatchResponseMarketingEventPublicDefaultResponse) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


