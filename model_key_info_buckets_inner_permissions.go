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

// checks if the KeyInfoBucketsInnerPermissions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyInfoBucketsInnerPermissions{}

// KeyInfoBucketsInnerPermissions struct for KeyInfoBucketsInnerPermissions
type KeyInfoBucketsInnerPermissions struct {
	Read *bool `json:"read,omitempty"`
	Write *bool `json:"write,omitempty"`
	Owner *bool `json:"owner,omitempty"`
}

// NewKeyInfoBucketsInnerPermissions instantiates a new KeyInfoBucketsInnerPermissions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyInfoBucketsInnerPermissions() *KeyInfoBucketsInnerPermissions {
	this := KeyInfoBucketsInnerPermissions{}
	return &this
}

// NewKeyInfoBucketsInnerPermissionsWithDefaults instantiates a new KeyInfoBucketsInnerPermissions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyInfoBucketsInnerPermissionsWithDefaults() *KeyInfoBucketsInnerPermissions {
	this := KeyInfoBucketsInnerPermissions{}
	return &this
}

// GetRead returns the Read field value if set, zero value otherwise.
func (o *KeyInfoBucketsInnerPermissions) GetRead() bool {
	if o == nil || IsNil(o.Read) {
		var ret bool
		return ret
	}
	return *o.Read
}

// GetReadOk returns a tuple with the Read field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyInfoBucketsInnerPermissions) GetReadOk() (*bool, bool) {
	if o == nil || IsNil(o.Read) {
		return nil, false
	}
	return o.Read, true
}

// HasRead returns a boolean if a field has been set.
func (o *KeyInfoBucketsInnerPermissions) HasRead() bool {
	if o != nil && !IsNil(o.Read) {
		return true
	}

	return false
}

// SetRead gets a reference to the given bool and assigns it to the Read field.
func (o *KeyInfoBucketsInnerPermissions) SetRead(v bool) {
	o.Read = &v
}

// GetWrite returns the Write field value if set, zero value otherwise.
func (o *KeyInfoBucketsInnerPermissions) GetWrite() bool {
	if o == nil || IsNil(o.Write) {
		var ret bool
		return ret
	}
	return *o.Write
}

// GetWriteOk returns a tuple with the Write field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyInfoBucketsInnerPermissions) GetWriteOk() (*bool, bool) {
	if o == nil || IsNil(o.Write) {
		return nil, false
	}
	return o.Write, true
}

// HasWrite returns a boolean if a field has been set.
func (o *KeyInfoBucketsInnerPermissions) HasWrite() bool {
	if o != nil && !IsNil(o.Write) {
		return true
	}

	return false
}

// SetWrite gets a reference to the given bool and assigns it to the Write field.
func (o *KeyInfoBucketsInnerPermissions) SetWrite(v bool) {
	o.Write = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *KeyInfoBucketsInnerPermissions) GetOwner() bool {
	if o == nil || IsNil(o.Owner) {
		var ret bool
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyInfoBucketsInnerPermissions) GetOwnerOk() (*bool, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *KeyInfoBucketsInnerPermissions) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given bool and assigns it to the Owner field.
func (o *KeyInfoBucketsInnerPermissions) SetOwner(v bool) {
	o.Owner = &v
}

func (o KeyInfoBucketsInnerPermissions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyInfoBucketsInnerPermissions) ToMap() (map[string]interface{}, error) {
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

type NullableKeyInfoBucketsInnerPermissions struct {
	value *KeyInfoBucketsInnerPermissions
	isSet bool
}

func (v NullableKeyInfoBucketsInnerPermissions) Get() *KeyInfoBucketsInnerPermissions {
	return v.value
}

func (v *NullableKeyInfoBucketsInnerPermissions) Set(val *KeyInfoBucketsInnerPermissions) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyInfoBucketsInnerPermissions) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyInfoBucketsInnerPermissions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyInfoBucketsInnerPermissions(val *KeyInfoBucketsInnerPermissions) *NullableKeyInfoBucketsInnerPermissions {
	return &NullableKeyInfoBucketsInnerPermissions{value: val, isSet: true}
}

func (v NullableKeyInfoBucketsInnerPermissions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyInfoBucketsInnerPermissions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


