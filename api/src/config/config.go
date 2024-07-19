package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// DatabaseStringConnection is the string connection to MySQL
	DatabaseStringConnection = ""

	// Port where the API is working
	Port = 0

	// SecretKey is the key that will be used to sign the token
	SecretKey []byte
)

// Load initializes the environment variables
func Load() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	} 

	Port, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Port = 9000
	}

	DatabaseStringConnection = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
	 	os.Getenv("DB_USER"),
	  	os.Getenv("DB_PASSWORD"),
	   	os.Getenv("DB_NAME"),
	)

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}