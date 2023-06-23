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

func TestDonates(t *testing.T) {
	godotenv.Load()

	var (
		logger     = logrus.New()
		groupID, _ = strconv.Atoi(os.Getenv("GROUP_ID"))
	)

	logger.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	donuts := API.Create(os.Getenv("TOKEN"), groupID, logger)

	t.Run("GetDonates", func(t *testing.T) {
		donates := donuts.GetDonates(API.Count(2))
		assert.Equal(t, 0, donates.Error, donates.Message)
	})

	t.Run("GetLastDonates", func(t *testing.T) {
		donates := donuts.GetLastDonates(840182)
		assert.Equal(t, 0, donates.Error, donates.Message)
	})

	t.Run("ChangeStatus", func(t *testing.T) {
		donates := donuts.ChangeStatus(840185, API.PUBLIC)
		assert.Equal(t, 0, donates.Error, donates.Message)
	})

	t.Run("Answer", func(t *testing.T) {
		donates := donuts.Answer("840185", "thanks for!")
		assert.Equal(t, 0, donates.Error, donates.Message)
	})

	t.Run("ChangeReward", func(t *testing.T) {
		donates := donuts.ChangeReward("840185", API.SENDED)
		assert.Equal(t, 0, donates.Error, donates.Message)
	})
}
