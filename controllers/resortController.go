package controllers

import (
	"net/http"

	m "github.com/andrushk/goapi/models"
	u "github.com/andrushk/goapi/utils"
)

// ResortsAll - возвращает полный список "живых" горнолыжек
var ResortsAll = func(w http.ResponseWriter, r *http.Request) {
	result := []m.Resort{}

	for _, n := range ResortsAllData {
		if n.State >= 0 {
			result = append(result, n)
		}
	}

	u.Respond(w, result)
}

// ResortsAllData - список всех курортов, должен хранится в БД
var ResortsAllData []m.Resort = []m.Resort{
	m.Resort{ID: 1, Title: "FastRide", State: 1},
	m.Resort{ID: 2, Title: "Терраски", State: 1},
	m.Resort{ID: 3, Title: "Хабарское", State: 1},
	m.Resort{ID: 4, Title: "Новинки", State: 1},
	m.Resort{ID: 5, Title: "Свияжские холмы", State: 0},
	m.Resort{ID: 6, Title: "Пужалова гора", State: 0},
	m.Resort{ID: 7, Title: "Вестендорф", State: -1},
	m.Resort{ID: 8, Title: "Ваграйн", State: -1},
	m.Resort{ID: 9, Title: "Фюген", State: -1},
	m.Resort{ID: 10, Title: "Альпбах", State: -1},
	m.Resort{ID: 11, Title: "Оберстдорф", State: -1},
}
