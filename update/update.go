package update

type Update struct {
	Group   string  `json:"group,omitempty"`
	Type    string  `json:"type,omitempty"`
	Hash    string  `json:"hash,omitempty"`
	Donate  Donate  `json:"donate,omitempty"`
	Payment Payment `json:"payment,omitempty"`
}
