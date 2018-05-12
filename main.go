package main

import (
	"github.com/HackIllinois/api-auth/controller"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/HackIllinois/api-commons/middleware"
)

func main() {
	router := mux.NewRouter()
	controller.SetupController(router.PathPrefix("/auth"))

	router.Use(middleware.ErrorMiddleware)
	router.Use(middleware.ContentTypeMiddleware)
	log.Fatal(http.ListenAndServe(":8002", router))
}
