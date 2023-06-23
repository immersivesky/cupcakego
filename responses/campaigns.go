package responses

type Campaigns struct {
	Info
	Campaigns []Campaign `json:"list"`
}

type ActiveCampaign struct {
	Info
	Campaign Campaign `json:"campaign"`
}

type Campaign struct {
	ID            string `json:"id,omitempty"`
	Title         string `json:"title,omitempty"`
	Status        string `json:"status,omitempty"`
	Start         string `json:"start,omitempty"`
	End           string `json:"end,omitempty"`
	Point         string `json:"point,omitempty"`
	StartReceived string `json:"start_received,omitempty"`
	StartBackers  string `json:"start_backers,omitempty"`
	Received      string `json:"received,omitempty"`
	Backers       string `json:"backers,omitempty"`
}

type RewardsCampaign struct {
	Info
	Rewards []CampaignReward `json:"list,omitempty"`
}

type ChangeCampaign struct {
	Info
}
