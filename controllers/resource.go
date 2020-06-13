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

// DeleteResources deletes a particluar resource by name
func DeleteResources(c *gin.Context) {
	path := "." + os.Getenv("STORE")
	fileName := c.Param("filename")

	folder := services.Folder{
		Path: path,
	}

	files, err := folder.DeleteFile(fileName)
	if err != nil {
		c.JSON(http.StatusBadRequest, structs.Error{Code: http.StatusNotFound, Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, structs.Success{Code: http.StatusOK, Data: files})
}
