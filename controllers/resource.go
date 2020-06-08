package controllers

import (
	"log"
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
	q := c.Request.URL.Query()
	path := "." + os.Getenv("STORE")
	fileName := c.Param("filename")
	selfLink := q["selfLink"][0]

	folder := services.Folder{
		Path: path,
	}

	err := folder.DeleteFile(selfLink)
	if err != nil {
		c.JSON(http.StatusNotFound, structs.Error{Code: http.StatusNotFound, Error: err.Error()})
		return
	}

	log.Printf("file %s was successfully deleted", fileName)
	files := folder.Interprete()

	c.JSON(http.StatusOK, structs.Success{Code: http.StatusOK, Data: files})
}
