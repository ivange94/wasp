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

// ContractCallViewRequest struct for ContractCallViewRequest
type ContractCallViewRequest struct {
	Arguments *JSONDict `json:"arguments,omitempty"`
	// The chain id
	ChainId *string `json:"chainId,omitempty"`
	// The contract name as HName (Hex)
	ContractHName *string `json:"contractHName,omitempty"`
	// The contract name
	ContractName *string `json:"contractName,omitempty"`
	// The function name as HName (Hex)
	FunctionHName *string `json:"functionHName,omitempty"`
	// The function name
	FunctionName *string `json:"functionName,omitempty"`
}

// NewContractCallViewRequest instantiates a new ContractCallViewRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContractCallViewRequest() *ContractCallViewRequest {
	this := ContractCallViewRequest{}
	return &this
}

// NewContractCallViewRequestWithDefaults instantiates a new ContractCallViewRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractCallViewRequestWithDefaults() *ContractCallViewRequest {
	this := ContractCallViewRequest{}
	return &this
}

// GetArguments returns the Arguments field value if set, zero value otherwise.
func (o *ContractCallViewRequest) GetArguments() JSONDict {
	if o == nil || isNil(o.Arguments) {
		var ret JSONDict
		return ret
	}
	return *o.Arguments
}

// GetArgumentsOk returns a tuple with the Arguments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractCallViewRequest) GetArgumentsOk() (*JSONDict, bool) {
	if o == nil || isNil(o.Arguments) {
    return nil, false
	}
	return o.Arguments, true
}

// HasArguments returns a boolean if a field has been set.
func (o *ContractCallViewRequest) HasArguments() bool {
	if o != nil && !isNil(o.Arguments) {
		return true
	}

	return false
}

// SetArguments gets a reference to the given JSONDict and assigns it to the Arguments field.
func (o *ContractCallViewRequest) SetArguments(v JSONDict) {
	o.Arguments = &v
}

// GetChainId returns the ChainId field value if set, zero value otherwise.
func (o *ContractCallViewRequest) GetChainId() string {
	if o == nil || isNil(o.ChainId) {
		var ret string
		return ret
	}
	return *o.ChainId
}

// GetChainIdOk returns a tuple with the ChainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractCallViewRequest) GetChainIdOk() (*string, bool) {
	if o == nil || isNil(o.ChainId) {
    return nil, false
	}
	return o.ChainId, true
}

// HasChainId returns a boolean if a field has been set.
func (o *ContractCallViewRequest) HasChainId() bool {
	if o != nil && !isNil(o.ChainId) {
		return true
	}

	return false
}

// SetChainId gets a reference to the given string and assigns it to the ChainId field.
func (o *ContractCallViewRequest) SetChainId(v string) {
	o.ChainId = &v
}

// GetContractHName returns the ContractHName field value if set, zero value otherwise.
func (o *ContractCallViewRequest) GetContractHName() string {
	if o == nil || isNil(o.ContractHName) {
		var ret string
		return ret
	}
	return *o.ContractHName
}

// GetContractHNameOk returns a tuple with the ContractHName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractCallViewRequest) GetContractHNameOk() (*string, bool) {
	if o == nil || isNil(o.ContractHName) {
    return nil, false
	}
	return o.ContractHName, true
}

// HasContractHName returns a boolean if a field has been set.
func (o *ContractCallViewRequest) HasContractHName() bool {
	if o != nil && !isNil(o.ContractHName) {
		return true
	}

	return false
}

// SetContractHName gets a reference to the given string and assigns it to the ContractHName field.
func (o *ContractCallViewRequest) SetContractHName(v string) {
	o.ContractHName = &v
}

// GetContractName returns the ContractName field value if set, zero value otherwise.
func (o *ContractCallViewRequest) GetContractName() string {
	if o == nil || isNil(o.ContractName) {
		var ret string
		return ret
	}
	return *o.ContractName
}

// GetContractNameOk returns a tuple with the ContractName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractCallViewRequest) GetContractNameOk() (*string, bool) {
	if o == nil || isNil(o.ContractName) {
    return nil, false
	}
	return o.ContractName, true
}

// HasContractName returns a boolean if a field has been set.
func (o *ContractCallViewRequest) HasContractName() bool {
	if o != nil && !isNil(o.ContractName) {
		return true
	}

	return false
}

// SetContractName gets a reference to the given string and assigns it to the ContractName field.
func (o *ContractCallViewRequest) SetContractName(v string) {
	o.ContractName = &v
}

// GetFunctionHName returns the FunctionHName field value if set, zero value otherwise.
func (o *ContractCallViewRequest) GetFunctionHName() string {
	if o == nil || isNil(o.FunctionHName) {
		var ret string
		return ret
	}
	return *o.FunctionHName
}

// GetFunctionHNameOk returns a tuple with the FunctionHName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractCallViewRequest) GetFunctionHNameOk() (*string, bool) {
	if o == nil || isNil(o.FunctionHName) {
    return nil, false
	}
	return o.FunctionHName, true
}

// HasFunctionHName returns a boolean if a field has been set.
func (o *ContractCallViewRequest) HasFunctionHName() bool {
	if o != nil && !isNil(o.FunctionHName) {
		return true
	}

	return false
}

// SetFunctionHName gets a reference to the given string and assigns it to the FunctionHName field.
func (o *ContractCallViewRequest) SetFunctionHName(v string) {
	o.FunctionHName = &v
}

// GetFunctionName returns the FunctionName field value if set, zero value otherwise.
func (o *ContractCallViewRequest) GetFunctionName() string {
	if o == nil || isNil(o.FunctionName) {
		var ret string
		return ret
	}
	return *o.FunctionName
}

// GetFunctionNameOk returns a tuple with the FunctionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractCallViewRequest) GetFunctionNameOk() (*string, bool) {
	if o == nil || isNil(o.FunctionName) {
    return nil, false
	}
	return o.FunctionName, true
}

// HasFunctionName returns a boolean if a field has been set.
func (o *ContractCallViewRequest) HasFunctionName() bool {
	if o != nil && !isNil(o.FunctionName) {
		return true
	}

	return false
}

// SetFunctionName gets a reference to the given string and assigns it to the FunctionName field.
func (o *ContractCallViewRequest) SetFunctionName(v string) {
	o.FunctionName = &v
}

func (o ContractCallViewRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Arguments) {
		toSerialize["arguments"] = o.Arguments
	}
	if !isNil(o.ChainId) {
		toSerialize["chainId"] = o.ChainId
	}
	if !isNil(o.ContractHName) {
		toSerialize["contractHName"] = o.ContractHName
	}
	if !isNil(o.ContractName) {
		toSerialize["contractName"] = o.ContractName
	}
	if !isNil(o.FunctionHName) {
		toSerialize["functionHName"] = o.FunctionHName
	}
	if !isNil(o.FunctionName) {
		toSerialize["functionName"] = o.FunctionName
	}
	return json.Marshal(toSerialize)
}

type NullableContractCallViewRequest struct {
	value *ContractCallViewRequest
	isSet bool
}

func (v NullableContractCallViewRequest) Get() *ContractCallViewRequest {
	return v.value
}

func (v *NullableContractCallViewRequest) Set(val *ContractCallViewRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableContractCallViewRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableContractCallViewRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractCallViewRequest(val *ContractCallViewRequest) *NullableContractCallViewRequest {
	return &NullableContractCallViewRequest{value: val, isSet: true}
}

func (v NullableContractCallViewRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractCallViewRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


