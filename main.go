package main

import (
	"chat/handlers"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func registerHandlers() *httprouter.Router {
	router := httprouter.New()

	router.GET("/ws", handlers.WS)

	return router
}

func main() {
	r := registerHandlers()
	http.ListenAndServe(":8080", r)
}
