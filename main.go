package main

import (
	"os"
	//"log"
	"net/http"

	"context"
	"time"
	"flag"
	"os/signal"

	log "github.com/sirupsen/logrus"

	"github.com/johnosullivan/go-fun/routes"
	"github.com/johnosullivan/go-fun/utilities"
	//"github.com/johnosullivan/go-fun/db"
)

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(log.Fields{
	       "remote_addr": r.RemoteAddr,
	       "method": r.Method,
				 "url": r.URL,
				 "userAgent": r.Header.Get("User-Agent")}).Info("request")

		handler.ServeHTTP(w, r)
	})
}

func main() {
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
  if err != nil {
      log.Fatal(err)
  }
  defer file.Close()
	log.SetOutput(file)
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
      Handler: logRequest(router),
  }

  go func() {
      if err := srv.ListenAndServe(); err != nil {
          log.Println(err)
      }
  }()

  c := make(chan os.Signal, 1)
  // SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
  signal.Notify(c, os.Interrupt)
  // Block until we receive our signal.
  <-c

  ctx, cancel := context.WithTimeout(context.Background(), wait)
  defer cancel()
  srv.Shutdown(ctx)
  os.Exit(0)
}
