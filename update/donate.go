package update

import "github.com/immersivesky/cupcakego/responses"

type Donate struct {
	ID        string                   `json:"id,omitempty"`
	UserID    string                   `json:"user,omitempty"`
	Date      string                   `json:"date,omitempty"`
	Amount    string                   `json:"amount,omitempty"`
	Total     string                   `json:"total,omitempty"`
	Message   string                   `json:"msg,omitempty"`
	Anonymous bool                     `json:"anonym,omitempty"`
	Answer    string                   `json:"answer,omitempty"`
	VKPay     bool                     `json:"vkpay,omitempty"`
	Status    string                   `json:"status,omitempty"`
	Reward    responses.DonationReward `json:"reward,omitempty"`
	Payload   string                   `json:"op,omitempty"`
}
