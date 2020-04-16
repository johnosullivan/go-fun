package main

import (
	"net/http"
  "encoding/json"
	"github.com/gorilla/mux"
)

type User struct {
     Id    int     `json:"id"`
     Name  string  `json:"name"`
     Email string  `json:"email"`
     Phone string  `json:"phone"`
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user := User{Id: 1, Name: "John Doe", Email: "johndoe@gmail.com", Phone: "000099999"}
	json.NewEncoder(w).Encode(user)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/ping", homeLink)

	http.ListenAndServe(":8080", router)
}
