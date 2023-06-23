package responses

type Balance struct {
	Info
	Balance string `json:"balance,omitempty"`
}
