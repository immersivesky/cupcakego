package update

type Payment struct {
	ID        string `json:"id,omitempty"`
	Status    string `json:"status,omitempty"`
	Created   string `json:"created,omitempty"`
	Processed string `json:"processed,omitempty"`
	Purse     string `json:"purse,omitempty"`
	Amount    string `json:"amount,omitempty"`
	UserID    string `json:"user,omitempty"`
}
