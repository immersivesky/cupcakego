package API

import (
	"fmt"
	"github.com/immersivesky/cupcakego/responses"
)

type GetPayments struct {
	IDs []int
}

func (cupcake *Options) GetPayments(properties ...any) (payments responses.Payments) {
	var params = map[string]interface{}{
		"group": cupcake.ID,
		"token": cupcake.Token,
		"v":     cupcake.Version,
	}

	for _, property := range properties {
		switch property.(type) {
		case GetPayments:
			payments := property.(GetPayments)

			if len(payments.IDs) > 0 {
				params["ids"] = payments.IDs
			}
		case []int:
			params["ids"] = property.([]int)
		}
	}

	if err := cupcake.Call("payments/get", params, &payments); err != nil {
		fmt.Println(err)
	}

	return payments
}

type CreatePayments struct {
	System string
	Purse  string
	Name   string
	Amount int
}

func (cupcake *Options) CreatePayments(properties ...any) (payments responses.CreatePayments) {
	var (
		params = map[string]interface{}{
			"group": cupcake.ID,
			"token": cupcake.Token,
			"v":     cupcake.Version,
		}
		system bool
		purse  bool
	)

	for _, property := range properties {
		switch property.(type) {
		case CreatePayments:
			payments := property.(CreatePayments)

			if payments.System != "" {
				params["system"] = payments.System
				system = true
			}
			if payments.Purse != "" {
				params["purse"] = payments.Purse
				purse = true
			}
			if payments.Amount != 0 {
				params["amount"] = payments.Amount
			}

			if payments.Name != "" {
				params["name"] = payments.Name
			}
		case int:
			params["amount"] = property.(int)
		case float32:
			params["amount"] = property.(float32)
		case float64:
			params["amount"] = property.(float64)
		case string:
			value := property.(string)
			if !system {
				params["system"] = value
				system = true
			} else if !purse {
				params["purse"] = value
				purse = true
			} else {
				params["name"] = value
			}
		}
	}

	if err := cupcake.Call("payments/create", params, &payments); err != nil {
		fmt.Println(err)
	}

	return payments
}
