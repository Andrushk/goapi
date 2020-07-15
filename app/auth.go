package app

import (
	"net/http"
)

// Authentication - когда-нибудь тут будет аутентификация
var Authentication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//TODO
		next.ServeHTTP(w, r) //передать управление следующему обработчику!
	})
}