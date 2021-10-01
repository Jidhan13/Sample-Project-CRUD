package models

type Response struct {
	Status  int         `json:"id"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
