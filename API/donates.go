package API

import (
	"fmt"

	"github.com/immersivesky/cupcakego/responses"
)

type GetDonations struct {
	Count     int
	Offset    int
	StartDate int
	EndDate   int
	Sort      string
	Reverse   bool
}

func (cupcake Options) GetDonates(properties ...any) (donations responses.Donations) {
	return cupcake.GetDonations(properties...)
}

func (cupcake Options) GetDonations(properties ...any) (donations responses.Donations) {
	var (
		params = map[string]interface{}{
			"group": cupcake.ID,
			"token": cupcake.Token,
			"v":     cupcake.Version,
		}
		offset bool
	)

	for _, property := range properties {
		switch property.(type) {
		case GetDonations:
			donations := property.(GetDonations)

			if donations.Count != 0 {
				params["len"] = donations.Count
			}
			if donations.Offset != 0 {
				params["offset"] = donations.Offset
				offset = true
			}

			if donations.StartDate != 0 {
				params["start_date"] = donations.StartDate
			}
			if donations.EndDate != 0 {
				params["end_date"] = donations.EndDate
			}

			if donations.Sort != "" {
				params["sort"] = donations.Sort
			}
			if donations.Reverse {
				params["reverse"] = donations.Reverse
			}
		case Count:
			params["len"] = property.(Count)
		case Offset:
			params["offset"] = property.(Offset)
			offset = true
		case StartDate:
			params["start_date"] = property.(StartDate)
		case EndDate:
			params["end_date"] = property.(EndDate)
		case Sort:
			params["sort"] = property.(Sort)
		case Reverse:
			params["reverse"] = property.(Reverse)
		case int:
			value := property.(int)

			if !offset {
				params["offset"] = value
				offset = true
			} else {
				params["len"] = value
			}
		case string:
			value := property.(int)
			params["sort"] = value
		case bool:
			params["reverse"] = property.(bool)
		}
	}

	if err := cupcake.Call("donates/get", params, &donations); err != nil {
		fmt.Println(err)
	}

	return donations
}

type GetLastDonations struct {
	OffsetID int
}

func (cupcake Options) GetLastDonates(properties ...any) (donations responses.Donations) {
	return cupcake.GetLastDonations(properties...)
}

func (cupcake Options) GetLastDonations(properties ...any) (donations responses.Donations) {
	var params = map[string]interface{}{
		"group": cupcake.ID,
		"token": cupcake.Token,
		"v":     cupcake.Version,
	}

	for _, property := range properties {
		switch property.(type) {
		case GetLastDonations:
			donations := property.(GetLastDonations)

			if donations.OffsetID != 0 {
				params["last"] = donations.OffsetID
			}
		case OffsetID:
			params["last"] = property.(OffsetID)
		case int:
			params["last"] = property.(int)
		case string:
			params["last"] = property.(string)
		}
	}

	if err := cupcake.Call("donates/get-last", params, &donations); err != nil {
		fmt.Println(err)
	}

	return donations
}

type ChangeStatus struct {
	ID     int
	Status string
}

func (cupcake *Options) SetStatus(properties ...any) (status responses.Status) {
	return cupcake.ChangeStatus(properties...)
}

func (cupcake *Options) ChangeStatus(properties ...any) (status responses.Status) {
	var (
		params = map[string]interface{}{
			"group": cupcake.ID,
			"token": cupcake.Token,
			"v":     cupcake.Version,
		}
		donateID bool
	)

	for _, property := range properties {
		switch property.(type) {
		case ChangeStatus:
			status := property.(ChangeStatus)

			if status.ID != 0 {
				params["id"] = status.ID
				donateID = true
			}

			if status.Status != "" {
				params["status"] = status.Status
			}
		case ID:
			params["id"] = property.(ID)
		case Status:
			params["status"] = property.(Status)
		case int:
			params["id"] = property.(int)
			donateID = true
		case string:
			value := property.(string)
			if !donateID {
				params["id"] = value
				donateID = true
			} else {
				params["status"] = value
			}
		}
	}

	if err := cupcake.Call("donates/change-status", params, &status); err != nil {
		fmt.Println(err)
	}

	return status
}

type Answer struct {
	ID     int
	Answer string
}

func (cupcake *Options) Answer(properties ...any) (answer responses.Answer) {
	var (
		params = map[string]interface{}{
			"group": cupcake.ID,
			"token": cupcake.Token,
			"v":     cupcake.Version,
		}
		donateID bool
	)

	for _, property := range properties {
		switch property.(type) {
		case Answer:
			answer := property.(Answer)

			if answer.ID != 0 {
				params["id"] = answer.ID
				donateID = true
			}
			if answer.Answer != "" {
				params["answer"] = answer.Answer
			}
		case int:
			params["id"] = property.(int)
			donateID = true
		case string:
			value := property.(string)
			if !donateID {
				params["id"] = value
				donateID = true
			} else {
				params["answer"] = value
			}
		}
	}

	if err := cupcake.Call("donates/answer", params, &answer); err != nil {
		fmt.Println(err)
	}

	return answer
}

type ChangeReward struct {
	ID     int
	Status string
}

func (cupcake *Options) SetRewardStatus(properties ...any) (reward responses.ChangeReward) {
	return cupcake.ChangeReward(properties...)
}

func (cupcake *Options) ChangeRewardStatus(properties ...any) (reward responses.ChangeReward) {
	return cupcake.ChangeReward(properties...)
}

func (cupcake *Options) SetReward(properties ...any) (reward responses.ChangeReward) {
	return cupcake.ChangeReward(properties...)
}

func (cupcake *Options) ChangeReward(properties ...any) (reward responses.ChangeReward) {
	var (
		params = map[string]interface{}{
			"group": cupcake.ID,
			"token": cupcake.Token,
			"v":     cupcake.Version,
		}
		donateID bool
	)

	for _, property := range properties {
		switch property.(type) {
		case ChangeReward:
			reward := property.(ChangeReward)

			if reward.ID != 0 {
				params["id"] = reward.ID
				donateID = true
			}
			if reward.Status != "" {
				params["status"] = reward.Status
			}
		case int:
			params["id"] = property.(int)
			donateID = true
		case string:
			value := property.(string)
			if !donateID {
				params["id"] = value
				donateID = true
			} else {
				params["status"] = value
			}
		}
	}

	if err := cupcake.Call("donates/change-reward-status", params, &reward); err != nil {
		fmt.Println(err)
	}

	return reward
}
