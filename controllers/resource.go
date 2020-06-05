package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jesseokeya/nightowl/internal"
	"github.com/jesseokeya/nightowl/structs"
)

// GetResources retrieves all the resources
func GetResources(c *gin.Context) {
	folder := internal.Folder{
		Path: "/Users/jesseokeya/go/src/github.com/jesseokeya/nightowl/client",
	}

	files := folder.Interprete()

	reference := &files
	response, err := json.Marshal(reference)

	if err != nil {
		c.JSON(http.StatusNotFound, structs.Error{Code: http.StatusNotFound, Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
