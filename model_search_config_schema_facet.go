/*
Search Service

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
Contact: info@gemini-commerce.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package search

import (
	"encoding/json"
	"fmt"
)

// checks if the SearchConfigSchemaFacet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchConfigSchemaFacet{}

// SearchConfigSchemaFacet contains the configurations needed to define an aggregation (Facet) using defined index attributes
type SearchConfigSchemaFacet struct {
	// json array of attributes of index that you can use as aggregation to make facets
	Attribute            string `json:"attribute"`
	AdditionalProperties map[string]interface{}
}

type _SearchConfigSchemaFacet SearchConfigSchemaFacet

// NewSearchConfigSchemaFacet instantiates a new SearchConfigSchemaFacet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchConfigSchemaFacet(attribute string) *SearchConfigSchemaFacet {
	this := SearchConfigSchemaFacet{}
	this.Attribute = attribute
	return &this
}

// NewSearchConfigSchemaFacetWithDefaults instantiates a new SearchConfigSchemaFacet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchConfigSchemaFacetWithDefaults() *SearchConfigSchemaFacet {
	this := SearchConfigSchemaFacet{}
	return &this
}

// GetAttribute returns the Attribute field value
func (o *SearchConfigSchemaFacet) GetAttribute() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Attribute
}

// GetAttributeOk returns a tuple with the Attribute field value
// and a boolean to check if the value has been set.
func (o *SearchConfigSchemaFacet) GetAttributeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attribute, true
}

// SetAttribute sets field value
func (o *SearchConfigSchemaFacet) SetAttribute(v string) {
	o.Attribute = v
}

func (o SearchConfigSchemaFacet) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchConfigSchemaFacet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["attribute"] = o.Attribute

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SearchConfigSchemaFacet) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"attribute",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSearchConfigSchemaFacet := _SearchConfigSchemaFacet{}

	err = json.Unmarshal(data, &varSearchConfigSchemaFacet)

	if err != nil {
		return err
	}

	*o = SearchConfigSchemaFacet(varSearchConfigSchemaFacet)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "attribute")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *SearchConfigSchemaFacet) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *SearchConfigSchemaFacet) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableSearchConfigSchemaFacet struct {
	value *SearchConfigSchemaFacet
	isSet bool
}

func (v NullableSearchConfigSchemaFacet) Get() *SearchConfigSchemaFacet {
	return v.value
}

func (v *NullableSearchConfigSchemaFacet) Set(val *SearchConfigSchemaFacet) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchConfigSchemaFacet) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchConfigSchemaFacet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchConfigSchemaFacet(val *SearchConfigSchemaFacet) *NullableSearchConfigSchemaFacet {
	return &NullableSearchConfigSchemaFacet{value: val, isSet: true}
}

func (v NullableSearchConfigSchemaFacet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchConfigSchemaFacet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}