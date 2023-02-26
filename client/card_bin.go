package client

import (
	"fmt"
	"github.com/bpalermo/cielo-go/types"
)

const (
	cardBinBasePath = "/1/cardBin"
)

func (c *CieloClient) CardBin(bin string) (*types.CardBinResponse, error) {
	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s/%s", c.environment.ApiUrl, cardBinBasePath, bin), nil)
	if err != nil {
		return nil, err
	}

	response := &types.CardBinResponse{}

	err = c.send(req, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
