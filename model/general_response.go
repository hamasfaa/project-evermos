package model

type GeneralResponse struct {
	Status  bool        `json:"status"`
	Errors  interface{} `json:"errors"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
