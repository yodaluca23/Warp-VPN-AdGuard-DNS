/*
 * untitled API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 536
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// UpdateAccountRequest struct for UpdateAccountRequest
type UpdateAccountRequest struct {
	License string `json:"license"`
}

// NewUpdateAccountRequest instantiates a new UpdateAccountRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAccountRequest(license string, ) *UpdateAccountRequest {
	this := UpdateAccountRequest{}
	this.License = license
	return &this
}

// NewUpdateAccountRequestWithDefaults instantiates a new UpdateAccountRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAccountRequestWithDefaults() *UpdateAccountRequest {
	this := UpdateAccountRequest{}
	return &this
}

// GetLicense returns the License field value
func (o *UpdateAccountRequest) GetLicense() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.License
}

// GetLicenseOk returns a tuple with the License field value
// and a boolean to check if the value has been set.
func (o *UpdateAccountRequest) GetLicenseOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.License, true
}

// SetLicense sets field value
func (o *UpdateAccountRequest) SetLicense(v string) {
	o.License = v
}

func (o UpdateAccountRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["license"] = o.License
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateAccountRequest struct {
	value *UpdateAccountRequest
	isSet bool
}

func (v NullableUpdateAccountRequest) Get() *UpdateAccountRequest {
	return v.value
}

func (v *NullableUpdateAccountRequest) Set(val *UpdateAccountRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAccountRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAccountRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAccountRequest(val *UpdateAccountRequest) *NullableUpdateAccountRequest {
	return &NullableUpdateAccountRequest{value: val, isSet: true}
}

func (v NullableUpdateAccountRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAccountRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

