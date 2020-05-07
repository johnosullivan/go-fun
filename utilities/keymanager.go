package utilities

import (
  "io/ioutil"
  "os"
)

const SECRET_VOLUME_PATH = "SECRET_VOLUME_PATH"

var jwtSecret = ""

func GetJWTSecret() string {
  return jwtSecret
}

func InitKeys() {
  secret, err := ioutil.ReadFile(os.Getenv(SECRET_VOLUME_PATH) + "/secret")
  CheckError(err)
	jwtSecret = string(secret)
}
