package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/marufboy/golang-rest-api-postgres/models"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

func NewUserController(DB *gorm.DB) UserController {
	return UserController{DB}
}

// TODO: 1. Get user
func (uc *UserController) GetUser(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)

	response := &models.UserResponse{
		ID:        currentUser.ID,
		Name:      currentUser.Name,
		Email:     currentUser.Email,
		Photo:     currentUser.Photo,
		CreatedAt: currentUser.CreatedAt,
		UpdatedAt: currentUser.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"user": response}})
}

// TODO: 2. Update user data
func (uc *UserController) UpdateUser(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)

	var payload *models.UpdateUser
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	//generate new data user
	now := time.Now()
	postUpdateUser := models.User{
		Name:      payload.Name,
		Photo:     payload.Photo,
		CreatedAt: currentUser.CreatedAt,
		UpdatedAt: now,
	}

	uc.DB.Model(&currentUser).Updates(postUpdateUser)
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": currentUser})
}
