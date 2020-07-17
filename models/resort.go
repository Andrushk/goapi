package models

// Resort - один курорт
type Resort struct{
	ID int `json:"id"`
	Title string `json:"title"`
	State int  `json:"state"` // Удалено: -1, Неактивно: 0, Активно: 1
}