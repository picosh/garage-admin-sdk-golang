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

// checks if the BucketInfoQuotas type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BucketInfoQuotas{}

// BucketInfoQuotas struct for BucketInfoQuotas
type BucketInfoQuotas struct {
	MaxSize NullableInt64 `json:"maxSize,omitempty"`
	MaxObjects NullableInt64 `json:"maxObjects,omitempty"`
}

// NewBucketInfoQuotas instantiates a new BucketInfoQuotas object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBucketInfoQuotas() *BucketInfoQuotas {
	this := BucketInfoQuotas{}
	return &this
}

// NewBucketInfoQuotasWithDefaults instantiates a new BucketInfoQuotas object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBucketInfoQuotasWithDefaults() *BucketInfoQuotas {
	this := BucketInfoQuotas{}
	return &this
}

// GetMaxSize returns the MaxSize field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BucketInfoQuotas) GetMaxSize() int64 {
	if o == nil || IsNil(o.MaxSize.Get()) {
		var ret int64
		return ret
	}
	return *o.MaxSize.Get()
}

// GetMaxSizeOk returns a tuple with the MaxSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BucketInfoQuotas) GetMaxSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxSize.Get(), o.MaxSize.IsSet()
}

// HasMaxSize returns a boolean if a field has been set.
func (o *BucketInfoQuotas) HasMaxSize() bool {
	if o != nil && o.MaxSize.IsSet() {
		return true
	}

	return false
}

// SetMaxSize gets a reference to the given NullableInt64 and assigns it to the MaxSize field.
func (o *BucketInfoQuotas) SetMaxSize(v int64) {
	o.MaxSize.Set(&v)
}
// SetMaxSizeNil sets the value for MaxSize to be an explicit nil
func (o *BucketInfoQuotas) SetMaxSizeNil() {
	o.MaxSize.Set(nil)
}

// UnsetMaxSize ensures that no value is present for MaxSize, not even an explicit nil
func (o *BucketInfoQuotas) UnsetMaxSize() {
	o.MaxSize.Unset()
}

// GetMaxObjects returns the MaxObjects field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BucketInfoQuotas) GetMaxObjects() int64 {
	if o == nil || IsNil(o.MaxObjects.Get()) {
		var ret int64
		return ret
	}
	return *o.MaxObjects.Get()
}

// GetMaxObjectsOk returns a tuple with the MaxObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BucketInfoQuotas) GetMaxObjectsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxObjects.Get(), o.MaxObjects.IsSet()
}

// HasMaxObjects returns a boolean if a field has been set.
func (o *BucketInfoQuotas) HasMaxObjects() bool {
	if o != nil && o.MaxObjects.IsSet() {
		return true
	}

	return false
}

// SetMaxObjects gets a reference to the given NullableInt64 and assigns it to the MaxObjects field.
func (o *BucketInfoQuotas) SetMaxObjects(v int64) {
	o.MaxObjects.Set(&v)
}
// SetMaxObjectsNil sets the value for MaxObjects to be an explicit nil
func (o *BucketInfoQuotas) SetMaxObjectsNil() {
	o.MaxObjects.Set(nil)
}

// UnsetMaxObjects ensures that no value is present for MaxObjects, not even an explicit nil
func (o *BucketInfoQuotas) UnsetMaxObjects() {
	o.MaxObjects.Unset()
}

func (o BucketInfoQuotas) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BucketInfoQuotas) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.MaxSize.IsSet() {
		toSerialize["maxSize"] = o.MaxSize.Get()
	}
	if o.MaxObjects.IsSet() {
		toSerialize["maxObjects"] = o.MaxObjects.Get()
	}
	return toSerialize, nil
}

type NullableBucketInfoQuotas struct {
	value *BucketInfoQuotas
	isSet bool
}

func (v NullableBucketInfoQuotas) Get() *BucketInfoQuotas {
	return v.value
}

func (v *NullableBucketInfoQuotas) Set(val *BucketInfoQuotas) {
	v.value = val
	v.isSet = true
}

func (v NullableBucketInfoQuotas) IsSet() bool {
	return v.isSet
}

func (v *NullableBucketInfoQuotas) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBucketInfoQuotas(val *BucketInfoQuotas) *NullableBucketInfoQuotas {
	return &NullableBucketInfoQuotas{value: val, isSet: true}
}

func (v NullableBucketInfoQuotas) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBucketInfoQuotas) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


