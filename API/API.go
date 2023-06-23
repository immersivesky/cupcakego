package API

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type Options struct {
	ID      int    `json:"id"`
	Token   string `json:"token"`
	Version int    `json:"version"`
	Client  *http.Client
	Logger  *log.Logger
}

func Create(properties ...any) *Options {
	var (
		groupID bool
		options = Options{
			Client:  http.DefaultClient,
			Version: 1,
		}
	)

	for _, property := range properties {
		switch property.(type) {
		case int:
			value := property.(int)
			if groupID {
				options.Version = value
			} else {
				options.ID = value
			}
		case string:
			options.Token = property.(string)
		case *log.Logger:
			options.Logger = property.(*log.Logger)
		}
	}

	if options.Token == "" {
		options.Logger.Error("Token is required")
	}
	if options.ID == 0 {
		options.Logger.Error("ID is required")
	}

	return &options
}

func (cupcake Options) Call(method string, params map[string]interface{}, data any) error {
	body, err := json.Marshal(params)
	if err != nil {
		return err
	}

	logger := cupcake.Logger.WithFields(log.Fields{
		"method": method,
		"params": string(body),
	})

	request, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://api.keksik.io/"+method, bytes.NewBuffer(body))
	if err != nil {
		logger.Error("HTTP NewRequest error is  ", err)
		return err
	}

	response, err := cupcake.Client.Do(request)
	if err != nil {
		logger.Error("HTTP Client Do error is  ", err)
		return err
	}

	defer response.Body.Close()

	if err = json.NewDecoder(response.Body).Decode(&data); err != nil {
		logger.Error("HTTP DecodeJSON error is  ", err)
		return err
	}

	if cupcake.Logger != nil {
		response, _ := json.Marshal(&data)
		logger.Info("Response from API is ", string(response))
	}

	return nil
}
