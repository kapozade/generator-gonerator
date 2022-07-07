package models

type ApplicationErrors struct {
	Errors []ApplicationError `json:"errors"`
}

type ApplicationError struct {
	Message string `json:"message"`
	Data    string `json:"data"`
}
