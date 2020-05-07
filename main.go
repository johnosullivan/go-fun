package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/johnosullivan/go-fun/middlewares"
	"github.com/johnosullivan/go-fun/controllers"
	"github.com/johnosullivan/go-fun/utilities"
)

func main() {

	utilities.InitKeys()

	router := mux.NewRouter().StrictSlash(false)

	router.HandleFunc("/ping", controllers.PingLink)

	router.HandleFunc("/token", controllers.TokenHandler)

	router.Handle("/", middlewares.AuthMiddleware(http.HandlerFunc(controllers.ExampleHandler)))

	// Start a basic HTTP server
  if err := http.ListenAndServe(":8080", router); err != nil {
      log.Fatal(err)
  }
}
