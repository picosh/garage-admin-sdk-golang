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

// checks if the UpdateBucketRequestQuotas type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateBucketRequestQuotas{}

// UpdateBucketRequestQuotas struct for UpdateBucketRequestQuotas
type UpdateBucketRequestQuotas struct {
	MaxSize NullableInt64 `json:"maxSize,omitempty"`
	MaxObjects NullableInt64 `json:"maxObjects,omitempty"`
}

// NewUpdateBucketRequestQuotas instantiates a new UpdateBucketRequestQuotas object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateBucketRequestQuotas() *UpdateBucketRequestQuotas {
	this := UpdateBucketRequestQuotas{}
	return &this
}

// NewUpdateBucketRequestQuotasWithDefaults instantiates a new UpdateBucketRequestQuotas object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateBucketRequestQuotasWithDefaults() *UpdateBucketRequestQuotas {
	this := UpdateBucketRequestQuotas{}
	return &this
}

// GetMaxSize returns the MaxSize field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateBucketRequestQuotas) GetMaxSize() int64 {
	if o == nil || IsNil(o.MaxSize.Get()) {
		var ret int64
		return ret
	}
	return *o.MaxSize.Get()
}

// GetMaxSizeOk returns a tuple with the MaxSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateBucketRequestQuotas) GetMaxSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxSize.Get(), o.MaxSize.IsSet()
}

// HasMaxSize returns a boolean if a field has been set.
func (o *UpdateBucketRequestQuotas) HasMaxSize() bool {
	if o != nil && o.MaxSize.IsSet() {
		return true
	}

	return false
}

// SetMaxSize gets a reference to the given NullableInt64 and assigns it to the MaxSize field.
func (o *UpdateBucketRequestQuotas) SetMaxSize(v int64) {
	o.MaxSize.Set(&v)
}
// SetMaxSizeNil sets the value for MaxSize to be an explicit nil
func (o *UpdateBucketRequestQuotas) SetMaxSizeNil() {
	o.MaxSize.Set(nil)
}

// UnsetMaxSize ensures that no value is present for MaxSize, not even an explicit nil
func (o *UpdateBucketRequestQuotas) UnsetMaxSize() {
	o.MaxSize.Unset()
}

// GetMaxObjects returns the MaxObjects field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateBucketRequestQuotas) GetMaxObjects() int64 {
	if o == nil || IsNil(o.MaxObjects.Get()) {
		var ret int64
		return ret
	}
	return *o.MaxObjects.Get()
}

// GetMaxObjectsOk returns a tuple with the MaxObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateBucketRequestQuotas) GetMaxObjectsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxObjects.Get(), o.MaxObjects.IsSet()
}

// HasMaxObjects returns a boolean if a field has been set.
func (o *UpdateBucketRequestQuotas) HasMaxObjects() bool {
	if o != nil && o.MaxObjects.IsSet() {
		return true
	}

	return false
}

// SetMaxObjects gets a reference to the given NullableInt64 and assigns it to the MaxObjects field.
func (o *UpdateBucketRequestQuotas) SetMaxObjects(v int64) {
	o.MaxObjects.Set(&v)
}
// SetMaxObjectsNil sets the value for MaxObjects to be an explicit nil
func (o *UpdateBucketRequestQuotas) SetMaxObjectsNil() {
	o.MaxObjects.Set(nil)
}

// UnsetMaxObjects ensures that no value is present for MaxObjects, not even an explicit nil
func (o *UpdateBucketRequestQuotas) UnsetMaxObjects() {
	o.MaxObjects.Unset()
}

func (o UpdateBucketRequestQuotas) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateBucketRequestQuotas) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.MaxSize.IsSet() {
		toSerialize["maxSize"] = o.MaxSize.Get()
	}
	if o.MaxObjects.IsSet() {
		toSerialize["maxObjects"] = o.MaxObjects.Get()
	}
	return toSerialize, nil
}

type NullableUpdateBucketRequestQuotas struct {
	value *UpdateBucketRequestQuotas
	isSet bool
}

func (v NullableUpdateBucketRequestQuotas) Get() *UpdateBucketRequestQuotas {
	return v.value
}

func (v *NullableUpdateBucketRequestQuotas) Set(val *UpdateBucketRequestQuotas) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateBucketRequestQuotas) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateBucketRequestQuotas) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateBucketRequestQuotas(val *UpdateBucketRequestQuotas) *NullableUpdateBucketRequestQuotas {
	return &NullableUpdateBucketRequestQuotas{value: val, isSet: true}
}

func (v NullableUpdateBucketRequestQuotas) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateBucketRequestQuotas) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


