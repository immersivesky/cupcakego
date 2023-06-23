package tests

import (
	"os"
	"strconv"
	"testing"

	"github.com/immersivesky/cupcakego/API"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestBalance(t *testing.T) {
	godotenv.Load()

	var (
		logger     = logrus.New()
		groupID, _ = strconv.Atoi(os.Getenv("GROUP_ID"))
	)

	logger.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	donuts := API.Create(os.Getenv("TOKEN"), groupID, logger)

	t.Run("Balance", func(t *testing.T) {
		donates := donuts.Balance()
		assert.Equal(t, 0, donates.Error, donates.Message)
	})
}
