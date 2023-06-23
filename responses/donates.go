package responses

type Donations struct {
	Info
	Donations []Donation `json:"list,omitempty"`
}

type Donation struct {
	ID        string         `json:"id,omitempty"`
	User      string         `json:"user,omitempty"`
	Date      string         `json:"date,omitempty"`
	Amount    string         `json:"amount,omitempty"`
	Total     string         `json:"total,omitempty"`
	Message   string         `json:"msg,omitempty"`
	Anonymous bool           `json:"anonym,omitempty"`
	Answer    string         `json:"answer,omitempty"`
	VKPay     bool           `json:"vkpay,omitempty"`
	Status    string         `json:"status,omitempty"`
	Reward    DonationReward `json:"reward,omitempty"`
	Payload   string         `json:"op,omitempty"`
}

type Status struct {
	Info
}

type Answer struct {
	Info
}

type ChangeReward struct {
	Info
}
