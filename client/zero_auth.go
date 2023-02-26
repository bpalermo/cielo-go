package client

import (
	"fmt"
	"github.com/bpalermo/cielo-go/types"
)

const (
	zeroAuthBasePath = "/1/zeroauth"
)

func (c *CieloClient) ZeroAuth(request *types.ZeroAuthRequest) (*types.ZeroAuthResponse, error) {
	req, err := c.NewRequest("POST", fmt.Sprintf("%s%s", c.environment.ApiUrl, zeroAuthBasePath), request)
	if err != nil {
		return nil, err
	}

	response := &types.ZeroAuthResponse{}

	err = c.send(req, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
