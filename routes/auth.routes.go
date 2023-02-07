package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/marufboy/golang-rest-api-postgres/controllers"
	"github.com/marufboy/golang-rest-api-postgres/middleware"
)

type AuthRouteController struct {
	authController controllers.AuthController
}

func NewAuthRouteController(authController controllers.AuthController) AuthRouteController {
	return AuthRouteController{authController}
}

func (ac *AuthRouteController) AuthRoute(rg *gin.RouterGroup) {
	router := rg.Group("auth")

	router.POST("/register", ac.authController.SignUpUser)
	router.POST("/login", ac.authController.SignInUser)
	router.POST("/refresh", ac.authController.RefreshAccessToken)
	router.POST("/logout", middleware.DeserializeUser(), ac.authController.LogoutUser)
}
