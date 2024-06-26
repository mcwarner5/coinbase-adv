package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/mcwarner5/BlockBot8000/libraries/coinbase-adv/model"
)

const (
	DefaultLimit = 50
	MaxLimit     = 300
	BatchLimit   = 1000
)

// GetOrder -- get order
func (c *Client) GetOrder(ctx context.Context, id string) (*model.GetOrderResponse, error) {
	var (
		u, _        = url.Parse(CoinbaseAdvV3endpoint + fmt.Sprintf("/brokerage/orders/historical/%s", id))
		response    model.GetOrderResponse
		headersMap  = make(map[string]string)
		queryParams = make(map[string]string)
	)

	err := c.GetAndDecode(ctx, *u, &response, &headersMap, &queryParams)
	if err != nil {
		return nil, err
	}
	return &response, err
}

// CancelOrders -- cancel orders
func (c *Client) CancelOrders(ctx context.Context, ids []string) (*model.CancelOrderResponse, error) {
	var (
		u, _        = url.Parse(CoinbaseAdvV3endpoint + "/brokerage/orders/batch_cancel")
		response    model.CancelOrderResponse
		headersMap  = make(map[string]string)
		queryParams = make(map[string]string)
	)

	p := model.CancelOrderRequest{OrderIds: ids}

	payload, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	err = c.PostAndDecode(ctx, *u, &response, &headersMap, &queryParams, payload)
	if err != nil {
		return nil, err
	}
	return &response, err

}

// CreateOrder -- create order
func (c *Client) CreateOrder(ctx context.Context, p *model.CreateOrderRequest) (*model.CreateOrderResponse, error) {

	var (
		u, _        = url.Parse(CoinbaseAdvV3endpoint + "/brokerage/orders")
		response    model.CreateOrderResponse
		headersMap  = make(map[string]string)
		queryParams = make(map[string]string)
	)

	payload, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	err = c.PostAndDecode(ctx, *u, &response, &headersMap, &queryParams, payload)
	if err != nil {
		return nil, err
	}
	return &response, err
}

type ListOrdersParams struct {
	ProductId          string
	OrderStatus        []string
	Limit              int32
	StartDate          time.Time
	EndDate            time.Time
	UserNativeCurrency string
	OrderType          model.OrderType
	OrderSide          model.OrderSide
	ProductType        model.ProductType

	Cursor *string
}

// ListOrders -- list orders
func (c *Client) ListOrders(ctx context.Context, p *ListOrdersParams) (*model.ListOrdersResponse, error) {

	var (
		u, _        = url.Parse(CoinbaseAdvV3endpoint + "/brokerage/orders/historical/batch")
		response    model.ListOrdersResponse
		headersMap  = make(map[string]string)
		queryParams = make(map[string]string)
	)
	if p != nil {
		if p.Limit > 0 {
			c.addInt32Param(queryParams, "limit", p.Limit)
		} else {
			c.addInt32Param(queryParams, "limit", BatchLimit)
		}
		c.addStringParam(queryParams, "product_id", p.ProductId)
		if len(p.OrderStatus) > 0 {
			orderStatus := strings.Join(p.OrderStatus, ",")
			c.addStringParam(queryParams, "order_status", orderStatus)
		}

		if !p.StartDate.IsZero() {
			c.addStringParam(queryParams, "start_date", p.StartDate.Format(time.RFC3339))
		}
		if !p.EndDate.IsZero() {
			c.addStringParam(queryParams, "end_date", p.EndDate.Format(time.RFC3339))
		}
		c.addStringParam(queryParams, "user_native_currency", p.UserNativeCurrency)

		c.addStringParam(queryParams, "order_type", string(p.OrderType))
		c.addStringParam(queryParams, "order_side", string(p.OrderSide))
		c.addStringParam(queryParams, "product_type", string(p.ProductType))

		if p.Cursor != nil {
			c.addStringParam(queryParams, "cursor", *p.Cursor)
		}

	} else {
		c.addInt32Param(queryParams, "limit", BatchLimit)
	}

	err := c.GetAndDecode(ctx, *u, &response, &headersMap, &queryParams)
	if err != nil {
		return nil, err
	}
	return &response, err

}
