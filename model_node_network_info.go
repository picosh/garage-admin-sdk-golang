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

// checks if the NodeNetworkInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NodeNetworkInfo{}

// NodeNetworkInfo struct for NodeNetworkInfo
type NodeNetworkInfo struct {
	Id *string `json:"id,omitempty"`
	Addr string `json:"addr"`
	IsUp bool `json:"isUp"`
	LastSeenSecsAgo NullableInt32 `json:"lastSeenSecsAgo"`
	Hostname string `json:"hostname"`
}

type _NodeNetworkInfo NodeNetworkInfo

// NewNodeNetworkInfo instantiates a new NodeNetworkInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNodeNetworkInfo(addr string, isUp bool, lastSeenSecsAgo NullableInt32, hostname string) *NodeNetworkInfo {
	this := NodeNetworkInfo{}
	this.Addr = addr
	this.IsUp = isUp
	this.LastSeenSecsAgo = lastSeenSecsAgo
	this.Hostname = hostname
	return &this
}

// NewNodeNetworkInfoWithDefaults instantiates a new NodeNetworkInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodeNetworkInfoWithDefaults() *NodeNetworkInfo {
	this := NodeNetworkInfo{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NodeNetworkInfo) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeNetworkInfo) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NodeNetworkInfo) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NodeNetworkInfo) SetId(v string) {
	o.Id = &v
}

// GetAddr returns the Addr field value
func (o *NodeNetworkInfo) GetAddr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Addr
}

// GetAddrOk returns a tuple with the Addr field value
// and a boolean to check if the value has been set.
func (o *NodeNetworkInfo) GetAddrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Addr, true
}

// SetAddr sets field value
func (o *NodeNetworkInfo) SetAddr(v string) {
	o.Addr = v
}

// GetIsUp returns the IsUp field value
func (o *NodeNetworkInfo) GetIsUp() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsUp
}

// GetIsUpOk returns a tuple with the IsUp field value
// and a boolean to check if the value has been set.
func (o *NodeNetworkInfo) GetIsUpOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsUp, true
}

// SetIsUp sets field value
func (o *NodeNetworkInfo) SetIsUp(v bool) {
	o.IsUp = v
}

// GetLastSeenSecsAgo returns the LastSeenSecsAgo field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *NodeNetworkInfo) GetLastSeenSecsAgo() int32 {
	if o == nil || o.LastSeenSecsAgo.Get() == nil {
		var ret int32
		return ret
	}

	return *o.LastSeenSecsAgo.Get()
}

// GetLastSeenSecsAgoOk returns a tuple with the LastSeenSecsAgo field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NodeNetworkInfo) GetLastSeenSecsAgoOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastSeenSecsAgo.Get(), o.LastSeenSecsAgo.IsSet()
}

// SetLastSeenSecsAgo sets field value
func (o *NodeNetworkInfo) SetLastSeenSecsAgo(v int32) {
	o.LastSeenSecsAgo.Set(&v)
}

// GetHostname returns the Hostname field value
func (o *NodeNetworkInfo) GetHostname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value
// and a boolean to check if the value has been set.
func (o *NodeNetworkInfo) GetHostnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hostname, true
}

// SetHostname sets field value
func (o *NodeNetworkInfo) SetHostname(v string) {
	o.Hostname = v
}

func (o NodeNetworkInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NodeNetworkInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["addr"] = o.Addr
	toSerialize["isUp"] = o.IsUp
	toSerialize["lastSeenSecsAgo"] = o.LastSeenSecsAgo.Get()
	toSerialize["hostname"] = o.Hostname
	return toSerialize, nil
}

func (o *NodeNetworkInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"addr",
		"isUp",
		"lastSeenSecsAgo",
		"hostname",
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

	varNodeNetworkInfo := _NodeNetworkInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNodeNetworkInfo)

	if err != nil {
		return err
	}

	*o = NodeNetworkInfo(varNodeNetworkInfo)

	return err
}

type NullableNodeNetworkInfo struct {
	value *NodeNetworkInfo
	isSet bool
}

func (v NullableNodeNetworkInfo) Get() *NodeNetworkInfo {
	return v.value
}

func (v *NullableNodeNetworkInfo) Set(val *NodeNetworkInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableNodeNetworkInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableNodeNetworkInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodeNetworkInfo(val *NodeNetworkInfo) *NullableNodeNetworkInfo {
	return &NullableNodeNetworkInfo{value: val, isSet: true}
}

func (v NullableNodeNetworkInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodeNetworkInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


