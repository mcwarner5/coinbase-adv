package client

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/mcwarner5/coinbase-adv/model"
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
		u, _        = url.Parse(CoinbaseAdvV3endpoint + fmt.Sprintf("/brokerage/products"))
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

	} else {
		c.addInt32Param(queryParams, "limit", DefaultLimit)
	}

	err := c.GetAndDecode(ctx, *u, &response, &headersMap, &queryParams)
	if err != nil {
		return nil, err
	}
	return &response, err
}
