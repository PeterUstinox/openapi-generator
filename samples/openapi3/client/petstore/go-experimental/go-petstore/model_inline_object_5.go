/*
 * OpenAPI Petstore
 *
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package petstore

import (
	"bytes"
	"encoding/json"
	"os"
)

// InlineObject5 struct for InlineObject5
type InlineObject5 struct {
	// Additional data to pass to server
	AdditionalMetadata *string `json:"additionalMetadata,omitempty"`
	// file to upload
	RequiredFile *os.File `json:"requiredFile"`
}

// GetAdditionalMetadata returns the AdditionalMetadata field value if set, zero value otherwise.
func (o *InlineObject5) GetAdditionalMetadata() string {
	if o == nil || o.AdditionalMetadata == nil {
		var ret string
		return ret
	}
	return *o.AdditionalMetadata
}

// GetAdditionalMetadataOk returns a tuple with the AdditionalMetadata field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject5) GetAdditionalMetadataOk() (string, bool) {
	if o == nil || o.AdditionalMetadata == nil {
		var ret string
		return ret, false
	}
	return *o.AdditionalMetadata, true
}

// HasAdditionalMetadata returns a boolean if a field has been set.
func (o *InlineObject5) HasAdditionalMetadata() bool {
	if o != nil && o.AdditionalMetadata != nil {
		return true
	}

	return false
}

// SetAdditionalMetadata gets a reference to the given string and assigns it to the AdditionalMetadata field.
func (o *InlineObject5) SetAdditionalMetadata(v string) {
	o.AdditionalMetadata = &v
}

// GetRequiredFile returns the RequiredFile field value
func (o *InlineObject5) GetRequiredFile() *os.File {
	if o == nil {
		var ret *os.File
		return ret
	}

	return o.RequiredFile
}

// SetRequiredFile sets field value
func (o *InlineObject5) SetRequiredFile(v *os.File) {
	o.RequiredFile = v
}

type NullableInlineObject5 struct {
	Value InlineObject5
	ExplicitNull bool
}

func (v NullableInlineObject5) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableInlineObject5) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
