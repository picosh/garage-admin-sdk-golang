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

// checks if the CreateBucketRequestLocalAliasAllow type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateBucketRequestLocalAliasAllow{}

// CreateBucketRequestLocalAliasAllow struct for CreateBucketRequestLocalAliasAllow
type CreateBucketRequestLocalAliasAllow struct {
	Read *bool `json:"read,omitempty"`
	Write *bool `json:"write,omitempty"`
	Owner *bool `json:"owner,omitempty"`
}

// NewCreateBucketRequestLocalAliasAllow instantiates a new CreateBucketRequestLocalAliasAllow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateBucketRequestLocalAliasAllow() *CreateBucketRequestLocalAliasAllow {
	this := CreateBucketRequestLocalAliasAllow{}
	return &this
}

// NewCreateBucketRequestLocalAliasAllowWithDefaults instantiates a new CreateBucketRequestLocalAliasAllow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateBucketRequestLocalAliasAllowWithDefaults() *CreateBucketRequestLocalAliasAllow {
	this := CreateBucketRequestLocalAliasAllow{}
	return &this
}

// GetRead returns the Read field value if set, zero value otherwise.
func (o *CreateBucketRequestLocalAliasAllow) GetRead() bool {
	if o == nil || IsNil(o.Read) {
		var ret bool
		return ret
	}
	return *o.Read
}

// GetReadOk returns a tuple with the Read field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBucketRequestLocalAliasAllow) GetReadOk() (*bool, bool) {
	if o == nil || IsNil(o.Read) {
		return nil, false
	}
	return o.Read, true
}

// HasRead returns a boolean if a field has been set.
func (o *CreateBucketRequestLocalAliasAllow) HasRead() bool {
	if o != nil && !IsNil(o.Read) {
		return true
	}

	return false
}

// SetRead gets a reference to the given bool and assigns it to the Read field.
func (o *CreateBucketRequestLocalAliasAllow) SetRead(v bool) {
	o.Read = &v
}

// GetWrite returns the Write field value if set, zero value otherwise.
func (o *CreateBucketRequestLocalAliasAllow) GetWrite() bool {
	if o == nil || IsNil(o.Write) {
		var ret bool
		return ret
	}
	return *o.Write
}

// GetWriteOk returns a tuple with the Write field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBucketRequestLocalAliasAllow) GetWriteOk() (*bool, bool) {
	if o == nil || IsNil(o.Write) {
		return nil, false
	}
	return o.Write, true
}

// HasWrite returns a boolean if a field has been set.
func (o *CreateBucketRequestLocalAliasAllow) HasWrite() bool {
	if o != nil && !IsNil(o.Write) {
		return true
	}

	return false
}

// SetWrite gets a reference to the given bool and assigns it to the Write field.
func (o *CreateBucketRequestLocalAliasAllow) SetWrite(v bool) {
	o.Write = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *CreateBucketRequestLocalAliasAllow) GetOwner() bool {
	if o == nil || IsNil(o.Owner) {
		var ret bool
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBucketRequestLocalAliasAllow) GetOwnerOk() (*bool, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *CreateBucketRequestLocalAliasAllow) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given bool and assigns it to the Owner field.
func (o *CreateBucketRequestLocalAliasAllow) SetOwner(v bool) {
	o.Owner = &v
}

func (o CreateBucketRequestLocalAliasAllow) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateBucketRequestLocalAliasAllow) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Read) {
		toSerialize["read"] = o.Read
	}
	if !IsNil(o.Write) {
		toSerialize["write"] = o.Write
	}
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	return toSerialize, nil
}

type NullableCreateBucketRequestLocalAliasAllow struct {
	value *CreateBucketRequestLocalAliasAllow
	isSet bool
}

func (v NullableCreateBucketRequestLocalAliasAllow) Get() *CreateBucketRequestLocalAliasAllow {
	return v.value
}

func (v *NullableCreateBucketRequestLocalAliasAllow) Set(val *CreateBucketRequestLocalAliasAllow) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateBucketRequestLocalAliasAllow) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateBucketRequestLocalAliasAllow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateBucketRequestLocalAliasAllow(val *CreateBucketRequestLocalAliasAllow) *NullableCreateBucketRequestLocalAliasAllow {
	return &NullableCreateBucketRequestLocalAliasAllow{value: val, isSet: true}
}

func (v NullableCreateBucketRequestLocalAliasAllow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateBucketRequestLocalAliasAllow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


