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

type Trade struct{
	TradeId *string `json:"trade_id,omitempty"`
	ProductId *string `json:"product_id,omitempty"`
	Price *float64 `json:"price,omitempty,string"`
	Size *float64 `json:"size,omitempty,string"`
	Time *string `json:"time,omitempty"`
	Side *string `json:"side,omitempty"`
	Ask *string `json:"ask,omitempty"`
	Bid *string `json:"bid,omitempty"`
}

// ListOrdersResponse struct for ListOrdersResponse
type ListProductsTickerResponse struct {
	Trades  []Trade `json:"trades,omitempty"`
	BestAsk *string `json:"best_ask,omitempty"`
	BestBid *string `json:"best_bid,omitempty"`
	Sequence *int32 `json:"sequence,omitempty,string"`
	HasNext *bool `json:"has_next,omitempty"`
	Cursor *string `json:"cursor,omitempty"`
}

// NewListOrdersResponse instantiates a new ListOrdersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListProductsTickerResponse() *ListProductsTickerResponse {
	this := ListProductsTickerResponse{}
	return &this
}

// NewListOrdersResponseWithDefaults instantiates a new ListOrdersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListProductsTickerResponseWithDefaults() *ListProductsTickerResponse {
	this := ListProductsTickerResponse{}
	return &this
}

// GetOrders returns the Orders field value if set, zero value otherwise.
func (o *ListProductsTickerResponse) GetTrades() []Trade {
	if o == nil || isNil(o.Trades) {
		var ret []Trade
		return ret
	}
	return o.Trades
}

// GetOrdersOk returns a tuple with the Orders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListProductsTickerResponse) GetTradesOk() ([]Trade, bool) {
	if o == nil || isNil(o.Trades) {
    return nil, false
	}
	return o.Trades, true
}

// HasOrders returns a boolean if a field has been set.
func (o *ListProductsTickerResponse) HasTrades() bool {
	if o != nil && !isNil(o.Trades) {
		return true
	}

	return false
}

// SetOrders gets a reference to the given []Order and assigns it to the Orders field.
func (o *ListProductsTickerResponse) SetTrades(v []Trade) {
	o.Trades = v
}

// GetSequence returns the Sequence field value if set, zero value otherwise.
func (o *ListProductsTickerResponse) GetSequence() int32 {
	if o == nil || isNil(o.Sequence) {
		var ret int32
		return ret
	}
	return *o.Sequence
}

// GetSequenceOk returns a tuple with the Sequence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListProductsTickerResponse) GetSequenceOk() (*int32, bool) {
	if o == nil || isNil(o.Sequence) {
    return nil, false
	}
	return o.Sequence, true
}

// HasSequence returns a boolean if a field has been set.
func (o *ListProductsTickerResponse) HasSequence() bool {
	if o != nil && !isNil(o.Sequence) {
		return true
	}

	return false
}

// SetSequence gets a reference to the given int32 and assigns it to the Sequence field.
func (o *ListProductsTickerResponse) SetSequence(v int32) {
	o.Sequence = &v
}

// GetHasNext returns the HasNext field value if set, zero value otherwise.
func (o *ListProductsTickerResponse) GetHasNext() bool {
	if o == nil || isNil(o.HasNext) {
		var ret bool
		return ret
	}
	return *o.HasNext
}

// GetHasNextOk returns a tuple with the HasNext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListProductsTickerResponse) GetHasNextOk() (*bool, bool) {
	if o == nil || isNil(o.HasNext) {
    return nil, false
	}
	return o.HasNext, true
}

// HasHasNext returns a boolean if a field has been set.
func (o *ListProductsTickerResponse) HasHasNext() bool {
	if o != nil && !isNil(o.HasNext) {
		return true
	}

	return false
}

// SetHasNext gets a reference to the given bool and assigns it to the HasNext field.
func (o *ListProductsTickerResponse) SetHasNext(v bool) {
	o.HasNext = &v
}

// GetCursor returns the Cursor field value if set, zero value otherwise.
func (o *ListProductsTickerResponse) GetCursor() string {
	if o == nil || isNil(o.Cursor) {
		var ret string
		return ret
	}
	return *o.Cursor
}

// GetCursorOk returns a tuple with the Cursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListProductsTickerResponse) GetCursorOk() (*string, bool) {
	if o == nil || isNil(o.Cursor) {
    return nil, false
	}
	return o.Cursor, true
}

// HasCursor returns a boolean if a field has been set.
func (o *ListProductsTickerResponse) HasCursor() bool {
	if o != nil && !isNil(o.Cursor) {
		return true
	}

	return false
}

// SetCursor gets a reference to the given string and assigns it to the Cursor field.
func (o *ListProductsTickerResponse) SetCursor(v string) {
	o.Cursor = &v
}

func (o ListProductsTickerResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Trades) {
		toSerialize["candles"] = o.Trades
	}
	if !isNil(o.Sequence) {
		toSerialize["sequence"] = o.Sequence
	}
	if !isNil(o.HasNext) {
		toSerialize["has_next"] = o.HasNext
	}
	if !isNil(o.Cursor) {
		toSerialize["cursor"] = o.Cursor
	}
	return json.Marshal(toSerialize)
}

type NullableListProductsTickerResponse struct {
	value *ListProductsTickerResponse
	isSet bool
}

func (v NullableListProductsTickerResponse) Get() *ListProductsTickerResponse {
	return v.value
}

func (v *NullableListProductsTickerResponse) Set(val *ListProductsTickerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListProductsTickerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListProductsTickerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListProductsTickerResponse(val *ListProductsTickerResponse) *NullableListProductsTickerResponse {
	return &NullableListProductsTickerResponse{value: val, isSet: true}
}

func (v NullableListProductsTickerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListProductsTickerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


