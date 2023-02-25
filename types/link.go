package types

type Link struct {
	Method string `json:",omitempty"`
	Rel    string `json:",omitempty"`
	Href   string `json:",omitempty"`
}
