# Column

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForeignTableId** | Pointer to **int64** | Foreign table id referenced | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Label** | **string** | Label of the column | 
**Type** | **string** | Type of the column | 
**OptionCount** | Pointer to **int32** | Number of options available | [optional] 
**ForeignIds** | Pointer to [**[]ForeignId**](ForeignId.md) | Foreign Ids | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**Name** | **string** | Name of the column | 
**Options** | Pointer to [**[]Option**](Option.md) | Options to choose for select and multi-select columns | [optional] 
**Width** | Pointer to **int32** | Column width for HubDB UI | [optional] 
**Id** | Pointer to **string** | Column Id | [optional] 
**ForeignIdsById** | Pointer to [**map[string]ForeignId**](ForeignId.md) | Foreign ids | [optional] 
**ForeignColumnId** | Pointer to **int32** | Foreign Column id | [optional] 
**ForeignIdsByName** | Pointer to [**map[string]ForeignId**](ForeignId.md) | Foreign ids by name | [optional] 

## Methods

### NewColumn

`func NewColumn(label string, type_ string, name string, ) *Column`

NewColumn instantiates a new Column object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewColumnWithDefaults

`func NewColumnWithDefaults() *Column`

NewColumnWithDefaults instantiates a new Column object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForeignTableId

`func (o *Column) GetForeignTableId() int64`

GetForeignTableId returns the ForeignTableId field if non-nil, zero value otherwise.

### GetForeignTableIdOk

`func (o *Column) GetForeignTableIdOk() (*int64, bool)`

GetForeignTableIdOk returns a tuple with the ForeignTableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignTableId

`func (o *Column) SetForeignTableId(v int64)`

SetForeignTableId sets ForeignTableId field to given value.

### HasForeignTableId

`func (o *Column) HasForeignTableId() bool`

HasForeignTableId returns a boolean if a field has been set.

### GetDescription

`func (o *Column) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Column) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Column) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Column) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabel

`func (o *Column) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Column) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Column) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetType

`func (o *Column) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Column) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Column) SetType(v string)`

SetType sets Type field to given value.


### GetOptionCount

`func (o *Column) GetOptionCount() int32`

GetOptionCount returns the OptionCount field if non-nil, zero value otherwise.

### GetOptionCountOk

`func (o *Column) GetOptionCountOk() (*int32, bool)`

GetOptionCountOk returns a tuple with the OptionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionCount

`func (o *Column) SetOptionCount(v int32)`

SetOptionCount sets OptionCount field to given value.

### HasOptionCount

`func (o *Column) HasOptionCount() bool`

HasOptionCount returns a boolean if a field has been set.

### GetForeignIds

`func (o *Column) GetForeignIds() []ForeignId`

GetForeignIds returns the ForeignIds field if non-nil, zero value otherwise.

### GetForeignIdsOk

`func (o *Column) GetForeignIdsOk() (*[]ForeignId, bool)`

GetForeignIdsOk returns a tuple with the ForeignIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignIds

`func (o *Column) SetForeignIds(v []ForeignId)`

SetForeignIds sets ForeignIds field to given value.

### HasForeignIds

`func (o *Column) HasForeignIds() bool`

HasForeignIds returns a boolean if a field has been set.

### GetDeleted

`func (o *Column) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Column) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Column) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *Column) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetName

`func (o *Column) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Column) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Column) SetName(v string)`

SetName sets Name field to given value.


### GetOptions

`func (o *Column) GetOptions() []Option`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *Column) GetOptionsOk() (*[]Option, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *Column) SetOptions(v []Option)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *Column) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetWidth

`func (o *Column) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *Column) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *Column) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *Column) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetId

`func (o *Column) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Column) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Column) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Column) HasId() bool`

HasId returns a boolean if a field has been set.

### GetForeignIdsById

`func (o *Column) GetForeignIdsById() map[string]ForeignId`

GetForeignIdsById returns the ForeignIdsById field if non-nil, zero value otherwise.

### GetForeignIdsByIdOk

`func (o *Column) GetForeignIdsByIdOk() (*map[string]ForeignId, bool)`

GetForeignIdsByIdOk returns a tuple with the ForeignIdsById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignIdsById

`func (o *Column) SetForeignIdsById(v map[string]ForeignId)`

SetForeignIdsById sets ForeignIdsById field to given value.

### HasForeignIdsById

`func (o *Column) HasForeignIdsById() bool`

HasForeignIdsById returns a boolean if a field has been set.

### GetForeignColumnId

`func (o *Column) GetForeignColumnId() int32`

GetForeignColumnId returns the ForeignColumnId field if non-nil, zero value otherwise.

### GetForeignColumnIdOk

`func (o *Column) GetForeignColumnIdOk() (*int32, bool)`

GetForeignColumnIdOk returns a tuple with the ForeignColumnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignColumnId

`func (o *Column) SetForeignColumnId(v int32)`

SetForeignColumnId sets ForeignColumnId field to given value.

### HasForeignColumnId

`func (o *Column) HasForeignColumnId() bool`

HasForeignColumnId returns a boolean if a field has been set.

### GetForeignIdsByName

`func (o *Column) GetForeignIdsByName() map[string]ForeignId`

GetForeignIdsByName returns the ForeignIdsByName field if non-nil, zero value otherwise.

### GetForeignIdsByNameOk

`func (o *Column) GetForeignIdsByNameOk() (*map[string]ForeignId, bool)`

GetForeignIdsByNameOk returns a tuple with the ForeignIdsByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignIdsByName

`func (o *Column) SetForeignIdsByName(v map[string]ForeignId)`

SetForeignIdsByName sets ForeignIdsByName field to given value.

### HasForeignIdsByName

`func (o *Column) HasForeignIdsByName() bool`

HasForeignIdsByName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


