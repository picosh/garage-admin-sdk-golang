/*
Garage Administration API v0+garage-v0.9.0

Administrate your Garage cluster programatically, including status, layout, keys, buckets, and maintainance tasks.  *Disclaimer: The API is not stable yet, hence its v0 tag. The API can change at any time, and changes can include breaking backward compatibility. Read the changelog and upgrade your scripts before upgrading. Additionnaly, this specification is very early stage and can contain bugs, especially on error return codes/types that are not tested yet. Do not expect a well finished and polished product!* 

API version: v0.9.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package garage

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ImportKeyRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImportKeyRequest{}

// ImportKeyRequest struct for ImportKeyRequest
type ImportKeyRequest struct {
	Name NullableString `json:"name"`
	AccessKeyId string `json:"accessKeyId"`
	SecretAccessKey string `json:"secretAccessKey"`
}

type _ImportKeyRequest ImportKeyRequest

// NewImportKeyRequest instantiates a new ImportKeyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportKeyRequest(name NullableString, accessKeyId string, secretAccessKey string) *ImportKeyRequest {
	this := ImportKeyRequest{}
	this.Name = name
	this.AccessKeyId = accessKeyId
	this.SecretAccessKey = secretAccessKey
	return &this
}

// NewImportKeyRequestWithDefaults instantiates a new ImportKeyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportKeyRequestWithDefaults() *ImportKeyRequest {
	this := ImportKeyRequest{}
	return &this
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ImportKeyRequest) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportKeyRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *ImportKeyRequest) SetName(v string) {
	o.Name.Set(&v)
}

// GetAccessKeyId returns the AccessKeyId field value
func (o *ImportKeyRequest) GetAccessKeyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessKeyId
}

// GetAccessKeyIdOk returns a tuple with the AccessKeyId field value
// and a boolean to check if the value has been set.
func (o *ImportKeyRequest) GetAccessKeyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessKeyId, true
}

// SetAccessKeyId sets field value
func (o *ImportKeyRequest) SetAccessKeyId(v string) {
	o.AccessKeyId = v
}

// GetSecretAccessKey returns the SecretAccessKey field value
func (o *ImportKeyRequest) GetSecretAccessKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecretAccessKey
}

// GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field value
// and a boolean to check if the value has been set.
func (o *ImportKeyRequest) GetSecretAccessKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecretAccessKey, true
}

// SetSecretAccessKey sets field value
func (o *ImportKeyRequest) SetSecretAccessKey(v string) {
	o.SecretAccessKey = v
}

func (o ImportKeyRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImportKeyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name.Get()
	toSerialize["accessKeyId"] = o.AccessKeyId
	toSerialize["secretAccessKey"] = o.SecretAccessKey
	return toSerialize, nil
}

func (o *ImportKeyRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"accessKeyId",
		"secretAccessKey",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varImportKeyRequest := _ImportKeyRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varImportKeyRequest)

	if err != nil {
		return err
	}

	*o = ImportKeyRequest(varImportKeyRequest)

	return err
}

type NullableImportKeyRequest struct {
	value *ImportKeyRequest
	isSet bool
}

func (v NullableImportKeyRequest) Get() *ImportKeyRequest {
	return v.value
}

func (v *NullableImportKeyRequest) Set(val *ImportKeyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableImportKeyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableImportKeyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportKeyRequest(val *ImportKeyRequest) *NullableImportKeyRequest {
	return &NullableImportKeyRequest{value: val, isSet: true}
}

func (v NullableImportKeyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportKeyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


