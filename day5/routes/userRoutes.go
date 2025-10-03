package routes

import (
	"cdac.com/day5/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, userController *controllers.UserController) {
	api := r.Group("/api")
	userController.RegisterRoutes(api)
}
