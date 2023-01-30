package client

import (
	"coinbase-adv/model"
	"fmt"
	"net/url"
	"strings"
)

type Quote struct {
	Buy  float64
	Sell float64
}

// GetPrice -- get price
func (c *Client) GetPrice(currency string, side string) (*float64, error) {
	var (
		u, _        = url.Parse(CoinbaseAdv_V2Endpoint + fmt.Sprintf("/prices/%s/%s", currency, strings.ToLower(side)))
		response    model.GetPriceResponse
		headersMap  = make(map[string]string)
		queryParams = make(map[string]string)
	)

	err := c.GetAndDecode(*u, &response, &headersMap, &queryParams)
	if err != nil {
		return nil, err
	}
	return response.Data.Amount, err
}

// GetQuote -- get both sides of price
func (c *Client) GetQuote(currency string) (*Quote, error) {

	qt := Quote{}

	buyRv, err := c.GetPrice(currency, string(model.BUY))
	if err != nil {
		return nil, err
	}
	sellRv, err := c.GetPrice(currency, string(model.SELL))
	if err != nil {
		return nil, err
	}
	qt.Buy = *buyRv
	qt.Sell = *sellRv
	return &qt, nil
}
