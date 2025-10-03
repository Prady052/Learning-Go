package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"cdac.com/day5/models"
	"cdac.com/day5/service"

	// go get github.com/gin-gonic/gin
	"github.com/gin-gonic/gin"
)

type UserController struct {
	service service.UserService
}

func NewUserController(s service.UserService) *UserController {
	return &UserController{service: s}
}

func (uc *UserController) RegisterRoutes(rg *gin.RouterGroup) {
	rg.POST("/users/register", uc.CreateUser)
	rg.GET("/users", uc.ListUsers)
	rg.GET("/users/:id", uc.GetUser)
	rg.PUT("/users/:id", uc.UpdateUser)
	rg.DELETE("/users/:id", uc.DeleteUser)
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var input models.User
	err := c.ShouldBindJSON(&input)
	fmt.Println(input)
	if err != nil {
		fmt.Print("------------invalid create payload", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}
	err = uc.service.CreateUser(&input)
	if err != nil {
		fmt.Print("-----------------create user failed", err)
		if err == service.ErrInvalidData {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "create failed"})
		return
	}
	c.JSON(http.StatusCreated, input)
}

func (uc *UserController) ListUsers(c *gin.Context) {

	users, err := uc.service.GetAllUsers()
	if err != nil {
		fmt.Print("-----------------------list users failed", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "list failed"})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (uc *UserController) GetUser(c *gin.Context) {
	idParam := c.Param("id")
	id64, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	user, err := uc.service.GetUser(uint(id64))
	if err != nil {
		if err == service.ErrNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
			return
		}
		fmt.Print("--------------get user failed", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "get failed"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (uc *UserController) UpdateUser(c *gin.Context) {
	idParam := c.Param("id")
	id64, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	fmt.Println("---------------------------------", input, "-----------------", id64)
	updated, err := uc.service.UpdateUser(uint(id64), &input)
	if err != nil {
		if err == service.ErrNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
			return
		}
		if err == service.ErrInvalidData {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid data"})
			return
		}
		fmt.Print("------------update user failed", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "update failed"})
		return
	}
	c.JSON(http.StatusOK, updated)
}

func (uc *UserController) DeleteUser(c *gin.Context) {
	idParam := c.Param("id")
	id64, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	if err := uc.service.DeleteUser(uint(id64)); err != nil {
		if err == service.ErrNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
			return
		}
		fmt.Print("delete user failed", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "delete failed"})
		return
	}
	c.Status(http.StatusNoContent)
}
