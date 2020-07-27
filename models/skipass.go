package models

// Skipass - один скипас
type Skipass struct{
	ID int `json:"id"`
	Description string `json:"description"`
	Resort Resort `json:"resort"`
	State int `json:"state"` // Удалено: -1, Неактивно: 0, Активно: 1
}