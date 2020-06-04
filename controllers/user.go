package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jesseokeya/nightowl/database"
	"github.com/jesseokeya/nightowl/models"
	"github.com/jesseokeya/nightowl/structs"
)

// GetUsers retrieves all the users
func GetUsers(c *gin.Context) {
	users := []models.User{}
	if err := database.DBCon.Find(&users).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.Error{Code: http.StatusNotFound, Error: "getting error while processing db request"})
		return
	}

	c.JSON(http.StatusOK, users)
}

// GetUser retrieves a particular user
func GetUser(c *gin.Context) {}

// CreateUser creates a new user
func CreateUser(c *gin.Context) {}

// UpdateUser updates a particular user
func UpdateUser(c *gin.Context) {}

// DeleteUser deletes a particular user
func DeleteUser(c *gin.Context) {}
