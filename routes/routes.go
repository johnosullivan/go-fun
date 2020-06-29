package routes

import (
  "os"
  "net/http"

	"github.com/johnosullivan/go-fun/middlewares"
	"github.com/johnosullivan/go-fun/controllers"
  "github.com/johnosullivan/go-fun/websockets"
  "github.com/gorilla/handlers"
)

func GetRoutes() *http.ServeMux {
  router := http.NewServeMux()

  hub := websockets.NewHub()
  go hub.Run()

  router.Handle("/ping", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(controllers.PingLink)))

  router.Handle("/authenticate", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(controllers.AuthenticateHandler)))

  router.Handle("/users", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(controllers.UsersHandler)))

  router.Handle("/authping", handlers.LoggingHandler(os.Stdout,  middlewares.AuthMiddleware(http.HandlerFunc(controllers.AuthPingHandler))))

  router.HandleFunc("/ws_test", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "ws_test.html")
  })

	router.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		    websockets.ServeWebSocket(hub, w, r)
	})

  return router
}
