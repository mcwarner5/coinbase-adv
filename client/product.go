package client

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/mcwarner5/BlockBot8000/libraries/coinbase-adv/model"
)

// GetProduct -- get product
func (c *Client) GetProduct(ctx context.Context, productId string) (*model.GetProductResponse, error) {
	var (
		u, _        = url.Parse(CoinbaseAdvV3endpoint + fmt.Sprintf("/brokerage/products/%s", productId))
		response    model.GetProductResponse
		headersMap  = make(map[string]string)
		queryParams = make(map[string]string)
	)

	err := c.GetAndDecode(ctx, *u, &response, &headersMap, &queryParams)
	if err != nil {
		return nil, err
	}
	return &response, err
}

type ListProductsTickerHistoryParams struct {
	Limit     int32
	ProductId string
	StartTime time.Time
	EndTime   time.Time

	Cursor *string
}

func (c *Client) ListProductsTickerHistory(ctx context.Context, p *ListProductsTickerHistoryParams) (*model.ListProductsTickerResponse, error) {
	if p == nil || len(p.ProductId) == 0 {
		return nil, errors.New("no product sent to /product/{product_id}/ticker request")
	}

	var (
		u, _        = url.Parse(CoinbaseAdvV3endpoint + fmt.Sprintf("/brokerage/products/%s/ticker", p.ProductId))
		response    model.ListProductsTickerResponse
		headersMap  = make(map[string]string)
		queryParams = make(map[string]string)
	)

	if p.Limit > 0 {
		c.addInt32Param(queryParams, "limit", p.Limit)
	} else {
		c.addInt32Param(queryParams, "limit", BatchLimit)
	}

	c.addStringParam(queryParams, "start", fmt.Sprint(p.StartTime.Unix()))
	c.addStringParam(queryParams, "end", fmt.Sprint(p.EndTime.Unix()))

	if p.Cursor != nil {
		c.addStringParam(queryParams, "cursor", *p.Cursor)
	}

	err := c.GetAndDecode(ctx, *u, &response, &headersMap, &queryParams)
	if err != nil {
		return nil, err
	}
	return &response, err
}

type ListProductsParams struct {
	Limit                  int32
	Offset                 int32
	ProductType            model.ProductType
	ProductIds             []string
	ContractExpiryType     string
	ExpiringContractStatus string

	Cursor *string
}

func (c *Client) ListProducts(ctx context.Context, p *ListProductsParams) (*model.ListProductsResponse, error) {
	var (
		u, _        = url.Parse(CoinbaseAdvV3endpoint + "/brokerage/products")
		response    model.ListProductsResponse
		headersMap  = make(map[string]string)
		queryParams = make(map[string]string)
	)

	if p != nil {
		if p.Limit > 0 {
			c.addInt32Param(queryParams, "limit", p.Limit)
		} else {
			c.addInt32Param(queryParams, "limit", DefaultLimit)
		}

		c.addStringParam(queryParams, "product_type", string(p.ProductType))

		if len(p.ProductIds) > 0 {
			ProductIds := strings.Join(p.ProductIds, ",")
			c.addStringParam(queryParams, "product_ids", ProductIds)
		}

		c.addStringParam(queryParams, "contract_expiry_type", string(p.ContractExpiryType))
		c.addStringParam(queryParams, "expiring_contract_status", string(p.ExpiringContractStatus))

		if p.Cursor != nil {
			c.addStringParam(queryParams, "cursor", *p.Cursor)
		}

		if p.Offset != 0 {
			c.addInt32Param(queryParams, "offset", p.Offset)
		}

	} else {
		c.addInt32Param(queryParams, "limit", DefaultLimit)
	}

	err := c.GetAndDecode(ctx, *u, &response, &headersMap, &queryParams)
	if err != nil {
		return nil, err
	}
	return &response, err
}

type ListProductsCandlesParams struct {
	Product   string
	StartTime string
	EndTime   string
	Interval  int

	Cursor *string
}

func (c *Client) GetProductCandles(ctx context.Context, p *ListProductsCandlesParams) (*model.ListProductsCandlestickResponse, error) {
	if p == nil || len(p.Product) == 0 {
		return nil, errors.New("no product sent to /product/{product_id}/candles request")
	}

	var (
		u, _        = url.Parse(CoinbaseAdvV3endpoint + fmt.Sprintf("/brokerage/products/%s/candles", p.Product))
		response    model.ListProductsCandlestickResponse
		headersMap  = make(map[string]string)
		queryParams = make(map[string]string)
	)

	c.addStringParam(queryParams, "start", p.StartTime)
	c.addStringParam(queryParams, "end", p.EndTime)
	var interval_str string
	if p.Interval > 0 {
		switch os := p.Interval; os {
		case 1:
			interval_str = "ONE_MINUTE"
		case 5:
			interval_str = "FIVE_MINUTE"
		case 15:
			interval_str = "FIFTEEN_MINUTE"
		case 30:
			interval_str = "THIRTY_MINUTE"
		case 60:
			interval_str = "ONE_HOUR"
		case 120:
			interval_str = "TWO_HOUR"
		case 360:
			interval_str = "SIX_HOUR"
		case 1440:
			interval_str = "ONE_DAY"

		default:
			interval_str = "UNKNOWN_GRANULARITY"
		}
		c.addStringParam(queryParams, "granularity", interval_str)
	}
	if p.Cursor != nil {
		c.addStringParam(queryParams, "cursor", *p.Cursor)
	}

	err := c.GetAndDecode(ctx, *u, &response, &headersMap, &queryParams)
	if err != nil {
		return nil, err
	}
	return &response, err
}
