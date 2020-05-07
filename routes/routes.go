package routes

import (
  "net/http"
	"github.com/gorilla/mux"

	"github.com/johnosullivan/go-fun/middlewares"
	"github.com/johnosullivan/go-fun/controllers"
)

func GetRoutes() *mux.Router {
  router := mux.NewRouter().StrictSlash(false)

	router.HandleFunc("/ping", controllers.PingLink)

	router.HandleFunc("/token", controllers.TokenHandler)

	router.Handle("/", middlewares.AuthMiddleware(http.HandlerFunc(controllers.ExampleHandler)))

  return router
}
