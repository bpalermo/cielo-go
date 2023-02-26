package types

type RecurrentPayment struct {
	AuthorizeNow       bool   `json:",omitempty"`
	EndDate            string `json:",omitempty"`
	Interval           uint32 `json:",omitempty"`
	RecurrentPaymentId string `json:",omitempty"`
	ReasonCode         int32  `json:",omitempty"`
	ReasonMessage      string `json:",omitempty"`
	NextRecurrency     string `json:",omitempty"`
}
