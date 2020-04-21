package main

import (
	"net/http"
  "encoding/json"
	"github.com/gorilla/mux"
)

type User struct {
     Name  string  `json:"name"`
     Email string  `json:"email"`
}

func pingLink(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user := User{Name: "John Doe Test", Email: "johndoe@gmail.com"}
	json.NewEncoder(w).Encode(user)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/ping", pingLink)

	http.ListenAndServe(":8080", router)
}
