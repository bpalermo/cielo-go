package types

type Customer struct {
	Name            string   `json:",omitempty"`
	Email           string   `json:",omitempty"`
	BirthDate       string   `json:",omitempty"`
	Identity        string   `json:",omitempty"`
	IdentityType    string   `json:",omitempty"`
	Address         *Address `json:",omitempty"`
	DeliveryAddress *Address `json:",omitempty"`
}
