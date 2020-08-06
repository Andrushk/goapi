package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	m "github.com/andrushk/goapi/models"
	u "github.com/andrushk/goapi/utils"
)

// SkipassesAll - возвращает список "живых" скипасов пользователя
var SkipassesAll = func(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(int) //Получение идентификатора пользователя, отправившего запрос

	result := []m.Skipass{}

	for _, n := range SkipassesAllData {
		if n.UserID == user && n.State >= 0 {
			result = append(result, n)
		}
	}

	u.Respond(w, result)
}

// SkipassPut - создать или изменить скипас
var SkipassPut = func(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(int) //Получение идентификатора пользователя, отправившего запрос

	var skipass m.Skipass
	_ = json.NewDecoder(r.Body).Decode(&skipass)

		// нет ID - это инсерт
		if skipass.ID == 0 {
		// сэрвер должен игнорить некоторые данные, например, сам должен
		// проставить UserID
		skipass.ID = getID()
		skipass.Resort = resortByID(skipass.Resort.ID)
		skipass.State = 1
		skipass.UserID = user

		SkipassesAllData = append(SkipassesAllData, skipass)

		u.Respond(w, skipass)
	} else {
		fmt.Println("Skipass ID: " + strconv.Itoa(skipass.ID))
		u.Respond(w, "update не реализован")
	}

}

// SkipassDelete - удаляет скипас
var SkipassDelete = func(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(int) //Получение идентификатора пользователя, отправившего запрос
	vars := mux.Vars(r)
	key := vars["id"]

	newAll := []m.Skipass{}
	deleteOk := false

	for _, n := range SkipassesAllData {
		if n.UserID == user && key == strconv.Itoa(n.ID) {
			deleteOk = true
			continue
		}
		newAll = append(newAll, n)
	}
	SkipassesAllData = newAll

	u.Respond(w, deleteOk)
}

func getID() int {
	newID := 0
	for _, n := range SkipassesAllData {
		if n.ID > newID {
			newID = n.ID
		}
	}
	return newID + 1
}

// SkipassesAllData - список всех скипасов, должен хранится в БД
var SkipassesAllData []m.Skipass = []m.Skipass{
	m.Skipass{ID: 1, Description: "", Resort: resortByID(1), State: 1, UserID: 42},
	m.Skipass{ID: 2, Description: "wife", Resort: resortByID(1), State: 1, UserID: 42},
	m.Skipass{ID: 3, Description: "dog", Resort: resortByID(1), State: 1, UserID: 33},
}
