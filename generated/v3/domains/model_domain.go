/*
CMS Domains

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package domains

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// checks if the Domain type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Domain{}

// Domain struct for Domain
type Domain struct {
	// Whether the domain is used for CMS landing pages.
	IsUsedForLandingPage      bool       `json:"isUsedForLandingPage"`
	PrimaryBlogPost           *bool      `json:"primaryBlogPost,omitempty"`
	PrimaryKnowledge          *bool      `json:"primaryKnowledge,omitempty"`
	Created                   *time.Time `json:"created,omitempty"`
	SecondaryToDomain         *string    `json:"secondaryToDomain,omitempty"`
	ManuallyMarkedAsResolving *bool      `json:"manuallyMarkedAsResolving,omitempty"`
	// Whether the domain is used for CMS knowledge pages.
	IsUsedForKnowledge bool `json:"isUsedForKnowledge"`
	// Whether the domain is used for CMS blog posts.
	IsUsedForBlogPost bool `json:"isUsedForBlogPost"`
	// Whether the domain is used for CMS site pages.
	IsUsedForSitePage bool `json:"isUsedForSitePage"`
	// Whether the DNS for this domain is optimally configured for use with HubSpot.
	IsResolving  bool  `json:"isResolving"`
	IsSslEnabled *bool `json:"isSslEnabled,omitempty"`
	// Whether the domain is used for CMS email web pages.
	IsUsedForEmail bool `json:"isUsedForEmail"`
	// The actual domain or sub-domain. e.g. www.hubspot.com
	Domain             string `json:"domain"`
	PrimarySitePage    *bool  `json:"primarySitePage,omitempty"`
	PrimaryLandingPage *bool  `json:"primaryLandingPage,omitempty"`
	// The unique ID of this domain.
	Id           string     `json:"id"`
	CorrectCname *string    `json:"correctCname,omitempty"`
	IsSslOnly    *bool      `json:"isSslOnly,omitempty"`
	Updated      *time.Time `json:"updated,omitempty"`
	PrimaryEmail *bool      `json:"primaryEmail,omitempty"`
}

type _Domain Domain

// NewDomain instantiates a new Domain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomain(isUsedForLandingPage bool, isUsedForKnowledge bool, isUsedForBlogPost bool, isUsedForSitePage bool, isResolving bool, isUsedForEmail bool, domain string, id string) *Domain {
	this := Domain{}
	this.IsUsedForLandingPage = isUsedForLandingPage
	this.IsUsedForKnowledge = isUsedForKnowledge
	this.IsUsedForBlogPost = isUsedForBlogPost
	this.IsUsedForSitePage = isUsedForSitePage
	this.IsResolving = isResolving
	this.IsUsedForEmail = isUsedForEmail
	this.Domain = domain
	this.Id = id
	return &this
}

// NewDomainWithDefaults instantiates a new Domain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainWithDefaults() *Domain {
	this := Domain{}
	return &this
}

// GetIsUsedForLandingPage returns the IsUsedForLandingPage field value
func (o *Domain) GetIsUsedForLandingPage() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsUsedForLandingPage
}

// GetIsUsedForLandingPageOk returns a tuple with the IsUsedForLandingPage field value
// and a boolean to check if the value has been set.
func (o *Domain) GetIsUsedForLandingPageOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsUsedForLandingPage, true
}

// SetIsUsedForLandingPage sets field value
func (o *Domain) SetIsUsedForLandingPage(v bool) {
	o.IsUsedForLandingPage = v
}

// GetPrimaryBlogPost returns the PrimaryBlogPost field value if set, zero value otherwise.
func (o *Domain) GetPrimaryBlogPost() bool {
	if o == nil || IsNil(o.PrimaryBlogPost) {
		var ret bool
		return ret
	}
	return *o.PrimaryBlogPost
}

// GetPrimaryBlogPostOk returns a tuple with the PrimaryBlogPost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetPrimaryBlogPostOk() (*bool, bool) {
	if o == nil || IsNil(o.PrimaryBlogPost) {
		return nil, false
	}
	return o.PrimaryBlogPost, true
}

// HasPrimaryBlogPost returns a boolean if a field has been set.
func (o *Domain) HasPrimaryBlogPost() bool {
	if o != nil && !IsNil(o.PrimaryBlogPost) {
		return true
	}

	return false
}

// SetPrimaryBlogPost gets a reference to the given bool and assigns it to the PrimaryBlogPost field.
func (o *Domain) SetPrimaryBlogPost(v bool) {
	o.PrimaryBlogPost = &v
}

// GetPrimaryKnowledge returns the PrimaryKnowledge field value if set, zero value otherwise.
func (o *Domain) GetPrimaryKnowledge() bool {
	if o == nil || IsNil(o.PrimaryKnowledge) {
		var ret bool
		return ret
	}
	return *o.PrimaryKnowledge
}

// GetPrimaryKnowledgeOk returns a tuple with the PrimaryKnowledge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetPrimaryKnowledgeOk() (*bool, bool) {
	if o == nil || IsNil(o.PrimaryKnowledge) {
		return nil, false
	}
	return o.PrimaryKnowledge, true
}

// HasPrimaryKnowledge returns a boolean if a field has been set.
func (o *Domain) HasPrimaryKnowledge() bool {
	if o != nil && !IsNil(o.PrimaryKnowledge) {
		return true
	}

	return false
}

// SetPrimaryKnowledge gets a reference to the given bool and assigns it to the PrimaryKnowledge field.
func (o *Domain) SetPrimaryKnowledge(v bool) {
	o.PrimaryKnowledge = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *Domain) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *Domain) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *Domain) SetCreated(v time.Time) {
	o.Created = &v
}

// GetSecondaryToDomain returns the SecondaryToDomain field value if set, zero value otherwise.
func (o *Domain) GetSecondaryToDomain() string {
	if o == nil || IsNil(o.SecondaryToDomain) {
		var ret string
		return ret
	}
	return *o.SecondaryToDomain
}

// GetSecondaryToDomainOk returns a tuple with the SecondaryToDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetSecondaryToDomainOk() (*string, bool) {
	if o == nil || IsNil(o.SecondaryToDomain) {
		return nil, false
	}
	return o.SecondaryToDomain, true
}

// HasSecondaryToDomain returns a boolean if a field has been set.
func (o *Domain) HasSecondaryToDomain() bool {
	if o != nil && !IsNil(o.SecondaryToDomain) {
		return true
	}

	return false
}

// SetSecondaryToDomain gets a reference to the given string and assigns it to the SecondaryToDomain field.
func (o *Domain) SetSecondaryToDomain(v string) {
	o.SecondaryToDomain = &v
}

// GetManuallyMarkedAsResolving returns the ManuallyMarkedAsResolving field value if set, zero value otherwise.
func (o *Domain) GetManuallyMarkedAsResolving() bool {
	if o == nil || IsNil(o.ManuallyMarkedAsResolving) {
		var ret bool
		return ret
	}
	return *o.ManuallyMarkedAsResolving
}

// GetManuallyMarkedAsResolvingOk returns a tuple with the ManuallyMarkedAsResolving field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetManuallyMarkedAsResolvingOk() (*bool, bool) {
	if o == nil || IsNil(o.ManuallyMarkedAsResolving) {
		return nil, false
	}
	return o.ManuallyMarkedAsResolving, true
}

// HasManuallyMarkedAsResolving returns a boolean if a field has been set.
func (o *Domain) HasManuallyMarkedAsResolving() bool {
	if o != nil && !IsNil(o.ManuallyMarkedAsResolving) {
		return true
	}

	return false
}

// SetManuallyMarkedAsResolving gets a reference to the given bool and assigns it to the ManuallyMarkedAsResolving field.
func (o *Domain) SetManuallyMarkedAsResolving(v bool) {
	o.ManuallyMarkedAsResolving = &v
}

// GetIsUsedForKnowledge returns the IsUsedForKnowledge field value
func (o *Domain) GetIsUsedForKnowledge() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsUsedForKnowledge
}

// GetIsUsedForKnowledgeOk returns a tuple with the IsUsedForKnowledge field value
// and a boolean to check if the value has been set.
func (o *Domain) GetIsUsedForKnowledgeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsUsedForKnowledge, true
}

// SetIsUsedForKnowledge sets field value
func (o *Domain) SetIsUsedForKnowledge(v bool) {
	o.IsUsedForKnowledge = v
}

// GetIsUsedForBlogPost returns the IsUsedForBlogPost field value
func (o *Domain) GetIsUsedForBlogPost() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsUsedForBlogPost
}

// GetIsUsedForBlogPostOk returns a tuple with the IsUsedForBlogPost field value
// and a boolean to check if the value has been set.
func (o *Domain) GetIsUsedForBlogPostOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsUsedForBlogPost, true
}

// SetIsUsedForBlogPost sets field value
func (o *Domain) SetIsUsedForBlogPost(v bool) {
	o.IsUsedForBlogPost = v
}

// GetIsUsedForSitePage returns the IsUsedForSitePage field value
func (o *Domain) GetIsUsedForSitePage() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsUsedForSitePage
}

// GetIsUsedForSitePageOk returns a tuple with the IsUsedForSitePage field value
// and a boolean to check if the value has been set.
func (o *Domain) GetIsUsedForSitePageOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsUsedForSitePage, true
}

// SetIsUsedForSitePage sets field value
func (o *Domain) SetIsUsedForSitePage(v bool) {
	o.IsUsedForSitePage = v
}

// GetIsResolving returns the IsResolving field value
func (o *Domain) GetIsResolving() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsResolving
}

// GetIsResolvingOk returns a tuple with the IsResolving field value
// and a boolean to check if the value has been set.
func (o *Domain) GetIsResolvingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsResolving, true
}

// SetIsResolving sets field value
func (o *Domain) SetIsResolving(v bool) {
	o.IsResolving = v
}

// GetIsSslEnabled returns the IsSslEnabled field value if set, zero value otherwise.
func (o *Domain) GetIsSslEnabled() bool {
	if o == nil || IsNil(o.IsSslEnabled) {
		var ret bool
		return ret
	}
	return *o.IsSslEnabled
}

// GetIsSslEnabledOk returns a tuple with the IsSslEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetIsSslEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSslEnabled) {
		return nil, false
	}
	return o.IsSslEnabled, true
}

// HasIsSslEnabled returns a boolean if a field has been set.
func (o *Domain) HasIsSslEnabled() bool {
	if o != nil && !IsNil(o.IsSslEnabled) {
		return true
	}

	return false
}

// SetIsSslEnabled gets a reference to the given bool and assigns it to the IsSslEnabled field.
func (o *Domain) SetIsSslEnabled(v bool) {
	o.IsSslEnabled = &v
}

// GetIsUsedForEmail returns the IsUsedForEmail field value
func (o *Domain) GetIsUsedForEmail() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsUsedForEmail
}

// GetIsUsedForEmailOk returns a tuple with the IsUsedForEmail field value
// and a boolean to check if the value has been set.
func (o *Domain) GetIsUsedForEmailOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsUsedForEmail, true
}

// SetIsUsedForEmail sets field value
func (o *Domain) SetIsUsedForEmail(v bool) {
	o.IsUsedForEmail = v
}

// GetDomain returns the Domain field value
func (o *Domain) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *Domain) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value
func (o *Domain) SetDomain(v string) {
	o.Domain = v
}

// GetPrimarySitePage returns the PrimarySitePage field value if set, zero value otherwise.
func (o *Domain) GetPrimarySitePage() bool {
	if o == nil || IsNil(o.PrimarySitePage) {
		var ret bool
		return ret
	}
	return *o.PrimarySitePage
}

// GetPrimarySitePageOk returns a tuple with the PrimarySitePage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetPrimarySitePageOk() (*bool, bool) {
	if o == nil || IsNil(o.PrimarySitePage) {
		return nil, false
	}
	return o.PrimarySitePage, true
}

// HasPrimarySitePage returns a boolean if a field has been set.
func (o *Domain) HasPrimarySitePage() bool {
	if o != nil && !IsNil(o.PrimarySitePage) {
		return true
	}

	return false
}

// SetPrimarySitePage gets a reference to the given bool and assigns it to the PrimarySitePage field.
func (o *Domain) SetPrimarySitePage(v bool) {
	o.PrimarySitePage = &v
}

// GetPrimaryLandingPage returns the PrimaryLandingPage field value if set, zero value otherwise.
func (o *Domain) GetPrimaryLandingPage() bool {
	if o == nil || IsNil(o.PrimaryLandingPage) {
		var ret bool
		return ret
	}
	return *o.PrimaryLandingPage
}

// GetPrimaryLandingPageOk returns a tuple with the PrimaryLandingPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetPrimaryLandingPageOk() (*bool, bool) {
	if o == nil || IsNil(o.PrimaryLandingPage) {
		return nil, false
	}
	return o.PrimaryLandingPage, true
}

// HasPrimaryLandingPage returns a boolean if a field has been set.
func (o *Domain) HasPrimaryLandingPage() bool {
	if o != nil && !IsNil(o.PrimaryLandingPage) {
		return true
	}

	return false
}

// SetPrimaryLandingPage gets a reference to the given bool and assigns it to the PrimaryLandingPage field.
func (o *Domain) SetPrimaryLandingPage(v bool) {
	o.PrimaryLandingPage = &v
}

// GetId returns the Id field value
func (o *Domain) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Domain) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Domain) SetId(v string) {
	o.Id = v
}

// GetCorrectCname returns the CorrectCname field value if set, zero value otherwise.
func (o *Domain) GetCorrectCname() string {
	if o == nil || IsNil(o.CorrectCname) {
		var ret string
		return ret
	}
	return *o.CorrectCname
}

// GetCorrectCnameOk returns a tuple with the CorrectCname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetCorrectCnameOk() (*string, bool) {
	if o == nil || IsNil(o.CorrectCname) {
		return nil, false
	}
	return o.CorrectCname, true
}

// HasCorrectCname returns a boolean if a field has been set.
func (o *Domain) HasCorrectCname() bool {
	if o != nil && !IsNil(o.CorrectCname) {
		return true
	}

	return false
}

// SetCorrectCname gets a reference to the given string and assigns it to the CorrectCname field.
func (o *Domain) SetCorrectCname(v string) {
	o.CorrectCname = &v
}

// GetIsSslOnly returns the IsSslOnly field value if set, zero value otherwise.
func (o *Domain) GetIsSslOnly() bool {
	if o == nil || IsNil(o.IsSslOnly) {
		var ret bool
		return ret
	}
	return *o.IsSslOnly
}

// GetIsSslOnlyOk returns a tuple with the IsSslOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetIsSslOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSslOnly) {
		return nil, false
	}
	return o.IsSslOnly, true
}

// HasIsSslOnly returns a boolean if a field has been set.
func (o *Domain) HasIsSslOnly() bool {
	if o != nil && !IsNil(o.IsSslOnly) {
		return true
	}

	return false
}

// SetIsSslOnly gets a reference to the given bool and assigns it to the IsSslOnly field.
func (o *Domain) SetIsSslOnly(v bool) {
	o.IsSslOnly = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *Domain) GetUpdated() time.Time {
	if o == nil || IsNil(o.Updated) {
		var ret time.Time
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Updated) {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *Domain) HasUpdated() bool {
	if o != nil && !IsNil(o.Updated) {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given time.Time and assigns it to the Updated field.
func (o *Domain) SetUpdated(v time.Time) {
	o.Updated = &v
}

// GetPrimaryEmail returns the PrimaryEmail field value if set, zero value otherwise.
func (o *Domain) GetPrimaryEmail() bool {
	if o == nil || IsNil(o.PrimaryEmail) {
		var ret bool
		return ret
	}
	return *o.PrimaryEmail
}

// GetPrimaryEmailOk returns a tuple with the PrimaryEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetPrimaryEmailOk() (*bool, bool) {
	if o == nil || IsNil(o.PrimaryEmail) {
		return nil, false
	}
	return o.PrimaryEmail, true
}

// HasPrimaryEmail returns a boolean if a field has been set.
func (o *Domain) HasPrimaryEmail() bool {
	if o != nil && !IsNil(o.PrimaryEmail) {
		return true
	}

	return false
}

// SetPrimaryEmail gets a reference to the given bool and assigns it to the PrimaryEmail field.
func (o *Domain) SetPrimaryEmail(v bool) {
	o.PrimaryEmail = &v
}

func (o Domain) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Domain) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["isUsedForLandingPage"] = o.IsUsedForLandingPage
	if !IsNil(o.PrimaryBlogPost) {
		toSerialize["primaryBlogPost"] = o.PrimaryBlogPost
	}
	if !IsNil(o.PrimaryKnowledge) {
		toSerialize["primaryKnowledge"] = o.PrimaryKnowledge
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.SecondaryToDomain) {
		toSerialize["secondaryToDomain"] = o.SecondaryToDomain
	}
	if !IsNil(o.ManuallyMarkedAsResolving) {
		toSerialize["manuallyMarkedAsResolving"] = o.ManuallyMarkedAsResolving
	}
	toSerialize["isUsedForKnowledge"] = o.IsUsedForKnowledge
	toSerialize["isUsedForBlogPost"] = o.IsUsedForBlogPost
	toSerialize["isUsedForSitePage"] = o.IsUsedForSitePage
	toSerialize["isResolving"] = o.IsResolving
	if !IsNil(o.IsSslEnabled) {
		toSerialize["isSslEnabled"] = o.IsSslEnabled
	}
	toSerialize["isUsedForEmail"] = o.IsUsedForEmail
	toSerialize["domain"] = o.Domain
	if !IsNil(o.PrimarySitePage) {
		toSerialize["primarySitePage"] = o.PrimarySitePage
	}
	if !IsNil(o.PrimaryLandingPage) {
		toSerialize["primaryLandingPage"] = o.PrimaryLandingPage
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.CorrectCname) {
		toSerialize["correctCname"] = o.CorrectCname
	}
	if !IsNil(o.IsSslOnly) {
		toSerialize["isSslOnly"] = o.IsSslOnly
	}
	if !IsNil(o.Updated) {
		toSerialize["updated"] = o.Updated
	}
	if !IsNil(o.PrimaryEmail) {
		toSerialize["primaryEmail"] = o.PrimaryEmail
	}
	return toSerialize, nil
}

func (o *Domain) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"isUsedForLandingPage",
		"isUsedForKnowledge",
		"isUsedForBlogPost",
		"isUsedForSitePage",
		"isResolving",
		"isUsedForEmail",
		"domain",
		"id",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varDomain := _Domain{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDomain)

	if err != nil {
		return err
	}

	*o = Domain(varDomain)

	return err
}

type NullableDomain struct {
	value *Domain
	isSet bool
}

func (v NullableDomain) Get() *Domain {
	return v.value
}

func (v *NullableDomain) Set(val *Domain) {
	v.value = val
	v.isSet = true
}

func (v NullableDomain) IsSet() bool {
	return v.isSet
}

func (v *NullableDomain) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomain(val *Domain) *NullableDomain {
	return &NullableDomain{value: val, isSet: true}
}

func (v NullableDomain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
