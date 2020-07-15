package models

// News - одна новость
type News struct {
	ResortID int `json:"resort_id"`
	Title string `json:"title"`
	Body string `json:"body"`
}