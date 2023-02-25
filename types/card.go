package types

type CreditCard struct {
	CardNumber              string  `json:",omitempty"`
	CustomerName            string  `json:",omitempty"`
	Holder                  string  `json:",omitempty"`
	ExpirationDate          string  `json:",omitempty"`
	SecurityCode            string  `json:",omitempty"`
	SaveCard                bool    `json:",omitempty"`
	Brand                   string  `json:",omitempty"`
	CardToken               string  `json:",omitempty"`
	PaymentAccountReference string  `json:",omitempty"`
	Links                   []*Link `json:"-"`
}

type DebitCard struct {
	CardNumber                  string           `json:",omitempty"`
	CustomerName                string           `json:",omitempty"`
	Authenticate                bool             `json:",omitempty"`
	ReturnUrl                   string           `json:",omitempty"`
	IsCryptoCurrencyNegotiation bool             `json:",omitempty"`
	Holder                      string           `json:",omitempty"`
	ExpirationDate              string           `json:",omitempty"`
	RecurrentPayment            RecurrentPayment `json:",omitempty"`
	SecurityCode                string           `json:",omitempty"`
	CardOnFile                  CardOnFile       `json:",omitempty"`
	Brand                       string           `json:",omitempty"`
	CardToken                   string           `json:",omitempty"`
	Links                       []*Link          `json:",omitempty"`
}

type CardOnFile struct {
	Usage  string `json:",omitempty"`
	Reason string `json:",omitempty"`
}

type CardBinResponse struct {
	Status        string `json:",omitempty"`
	Provider      string `json:",omitempty"`
	CardType      string `json:",omitempty"`
	ForeignCard   bool   `json:""`
	CorporateCard bool   `json:""`
	Issuer        string `json:",omitempty"`
	IssuerCode    string `json:",omitempty"`
	Prepaid       bool   `json:""`
}
