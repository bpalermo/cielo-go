package types

import (
	"github.com/google/uuid"
)

type Payment struct {
	Address                     string            `json:",omitempty"`
	Amount                      uint32            `json:",omitempty"`
	Authenticate                bool              `json:",omitempty"`
	AuthorizationCode           string            `json:",omitempty"`
	BarCodeNumber               string            `json:",omitempty"`
	Capture                     bool              `json:",omitempty"`
	CapturedAmount              uint32            `json:",omitempty"`
	CapturedDate                string            `json:",omitempty"`
	Country                     string            `json:",omitempty"`
	CreditCard                  *CreditCard       `json:",omitempty"`
	Currency                    string            `json:",omitempty"`
	DebitCard                   *DebitCard        `json:",omitempty"`
	DigitableLine               string            `json:",omitempty"`
	DynamicCurrencyConversion   bool              `json:",omitempty"`
	ExpirationDate              string            `json:",omitempty"`
	ExtraDataCollection         []string          `json:",omitempty"`
	Installments                uint32            `json:",omitempty"`
	Interest                    float64           `json:",omitempty"`
	IsCryptoCurrencyNegotiation bool              `json:",omitempty"`
	IsQrCode                    bool              `json:",omitempty"`
	IsSplitted                  bool              `json:",omitempty"`
	Links                       []*Link           `json:",omitempty"`
	Number                      string            `json:",omitempty"`
	PaymentID                   uuid.UUID         `json:",omitempty"`
	AcquirerTransactionId       uuid.UUID         `json:",omitempty"`
	QrcodeBase64Image           string            `json:",omitempty"`
	QrCodeString                string            `json:",omitempty"`
	ProofOfSale                 string            `json:",omitempty"`
	Provider                    string            `json:",omitempty"`
	ReceivedDate                DateTime          `json:",omitempty"`
	Recurrent                   bool              `json:",omitempty"`
	RecurrentPayment            *RecurrentPayment `json:",omitempty"`
	ReturnCode                  string            `json:",omitempty"`
	ReturnMessage               string            `json:",omitempty"`
	ReturnURL                   string            `json:",omitempty"`
	ServiceTaxAmount            uint32            `json:",omitempty"`
	SoftDescriptor              string            `json:",omitempty"`
	Status                      uint32            `json:",omitempty"`
	Tid                         string            `json:",omitempty"`
	TryAutomaticCancellation    bool              `json:",omitempty"`
	Type                        string            `json:",omitempty"`
	URL                         string            `json:",omitempty"`
}
