package handlers

import (
	"github.com/Venukishore-R/CODECRAFT_BW_01/internal/models"
	"github.com/Venukishore-R/CODECRAFT_BW_01/internal/services"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *services.User
}

func NewUserHandler(service *services.User) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) CreateUser(ctx *gin.Context) {
	var user *models.User

	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(404, gin.H{"error": err.Error()})
		return
	}

	user.Id = h.service.CreateUUID()

	result, err := h.service.CreateUserService(user)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, result)
}

func (h *UserHandler) GetAll(ctx *gin.Context) {
	users, err := h.service.GetAllUsersService()
	if err != nil {
		ctx.JSON(404, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, users)
}

func (h *UserHandler) GetOne(ctx *gin.Context) {
	email := ctx.Param("email")
	user, err := h.service.GetOneUserService(email)
	if err != nil {
		ctx.JSON(404, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, user)
}

func (h *UserHandler) Update(ctx *gin.Context) {
	email := ctx.Param("email")

	var user *models.User

	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(404, gin.H{"error": err.Error()})
		return
	}

	updatedUser, err := h.service.UpdateUserService(user, email)
	if err != nil {
		ctx.JSON(404, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, updatedUser)
}

func (h *UserHandler) Delete(ctx *gin.Context) {
	email := ctx.Param("email")

	err := h.service.DeleteUserService(email)
	if err != nil {
		ctx.JSON(404, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(204, gin.H{"message": "user deleted"})
}
