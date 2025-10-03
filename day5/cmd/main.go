package main

import (
	"fmt"

	"cdac.com/day5/controllers"
	"cdac.com/day5/db"
	"cdac.com/day5/repository"
	"cdac.com/day5/routes"
	"cdac.com/day5/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect DB
	db.Connect()

	// Build components
	userRepository := repository.NewUserRepository(db.DB)
	userService := service.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	// Gin setup
	//1. create a new router
	r := gin.Default()

	//2. Register routes
	routes.RegisterRoutes(r, userController)

	//3. Start server
	r.Run(fmt.Sprintf(":%d", 3000))
}
