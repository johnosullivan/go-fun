package controllers

import (
  "io"
  "time"
	"net/http"

  "github.com/dgrijalva/jwt-go"

  "github.com/johnosullivan/go-fun/utilities"
)

func TokenHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Content-Type", "application/json")
    r.ParseForm()

    var APP_KEY = utilities.GetJWTSecret()

    username := r.Form.Get("username")
    password := r.Form.Get("password")

    if username != "myusername" && password != "mypassword" {
        w.WriteHeader(http.StatusUnauthorized)
        io.WriteString(w, `{"error":"invalid_credentials"}`)
        return
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user": username,
        "exp":  time.Now().Add(time.Hour * time.Duration(1)).Unix(),
        "iat":  time.Now().Unix(),
    })

    tokenString, err := token.SignedString([]byte(APP_KEY))
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        io.WriteString(w, `{"error":"token_generation_failed"}`)
        return
    }

    io.WriteString(w, `{"token":"`+tokenString+`"}`)
    return
}
