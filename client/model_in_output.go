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

// InOutput struct for InOutput
type InOutput struct {
	Output *Output `json:"output,omitempty"`
	// The output ID
	OutputId *string `json:"outputId,omitempty"`
}

// NewInOutput instantiates a new InOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInOutput() *InOutput {
	this := InOutput{}
	return &this
}

// NewInOutputWithDefaults instantiates a new InOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInOutputWithDefaults() *InOutput {
	this := InOutput{}
	return &this
}

// GetOutput returns the Output field value if set, zero value otherwise.
func (o *InOutput) GetOutput() Output {
	if o == nil || isNil(o.Output) {
		var ret Output
		return ret
	}
	return *o.Output
}

// GetOutputOk returns a tuple with the Output field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InOutput) GetOutputOk() (*Output, bool) {
	if o == nil || isNil(o.Output) {
    return nil, false
	}
	return o.Output, true
}

// HasOutput returns a boolean if a field has been set.
func (o *InOutput) HasOutput() bool {
	if o != nil && !isNil(o.Output) {
		return true
	}

	return false
}

// SetOutput gets a reference to the given Output and assigns it to the Output field.
func (o *InOutput) SetOutput(v Output) {
	o.Output = &v
}

// GetOutputId returns the OutputId field value if set, zero value otherwise.
func (o *InOutput) GetOutputId() string {
	if o == nil || isNil(o.OutputId) {
		var ret string
		return ret
	}
	return *o.OutputId
}

// GetOutputIdOk returns a tuple with the OutputId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InOutput) GetOutputIdOk() (*string, bool) {
	if o == nil || isNil(o.OutputId) {
    return nil, false
	}
	return o.OutputId, true
}

// HasOutputId returns a boolean if a field has been set.
func (o *InOutput) HasOutputId() bool {
	if o != nil && !isNil(o.OutputId) {
		return true
	}

	return false
}

// SetOutputId gets a reference to the given string and assigns it to the OutputId field.
func (o *InOutput) SetOutputId(v string) {
	o.OutputId = &v
}

func (o InOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Output) {
		toSerialize["output"] = o.Output
	}
	if !isNil(o.OutputId) {
		toSerialize["outputId"] = o.OutputId
	}
	return json.Marshal(toSerialize)
}

type NullableInOutput struct {
	value *InOutput
	isSet bool
}

func (v NullableInOutput) Get() *InOutput {
	return v.value
}

func (v *NullableInOutput) Set(val *InOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableInOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableInOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInOutput(val *InOutput) *NullableInOutput {
	return &NullableInOutput{value: val, isSet: true}
}

func (v NullableInOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


