package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/marufboy/golang-rest-api-postgres/controllers"
	"github.com/marufboy/golang-rest-api-postgres/middleware"
)

type UserRouteController struct {
	userController controllers.UserController
}

func NewUserRouteController(userController controllers.UserController) UserRouteController {
	return UserRouteController{userController}
}

func (uc *UserRouteController) UserRoute(rg *gin.RouterGroup) {
	router := rg.Group("users")

	router.GET("/me", middleware.DeserializeUser(), uc.userController.GetUser)
	router.PUT("/update", middleware.DeserializeUser(), uc.userController.UpdateUser)
}
