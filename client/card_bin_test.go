package client

import (
	"fmt"
	"github.com/bpalermo/cielo-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	cardBinGoodResponse = `
{
    "Status": "00",
    "Provider": "MASTERCARD",
    "CardType": "Crédito",
    "ForeignCard": true,
    "Crédito": true,
    "CorporateCard": true,
    "Issuer": "Bradesco",
    "IssuerCode": "237",
    "Prepaid": true
}
`
)

func TestCieloClient_CardBin(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		cardBin  string
		response string
		want     *types.CardBinResponse
		wantErr  error
	}{
		{
			name:     "should succeed",
			path:     cardBinBasePath,
			cardBin:  "132456",
			response: cardBinGoodResponse,
			want: &types.CardBinResponse{
				Status:        "00",
				Provider:      "MASTERCARD",
				CardType:      "Crédito",
				ForeignCard:   true,
				CorporateCard: true,
				Issuer:        "Bradesco",
				IssuerCode:    "237",
				Prepaid:       true,
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server, c := NewTestServer(t, fmt.Sprintf("%s/%s", tt.path, tt.cardBin), tt.response)
			defer server.Close()

			got, err := c.CardBin(tt.cardBin)

			assert.NotNil(t, got)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
