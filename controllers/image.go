package controllers

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetImages retrieves all the images
func GetImages(c *gin.Context) {}

// GetImage retrieves a particular image
func GetImage(c *gin.Context) {}

// CreateImage creates a new image
func CreateImage(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)
	println(string(body))
	c.JSON(http.StatusOK, gin.H{"msg": "this worked"})
}

// UpdateImage updates a particular image
func UpdateImage(c *gin.Context) {}

// DeleteImage deletes a particular image
func DeleteImage(c *gin.Context) {}
