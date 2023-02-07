package main

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/marufboy/golang-rest-api-postgres/config"
	"github.com/marufboy/golang-rest-api-postgres/controllers"
	"github.com/marufboy/golang-rest-api-postgres/routes"
	// "github.com/marufboy/golang-rest-api-postgres/routes"
)

var (
	server              *gin.Engine
	AuthController      controllers.AuthController
	AuthRouteController routes.AuthRouteController

	UserController      controllers.UserController
	UserRouteController routes.UserRouteController
)

func init() {
	setup, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	config.ConnectDB(&setup)
	//initialize controller and route
	AuthController = controllers.NewAuthController(config.DB)
	AuthRouteController = routes.NewAuthRouteController(AuthController)

	UserController = controllers.NewUserController(config.DB)
	UserRouteController = routes.NewUserRouteController(UserController)

	//initialize gin
	// gin.SetMode(gin.ReleaseMode)
	// server = gin.New() //empty engine
	server = gin.Default() //empty engine
}

func main() {
	setup, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	//cors setup
	corsSetup := cors.DefaultConfig()
	corsSetup.AllowAllOrigins = true
	corsSetup.AllowCredentials = true

	server.Use(cors.New(corsSetup))
	//setup logger
	// server.Use(middleware.DefaultStructuredLogger()) //adds our new middleware logger
	// server.Use(gin.Recovery())

	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to Golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	AuthRouteController.AuthRoute(router)
	UserRouteController.UserRoute(router)

	log.Fatal(server.Run(":" + setup.ServerPort))
}
