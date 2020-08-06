package main

import (
	"github.com/gorilla/mux"
	"fmt"
	"net/http"
	"github.com/andrushk/goapi/app"
	"github.com/andrushk/goapi/controllers"
)

func main() {

	router := mux.NewRouter()
	router.Use(app.Authentication) // добавляем middleware аутентификации

	router.HandleFunc("/api/resorts/all", controllers.ResortsAll).Methods("GET")
	router.HandleFunc("/api/news/resort", controllers.NewsByResort).Queries("id", "{id}").Methods("GET")
	router.HandleFunc("/api/skipasses/all", controllers.SkipassesAll).Methods("GET")
	router.HandleFunc("/api/skipasses", controllers.SkipassPut).Methods("PUT")
	router.HandleFunc("/api/skipasses", controllers.SkipassDelete).Queries("id", "{id}").Methods("DELETE")
	//TODO GET: \price?resort_id=GUID - лист услуг конкретной горнолыжки
	//TODO POST: \acquirer_notification - статусы от Тинька

	port := "8000"
	// port := os.Getenv("PORT") //Получить порт из файла .env; мы не указали порт, поэтому при локальном тестировании должна возвращаться пустая строка
	// if port == "" {
	// 	port = "8000" //localhost
	// }

	fmt.Println(port)

	err := http.ListenAndServe(":" + port, router)

	if err != nil {
		fmt.Print(err)
	}
}