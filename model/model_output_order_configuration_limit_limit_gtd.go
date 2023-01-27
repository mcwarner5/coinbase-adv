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

// OutputOrderConfigurationLimitLimitGtd struct for OutputOrderConfigurationLimitLimitGtd
type OutputOrderConfigurationLimitLimitGtd struct {
	BaseSize *float64 `json:"base_size,omitempty,string"`
	LimitPrice *float64 `json:"limit_price,omitempty,string"`
	EndTime *string `json:"end_time,omitempty"`
	PostOnly *bool `json:"post_only,omitempty"`
}

// NewOutputOrderConfigurationLimitLimitGtd instantiates a new OutputOrderConfigurationLimitLimitGtd object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutputOrderConfigurationLimitLimitGtd() *OutputOrderConfigurationLimitLimitGtd {
	this := OutputOrderConfigurationLimitLimitGtd{}
	return &this
}

// NewOutputOrderConfigurationLimitLimitGtdWithDefaults instantiates a new OutputOrderConfigurationLimitLimitGtd object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutputOrderConfigurationLimitLimitGtdWithDefaults() *OutputOrderConfigurationLimitLimitGtd {
	this := OutputOrderConfigurationLimitLimitGtd{}
	return &this
}

// GetBaseSize returns the BaseSize field value if set, zero value otherwise.
func (o *OutputOrderConfigurationLimitLimitGtd) GetBaseSize() float64 {
	if o == nil || isNil(o.BaseSize) {
		var ret float64
		return ret
	}
	return *o.BaseSize
}

// GetBaseSizeOk returns a tuple with the BaseSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputOrderConfigurationLimitLimitGtd) GetBaseSizeOk() (*float64, bool) {
	if o == nil || isNil(o.BaseSize) {
    return nil, false
	}
	return o.BaseSize, true
}

// HasBaseSize returns a boolean if a field has been set.
func (o *OutputOrderConfigurationLimitLimitGtd) HasBaseSize() bool {
	if o != nil && !isNil(o.BaseSize) {
		return true
	}

	return false
}

// SetBaseSize gets a reference to the given float64 and assigns it to the BaseSize field.
func (o *OutputOrderConfigurationLimitLimitGtd) SetBaseSize(v float64) {
	o.BaseSize = &v
}

// GetLimitPrice returns the LimitPrice field value if set, zero value otherwise.
func (o *OutputOrderConfigurationLimitLimitGtd) GetLimitPrice() float64 {
	if o == nil || isNil(o.LimitPrice) {
		var ret float64
		return ret
	}
	return *o.LimitPrice
}

// GetLimitPriceOk returns a tuple with the LimitPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputOrderConfigurationLimitLimitGtd) GetLimitPriceOk() (*float64, bool) {
	if o == nil || isNil(o.LimitPrice) {
    return nil, false
	}
	return o.LimitPrice, true
}

// HasLimitPrice returns a boolean if a field has been set.
func (o *OutputOrderConfigurationLimitLimitGtd) HasLimitPrice() bool {
	if o != nil && !isNil(o.LimitPrice) {
		return true
	}

	return false
}

// SetLimitPrice gets a reference to the given float64 and assigns it to the LimitPrice field.
func (o *OutputOrderConfigurationLimitLimitGtd) SetLimitPrice(v float64) {
	o.LimitPrice = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *OutputOrderConfigurationLimitLimitGtd) GetEndTime() string {
	if o == nil || isNil(o.EndTime) {
		var ret string
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputOrderConfigurationLimitLimitGtd) GetEndTimeOk() (*string, bool) {
	if o == nil || isNil(o.EndTime) {
    return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *OutputOrderConfigurationLimitLimitGtd) HasEndTime() bool {
	if o != nil && !isNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given string and assigns it to the EndTime field.
func (o *OutputOrderConfigurationLimitLimitGtd) SetEndTime(v string) {
	o.EndTime = &v
}

// GetPostOnly returns the PostOnly field value if set, zero value otherwise.
func (o *OutputOrderConfigurationLimitLimitGtd) GetPostOnly() bool {
	if o == nil || isNil(o.PostOnly) {
		var ret bool
		return ret
	}
	return *o.PostOnly
}

// GetPostOnlyOk returns a tuple with the PostOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputOrderConfigurationLimitLimitGtd) GetPostOnlyOk() (*bool, bool) {
	if o == nil || isNil(o.PostOnly) {
    return nil, false
	}
	return o.PostOnly, true
}

// HasPostOnly returns a boolean if a field has been set.
func (o *OutputOrderConfigurationLimitLimitGtd) HasPostOnly() bool {
	if o != nil && !isNil(o.PostOnly) {
		return true
	}

	return false
}

// SetPostOnly gets a reference to the given bool and assigns it to the PostOnly field.
func (o *OutputOrderConfigurationLimitLimitGtd) SetPostOnly(v bool) {
	o.PostOnly = &v
}

func (o OutputOrderConfigurationLimitLimitGtd) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.BaseSize) {
		toSerialize["base_size"] = o.BaseSize
	}
	if !isNil(o.LimitPrice) {
		toSerialize["limit_price"] = o.LimitPrice
	}
	if !isNil(o.EndTime) {
		toSerialize["end_time"] = o.EndTime
	}
	if !isNil(o.PostOnly) {
		toSerialize["post_only"] = o.PostOnly
	}
	return json.Marshal(toSerialize)
}

type NullableOutputOrderConfigurationLimitLimitGtd struct {
	value *OutputOrderConfigurationLimitLimitGtd
	isSet bool
}

func (v NullableOutputOrderConfigurationLimitLimitGtd) Get() *OutputOrderConfigurationLimitLimitGtd {
	return v.value
}

func (v *NullableOutputOrderConfigurationLimitLimitGtd) Set(val *OutputOrderConfigurationLimitLimitGtd) {
	v.value = val
	v.isSet = true
}

func (v NullableOutputOrderConfigurationLimitLimitGtd) IsSet() bool {
	return v.isSet
}

func (v *NullableOutputOrderConfigurationLimitLimitGtd) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutputOrderConfigurationLimitLimitGtd(val *OutputOrderConfigurationLimitLimitGtd) *NullableOutputOrderConfigurationLimitLimitGtd {
	return &NullableOutputOrderConfigurationLimitLimitGtd{value: val, isSet: true}
}

func (v NullableOutputOrderConfigurationLimitLimitGtd) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutputOrderConfigurationLimitLimitGtd) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


