package types

type Sale struct {
	MerchantOrderID string    `json:",omitempty"`
	Customer        *Customer `json:",omitempty"`
	Payment         *Payment  `json:",omitempty"`
}
