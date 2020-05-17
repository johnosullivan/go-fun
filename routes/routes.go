package routes

import (
  "net/http"
	"github.com/gorilla/mux"

	"github.com/johnosullivan/go-fun/middlewares"
	"github.com/johnosullivan/go-fun/controllers"
)

func GetRoutes() *mux.Router {
  router := mux.NewRouter().StrictSlash(false)

	router.HandleFunc("/ping", controllers.PingLink).Methods("GET")

	router.HandleFunc("/authenticate", controllers.AuthenticateHandler).Methods("POST")

	router.Handle("/authping", middlewares.AuthMiddleware(http.HandlerFunc(controllers.AuthPingHandler)))

  return router
}
