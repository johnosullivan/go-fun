package controllers

import (
  "io"
	"net/http"
  "time"
  "encoding/json"
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
    w.Header().Set("Content-Type", "application/json")
    w.Write(js)
}

func AuthPingHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Content-Type", "application/json")
    io.WriteString(w, `{"status":"ok"}`)
}
