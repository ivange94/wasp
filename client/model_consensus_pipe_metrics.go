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

// ConsensusPipeMetrics struct for ConsensusPipeMetrics
type ConsensusPipeMetrics struct {
	EventACSMsgPipeSize *int32 `json:"eventACSMsgPipeSize,omitempty"`
	EventPeerLogIndexMsgPipeSize *int32 `json:"eventPeerLogIndexMsgPipeSize,omitempty"`
	EventStateTransitionMsgPipeSize *int32 `json:"eventStateTransitionMsgPipeSize,omitempty"`
	EventTimerMsgPipeSize *int32 `json:"eventTimerMsgPipeSize,omitempty"`
	EventVMResultMsgPipeSize *int32 `json:"eventVMResultMsgPipeSize,omitempty"`
}

// NewConsensusPipeMetrics instantiates a new ConsensusPipeMetrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsensusPipeMetrics() *ConsensusPipeMetrics {
	this := ConsensusPipeMetrics{}
	return &this
}

// NewConsensusPipeMetricsWithDefaults instantiates a new ConsensusPipeMetrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsensusPipeMetricsWithDefaults() *ConsensusPipeMetrics {
	this := ConsensusPipeMetrics{}
	return &this
}

// GetEventACSMsgPipeSize returns the EventACSMsgPipeSize field value if set, zero value otherwise.
func (o *ConsensusPipeMetrics) GetEventACSMsgPipeSize() int32 {
	if o == nil || isNil(o.EventACSMsgPipeSize) {
		var ret int32
		return ret
	}
	return *o.EventACSMsgPipeSize
}

// GetEventACSMsgPipeSizeOk returns a tuple with the EventACSMsgPipeSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsensusPipeMetrics) GetEventACSMsgPipeSizeOk() (*int32, bool) {
	if o == nil || isNil(o.EventACSMsgPipeSize) {
    return nil, false
	}
	return o.EventACSMsgPipeSize, true
}

// HasEventACSMsgPipeSize returns a boolean if a field has been set.
func (o *ConsensusPipeMetrics) HasEventACSMsgPipeSize() bool {
	if o != nil && !isNil(o.EventACSMsgPipeSize) {
		return true
	}

	return false
}

// SetEventACSMsgPipeSize gets a reference to the given int32 and assigns it to the EventACSMsgPipeSize field.
func (o *ConsensusPipeMetrics) SetEventACSMsgPipeSize(v int32) {
	o.EventACSMsgPipeSize = &v
}

// GetEventPeerLogIndexMsgPipeSize returns the EventPeerLogIndexMsgPipeSize field value if set, zero value otherwise.
func (o *ConsensusPipeMetrics) GetEventPeerLogIndexMsgPipeSize() int32 {
	if o == nil || isNil(o.EventPeerLogIndexMsgPipeSize) {
		var ret int32
		return ret
	}
	return *o.EventPeerLogIndexMsgPipeSize
}

// GetEventPeerLogIndexMsgPipeSizeOk returns a tuple with the EventPeerLogIndexMsgPipeSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsensusPipeMetrics) GetEventPeerLogIndexMsgPipeSizeOk() (*int32, bool) {
	if o == nil || isNil(o.EventPeerLogIndexMsgPipeSize) {
    return nil, false
	}
	return o.EventPeerLogIndexMsgPipeSize, true
}

// HasEventPeerLogIndexMsgPipeSize returns a boolean if a field has been set.
func (o *ConsensusPipeMetrics) HasEventPeerLogIndexMsgPipeSize() bool {
	if o != nil && !isNil(o.EventPeerLogIndexMsgPipeSize) {
		return true
	}

	return false
}

// SetEventPeerLogIndexMsgPipeSize gets a reference to the given int32 and assigns it to the EventPeerLogIndexMsgPipeSize field.
func (o *ConsensusPipeMetrics) SetEventPeerLogIndexMsgPipeSize(v int32) {
	o.EventPeerLogIndexMsgPipeSize = &v
}

// GetEventStateTransitionMsgPipeSize returns the EventStateTransitionMsgPipeSize field value if set, zero value otherwise.
func (o *ConsensusPipeMetrics) GetEventStateTransitionMsgPipeSize() int32 {
	if o == nil || isNil(o.EventStateTransitionMsgPipeSize) {
		var ret int32
		return ret
	}
	return *o.EventStateTransitionMsgPipeSize
}

// GetEventStateTransitionMsgPipeSizeOk returns a tuple with the EventStateTransitionMsgPipeSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsensusPipeMetrics) GetEventStateTransitionMsgPipeSizeOk() (*int32, bool) {
	if o == nil || isNil(o.EventStateTransitionMsgPipeSize) {
    return nil, false
	}
	return o.EventStateTransitionMsgPipeSize, true
}

// HasEventStateTransitionMsgPipeSize returns a boolean if a field has been set.
func (o *ConsensusPipeMetrics) HasEventStateTransitionMsgPipeSize() bool {
	if o != nil && !isNil(o.EventStateTransitionMsgPipeSize) {
		return true
	}

	return false
}

// SetEventStateTransitionMsgPipeSize gets a reference to the given int32 and assigns it to the EventStateTransitionMsgPipeSize field.
func (o *ConsensusPipeMetrics) SetEventStateTransitionMsgPipeSize(v int32) {
	o.EventStateTransitionMsgPipeSize = &v
}

// GetEventTimerMsgPipeSize returns the EventTimerMsgPipeSize field value if set, zero value otherwise.
func (o *ConsensusPipeMetrics) GetEventTimerMsgPipeSize() int32 {
	if o == nil || isNil(o.EventTimerMsgPipeSize) {
		var ret int32
		return ret
	}
	return *o.EventTimerMsgPipeSize
}

// GetEventTimerMsgPipeSizeOk returns a tuple with the EventTimerMsgPipeSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsensusPipeMetrics) GetEventTimerMsgPipeSizeOk() (*int32, bool) {
	if o == nil || isNil(o.EventTimerMsgPipeSize) {
    return nil, false
	}
	return o.EventTimerMsgPipeSize, true
}

// HasEventTimerMsgPipeSize returns a boolean if a field has been set.
func (o *ConsensusPipeMetrics) HasEventTimerMsgPipeSize() bool {
	if o != nil && !isNil(o.EventTimerMsgPipeSize) {
		return true
	}

	return false
}

// SetEventTimerMsgPipeSize gets a reference to the given int32 and assigns it to the EventTimerMsgPipeSize field.
func (o *ConsensusPipeMetrics) SetEventTimerMsgPipeSize(v int32) {
	o.EventTimerMsgPipeSize = &v
}

// GetEventVMResultMsgPipeSize returns the EventVMResultMsgPipeSize field value if set, zero value otherwise.
func (o *ConsensusPipeMetrics) GetEventVMResultMsgPipeSize() int32 {
	if o == nil || isNil(o.EventVMResultMsgPipeSize) {
		var ret int32
		return ret
	}
	return *o.EventVMResultMsgPipeSize
}

// GetEventVMResultMsgPipeSizeOk returns a tuple with the EventVMResultMsgPipeSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsensusPipeMetrics) GetEventVMResultMsgPipeSizeOk() (*int32, bool) {
	if o == nil || isNil(o.EventVMResultMsgPipeSize) {
    return nil, false
	}
	return o.EventVMResultMsgPipeSize, true
}

// HasEventVMResultMsgPipeSize returns a boolean if a field has been set.
func (o *ConsensusPipeMetrics) HasEventVMResultMsgPipeSize() bool {
	if o != nil && !isNil(o.EventVMResultMsgPipeSize) {
		return true
	}

	return false
}

// SetEventVMResultMsgPipeSize gets a reference to the given int32 and assigns it to the EventVMResultMsgPipeSize field.
func (o *ConsensusPipeMetrics) SetEventVMResultMsgPipeSize(v int32) {
	o.EventVMResultMsgPipeSize = &v
}

func (o ConsensusPipeMetrics) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.EventACSMsgPipeSize) {
		toSerialize["eventACSMsgPipeSize"] = o.EventACSMsgPipeSize
	}
	if !isNil(o.EventPeerLogIndexMsgPipeSize) {
		toSerialize["eventPeerLogIndexMsgPipeSize"] = o.EventPeerLogIndexMsgPipeSize
	}
	if !isNil(o.EventStateTransitionMsgPipeSize) {
		toSerialize["eventStateTransitionMsgPipeSize"] = o.EventStateTransitionMsgPipeSize
	}
	if !isNil(o.EventTimerMsgPipeSize) {
		toSerialize["eventTimerMsgPipeSize"] = o.EventTimerMsgPipeSize
	}
	if !isNil(o.EventVMResultMsgPipeSize) {
		toSerialize["eventVMResultMsgPipeSize"] = o.EventVMResultMsgPipeSize
	}
	return json.Marshal(toSerialize)
}

type NullableConsensusPipeMetrics struct {
	value *ConsensusPipeMetrics
	isSet bool
}

func (v NullableConsensusPipeMetrics) Get() *ConsensusPipeMetrics {
	return v.value
}

func (v *NullableConsensusPipeMetrics) Set(val *ConsensusPipeMetrics) {
	v.value = val
	v.isSet = true
}

func (v NullableConsensusPipeMetrics) IsSet() bool {
	return v.isSet
}

func (v *NullableConsensusPipeMetrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsensusPipeMetrics(val *ConsensusPipeMetrics) *NullableConsensusPipeMetrics {
	return &NullableConsensusPipeMetrics{value: val, isSet: true}
}

func (v NullableConsensusPipeMetrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsensusPipeMetrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


