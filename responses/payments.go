package responses

type Payments struct {
	Info
	Payments []Payment `json:"list,omitempty"`
}

type Payment struct {
	ID        string `json:"id,omitempty"`
	Status    string `json:"status,omitempty"`
	Processed string `json:"processed,omitempty"`
	System    string `json:"system,omitempty"`
	Purse     string `json:"purse,omitempty"`
	Amount    string `json:"amount,omitempty"`
	UserID    string `json:"user,omitempty"`
}

type CreatePayments struct {
	Info
	ID string `json:"id,omitempty"`
}
