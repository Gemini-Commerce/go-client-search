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

// checks if the SearchConfigSchemaSortable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchConfigSchemaSortable{}

// SearchConfigSchemaSortable contains the configurations needed to define which index attributes can be used for sort
type SearchConfigSchemaSortable struct {
	// json array of attributes of index which can be used for sort
	Attribute            string `json:"attribute"`
	AdditionalProperties map[string]interface{}
}

type _SearchConfigSchemaSortable SearchConfigSchemaSortable

// NewSearchConfigSchemaSortable instantiates a new SearchConfigSchemaSortable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchConfigSchemaSortable(attribute string) *SearchConfigSchemaSortable {
	this := SearchConfigSchemaSortable{}
	this.Attribute = attribute
	return &this
}

// NewSearchConfigSchemaSortableWithDefaults instantiates a new SearchConfigSchemaSortable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchConfigSchemaSortableWithDefaults() *SearchConfigSchemaSortable {
	this := SearchConfigSchemaSortable{}
	return &this
}

// GetAttribute returns the Attribute field value
func (o *SearchConfigSchemaSortable) GetAttribute() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Attribute
}

// GetAttributeOk returns a tuple with the Attribute field value
// and a boolean to check if the value has been set.
func (o *SearchConfigSchemaSortable) GetAttributeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attribute, true
}

// SetAttribute sets field value
func (o *SearchConfigSchemaSortable) SetAttribute(v string) {
	o.Attribute = v
}

func (o SearchConfigSchemaSortable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchConfigSchemaSortable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["attribute"] = o.Attribute

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SearchConfigSchemaSortable) UnmarshalJSON(data []byte) (err error) {
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

	varSearchConfigSchemaSortable := _SearchConfigSchemaSortable{}

	err = json.Unmarshal(data, &varSearchConfigSchemaSortable)

	if err != nil {
		return err
	}

	*o = SearchConfigSchemaSortable(varSearchConfigSchemaSortable)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "attribute")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *SearchConfigSchemaSortable) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *SearchConfigSchemaSortable) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableSearchConfigSchemaSortable struct {
	value *SearchConfigSchemaSortable
	isSet bool
}

func (v NullableSearchConfigSchemaSortable) Get() *SearchConfigSchemaSortable {
	return v.value
}

func (v *NullableSearchConfigSchemaSortable) Set(val *SearchConfigSchemaSortable) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchConfigSchemaSortable) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchConfigSchemaSortable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchConfigSchemaSortable(val *SearchConfigSchemaSortable) *NullableSearchConfigSchemaSortable {
	return &NullableSearchConfigSchemaSortable{value: val, isSet: true}
}

func (v NullableSearchConfigSchemaSortable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchConfigSchemaSortable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}