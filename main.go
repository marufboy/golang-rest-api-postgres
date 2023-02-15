package main

import (
	"net/http"

	_ "github.com/marufboy/golang-rest-api-postgres/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/marufboy/golang-rest-api-postgres/config"
	"github.com/marufboy/golang-rest-api-postgres/controllers"
	"github.com/marufboy/golang-rest-api-postgres/routes"
	"github.com/rs/zerolog/log"
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
		log.Fatal().Err(err).Msg("🚀 Could not load environment variables")
	}

	config.ConnectDB(&setup)
	//initialize controller and route
	AuthController = controllers.NewAuthController(config.DB)
	AuthRouteController = routes.NewAuthRouteController(AuthController)

	UserController = controllers.NewUserController(config.DB)
	UserRouteController = routes.NewUserRouteController(UserController)

	//initialize gin
	gin.SetMode(gin.ReleaseMode)
	// server = gin.New() //empty engine
	server = gin.Default() //empty engine
}

// @title 		Gin User Auth Service
// @version 	1.0
// @description A User Auth Service API in Go using Gin framework

// @contact.name   Muhammad Afif Ma'ruf
// @contact.url    http://www.swagger.io/support
// @contact.email  afif23ixa@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host 		localhost:8000
// @BasePath 	/api
func main() {
	setup, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("🚀 Could not load environment variables")
	}

	//cors setup
	corsSetup := cors.DefaultConfig()
	corsSetup.AllowAllOrigins = true
	corsSetup.AllowCredentials = true

	server.Use(cors.New(corsSetup))
	//setup logger
	// server.Use(middleware.DefaultStructuredLogger()) //adds our new middleware logger
	// server.Use(gin.Recovery())

	//config swagger to router
	server.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to Golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	AuthRouteController.AuthRoute(router)
	UserRouteController.UserRoute(router)

	err = server.Run(":" + setup.ServerPort)
	log.Fatal().Err(err).Msg("🚀 Could not run the server")
}
