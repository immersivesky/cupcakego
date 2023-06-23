package callback

import (
	"github.com/immersivesky/cupcakego/API"
	"github.com/immersivesky/cupcakego/update"
)

type Scenes struct {
	Donate  Donate
	Payment Payment
}

type Donate func(*API.Options, update.Donate)
type Payment func(*API.Options, update.Payment)

func (session *Options) Donate(donate Donate) {
	session.Scenes.Donate = donate
}

func (session *Options) Payment(payment Payment) {
	session.Scenes.Payment = payment
}
