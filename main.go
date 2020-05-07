package main

import (
	"os"
	"log"
	"net/http"

	"github.com/johnosullivan/go-fun/routes"
	"github.com/johnosullivan/go-fun/utilities"
	//"github.com/johnosullivan/go-fun/db"
)

func main() {
	utilities.InitEnvironment()

	//database.InitDB()

	router := routes.GetRoutes()

	var port = os.Getenv("PORT")
	if len(port) == 0 {
		panic("Please pick a port to listen/serve")
	}

  if err := http.ListenAndServe(":" + port, router); err != nil {
      log.Fatal(err)
  }
}
