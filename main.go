package main

import (
	"os"
	"net/http"

	"context"
	"time"
	"flag"
	"os/signal"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/johnosullivan/go-fun/routes"
	"github.com/johnosullivan/go-fun/utilities"
	//"github.com/johnosullivan/go-fun/db"
)


func main() {
	logfile, err := strconv.ParseBool(os.Getenv("LOGFILE"))
  utilities.CheckError(err)

	if logfile {
		file, err := os.OpenFile(os.Getenv("LOG_PATH"), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	  if err != nil {
	      log.Fatal(err)
	  }
	  defer file.Close()
		log.SetOutput(file)
	}
	log.SetFormatter(&log.JSONFormatter{})

	var wait time.Duration
  flag.DurationVar(&wait, "gto", time.Second * 15, "")
  flag.Parse()

	utilities.InitEnvironment()

	//database.InitDB()

	router := routes.GetRoutes()

	var port = os.Getenv("PORT")
	if len(port) == 0 {
		panic("env.PORT is required!")
	}

	srv := &http.Server{
      Addr:         ":" + port,
      WriteTimeout: time.Second * 15,
      ReadTimeout:  time.Second * 15,
      IdleTimeout:  time.Second * 60,
      Handler: utilities.LogRequest(router),
  }

  go func() {
      if err := srv.ListenAndServe(); err != nil {
          log.Println(err)
      }
  }()

  c := make(chan os.Signal, 1)
  signal.Notify(c, os.Interrupt)
  <-c

  ctx, cancel := context.WithTimeout(context.Background(), wait)
  defer cancel()
  srv.Shutdown(ctx)
  os.Exit(0)
}
