package main

import (
	"log"
	"yesiamdonation/server"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()
	server.RegisterRoute(r)
}
