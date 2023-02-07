package main

import (
	"fmt"
	"log"

	"github.com/marufboy/golang-rest-api-postgres/config"
	"github.com/marufboy/golang-rest-api-postgres/models"
)

func init() {
	setup, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	config.ConnectDB(&setup)
}

func main() {
	config.DB.AutoMigrate(&models.User{})
	fmt.Println("👍 Migration complete")
}
