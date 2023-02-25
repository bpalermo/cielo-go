package types

type ZeroAuthRequest struct {
	CardType       string      `json:"CardType,omitempty"`
	CardNumber     string      `json:"CardNumber,omitempty"`
	Holder         string      `json:"Holder,omitempty"`
	ExpirationDate string      `json:"ExpirationDate,omitempty"`
	SecurityCode   string      `json:"SecurityCode,omitempty"`
	SaveCard       bool        `json:"SaveCard"`
	Brand          string      `json:"Brand,omitempty"`
	CardToken      string      `json:"CardToken,omitempty"`
	CardOnFile     *CardOnFile `json:",omitempty"`
}

type ZeroAuthResponse struct {
	Valid               bool   `json:"Valid"`
	ReturnCode          string `json:"ReturnCode,omitempty"`
	ReturnMessage       string `json:"ReturnMessage,omitempty"`
	IssuerTransactionId string `json:"IssuerTransactionId,omitempty"`
}
