package client

import (
	"github.com/bpalermo/cielo-go/types"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const (
	saleMinimumResponse = `
{
    "MerchantOrderId": "2014111703",
    "Customer": {
        "Name": "[Guest]"
    },
    "Payment": {
        "ServiceTaxAmount": 0,
        "Installments": 1,
        "Interest": 0,
        "Capture": false,
        "Authenticate": false,
        "Recurrent": false,
        "CreditCard": {
            "CardNumber": "455187******0183",
            "Holder": "Teste Holder",
            "ExpirationDate": "12/2025",
            "SaveCard": false,
            "Brand": "Visa"
        },
        "Tid": "0226070625179",
        "SoftDescriptor": "123456789ABCD",
        "Provider": "Simulado",
        "IsQrCode": false,
        "DynamicCurrencyConversion": false,
        "Amount": 15700,
        "ReceivedDate": "2023-02-26 19:06:25",
        "Status": 3,
        "IsSplitted": false,
        "ReturnMessage": "Card Expired",
        "ReturnCode": "57",
        "PaymentId": "0b4bef34-dbd6-415e-b305-def739142378",
        "Type": "CreditCard",
        "Currency": "BRL",
        "Country": "BRA",
        "Links": [
            {
                "Method": "GET",
                "Rel": "self",
                "Href": "https://apiquerysandbox.cieloecommerce.cielo.com.br/1/sales/0b4bef34-dbd6-415e-b305-def739142378"
            }
        ]
    }
}
`
)

func TestCieloClient_Authorize(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		request  *types.Sale
		response string
		want     *types.Sale
		wantErr  error
	}{
		{
			name: "succeed minimum sale",
			path: saleBasePath,
			request: &types.Sale{
				MerchantOrderID: "2014111703",
				Payment: &types.Payment{
					Provider:       "Simulado",
					Type:           "CreditCard",
					Amount:         15700,
					Installments:   1,
					SoftDescriptor: "123456789ABCD",
					CreditCard: &types.CreditCard{
						CardNumber:     "4551870000000183",
						Holder:         "Teste Holder",
						ExpirationDate: "12/2021",
						SecurityCode:   "123",
						Brand:          "Visa",
					},
				},
			},
			response: saleMinimumResponse,
			want: &types.Sale{
				MerchantOrderID: "2014111703",
				Customer: &types.Customer{
					Name: "[Guest]",
				},
				Payment: &types.Payment{
					ServiceTaxAmount: 0,
					Installments:     1,
					Interest:         0.0,
					Capture:          false,
					Authenticate:     false,
					Recurrent:        false,
					CreditCard: &types.CreditCard{
						CardNumber:     "455187******0183",
						Holder:         "Teste Holder",
						ExpirationDate: "12/2025",
						SaveCard:       false,
						Brand:          "Visa",
					},
					Tid:                       "0226070625179",
					SoftDescriptor:            "123456789ABCD",
					Provider:                  "Simulado",
					IsQrCode:                  false,
					DynamicCurrencyConversion: false,
					Amount:                    15700,
					ReceivedDate:              types.DateTime(time.Date(2023, 2, 26, 19, 6, 25, 0, time.UTC)),
					Status:                    3,
					IsSplitted:                false,
					ReturnMessage:             "Card Expired",
					ReturnCode:                "57",
					PaymentID:                 uuid.Must(uuid.Parse("0b4bef34-dbd6-415e-b305-def739142378")),
					Type:                      "CreditCard",
					Currency:                  "BRL",
					Country:                   "BRA",
					Links: []*types.Link{
						{
							Method: "GET",
							Rel:    "self",
							Href:   "https://apiquerysandbox.cieloecommerce.cielo.com.br/1/sales/0b4bef34-dbd6-415e-b305-def739142378",
						},
					},
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server, c := NewTestServer(t, tt.path, tt.response)
			defer server.Close()

			got, err := c.Authorize(tt.request)

			assert.NotNil(t, got)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
