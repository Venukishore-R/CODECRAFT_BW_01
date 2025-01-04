package endpoints

import (
	"github.com/Venukishore-R/CODECRAFT_BW_01/pkg/handlers"
	"github.com/gin-gonic/gin"
)

func Endpoints(r *gin.Engine, handlers *handlers.UserHandler) {
	r.POST("/create", handlers.CreateUser)
	r.GET("/users", handlers.GetAll)
	r.GET("/user/:email", handlers.GetOne)
	r.PUT("/user/:email", handlers.Update)
	r.DELETE("/user/:email", handlers.Delete)
}
