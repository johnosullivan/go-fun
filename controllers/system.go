package controllers

import (
  "io"
	"net/http"
  "time"
  "encoding/json"

  //"fmt"
  //"github.com/johnosullivan/go-fun/websockets"
  //"github.com/johnosullivan/go-fun/db"
)

type SystemStatus struct {
  Date    time.Time
  Status  bool
}

func PingLink(w http.ResponseWriter, r *http.Request) {
    sysStatus := SystemStatus{time.Now(), true}
    js, err := json.Marshal(sysStatus)
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }

    /*hub := websockets.GetHub()
    hub.Broadcast <- []byte("Hello World!")
    fmt.Println(hub)*/

    w.Header().Set("Content-Type", "application/json")
    w.Write(js)
}

func AuthPingHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Content-Type", "application/json")
    io.WriteString(w, `{"status":"ok"}`)
}
