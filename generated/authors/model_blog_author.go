/*
Blog Post endpoints

Use these endpoints for interacting with Blog Posts, Blog Authors, and Blog Tags

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authors

import (
	"encoding/json"
	"time"
)

// BlogAuthor Model definition for a Blog Author.
type BlogAuthor struct {
	// The unique ID of the Blog Author.
	Id       string `json:"id"`
	FullName string `json:"fullName"`
	// Email address of the Blog Author.
	Email string `json:"email"`
	Slug  string `json:"slug"`
	// The explicitly defined ISO 639 language code of the blog author.
	Language string `json:"language"`
	// ID of the primary blog author this object was translated from.
	TranslatedFromId int64  `json:"translatedFromId"`
	Name             string `json:"name"`
	// The full name of the Blog Author to be displayed.
	DisplayName string `json:"displayName"`
	// A short biography of the blog author.
	Bio string `json:"bio"`
	// URL to the website of the Blog Author.
	Website string `json:"website"`
	// URL or username of the Twitter account associated with the Blog Author. This will be normalized into the Twitter url for said user.
	Twitter string `json:"twitter"`
	// URL to the Blog Author's Facebook page.
	Facebook string `json:"facebook"`
	// URL to the blog author's LinkedIn page.
	Linkedin string `json:"linkedin"`
	// URL to the blog author's avatar, if supplying a custom one.
	Avatar  string    `json:"avatar"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
	// The timestamp (ISO8601 format) when this Blog Author was deleted.
	DeletedAt time.Time `json:"deletedAt"`
}

// NewBlogAuthor instantiates a new BlogAuthor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlogAuthor(id string, fullName string, email string, slug string, language string, translatedFromId int64, name string, displayName string, bio string, website string, twitter string, facebook string, linkedin string, avatar string, created time.Time, updated time.Time, deletedAt time.Time) *BlogAuthor {
	this := BlogAuthor{}
	this.Id = id
	this.FullName = fullName
	this.Email = email
	this.Slug = slug
	this.Language = language
	this.TranslatedFromId = translatedFromId
	this.Name = name
	this.DisplayName = displayName
	this.Bio = bio
	this.Website = website
	this.Twitter = twitter
	this.Facebook = facebook
	this.Linkedin = linkedin
	this.Avatar = avatar
	this.Created = created
	this.Updated = updated
	this.DeletedAt = deletedAt
	return &this
}

// NewBlogAuthorWithDefaults instantiates a new BlogAuthor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlogAuthorWithDefaults() *BlogAuthor {
	this := BlogAuthor{}
	return &this
}

// GetId returns the Id field value
func (o *BlogAuthor) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BlogAuthor) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BlogAuthor) SetId(v string) {
	o.Id = v
}

// GetFullName returns the FullName field value
func (o *BlogAuthor) GetFullName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value
// and a boolean to check if the value has been set.
func (o *BlogAuthor) GetFullNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FullName, true
}

// SetFullName sets field value
func (o *BlogAuthor) SetFullName(v string) {
	o.FullName = v
}

// GetEmail returns the Email field value
func (o *BlogAuthor) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *BlogAuthor) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *BlogAuthor) SetEmail(v string) {
	o.Email = v
}

// GetSlug returns the Slug field value
func (o *BlogAuthor) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *BlogAuthor) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *BlogAuthor) SetSlug(v string) {
	o.Slug = v
}

// GetLanguage returns the Language field value
func (o *BlogAuthor) GetLanguage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Language
}

// GetLanguageOk returns a tuple with the Language field value
// and a boolean to check if the value has been set.
func (o *BlogAuthor) GetLanguageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Language, true
}

// SetLanguage sets field value
func (o *BlogAuthor) SetLanguage(v string) {
	o.Language = v
}

// GetTranslatedFromId returns the TranslatedFromId field value
func (o *BlogAuthor) GetTranslatedFromId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TranslatedFromId
}

// GetTranslatedFromIdOk returns a tuple with the TranslatedFromId field value
// and a boolean to check if the value has been set.
func (o *BlogAuthor) GetTranslatedFromIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TranslatedFromId, true
}

// SetTranslatedFromId sets field value
func (o *BlogAuthor) SetTranslatedFromId(v int64) {
	o.TranslatedFromId = v
}

// GetName returns the Name field value
func (o *BlogAuthor) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BlogAuthor) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BlogAuthor) SetName(v string) {
	o.Name = v
}

// GetDisplayName returns the DisplayName field value
func (o *BlogAuthor) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *BlogAuthor) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *BlogAuthor) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetBio returns the Bio field value
func (o *BlogAuthor) GetBio() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Bio
}

// GetBioOk returns a tuple with the Bio field value
// and a boolean to check if the value has been set.
func (o *BlogAuthor) GetBioOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bio, true
}

// SetBio sets field value
func (o *BlogAuthor) SetBio(v string) {
	o.Bio = v
}

// GetWebsite returns the Website field value
func (o *BlogAuthor) GetWebsite() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Website
}

// GetWebsiteOk returns a tuple with the Website field value
// and a boolean to check if the value has been set.
func (o *BlogAuthor) GetWebsiteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Website, true
}

// SetWebsite sets field value
func (o *BlogAuthor) SetWebsite(v string) {
	o.Website = v
}

// GetTwitter returns the Twitter field value
func (o *BlogAuthor) GetTwitter() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Twitter
}

// GetTwitterOk returns a tuple with the Twitter field value
// and a boolean to check if the value has been set.
func (o *BlogAuthor) GetTwitterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Twitter, true
}

// SetTwitter sets field value
func (o *BlogAuthor) SetTwitter(v string) {
	o.Twitter = v
}

// GetFacebook returns the Facebook field value
func (o *BlogAuthor) GetFacebook() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Facebook
}

// GetFacebookOk returns a tuple with the Facebook field value
// and a boolean to check if the value has been set.
func (o *BlogAuthor) GetFacebookOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Facebook, true
}

// SetFacebook sets field value
func (o *BlogAuthor) SetFacebook(v string) {
	o.Facebook = v
}

// GetLinkedin returns the Linkedin field value
func (o *BlogAuthor) GetLinkedin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Linkedin
}

// GetLinkedinOk returns a tuple with the Linkedin field value
// and a boolean to check if the value has been set.
func (o *BlogAuthor) GetLinkedinOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Linkedin, true
}

// SetLinkedin sets field value
func (o *BlogAuthor) SetLinkedin(v string) {
	o.Linkedin = v
}

// GetAvatar returns the Avatar field value
func (o *BlogAuthor) GetAvatar() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Avatar
}

// GetAvatarOk returns a tuple with the Avatar field value
// and a boolean to check if the value has been set.
func (o *BlogAuthor) GetAvatarOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Avatar, true
}

// SetAvatar sets field value
func (o *BlogAuthor) SetAvatar(v string) {
	o.Avatar = v
}

// GetCreated returns the Created field value
func (o *BlogAuthor) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *BlogAuthor) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *BlogAuthor) SetCreated(v time.Time) {
	o.Created = v
}

// GetUpdated returns the Updated field value
func (o *BlogAuthor) GetUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value
// and a boolean to check if the value has been set.
func (o *BlogAuthor) GetUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Updated, true
}

// SetUpdated sets field value
func (o *BlogAuthor) SetUpdated(v time.Time) {
	o.Updated = v
}

// GetDeletedAt returns the DeletedAt field value
func (o *BlogAuthor) GetDeletedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value
// and a boolean to check if the value has been set.
func (o *BlogAuthor) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeletedAt, true
}

// SetDeletedAt sets field value
func (o *BlogAuthor) SetDeletedAt(v time.Time) {
	o.DeletedAt = v
}

func (o BlogAuthor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["fullName"] = o.FullName
	}
	if true {
		toSerialize["email"] = o.Email
	}
	if true {
		toSerialize["slug"] = o.Slug
	}
	if true {
		toSerialize["language"] = o.Language
	}
	if true {
		toSerialize["translatedFromId"] = o.TranslatedFromId
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["displayName"] = o.DisplayName
	}
	if true {
		toSerialize["bio"] = o.Bio
	}
	if true {
		toSerialize["website"] = o.Website
	}
	if true {
		toSerialize["twitter"] = o.Twitter
	}
	if true {
		toSerialize["facebook"] = o.Facebook
	}
	if true {
		toSerialize["linkedin"] = o.Linkedin
	}
	if true {
		toSerialize["avatar"] = o.Avatar
	}
	if true {
		toSerialize["created"] = o.Created
	}
	if true {
		toSerialize["updated"] = o.Updated
	}
	if true {
		toSerialize["deletedAt"] = o.DeletedAt
	}
	return json.Marshal(toSerialize)
}

type NullableBlogAuthor struct {
	value *BlogAuthor
	isSet bool
}

func (v NullableBlogAuthor) Get() *BlogAuthor {
	return v.value
}

func (v *NullableBlogAuthor) Set(val *BlogAuthor) {
	v.value = val
	v.isSet = true
}

func (v NullableBlogAuthor) IsSet() bool {
	return v.isSet
}

func (v *NullableBlogAuthor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlogAuthor(val *BlogAuthor) *NullableBlogAuthor {
	return &NullableBlogAuthor{value: val, isSet: true}
}

func (v NullableBlogAuthor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlogAuthor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
