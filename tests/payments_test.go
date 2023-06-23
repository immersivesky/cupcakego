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

func TestPayments(t *testing.T) {
	godotenv.Load()

	var (
		logger     = logrus.New()
		groupID, _ = strconv.Atoi(os.Getenv("GROUP_ID"))
	)

	logger.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	donuts := API.Create(os.Getenv("TOKEN"), groupID, logger)

	t.Run("GetPayments", func(t *testing.T) {
		donates := donuts.GetPayments([]int{151312, 161722})
		assert.Equal(t, 0, donates.Error, donates.Message)
	})

	t.Run("CreatePayments", func(t *testing.T) {
		donates := donuts.CreatePayments(0.77, API.QIWI, "79780175583")
		assert.Equal(t, 0, donates.Error, donates.Message)
	})
}
