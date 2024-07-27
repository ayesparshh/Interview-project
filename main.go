package main

import (
	"os"
	"net/http"

	"interview-project/db"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {

	if err := godotenv.Load(); err != nil {
		log.Warn().Err(err).Msg("No .env file found")
	}

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

func main() {

	wait := make(chan bool)

	DatabaseUrl := os.Getenv("DATABASE_URL")

	DB := db.New(DatabaseUrl)
	DB.Seed()

}
