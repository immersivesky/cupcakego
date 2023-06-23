package callback

import (
	"github.com/immersivesky/cupcakego/API"
)

type Options struct {
	Application *API.Options
	Scenes      Scenes
}

func Create(properties ...any) *Options {
	options := Options{}

	for _, property := range properties {
		switch property.(type) {
		case *API.Options:
			options.Application = property.(*API.Options)
		}
	}

	return &options
}
