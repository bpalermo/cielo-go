package client

import (
	"github.com/bpalermo/cielo-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	expectedMerchantId  = "merchantId"
	expectedMerchantKey = "merchantKey"
)

func TestNewClient(t *testing.T) {
	logger := &TestLogger{}

	c := NewClient(expectedMerchantId, expectedMerchantKey, logger)
	assert.NotNil(t, c)
	assert.Equal(t, expectedMerchantId, c.merchantId)
	assert.Equal(t, expectedMerchantKey, c.merchantKey)
	assert.Equal(t, types.ProductionEnvironment, c.environment)
}

func TestWithSandboxUrl(t *testing.T) {
	logger := &TestLogger{}

	c := NewClient(expectedMerchantId, expectedMerchantKey, logger, WithEnvironment(types.SandboxEnvironment))
	assert.NotNil(t, c)
	assert.Equal(t, expectedMerchantId, c.merchantId)
	assert.Equal(t, expectedMerchantKey, c.merchantKey)
	assert.Equal(t, types.SandboxEnvironment, c.environment)
}
