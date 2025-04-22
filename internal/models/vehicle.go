package models

type Vehicle struct {
	Id    string `json:"id"`
	Brand string `json:"brand"`
	Model string `json:"model"`
	Year  int    `json:"year"`
}
