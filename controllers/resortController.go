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
	m.Resort{ID: 1, Title: "FastRide"},
	m.Resort{ID: 2, Title: "Терраски"},
}