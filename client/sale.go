package client

import (
	"fmt"
	"github.com/bpalermo/cielo-go/types"
)

const (
	saleBasePath = "/1/sales"
)

func (c *CieloClient) Authorize(sale *types.Sale) (*types.Sale, error) {
	req, err := c.NewRequest("POST", fmt.Sprintf("%s%s", c.environment.ApiUrl, saleBasePath), sale)
	if err != nil {
		return nil, err
	}

	response := &types.Sale{}

	err = c.send(req, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
