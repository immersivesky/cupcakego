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

func TestCampaigns(t *testing.T) {
	godotenv.Load()

	var (
		logger     = logrus.New()
		groupID, _ = strconv.Atoi(os.Getenv("GROUP_ID"))
	)

	logger.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	donuts := API.Create(os.Getenv("TOKEN"), groupID, logger)

	t.Run("GetCampaigns", func(t *testing.T) {
		campaigns := donuts.GetCampaigns()
		assert.Equal(t, 0, campaigns.Error, campaigns.Message)
	})

	t.Run("GetActiveCampaign", func(t *testing.T) {
		campaigns := donuts.GetActiveCampaign()
		assert.Equal(t, 0, campaigns.Error, campaigns.Message)
	})

	t.Run("GetCampaignRewards", func(t *testing.T) {
		campaigns := donuts.GetCampaignRewards(9056)
		assert.Equal(t, 0, campaigns.Error, campaigns.Message)
	})

	t.Run("ChangeCampaign", func(t *testing.T) {
		campaigns := donuts.ChangeCampaign(9056, API.ChangeCampaign{
			Title: "Promotion",
		})
		assert.Equal(t, 0, campaigns.Error, campaigns.Message)
	})

	t.Run("ChangeCampaignReward", func(t *testing.T) {
		campaigns := donuts.ChangeCampaignReward(8577, 149)
		assert.Equal(t, 0, campaigns.Error, campaigns.Message)
	})
}
