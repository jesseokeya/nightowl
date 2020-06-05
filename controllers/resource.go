package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jesseokeya/nightowl/internal"
	"github.com/jesseokeya/nightowl/structs"
)

// GetResources retrieves all the resources
func GetResources(c *gin.Context) {
	path := "/Users/jesseokeya/go/src/github.com/jesseokeya/nightowl/resources"
	// query := c.Request.URL.Query()["file"]

	folder := internal.Folder{
		Path: path,
	}

	files := folder.Interprete()

	c.JSON(http.StatusOK, structs.Success{Code: http.StatusOK, Data: files})
}
