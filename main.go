package main

import (
	"os"

	"github.com/ayesparshh/db"

	"github.com/ayesparshh/routes"
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
	// DB.Seed()

	router := gin.Default()

	//sales
	routes.CustomersRoutes(router, DB)

	//sales //accountant
	routes.BillingRoutes(router, DB)

	//accountant //hr
	routes.PayrollRoutes(router, DB)

	//admintrator
	routes.UserRoutes(router, DB)

	go func() {
		if err := router.Run(":8080"); err != nil {
			log.Fatal().Err(err).Msg("Failed to start server")
		}
	}()

	<-wait
}
