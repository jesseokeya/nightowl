package controllers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jesseokeya/nightowl/services"
	"github.com/jesseokeya/nightowl/structs"
)

// GetResources retrieves all the resources
func GetResources(c *gin.Context) {
	path := "." + os.Getenv("STORE")

	folder := services.Folder{
		Path: path,
	}

	files := folder.Interprete()

	c.JSON(http.StatusOK, structs.Success{Code: http.StatusOK, Data: files})
}
