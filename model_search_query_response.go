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
)

// checks if the SearchQueryResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchQueryResponse{}

// SearchQueryResponse contains the search results and errors
type SearchQueryResponse struct {
	// array of SearchResult that contains search results per index [#DOCGENBUG REQUIRED FIELD]
	Result []SearchResult `json:"result,omitempty"`
	// array of error obtained during search process per index [#DOCGENBUG REQUIRED FIELD]
	Errors               []SearchQueryError `json:"errors,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SearchQueryResponse SearchQueryResponse

// NewSearchQueryResponse instantiates a new SearchQueryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchQueryResponse() *SearchQueryResponse {
	this := SearchQueryResponse{}
	return &this
}

// NewSearchQueryResponseWithDefaults instantiates a new SearchQueryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchQueryResponseWithDefaults() *SearchQueryResponse {
	this := SearchQueryResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *SearchQueryResponse) GetResult() []SearchResult {
	if o == nil || IsNil(o.Result) {
		var ret []SearchResult
		return ret
	}
	return o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchQueryResponse) GetResultOk() ([]SearchResult, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *SearchQueryResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given []SearchResult and assigns it to the Result field.
func (o *SearchQueryResponse) SetResult(v []SearchResult) {
	o.Result = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SearchQueryResponse) GetErrors() []SearchQueryError {
	if o == nil || IsNil(o.Errors) {
		var ret []SearchQueryError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchQueryResponse) GetErrorsOk() ([]SearchQueryError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SearchQueryResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SearchQueryError and assigns it to the Errors field.
func (o *SearchQueryResponse) SetErrors(v []SearchQueryError) {
	o.Errors = v
}

func (o SearchQueryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchQueryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SearchQueryResponse) UnmarshalJSON(data []byte) (err error) {
	varSearchQueryResponse := _SearchQueryResponse{}

	err = json.Unmarshal(data, &varSearchQueryResponse)

	if err != nil {
		return err
	}

	*o = SearchQueryResponse(varSearchQueryResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		delete(additionalProperties, "errors")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *SearchQueryResponse) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *SearchQueryResponse) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableSearchQueryResponse struct {
	value *SearchQueryResponse
	isSet bool
}

func (v NullableSearchQueryResponse) Get() *SearchQueryResponse {
	return v.value
}

func (v *NullableSearchQueryResponse) Set(val *SearchQueryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchQueryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchQueryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchQueryResponse(val *SearchQueryResponse) *NullableSearchQueryResponse {
	return &NullableSearchQueryResponse{value: val, isSet: true}
}

func (v NullableSearchQueryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchQueryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
