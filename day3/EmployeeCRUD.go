package main

import (
	"fmt"
	"net/http"

	"cdac.com/myapp/Models"
	"github.com/gin-gonic/gin"
)

func main() {
	var employees []Models.Employee
	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()

	// Define a simple GET endpoint
	r.GET("/ping", func(c *gin.Context) {
		// Return JSON response
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/register", func(c *gin.Context) {
		var emp Models.Employee = c.Request.Context().Value("employee").(Models.Employee)
		fmt.Println(c.Request)
		employees = append(employees, emp)
		c.JSON(http.StatusCreated, gin.H{
			"message": "Employee registered successfully!",
		})
	})

	r.GET("/employees", func(c *gin.Context) {
		c.JSON(http.StatusOK, employees)
	})

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	r.Run(":8081")
}
