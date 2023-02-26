package client

import (
	"github.com/bpalermo/cielo-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	zeroAuthPositiveResponse = `
{
  "Valid": true,
  "ReturnCode": "00",
  "ReturnMessage": "Transacao autorizada",
  "IssuerTransactionId": "580027442382078"
}
`
	zeroAuthInvalidCardResponse = `
{
  "Valid": false,
  "ReturnCode": "57",
  "ReturnMessage": "Autorizacao negada"
}
`
	zeroAuthInvalidBrandResponse = `
{
  "Valid": false,
  "ReturnCode": "57",
  "ReturnMessage": "Bandeira inválida"
}
`
	zeroAuthRestrictionResponse = `
{
  "Valid": false,
  "ReturnCode": "389",
  "ReturnMessage": "Restrição Cadastral"
}
`
)

func TestCieloClient_ZeroAuth(t *testing.T) {
	request := &types.ZeroAuthRequest{
		CardNumber:     "1234123412341231",
		Holder:         "Alexsander Rosa",
		ExpirationDate: "12/2021",
		SecurityCode:   "123",
		SaveCard:       false,
		Brand:          "Visa",
		CardOnFile: &types.CardOnFile{
			Usage:  "First",
			Reason: "Recurring",
		},
	}

	tests := []struct {
		name     string
		path     string
		request  *types.ZeroAuthRequest
		response string
		want     *types.ZeroAuthResponse
		wantErr  error
	}{
		{
			name:     "should succeed",
			path:     zeroAuthBasePath,
			request:  request,
			response: zeroAuthPositiveResponse,
			want: &types.ZeroAuthResponse{
				Valid:               true,
				ReturnCode:          "00",
				ReturnMessage:       "Transacao autorizada",
				IssuerTransactionId: "580027442382078",
			},
			wantErr: nil,
		},
		{
			name:     "should fail with denied authorization",
			path:     zeroAuthBasePath,
			request:  request,
			response: zeroAuthInvalidCardResponse,
			want: &types.ZeroAuthResponse{
				Valid:         false,
				ReturnCode:    "57",
				ReturnMessage: "Autorizacao negada",
			},
			wantErr: nil,
		},
		{
			name:     "should fail with invalid brand",
			path:     zeroAuthBasePath,
			request:  request,
			response: zeroAuthInvalidBrandResponse,
			want: &types.ZeroAuthResponse{
				Valid:         false,
				ReturnCode:    "57",
				ReturnMessage: "Bandeira inválida",
			},
			wantErr: nil,
		},
		{
			name:     "should fail with restriction",
			path:     zeroAuthBasePath,
			request:  request,
			response: zeroAuthRestrictionResponse,
			want: &types.ZeroAuthResponse{
				Valid:         false,
				ReturnCode:    "389",
				ReturnMessage: "Restrição Cadastral",
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server, c := NewTestServer(t, tt.path, tt.response)
			defer server.Close()

			got, err := c.ZeroAuth(tt.request)

			assert.NotNil(t, got)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
