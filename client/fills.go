package client

import (
	"context"
	"net/url"
	"time"

	"github.com/mcwarner5/BlockBot8000/libraries/coinbase-adv/model"
)

type ListFillsParams struct {
	OrderId                string
	ProductId              string
	Limit                  int32
	StartSequenceTimestamp time.Time
	EndSequenceTimestamp   time.Time

	Cursor *string
}

// ListFills -- list fills
func (c *Client) ListFills(ctx context.Context, p *ListFillsParams) (*model.ListFillsResponse, error) {

	var (
		u, _        = url.Parse(CoinbaseAdvV3endpoint + "/brokerage/orders/historical/fills")
		response    model.ListFillsResponse
		headersMap  = make(map[string]string)
		queryParams = make(map[string]string)
	)
	if p != nil {
		if p.Limit > 0 {
			c.addInt32Param(queryParams, "limit", p.Limit)
		}
		c.addStringParam(queryParams, "order_id", p.OrderId)
		c.addStringParam(queryParams, "product_id", p.ProductId)

		if !p.StartSequenceTimestamp.IsZero() {
			c.addStringParam(queryParams, "start_sequence_timestamp", p.StartSequenceTimestamp.Format(time.RFC3339))
		}
		if !p.EndSequenceTimestamp.IsZero() {
			c.addStringParam(queryParams, "end_sequence_timestamp", p.EndSequenceTimestamp.Format(time.RFC3339))
		}
		if p.Cursor != nil {
			c.addStringParam(queryParams, "cursor", *p.Cursor)
		}
	}

	err := c.GetAndDecode(ctx, *u, &response, &headersMap, &queryParams)
	if err != nil {
		return nil, err
	}
	return &response, err

}
