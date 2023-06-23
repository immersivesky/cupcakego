package responses

type Info struct {
	Success bool   `json:"success,omitempty"`
	Error   int    `json:"error,omitempty"`
	Message string `json:"msg,omitempty"`
}
