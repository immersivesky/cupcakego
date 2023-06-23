package API

import (
	"fmt"
	"github.com/immersivesky/cupcakego/responses"
)

type GetCampaigns struct {
	IDs []int
}

func (cupcake *Options) GetCampaigns(properties ...any) (campaigns responses.Campaigns) {
	var params = map[string]interface{}{
		"group": cupcake.ID,
		"token": cupcake.Token,
		"v":     cupcake.Version,
	}

	for _, property := range properties {
		switch property.(type) {
		case GetCampaigns:
			payments := property.(GetCampaigns)
			if len(payments.IDs) > 0 {
				params["ids"] = payments.IDs
			}
		case []int:
			params["ids"] = property.([]int)
		}
	}

	if err := cupcake.Call("campaigns/get", params, &campaigns); err != nil {
		fmt.Println(err)
	}

	return campaigns
}

func (cupcake *Options) GetActiveCampaign(properties ...any) (campaigns responses.ActiveCampaign) {
	var params = map[string]interface{}{
		"group": cupcake.ID,
		"token": cupcake.Token,
		"v":     cupcake.Version,
	}

	if err := cupcake.Call("campaigns/get-active", params, &campaigns); err != nil {
		fmt.Println(err)
	}

	return campaigns
}

type GetRewardsCampaigns struct {
	Campaign int
}

func (cupcake *Options) GetCampaignRewards(properties ...any) (campaigns responses.RewardsCampaign) {
	var params = map[string]interface{}{
		"group": cupcake.ID,
		"token": cupcake.Token,
		"v":     cupcake.Version,
	}

	for _, property := range properties {
		switch property.(type) {
		case GetCampaigns:
			rewards := property.(GetRewardsCampaigns)
			if rewards.Campaign != 0 {
				params["campaign"] = rewards.Campaign
			}
		case int:
			params["campaign"] = property.(int)
		}
	}

	if err := cupcake.Call("campaigns/get-rewards", params, &campaigns); err != nil {
		fmt.Println(err)
	}

	return campaigns
}

type ChangeCampaign struct {
	ID            int
	Title         string
	Status        string
	End           int
	Point         int
	StartReceived int
	StartBackers  int
}

func (cupcake *Options) SetCampaign(properties ...any) (campaigns responses.ChangeCampaign) {
	return cupcake.ChangeCampaign(properties...)
}

func (cupcake *Options) ChangeCampaign(properties ...any) (campaigns responses.ChangeCampaign) {
	var (
		params = map[string]interface{}{
			"group": cupcake.ID,
			"token": cupcake.Token,
			"v":     cupcake.Version,
		}
		ID bool
	)

	for _, property := range properties {
		switch property.(type) {
		case ChangeCampaign:
			campaign := property.(ChangeCampaign)

			if campaign.ID != 0 {
				params["id"] = campaign.ID
				ID = true
			}

			if campaign.Title != "" {
				params["title"] = campaign.Title
			}
			if campaign.Status != "" {
				params["status"] = campaign.Status
			}

			if campaign.End != 0 {
				params["end"] = campaign.End
			}
			if campaign.Point != 0 {
				params["point"] = campaign.Point
			}

			if campaign.StartReceived != 0 {
				params["start_received"] = campaign.StartReceived
			}
			if campaign.StartBackers != 0 {
				params["start_backers"] = campaign.StartBackers
			}
		case int:
			value := property.(int)
			if !ID {
				params["id"] = value
				ID = true
			} else {
				params["point"] = value
			}
		case string:
			params["status"] = property.(string)
		}
	}

	if err := cupcake.Call("campaigns/change", params, &campaigns); err != nil {
		fmt.Println(err)
	}

	return campaigns
}

type ChangeCampaignReward struct {
	ID          int
	Title       string
	Description string
	Amount      int
	Limits      int
	Status      string
}

func (cupcake *Options) SetCampaignReward(properties ...any) (campaigns responses.ChangeCampaign) {
	return cupcake.ChangeCampaignReward(properties...)
}

func (cupcake *Options) ChangeCampaignReward(properties ...any) (campaigns responses.ChangeCampaign) {
	var (
		params = map[string]interface{}{
			"group": cupcake.ID,
			"token": cupcake.Token,
			"v":     cupcake.Version,
		}
		ID bool
	)

	for _, property := range properties {
		switch property.(type) {
		case ChangeCampaignReward:
			reward := property.(ChangeCampaignReward)

			if reward.ID != 0 {
				params["id"] = reward.ID
				ID = true
			}

			if reward.Title != "" {
				params["title"] = reward.Title
			}
			if reward.Description != "" {
				params["desc"] = reward.Description
			}
			if reward.Status != "" {
				params["status"] = reward.Status
			}

			if reward.Amount != 0 {
				params["min_donate"] = reward.Amount
			}
			if reward.Limits != 0 {
				params["limits"] = reward.Limits
			}
		case int:
			value := property.(int)
			if !ID {
				params["id"] = value
				ID = true
			} else {
				params["min_donate"] = value
			}
		case string:
			params["status"] = property.(string)
		}
	}

	if err := cupcake.Call("campaigns/change-reward", params, &campaigns); err != nil {
		fmt.Println(err)
	}

	return campaigns
}
