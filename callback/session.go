package callback

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/immersivesky/cupcakego/update"
)

func (session *Options) Fiber(context *fiber.Ctx) error {
	update := update.Update{}
	if err := json.Unmarshal(context.Body(), &update); err != nil {
		session.Application.Logger.Error("Update JSONMarshal error is ", err)
	}

	if session.Application.Logger != nil {
		session.Application.Logger.Info("Update response is ", string(context.Body()))
	}

	if update.Type == "confirmation" {
		return context.SendString(`{"status": "ok", "code": "` + context.Params("confirmation", "no_confirmation_param") + `"}`)
	} else {
		switch update.Type {
		case "new_donate":
			session.Scenes.Donate(session.Application, update.Donate)
		case "payment_status":
			session.Scenes.Payment(session.Application, update.Payment)
		}
	}

	return context.SendString(`{"status": "ok"}`)
}
