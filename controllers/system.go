package controllers

import (
  "io"
	"net/http"
)

func PingLink(w http.ResponseWriter, r *http.Request) {
  	w.Header().Set("Content-Type", "application/json")
  	//res := utilities.Response{Key: "", Secret: ""}
  	//json.NewEncoder(w).Encode(res)
    io.WriteString(w, `{"status":"ok"}`)
}

func AuthPingHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Content-Type", "application/json")
    io.WriteString(w, `{"status":"ok"}`)
}
