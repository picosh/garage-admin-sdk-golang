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

// checks if the GetNodes200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNodes200Response{}

// GetNodes200Response struct for GetNodes200Response
type GetNodes200Response struct {
	Node string `json:"node"`
	GarageVersion string `json:"garageVersion"`
	GarageFeatures []string `json:"garageFeatures"`
	RustVersion string `json:"rustVersion"`
	DbEngine string `json:"dbEngine"`
	KnownNodes []NodeNetworkInfo `json:"knownNodes"`
	Layout ClusterLayout `json:"layout"`
}

type _GetNodes200Response GetNodes200Response

// NewGetNodes200Response instantiates a new GetNodes200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNodes200Response(node string, garageVersion string, garageFeatures []string, rustVersion string, dbEngine string, knownNodes []NodeNetworkInfo, layout ClusterLayout) *GetNodes200Response {
	this := GetNodes200Response{}
	this.Node = node
	this.GarageVersion = garageVersion
	this.GarageFeatures = garageFeatures
	this.RustVersion = rustVersion
	this.DbEngine = dbEngine
	this.KnownNodes = knownNodes
	this.Layout = layout
	return &this
}

// NewGetNodes200ResponseWithDefaults instantiates a new GetNodes200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNodes200ResponseWithDefaults() *GetNodes200Response {
	this := GetNodes200Response{}
	return &this
}

// GetNode returns the Node field value
func (o *GetNodes200Response) GetNode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Node
}

// GetNodeOk returns a tuple with the Node field value
// and a boolean to check if the value has been set.
func (o *GetNodes200Response) GetNodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Node, true
}

// SetNode sets field value
func (o *GetNodes200Response) SetNode(v string) {
	o.Node = v
}

// GetGarageVersion returns the GarageVersion field value
func (o *GetNodes200Response) GetGarageVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GarageVersion
}

// GetGarageVersionOk returns a tuple with the GarageVersion field value
// and a boolean to check if the value has been set.
func (o *GetNodes200Response) GetGarageVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GarageVersion, true
}

// SetGarageVersion sets field value
func (o *GetNodes200Response) SetGarageVersion(v string) {
	o.GarageVersion = v
}

// GetGarageFeatures returns the GarageFeatures field value
func (o *GetNodes200Response) GetGarageFeatures() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.GarageFeatures
}

// GetGarageFeaturesOk returns a tuple with the GarageFeatures field value
// and a boolean to check if the value has been set.
func (o *GetNodes200Response) GetGarageFeaturesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GarageFeatures, true
}

// SetGarageFeatures sets field value
func (o *GetNodes200Response) SetGarageFeatures(v []string) {
	o.GarageFeatures = v
}

// GetRustVersion returns the RustVersion field value
func (o *GetNodes200Response) GetRustVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RustVersion
}

// GetRustVersionOk returns a tuple with the RustVersion field value
// and a boolean to check if the value has been set.
func (o *GetNodes200Response) GetRustVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RustVersion, true
}

// SetRustVersion sets field value
func (o *GetNodes200Response) SetRustVersion(v string) {
	o.RustVersion = v
}

// GetDbEngine returns the DbEngine field value
func (o *GetNodes200Response) GetDbEngine() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DbEngine
}

// GetDbEngineOk returns a tuple with the DbEngine field value
// and a boolean to check if the value has been set.
func (o *GetNodes200Response) GetDbEngineOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DbEngine, true
}

// SetDbEngine sets field value
func (o *GetNodes200Response) SetDbEngine(v string) {
	o.DbEngine = v
}

// GetKnownNodes returns the KnownNodes field value
func (o *GetNodes200Response) GetKnownNodes() []NodeNetworkInfo {
	if o == nil {
		var ret []NodeNetworkInfo
		return ret
	}

	return o.KnownNodes
}

// GetKnownNodesOk returns a tuple with the KnownNodes field value
// and a boolean to check if the value has been set.
func (o *GetNodes200Response) GetKnownNodesOk() ([]NodeNetworkInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.KnownNodes, true
}

// SetKnownNodes sets field value
func (o *GetNodes200Response) SetKnownNodes(v []NodeNetworkInfo) {
	o.KnownNodes = v
}

// GetLayout returns the Layout field value
func (o *GetNodes200Response) GetLayout() ClusterLayout {
	if o == nil {
		var ret ClusterLayout
		return ret
	}

	return o.Layout
}

// GetLayoutOk returns a tuple with the Layout field value
// and a boolean to check if the value has been set.
func (o *GetNodes200Response) GetLayoutOk() (*ClusterLayout, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Layout, true
}

// SetLayout sets field value
func (o *GetNodes200Response) SetLayout(v ClusterLayout) {
	o.Layout = v
}

func (o GetNodes200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNodes200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["node"] = o.Node
	toSerialize["garageVersion"] = o.GarageVersion
	toSerialize["garageFeatures"] = o.GarageFeatures
	toSerialize["rustVersion"] = o.RustVersion
	toSerialize["dbEngine"] = o.DbEngine
	toSerialize["knownNodes"] = o.KnownNodes
	toSerialize["layout"] = o.Layout
	return toSerialize, nil
}

func (o *GetNodes200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"node",
		"garageVersion",
		"garageFeatures",
		"rustVersion",
		"dbEngine",
		"knownNodes",
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

	varGetNodes200Response := _GetNodes200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetNodes200Response)

	if err != nil {
		return err
	}

	*o = GetNodes200Response(varGetNodes200Response)

	return err
}

type NullableGetNodes200Response struct {
	value *GetNodes200Response
	isSet bool
}

func (v NullableGetNodes200Response) Get() *GetNodes200Response {
	return v.value
}

func (v *NullableGetNodes200Response) Set(val *GetNodes200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNodes200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNodes200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNodes200Response(val *GetNodes200Response) *NullableGetNodes200Response {
	return &NullableGetNodes200Response{value: val, isSet: true}
}

func (v NullableGetNodes200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNodes200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


