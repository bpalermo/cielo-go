package types

type Payment struct {
	ServiceTaxAmount            uint32            `json:",omitempty"`
	Installments                uint32            `json:",omitempty"`
	Interest                    interface{}       `json:",omitempty"`
	Capture                     bool              `json:",omitempty"`
	Authenticate                bool              `json:",omitempty"`
	Recurrent                   bool              `json:",omitempty"`
	RecurrentPayment            *RecurrentPayment `json:",omitempty"`
	CreditCard                  *CreditCard       `json:",omitempty"`
	DebitCard                   *DebitCard        `json:",omitempty"`
	Tid                         string            `json:",omitempty"`
	ProofOfSale                 string            `json:",omitempty"`
	AuthorizationCode           string            `json:",omitempty"`
	SoftDescriptor              string            `json:",omitempty"`
	ReturnURL                   string            `json:",omitempty"`
	Provider                    string            `json:",omitempty"`
	PaymentID                   string            `json:",omitempty"`
	Type                        string            `json:",omitempty"`
	Amount                      uint32            `json:",omitempty"`
	ReceiveDate                 string            `json:",omitempty"`
	CapturedAmount              uint32            `json:",omitempty"`
	CapturedDate                string            `json:",omitempty"`
	Currency                    string            `json:",omitempty"`
	Country                     string            `json:",omitempty"`
	ReturnCode                  string            `json:",omitempty"`
	ReturnMessage               string            `json:",omitempty"`
	Status                      uint32            `json:",omitempty"`
	Links                       []*Link           `json:",omitempty"`
	ExtraDataCollection         []string          `json:",omitempty"`
	ExpirationDate              string            `json:",omitempty"`
	URL                         string            `json:",omitempty"`
	Number                      string            `json:",omitempty"`
	BarCodeNumber               string            `json:",omitempty"`
	DigitableLine               string            `json:",omitempty"`
	Address                     string            `json:",omitempty"`
	IsCryptoCurrencyNegotiation bool              `json:",omitempty"`
	TryAutomaticCancellation    bool              `json:",omitempty"`
}
