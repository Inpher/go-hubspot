# CollectionResponsePublicImportResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Paging** | Pointer to [**Paging**](Paging.md) |  | [optional] 
**Results** | [**[]PublicImportResponse**](PublicImportResponse.md) |  | 

## Methods

### NewCollectionResponsePublicImportResponse

`func NewCollectionResponsePublicImportResponse(results []PublicImportResponse, ) *CollectionResponsePublicImportResponse`

NewCollectionResponsePublicImportResponse instantiates a new CollectionResponsePublicImportResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionResponsePublicImportResponseWithDefaults

`func NewCollectionResponsePublicImportResponseWithDefaults() *CollectionResponsePublicImportResponse`

NewCollectionResponsePublicImportResponseWithDefaults instantiates a new CollectionResponsePublicImportResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaging

`func (o *CollectionResponsePublicImportResponse) GetPaging() Paging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *CollectionResponsePublicImportResponse) GetPagingOk() (*Paging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *CollectionResponsePublicImportResponse) SetPaging(v Paging)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *CollectionResponsePublicImportResponse) HasPaging() bool`

HasPaging returns a boolean if a field has been set.

### GetResults

`func (o *CollectionResponsePublicImportResponse) GetResults() []PublicImportResponse`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *CollectionResponsePublicImportResponse) GetResultsOk() (*[]PublicImportResponse, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *CollectionResponsePublicImportResponse) SetResults(v []PublicImportResponse)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


