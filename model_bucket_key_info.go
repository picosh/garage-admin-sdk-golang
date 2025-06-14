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

// checks if the BucketKeyInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BucketKeyInfo{}

// BucketKeyInfo struct for BucketKeyInfo
type BucketKeyInfo struct {
	AccessKeyId *string `json:"accessKeyId,omitempty"`
	Name *string `json:"name,omitempty"`
	Permissions *CreateBucketRequestLocalAliasAllow `json:"permissions,omitempty"`
	BucketLocalAliases []string `json:"bucketLocalAliases,omitempty"`
}

// NewBucketKeyInfo instantiates a new BucketKeyInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBucketKeyInfo() *BucketKeyInfo {
	this := BucketKeyInfo{}
	return &this
}

// NewBucketKeyInfoWithDefaults instantiates a new BucketKeyInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBucketKeyInfoWithDefaults() *BucketKeyInfo {
	this := BucketKeyInfo{}
	return &this
}

// GetAccessKeyId returns the AccessKeyId field value if set, zero value otherwise.
func (o *BucketKeyInfo) GetAccessKeyId() string {
	if o == nil || IsNil(o.AccessKeyId) {
		var ret string
		return ret
	}
	return *o.AccessKeyId
}

// GetAccessKeyIdOk returns a tuple with the AccessKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BucketKeyInfo) GetAccessKeyIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccessKeyId) {
		return nil, false
	}
	return o.AccessKeyId, true
}

// HasAccessKeyId returns a boolean if a field has been set.
func (o *BucketKeyInfo) HasAccessKeyId() bool {
	if o != nil && !IsNil(o.AccessKeyId) {
		return true
	}

	return false
}

// SetAccessKeyId gets a reference to the given string and assigns it to the AccessKeyId field.
func (o *BucketKeyInfo) SetAccessKeyId(v string) {
	o.AccessKeyId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BucketKeyInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BucketKeyInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BucketKeyInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BucketKeyInfo) SetName(v string) {
	o.Name = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *BucketKeyInfo) GetPermissions() CreateBucketRequestLocalAliasAllow {
	if o == nil || IsNil(o.Permissions) {
		var ret CreateBucketRequestLocalAliasAllow
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BucketKeyInfo) GetPermissionsOk() (*CreateBucketRequestLocalAliasAllow, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *BucketKeyInfo) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given CreateBucketRequestLocalAliasAllow and assigns it to the Permissions field.
func (o *BucketKeyInfo) SetPermissions(v CreateBucketRequestLocalAliasAllow) {
	o.Permissions = &v
}

// GetBucketLocalAliases returns the BucketLocalAliases field value if set, zero value otherwise.
func (o *BucketKeyInfo) GetBucketLocalAliases() []string {
	if o == nil || IsNil(o.BucketLocalAliases) {
		var ret []string
		return ret
	}
	return o.BucketLocalAliases
}

// GetBucketLocalAliasesOk returns a tuple with the BucketLocalAliases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BucketKeyInfo) GetBucketLocalAliasesOk() ([]string, bool) {
	if o == nil || IsNil(o.BucketLocalAliases) {
		return nil, false
	}
	return o.BucketLocalAliases, true
}

// HasBucketLocalAliases returns a boolean if a field has been set.
func (o *BucketKeyInfo) HasBucketLocalAliases() bool {
	if o != nil && !IsNil(o.BucketLocalAliases) {
		return true
	}

	return false
}

// SetBucketLocalAliases gets a reference to the given []string and assigns it to the BucketLocalAliases field.
func (o *BucketKeyInfo) SetBucketLocalAliases(v []string) {
	o.BucketLocalAliases = v
}

func (o BucketKeyInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BucketKeyInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessKeyId) {
		toSerialize["accessKeyId"] = o.AccessKeyId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}
	if !IsNil(o.BucketLocalAliases) {
		toSerialize["bucketLocalAliases"] = o.BucketLocalAliases
	}
	return toSerialize, nil
}

type NullableBucketKeyInfo struct {
	value *BucketKeyInfo
	isSet bool
}

func (v NullableBucketKeyInfo) Get() *BucketKeyInfo {
	return v.value
}

func (v *NullableBucketKeyInfo) Set(val *BucketKeyInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBucketKeyInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBucketKeyInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBucketKeyInfo(val *BucketKeyInfo) *NullableBucketKeyInfo {
	return &NullableBucketKeyInfo{value: val, isSet: true}
}

func (v NullableBucketKeyInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBucketKeyInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


