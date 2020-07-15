package controllers

import (
	"strconv"
	"net/http"
	"github.com/gorilla/mux"
	m "github.com/andrushk/goapi/models"
	u "github.com/andrushk/goapi/utils"
)

// NewsByResort - возвращает список действующих акций, актуальных новостей
var NewsByResort = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	resortNews := []m.News{}

	for _, n := range AllNews {
        if key == strconv.Itoa(n.ResortID) {
			resortNews = append(resortNews, n)
        }
	}
	
	u.Respond(w, resortNews)
}

// AllNews - новости по всем курортам, должны храниться в БД
var AllNews []m.News = []m.News {
	m.News{ResortID: 1, Title: "FastRide news 1", Body: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."},
	m.News{ResortID: 1, Title: "FastRide news 2", Body: "Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat."},
	m.News{ResortID: 2, Title: "Терраски news 1", Body: "Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur."},
}