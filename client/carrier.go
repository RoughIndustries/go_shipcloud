package client

import (
	"github.com/RoughIndustries/go_shipcloud/models"
	"net/http"
)

func (c *Client) GetAllCarriers() (*models.Carriers, error) {
	carriers := &models.Carriers{}
	err := c.do(http.MethodGet, "/v1/carriers", nil, &carriers.Carriers)
	return carriers, err
}
