package API

import (
	"fmt"
	"github.com/immersivesky/cupcakego/responses"
)

// Balance - бич я поймал свой balance
func (cupcake *Options) Balance() (balance responses.Balance) {
	var params = map[string]interface{}{
		"group": cupcake.ID,
		"token": cupcake.Token,
		"v":     cupcake.Version,
	}

	if err := cupcake.Call("balance", params, &balance); err != nil {
		fmt.Println(err)
	}

	return balance
}
