package main

import (
	"log"
	"statistics/config"
	"statistics/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	debug := config.Debug()
	if debug == true {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	} else {
		gin.SetMode("release")
	}

	database.Init()
	r := gin.New()
	r.Run(":4000")
}
