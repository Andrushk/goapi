package controllers

import (
	"net/http"

	m "github.com/andrushk/goapi/models"
	u "github.com/andrushk/goapi/utils"
)

// SkipassesAll - возвращает список "живых" скипасов
var SkipassesAll = func(w http.ResponseWriter, r *http.Request) {
	result := []m.Skipass{}

	for _, n := range SkipassesAllData {
		if n.State >= 0 {
			result = append(result, n)
		}
	}

	u.Respond(w, result)
}

// SkipassesAllData - список всех скипасов, должен хранится в БД
var SkipassesAllData []m.Skipass = []m.Skipass{
	m.Skipass{ID: 1, Description: "", Resort: resortByID(1), State: 1},
	m.Skipass{ID: 2, Description: "wife", Resort: resortByID(1), State: 1},
}
