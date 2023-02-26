package client

import (
	"encoding/json"
	"github.com/bpalermo/cielo-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	saleRequest = `
{
  "MerchantOrderId":"2014111701",
  "Customer":{
    "Name":"Comprador Teste",
    "Identity":"11225468954",
    "IdentityType":"CPF",
    "Email":"compradorteste@teste.com",
    "Birthdate":"1991-01-02",
    "Address":{
      "Street":"Rua Teste",
      "Number":"123",
      "Complement":"AP 123",
      "ZipCode":"12345987",
      "City":"Rio de Janeiro",
      "State":"RJ",
      "Country":"BRA"
    },
    "DeliveryAddress":{
      "Street":"Rua Teste",
      "Number":"123",
      "Complement":"AP 123",
      "ZipCode":"12345987",
      "City":"Rio de Janeiro",
      "State":"RJ",
      "Country":"BRA"
    }
  },
  "Payment":{
    "Type":"CreditCard",
    "Amount":15700,
    "Currency":"BRL",
    "Country":"BRA",
    "Provider":"Simulado",
    "ServiceTaxAmount":0,
    "Installments":1,
    "Interest":"ByMerchant",
    "Capture":false,
    "Authenticate":false,
    "Recurrent":false,
    "SoftDescriptor":"123456789ABCD",
    "CreditCard":{
      "CardNumber":"4024007197692931",
      "Holder":"Teste Holder",
      "ExpirationDate":"12/2021",
      "SecurityCode":"123",
      "SaveCard":false,
      "Brand":"Visa"
    }
  }
}
`
	saleGoodResponse = `
{
  "MerchantOrderId":"2014111701",
  "Customer":{
    "Name":"Comprador Teste",
    "Identity":"11225468954",
    "IdentityType":"CPF",
    "Email":"compradorteste@teste.com",
    "Birthdate":"1991-01-02",
    "Address":{
      "Street":"Rua Teste",
      "Number":"123",
      "Complement":"AP 123",
      "ZipCode":"12345987",
      "City":"Rio de Janeiro",
      "State":"RJ",
      "Country":"BRA",
      "AddressType":0
    },
    "DeliveryAddress":{
      "Street":"Rua Teste",
      "Number":"123",
      "Complement":"AP 123",
      "ZipCode":"12345987",
      "City":"Rio de Janeiro",
      "State":"RJ",
      "Country":"BRA",
      "AddressType":0
    }
  },
  "Payment":{
    "ServiceTaxAmount":0,
    "Installments":1,
    "Interest":0,
    "Capture":false,
    "Authenticate":false,
    "Recurrent":false,
    "CreditCard":{
      "CardNumber":"402400******2931",
      "Holder":"Teste Holder",
      "ExpirationDate":"12/2021",
      "SaveCard":false,
      "Brand":"Visa",
      "PaymentAccountReference":"4A3CT2PC6Z1APY356XQX2RDI28OG8"
    },
    "Tid":"0226090314511",
    "ProofOfSale":"130049",
    "AuthorizationCode":"184697",
    "SoftDescriptor":"123456789ABCD",
    "Provider":"Simulado",
    "IsQrCode":false,
    "DynamicCurrencyConversion":false,
    "Amount":15700,
    "ReceivedDate":"2023-02-26 09:03:14",
    "Status":1,
    "IsSplitted":false,
    "ReturnMessage":"Operation Successful",
    "ReturnCode":"4",
    "PaymentId":"129d7a53-04ff-4cd2-bbd3-f8455d057efe",
    "Type":"CreditCard",
    "Currency":"BRL",
    "Country":"BRA",
    "Links":[
      {
        "Method":"GET",
        "Rel":"self",
        "Href":"https://apiquerysandbox.cieloecommerce.cielo.com.br/1/sales/129d7a53-04ff-4cd2-bbd3-f8455d057efe"
      },
      {
        "Method":"PUT",
        "Rel":"capture",
        "Href":"https://apisandbox.cieloecommerce.cielo.com.br/1/sales/129d7a53-04ff-4cd2-bbd3-f8455d057efe/capture"
      },
      {
        "Method":"PUT",
        "Rel":"void",
        "Href":"https://apisandbox.cieloecommerce.cielo.com.br/1/sales/129d7a53-04ff-4cd2-bbd3-f8455d057efe/void"
      }
    ]
  }
}
`
)

func TestCieloClient_Authorize(t *testing.T) {
	server, c := NewTestServer(t, saleBasePath, saleGoodResponse)
	defer server.Close()

	req := &types.Sale{}

	err := json.Unmarshal([]byte(saleRequest), &req)
	if err != nil {
		t.Fatal(err)
	}

	res, err := c.Authorize(req)
	if err != nil {
		t.Fatal(err)
	}

	var expectedStatus uint32 = 1

	assert.NotNil(t, res)
	assert.Equal(t, "2014111701", res.MerchantOrderID)

	assert.NotNil(t, res.Customer)
	assert.Equal(t, "Comprador Teste", res.Customer.Name)
	assert.Equal(t, "11225468954", res.Customer.Identity)
	assert.Equal(t, "CPF", res.Customer.IdentityType)

	assert.NotNil(t, res.Payment)
	assert.Equal(t, "0226090314511", res.Payment.Tid)
	assert.Equal(t, "129d7a53-04ff-4cd2-bbd3-f8455d057efe", res.Payment.PaymentID)
	assert.Equal(t, "130049", res.Payment.ProofOfSale)
	assert.Equal(t, "184697", res.Payment.AuthorizationCode)
	assert.Equal(t, "123456789ABCD", res.Payment.SoftDescriptor)
	assert.Equal(t, "Operation Successful", res.Payment.ReturnMessage)
	assert.Equal(t, "4", res.Payment.ReturnCode)
	assert.Equal(t, expectedStatus, res.Payment.Status)
	assert.Equal(t, false, res.Payment.Authenticate)
	assert.Equal(t, false, res.Payment.IsCryptoCurrencyNegotiation)
	assert.Equal(t, false, res.Payment.TryAutomaticCancellation)

	assert.NotNil(t, res.Payment.CreditCard)
	assert.Equal(t, "4A3CT2PC6Z1APY356XQX2RDI28OG8", res.Payment.CreditCard.PaymentAccountReference)
}
