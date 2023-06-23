package tests

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/immersivesky/cupcakego/callback"
	"github.com/immersivesky/cupcakego/update"
	"os"
	"strconv"
	"testing"

	"github.com/immersivesky/cupcakego/API"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func TestCallback(t *testing.T) {
	godotenv.Load()

	var (
		logger     = logrus.New()
		groupID, _ = strconv.Atoi(os.Getenv("GROUP_ID"))
	)

	logger.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	donuts := API.Create(os.Getenv("TOKEN"), groupID, logger)

	t.Run("Callback", func(t *testing.T) {
		var (
			app     = fiber.New()
			session = callback.Create(donuts)
		)

		session.Donate(func(app *API.Options, donate update.Donate) {
			fmt.Printf("New donate: %+v\n", donate)
		})

		app.Post("/cupcake/:confirmation", session.Fiber)
		app.Listen(":3000")
	})
}
