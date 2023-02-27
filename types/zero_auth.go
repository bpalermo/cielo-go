package types

type ZeroAuthRequest struct {
	CardType       string      `json:",omitempty"`
	CardNumber     string      `json:",omitempty"`
	Holder         string      `json:",omitempty"`
	ExpirationDate string      `json:",omitempty"`
	SecurityCode   string      `json:",omitempty"`
	SaveCard       bool        `json:""`
	Brand          string      `json:",omitempty"`
	CardToken      string      `json:",omitempty"`
	CardOnFile     *CardOnFile `json:",omitempty"`
}

type ZeroAuthResponse struct {
	Valid               bool   `json:""`
	ReturnCode          string `json:",omitempty"`
	ReturnMessage       string `json:",omitempty"`
	IssuerTransactionId string `json:",omitempty"`
}
