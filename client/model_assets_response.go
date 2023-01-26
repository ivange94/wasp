/*
Wasp API

REST API for the Wasp node

API version: 0.4.0-alpha.2-402-g2adcf666b
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// AssetsResponse struct for AssetsResponse
type AssetsResponse struct {
	BaseTokens *int64 `json:"baseTokens,omitempty"`
	NativeTokens []NativeToken `json:"nativeTokens,omitempty"`
}

// NewAssetsResponse instantiates a new AssetsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetsResponse() *AssetsResponse {
	this := AssetsResponse{}
	return &this
}

// NewAssetsResponseWithDefaults instantiates a new AssetsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetsResponseWithDefaults() *AssetsResponse {
	this := AssetsResponse{}
	return &this
}

// GetBaseTokens returns the BaseTokens field value if set, zero value otherwise.
func (o *AssetsResponse) GetBaseTokens() int64 {
	if o == nil || isNil(o.BaseTokens) {
		var ret int64
		return ret
	}
	return *o.BaseTokens
}

// GetBaseTokensOk returns a tuple with the BaseTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsResponse) GetBaseTokensOk() (*int64, bool) {
	if o == nil || isNil(o.BaseTokens) {
    return nil, false
	}
	return o.BaseTokens, true
}

// HasBaseTokens returns a boolean if a field has been set.
func (o *AssetsResponse) HasBaseTokens() bool {
	if o != nil && !isNil(o.BaseTokens) {
		return true
	}

	return false
}

// SetBaseTokens gets a reference to the given int64 and assigns it to the BaseTokens field.
func (o *AssetsResponse) SetBaseTokens(v int64) {
	o.BaseTokens = &v
}

// GetNativeTokens returns the NativeTokens field value if set, zero value otherwise.
func (o *AssetsResponse) GetNativeTokens() []NativeToken {
	if o == nil || isNil(o.NativeTokens) {
		var ret []NativeToken
		return ret
	}
	return o.NativeTokens
}

// GetNativeTokensOk returns a tuple with the NativeTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsResponse) GetNativeTokensOk() ([]NativeToken, bool) {
	if o == nil || isNil(o.NativeTokens) {
    return nil, false
	}
	return o.NativeTokens, true
}

// HasNativeTokens returns a boolean if a field has been set.
func (o *AssetsResponse) HasNativeTokens() bool {
	if o != nil && !isNil(o.NativeTokens) {
		return true
	}

	return false
}

// SetNativeTokens gets a reference to the given []NativeToken and assigns it to the NativeTokens field.
func (o *AssetsResponse) SetNativeTokens(v []NativeToken) {
	o.NativeTokens = v
}

func (o AssetsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.BaseTokens) {
		toSerialize["baseTokens"] = o.BaseTokens
	}
	if !isNil(o.NativeTokens) {
		toSerialize["nativeTokens"] = o.NativeTokens
	}
	return json.Marshal(toSerialize)
}

type NullableAssetsResponse struct {
	value *AssetsResponse
	isSet bool
}

func (v NullableAssetsResponse) Get() *AssetsResponse {
	return v.value
}

func (v *NullableAssetsResponse) Set(val *AssetsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetsResponse(val *AssetsResponse) *NullableAssetsResponse {
	return &NullableAssetsResponse{value: val, isSet: true}
}

func (v NullableAssetsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


