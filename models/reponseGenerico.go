package models

type ResponseGenerico struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}
