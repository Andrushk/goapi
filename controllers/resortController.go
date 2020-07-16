package controllers

import (
	"net/http"
	m "github.com/andrushk/goapi/models"
	u "github.com/andrushk/goapi/utils"
)

// ResortsAll - возвращает полный список горнолыжек
var ResortsAll = func(w http.ResponseWriter, r *http.Request) {
	u.Respond(w, ResortsAllData)
}

// ResortsAllData - список всех курортов, должен хранится в БД
var ResortsAllData []m.Resort = []m.Resort{
	m.Resort{ID: 1, Title: "FastRide", Active: 1},
	m.Resort{ID: 2, Title: "Терраски", Active: 1},
	m.Resort{ID: 3, Title: "Хабарское", Active: 1},
	m.Resort{ID: 4, Title: "Новинки", Active: 1},
	m.Resort{ID: 5, Title: "Свияжские холмы", Active: 0},
	m.Resort{ID: 6, Title: "Пужалова гора", Active: 0},
}