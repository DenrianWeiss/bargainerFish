package env

import (
	"github.com/DenrianWeiss/barginerFish/utils"
	"log"
	"os"
)

var token string

func init() {
	tokenE, err := os.LookupEnv("ENV_TOKEN")
	if !err {
		token = utils.GenerateUUID()
		log.Println("Use random token: " + token)
	} else {
		token = tokenE
	}
}

func GetToken() string {
	return token
}
