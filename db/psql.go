package database

import (
  "database/sql"
	"fmt"

	_ "github.com/lib/pq"

  "github.com/johnosullivan/go-fun/utilities"

  "golang.org/x/crypto/bcrypt"
  //"log"
)

const (
    dbhost = "DBHOST"
    dbport = "DBPORT"
    dbuser = "DBUSER"
    dbpass = "DBPASS"
    dbname = "DBNAME"
)

var db *sql.DB

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func PingDB() {
  var err error
  err = db.Ping()
  if err != nil {
      panic(err)
  }
  fmt.Println("Ping DB");
}

func InitDB() {
    config := utilities.GetDBConfig()

    fmt.Println(config);

    var err error

    psqlInfo := fmt.Sprintf("host=%s port=%s user=%s " +
        "password=%s dbname=%s sslmode=disable",
        config[dbhost], config[dbport],
        config[dbuser], config[dbpass], config[dbname])

    db, err = sql.Open("postgres", psqlInfo)
    if err != nil {
        panic(err)
    }

    PingDB()

    /*
    password := "el2eRO01sVVnOg8EmRMRTEAl"
    hash, _ := HashPassword(password)
    fmt.Println("Password:", password)
    fmt.Println("Hash:    ", hash)
    match := CheckPasswordHash(password, hash)
    fmt.Println("Match:   ", match)
    */

    /*
    fmt.Println("# Querying")
    rows, err := db.Query("SELECT * FROM users;")
    utilities.CheckError(err)

    for rows.Next() {
        var id int
        var name string
        err = rows.Scan(&id, &name)
        utilities.CheckError(err)
        fmt.Println("id | name")
        fmt.Printf("%3v | %8v\n", id, name)
    }
    */
}
