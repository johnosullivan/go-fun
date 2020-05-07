package database

import (
  "database/sql"
	"fmt"

	_ "github.com/lib/pq"

  "github.com/johnosullivan/go-fun/utilities"
)

const (
    dbhost = "DBHOST"
    dbport = "DBPORT"
    dbuser = "DBUSER"
    dbpass = "DBPASS"
    dbname = "DBNAME"
)

var db *sql.DB

func InitDB() {
    config := utilities.DBConfig()

    var err error

    psqlInfo := fmt.Sprintf("host=%s port=%s user=%s " +
        "password=%s dbname=%s sslmode=disable",
        config[dbhost], config[dbport],
        config[dbuser], config[dbpass], config[dbname])

    db, err = sql.Open("postgres", psqlInfo)
    if err != nil {
        panic(err)
    }

    err = db.Ping()
    if err != nil {
        panic(err)
    }
}
