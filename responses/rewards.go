package responses

type DonationReward struct {
	ID     string `json:"id,omitempty"`
	Title  string `json:"title,omitempty"`
	Status string `json:"status,omitempty"`
}

type CampaignReward struct {
	ID          string `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"desc,omitempty"`
	Amount      string `json:"min_donate,omitempty"`
	Limits      string `json:"limits,omitempty"`
	Status      string `json:"status,omitempty"`
	Backers     string `json:"backers,omitempty"`
}
