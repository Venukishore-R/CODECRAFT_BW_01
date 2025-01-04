package main

import (
	"github.com/Venukishore-R/CODECRAFT_BW_01/endpoints"
	"github.com/Venukishore-R/CODECRAFT_BW_01/internal/models"
	"github.com/Venukishore-R/CODECRAFT_BW_01/internal/services"
	"github.com/Venukishore-R/CODECRAFT_BW_01/pkg/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	users := []*models.User{}
	userModel := models.NewUserModel(users)
	services := services.NewUserService(userModel)
	handlers := handlers.NewUserHandler(services)
	endpoints.Endpoints(r, handlers)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})

	r.Run(":5001")
}
