package client

import (
	"github.com/RoughIndustries/go_shipcloud/models"
	"net/http"
)

func (c *Client) CreateAShipmentQuote(input models.ShipmentQuoteInput) (*models.ShipmentQuoteResult, error) {
	quote := &models.ShipmentQuoteResult{}
	err := c.do(http.MethodPost, "/v1/shipment_quotes", input, &quote)
	return quote, err
}
