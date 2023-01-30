/*
Coinbase Advanced Trading API

OpenAPI 3.x specification for Coinbase Adavanced Trading

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package model

import (
	"encoding/json"
)

// CancelOrderResponseResultsInner struct for CancelOrderResponseResultsInner
type CancelOrderResponseResultsInner struct {
	Success *bool `json:"success,omitempty"`
	FailureReason *string `json:"failure_reason,omitempty"`
	OrderId *string `json:"order_id,omitempty"`
}

// NewCancelOrderResponseResultsInner instantiates a new CancelOrderResponseResultsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCancelOrderResponseResultsInner() *CancelOrderResponseResultsInner {
	this := CancelOrderResponseResultsInner{}
	return &this
}

// NewCancelOrderResponseResultsInnerWithDefaults instantiates a new CancelOrderResponseResultsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCancelOrderResponseResultsInnerWithDefaults() *CancelOrderResponseResultsInner {
	this := CancelOrderResponseResultsInner{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *CancelOrderResponseResultsInner) GetSuccess() bool {
	if o == nil || isNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelOrderResponseResultsInner) GetSuccessOk() (*bool, bool) {
	if o == nil || isNil(o.Success) {
    return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *CancelOrderResponseResultsInner) HasSuccess() bool {
	if o != nil && !isNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *CancelOrderResponseResultsInner) SetSuccess(v bool) {
	o.Success = &v
}

// GetFailureReason returns the FailureReason field value if set, zero value otherwise.
func (o *CancelOrderResponseResultsInner) GetFailureReason() string {
	if o == nil || isNil(o.FailureReason) {
		var ret string
		return ret
	}
	return *o.FailureReason
}

// GetFailureReasonOk returns a tuple with the FailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelOrderResponseResultsInner) GetFailureReasonOk() (*string, bool) {
	if o == nil || isNil(o.FailureReason) {
    return nil, false
	}
	return o.FailureReason, true
}

// HasFailureReason returns a boolean if a field has been set.
func (o *CancelOrderResponseResultsInner) HasFailureReason() bool {
	if o != nil && !isNil(o.FailureReason) {
		return true
	}

	return false
}

// SetFailureReason gets a reference to the given string and assigns it to the FailureReason field.
func (o *CancelOrderResponseResultsInner) SetFailureReason(v string) {
	o.FailureReason = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *CancelOrderResponseResultsInner) GetOrderId() string {
	if o == nil || isNil(o.OrderId) {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelOrderResponseResultsInner) GetOrderIdOk() (*string, bool) {
	if o == nil || isNil(o.OrderId) {
    return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *CancelOrderResponseResultsInner) HasOrderId() bool {
	if o != nil && !isNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *CancelOrderResponseResultsInner) SetOrderId(v string) {
	o.OrderId = &v
}

func (o CancelOrderResponseResultsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !isNil(o.FailureReason) {
		toSerialize["failure_reason"] = o.FailureReason
	}
	if !isNil(o.OrderId) {
		toSerialize["order_id"] = o.OrderId
	}
	return json.Marshal(toSerialize)
}

type NullableCancelOrderResponseResultsInner struct {
	value *CancelOrderResponseResultsInner
	isSet bool
}

func (v NullableCancelOrderResponseResultsInner) Get() *CancelOrderResponseResultsInner {
	return v.value
}

func (v *NullableCancelOrderResponseResultsInner) Set(val *CancelOrderResponseResultsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCancelOrderResponseResultsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCancelOrderResponseResultsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCancelOrderResponseResultsInner(val *CancelOrderResponseResultsInner) *NullableCancelOrderResponseResultsInner {
	return &NullableCancelOrderResponseResultsInner{value: val, isSet: true}
}

func (v NullableCancelOrderResponseResultsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCancelOrderResponseResultsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


