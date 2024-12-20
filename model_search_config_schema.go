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

// checks if the SearchConfigSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchConfigSchema{}

// SearchConfigSchema contains index configurations fields
type SearchConfigSchema struct {
	// index name, to create or where update configuration
	Index string `json:"index"`
	// field which value will be used as document id [#DOCGENBUG REQUIRED FIELD]
	DocumentKeyName *string `json:"documentKeyName,omitempty"`
	// fields that can be stored into index and later retrieved
	Attributes []SearchConfigSchemaAttribute `json:"attributes"`
	// fields that can be used for fulltext searches
	Searchables []SearchConfigSchemaSearchable `json:"searchables,omitempty"`
	// fields that can be used for aggregations
	Facets []SearchConfigSchemaFacet `json:"facets,omitempty"`
	// fields that can be used for filtering
	Filters []SearchConfigSchemaFilter `json:"filters,omitempty"`
	// fields that can be used for sorting
	Sortables            []SearchConfigSchemaSortable `json:"sortables,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SearchConfigSchema SearchConfigSchema

// NewSearchConfigSchema instantiates a new SearchConfigSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchConfigSchema(index string, attributes []SearchConfigSchemaAttribute) *SearchConfigSchema {
	this := SearchConfigSchema{}
	this.Index = index
	this.Attributes = attributes
	return &this
}

// NewSearchConfigSchemaWithDefaults instantiates a new SearchConfigSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchConfigSchemaWithDefaults() *SearchConfigSchema {
	this := SearchConfigSchema{}
	return &this
}

// GetIndex returns the Index field value
func (o *SearchConfigSchema) GetIndex() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *SearchConfigSchema) GetIndexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *SearchConfigSchema) SetIndex(v string) {
	o.Index = v
}

// GetDocumentKeyName returns the DocumentKeyName field value if set, zero value otherwise.
func (o *SearchConfigSchema) GetDocumentKeyName() string {
	if o == nil || IsNil(o.DocumentKeyName) {
		var ret string
		return ret
	}
	return *o.DocumentKeyName
}

// GetDocumentKeyNameOk returns a tuple with the DocumentKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchConfigSchema) GetDocumentKeyNameOk() (*string, bool) {
	if o == nil || IsNil(o.DocumentKeyName) {
		return nil, false
	}
	return o.DocumentKeyName, true
}

// HasDocumentKeyName returns a boolean if a field has been set.
func (o *SearchConfigSchema) HasDocumentKeyName() bool {
	if o != nil && !IsNil(o.DocumentKeyName) {
		return true
	}

	return false
}

// SetDocumentKeyName gets a reference to the given string and assigns it to the DocumentKeyName field.
func (o *SearchConfigSchema) SetDocumentKeyName(v string) {
	o.DocumentKeyName = &v
}

// GetAttributes returns the Attributes field value
func (o *SearchConfigSchema) GetAttributes() []SearchConfigSchemaAttribute {
	if o == nil {
		var ret []SearchConfigSchemaAttribute
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *SearchConfigSchema) GetAttributesOk() ([]SearchConfigSchemaAttribute, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attributes, true
}

// SetAttributes sets field value
func (o *SearchConfigSchema) SetAttributes(v []SearchConfigSchemaAttribute) {
	o.Attributes = v
}

// GetSearchables returns the Searchables field value if set, zero value otherwise.
func (o *SearchConfigSchema) GetSearchables() []SearchConfigSchemaSearchable {
	if o == nil || IsNil(o.Searchables) {
		var ret []SearchConfigSchemaSearchable
		return ret
	}
	return o.Searchables
}

// GetSearchablesOk returns a tuple with the Searchables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchConfigSchema) GetSearchablesOk() ([]SearchConfigSchemaSearchable, bool) {
	if o == nil || IsNil(o.Searchables) {
		return nil, false
	}
	return o.Searchables, true
}

// HasSearchables returns a boolean if a field has been set.
func (o *SearchConfigSchema) HasSearchables() bool {
	if o != nil && !IsNil(o.Searchables) {
		return true
	}

	return false
}

// SetSearchables gets a reference to the given []SearchConfigSchemaSearchable and assigns it to the Searchables field.
func (o *SearchConfigSchema) SetSearchables(v []SearchConfigSchemaSearchable) {
	o.Searchables = v
}

// GetFacets returns the Facets field value if set, zero value otherwise.
func (o *SearchConfigSchema) GetFacets() []SearchConfigSchemaFacet {
	if o == nil || IsNil(o.Facets) {
		var ret []SearchConfigSchemaFacet
		return ret
	}
	return o.Facets
}

// GetFacetsOk returns a tuple with the Facets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchConfigSchema) GetFacetsOk() ([]SearchConfigSchemaFacet, bool) {
	if o == nil || IsNil(o.Facets) {
		return nil, false
	}
	return o.Facets, true
}

// HasFacets returns a boolean if a field has been set.
func (o *SearchConfigSchema) HasFacets() bool {
	if o != nil && !IsNil(o.Facets) {
		return true
	}

	return false
}

// SetFacets gets a reference to the given []SearchConfigSchemaFacet and assigns it to the Facets field.
func (o *SearchConfigSchema) SetFacets(v []SearchConfigSchemaFacet) {
	o.Facets = v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *SearchConfigSchema) GetFilters() []SearchConfigSchemaFilter {
	if o == nil || IsNil(o.Filters) {
		var ret []SearchConfigSchemaFilter
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchConfigSchema) GetFiltersOk() ([]SearchConfigSchemaFilter, bool) {
	if o == nil || IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *SearchConfigSchema) HasFilters() bool {
	if o != nil && !IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given []SearchConfigSchemaFilter and assigns it to the Filters field.
func (o *SearchConfigSchema) SetFilters(v []SearchConfigSchemaFilter) {
	o.Filters = v
}

// GetSortables returns the Sortables field value if set, zero value otherwise.
func (o *SearchConfigSchema) GetSortables() []SearchConfigSchemaSortable {
	if o == nil || IsNil(o.Sortables) {
		var ret []SearchConfigSchemaSortable
		return ret
	}
	return o.Sortables
}

// GetSortablesOk returns a tuple with the Sortables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchConfigSchema) GetSortablesOk() ([]SearchConfigSchemaSortable, bool) {
	if o == nil || IsNil(o.Sortables) {
		return nil, false
	}
	return o.Sortables, true
}

// HasSortables returns a boolean if a field has been set.
func (o *SearchConfigSchema) HasSortables() bool {
	if o != nil && !IsNil(o.Sortables) {
		return true
	}

	return false
}

// SetSortables gets a reference to the given []SearchConfigSchemaSortable and assigns it to the Sortables field.
func (o *SearchConfigSchema) SetSortables(v []SearchConfigSchemaSortable) {
	o.Sortables = v
}

func (o SearchConfigSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchConfigSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["index"] = o.Index
	if !IsNil(o.DocumentKeyName) {
		toSerialize["documentKeyName"] = o.DocumentKeyName
	}
	toSerialize["attributes"] = o.Attributes
	if !IsNil(o.Searchables) {
		toSerialize["searchables"] = o.Searchables
	}
	if !IsNil(o.Facets) {
		toSerialize["facets"] = o.Facets
	}
	if !IsNil(o.Filters) {
		toSerialize["filters"] = o.Filters
	}
	if !IsNil(o.Sortables) {
		toSerialize["sortables"] = o.Sortables
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SearchConfigSchema) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"index",
		"attributes",
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

	varSearchConfigSchema := _SearchConfigSchema{}

	err = json.Unmarshal(data, &varSearchConfigSchema)

	if err != nil {
		return err
	}

	*o = SearchConfigSchema(varSearchConfigSchema)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "index")
		delete(additionalProperties, "documentKeyName")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "searchables")
		delete(additionalProperties, "facets")
		delete(additionalProperties, "filters")
		delete(additionalProperties, "sortables")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *SearchConfigSchema) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *SearchConfigSchema) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableSearchConfigSchema struct {
	value *SearchConfigSchema
	isSet bool
}

func (v NullableSearchConfigSchema) Get() *SearchConfigSchema {
	return v.value
}

func (v *NullableSearchConfigSchema) Set(val *SearchConfigSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchConfigSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchConfigSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchConfigSchema(val *SearchConfigSchema) *NullableSearchConfigSchema {
	return &NullableSearchConfigSchema{value: val, isSet: true}
}

func (v NullableSearchConfigSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchConfigSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
