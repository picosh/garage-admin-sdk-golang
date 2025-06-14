/*
Garage Administration API v0+garage-v0.9.0

Administrate your Garage cluster programatically, including status, layout, keys, buckets, and maintainance tasks.  *Disclaimer: The API is not stable yet, hence its v0 tag. The API can change at any time, and changes can include breaking backward compatibility. Read the changelog and upgrade your scripts before upgrading. Additionnaly, this specification is very early stage and can contain bugs, especially on error return codes/types that are not tested yet. Do not expect a well finished and polished product!* 

API version: v0.9.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package garage

import (
	"encoding/json"
)

// checks if the BucketInfoWebsiteConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BucketInfoWebsiteConfig{}

// BucketInfoWebsiteConfig struct for BucketInfoWebsiteConfig
type BucketInfoWebsiteConfig struct {
	IndexDocument *string `json:"indexDocument,omitempty"`
	ErrorDocument *string `json:"errorDocument,omitempty"`
}

// NewBucketInfoWebsiteConfig instantiates a new BucketInfoWebsiteConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBucketInfoWebsiteConfig() *BucketInfoWebsiteConfig {
	this := BucketInfoWebsiteConfig{}
	return &this
}

// NewBucketInfoWebsiteConfigWithDefaults instantiates a new BucketInfoWebsiteConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBucketInfoWebsiteConfigWithDefaults() *BucketInfoWebsiteConfig {
	this := BucketInfoWebsiteConfig{}
	return &this
}

// GetIndexDocument returns the IndexDocument field value if set, zero value otherwise.
func (o *BucketInfoWebsiteConfig) GetIndexDocument() string {
	if o == nil || IsNil(o.IndexDocument) {
		var ret string
		return ret
	}
	return *o.IndexDocument
}

// GetIndexDocumentOk returns a tuple with the IndexDocument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BucketInfoWebsiteConfig) GetIndexDocumentOk() (*string, bool) {
	if o == nil || IsNil(o.IndexDocument) {
		return nil, false
	}
	return o.IndexDocument, true
}

// HasIndexDocument returns a boolean if a field has been set.
func (o *BucketInfoWebsiteConfig) HasIndexDocument() bool {
	if o != nil && !IsNil(o.IndexDocument) {
		return true
	}

	return false
}

// SetIndexDocument gets a reference to the given string and assigns it to the IndexDocument field.
func (o *BucketInfoWebsiteConfig) SetIndexDocument(v string) {
	o.IndexDocument = &v
}

// GetErrorDocument returns the ErrorDocument field value if set, zero value otherwise.
func (o *BucketInfoWebsiteConfig) GetErrorDocument() string {
	if o == nil || IsNil(o.ErrorDocument) {
		var ret string
		return ret
	}
	return *o.ErrorDocument
}

// GetErrorDocumentOk returns a tuple with the ErrorDocument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BucketInfoWebsiteConfig) GetErrorDocumentOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorDocument) {
		return nil, false
	}
	return o.ErrorDocument, true
}

// HasErrorDocument returns a boolean if a field has been set.
func (o *BucketInfoWebsiteConfig) HasErrorDocument() bool {
	if o != nil && !IsNil(o.ErrorDocument) {
		return true
	}

	return false
}

// SetErrorDocument gets a reference to the given string and assigns it to the ErrorDocument field.
func (o *BucketInfoWebsiteConfig) SetErrorDocument(v string) {
	o.ErrorDocument = &v
}

func (o BucketInfoWebsiteConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BucketInfoWebsiteConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IndexDocument) {
		toSerialize["indexDocument"] = o.IndexDocument
	}
	if !IsNil(o.ErrorDocument) {
		toSerialize["errorDocument"] = o.ErrorDocument
	}
	return toSerialize, nil
}

type NullableBucketInfoWebsiteConfig struct {
	value *BucketInfoWebsiteConfig
	isSet bool
}

func (v NullableBucketInfoWebsiteConfig) Get() *BucketInfoWebsiteConfig {
	return v.value
}

func (v *NullableBucketInfoWebsiteConfig) Set(val *BucketInfoWebsiteConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableBucketInfoWebsiteConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableBucketInfoWebsiteConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBucketInfoWebsiteConfig(val *BucketInfoWebsiteConfig) *NullableBucketInfoWebsiteConfig {
	return &NullableBucketInfoWebsiteConfig{value: val, isSet: true}
}

func (v NullableBucketInfoWebsiteConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBucketInfoWebsiteConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


