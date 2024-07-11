package dto

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
	Error   string `json:"error"`
}
