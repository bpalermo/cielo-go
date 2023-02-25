package client

import (
	"bytes"
	"encoding/json"
	"github.com/bpalermo/cielo-go/log"
	"github.com/bpalermo/cielo-go/types"
	"github.com/google/uuid"
	"io"
	"net/http"
)

type CieloClient struct {
	merchantId  string
	merchantKey string
	environment *types.Environment
	logger      log.Logger
	HTTPClient  *http.Client
}

type CieloClientOption func(*CieloClient)

func NewClient(merchantId string, merchantKey string, logger log.Logger, opts ...CieloClientOption) *CieloClient {
	c := &CieloClient{
		merchantId:  merchantId,
		merchantKey: merchantKey,
		environment: types.ProductionEnvironment,
		logger:      logger,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func WithEnvironment(environment *types.Environment) CieloClientOption {
	return func(c *CieloClient) {
		c.environment = environment
	}
}

// NewRequest constructs a request
// Convert payload to a JSON
func (c *CieloClient) NewRequest(method, url string, payload interface{}) (*http.Request, error) {
	var buf io.Reader
	if payload != nil {
		b, err := json.Marshal(&payload)
		if err != nil {
			return nil, err
		}
		buf = bytes.NewBuffer(b)
	}

	return http.NewRequest(method, url, buf)
}

func (c *CieloClient) send(req *http.Request, v interface{}) error {
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("User-Agent", "CieloEcommerce/3.0 cielo-go")
	req.Header.Add("MerchantId", c.merchantId)
	req.Header.Add("MerchantKey", c.merchantKey)
	req.Header.Add("RequestId", uuid.New().String())

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer c.handleError(res.Body)

	if res.StatusCode < 200 || res.StatusCode > 299 {
		errResp := &types.ErrorResponse{
			Response: res,
		}
		data, err := io.ReadAll(res.Body)

		if err == nil && len(data) > 0 {
			err = json.Unmarshal(data, errResp)
			if err != nil {
				return err
			}
		}

		return errResp
	}

	if err = json.NewDecoder(res.Body).Decode(&v); err != nil {
		return err
	}

	return nil
}

func (c *CieloClient) handleError(body io.ReadCloser) {
	err := body.Close()
	if err != nil {
		c.logger.Error(err)
	}
}
