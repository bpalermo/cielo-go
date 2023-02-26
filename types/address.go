package types

type Address struct {
	Street     string `json:",omitempty"`
	Number     string `json:",omitempty"`
	Complement string `json:",omitempty"`
	ZipCode    string `json:",omitempty"`
	City       string `json:",omitempty"`
	State      string `json:",omitempty"`
	Country    string `json:",omitempty"`
}
