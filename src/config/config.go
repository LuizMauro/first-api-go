package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	StringConnectionDB = ""
	Port = 0
)

func InitConfig() {
	var err error

	if err  = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv("API_PORT")) 
	if err != nil {
		Port = 9000
	}


	StringConnectionDB = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", 
		os.Getenv("USER_DB"), 
		os.Getenv("PASSWORD_DB"), 
		os.Getenv("NAME_DB"),
	)
}