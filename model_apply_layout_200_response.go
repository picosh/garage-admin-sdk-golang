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

// checks if the ApplyLayout200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplyLayout200Response{}

// ApplyLayout200Response struct for ApplyLayout200Response
type ApplyLayout200Response struct {
	Message []string `json:"message"`
	Layout ClusterLayout `json:"layout"`
}

type _ApplyLayout200Response ApplyLayout200Response

// NewApplyLayout200Response instantiates a new ApplyLayout200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplyLayout200Response(message []string, layout ClusterLayout) *ApplyLayout200Response {
	this := ApplyLayout200Response{}
	this.Message = message
	this.Layout = layout
	return &this
}

// NewApplyLayout200ResponseWithDefaults instantiates a new ApplyLayout200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplyLayout200ResponseWithDefaults() *ApplyLayout200Response {
	this := ApplyLayout200Response{}
	return &this
}

// GetMessage returns the Message field value
func (o *ApplyLayout200Response) GetMessage() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ApplyLayout200Response) GetMessageOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message, true
}

// SetMessage sets field value
func (o *ApplyLayout200Response) SetMessage(v []string) {
	o.Message = v
}

// GetLayout returns the Layout field value
func (o *ApplyLayout200Response) GetLayout() ClusterLayout {
	if o == nil {
		var ret ClusterLayout
		return ret
	}

	return o.Layout
}

// GetLayoutOk returns a tuple with the Layout field value
// and a boolean to check if the value has been set.
func (o *ApplyLayout200Response) GetLayoutOk() (*ClusterLayout, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Layout, true
}

// SetLayout sets field value
func (o *ApplyLayout200Response) SetLayout(v ClusterLayout) {
	o.Layout = v
}

func (o ApplyLayout200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplyLayout200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message"] = o.Message
	toSerialize["layout"] = o.Layout
	return toSerialize, nil
}

func (o *ApplyLayout200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"message",
		"layout",
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

	varApplyLayout200Response := _ApplyLayout200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApplyLayout200Response)

	if err != nil {
		return err
	}

	*o = ApplyLayout200Response(varApplyLayout200Response)

	return err
}

type NullableApplyLayout200Response struct {
	value *ApplyLayout200Response
	isSet bool
}

func (v NullableApplyLayout200Response) Get() *ApplyLayout200Response {
	return v.value
}

func (v *NullableApplyLayout200Response) Set(val *ApplyLayout200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableApplyLayout200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableApplyLayout200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplyLayout200Response(val *ApplyLayout200Response) *NullableApplyLayout200Response {
	return &NullableApplyLayout200Response{value: val, isSet: true}
}

func (v NullableApplyLayout200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplyLayout200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


