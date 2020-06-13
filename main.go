package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jesseokeya/nightowl/routes"
	"github.com/jesseokeya/nightowl/services"
	"github.com/joho/godotenv"
	"github.com/martini-contrib/method"
)

var overrideHandler = method.Override()

// Override ...
func Override(c *gin.Context) {
	overrideHandler.ServeHTTP(c.Writer, c.Request)
}

func main() {
	var PORT string

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	store := os.Getenv("STORE")
	path := "." + store
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModePerm)
	}

	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "3000"
	}

	config := services.UserFTPConfig{
		User: os.Getenv("USER"),
		Pass: os.Getenv("PASS"),
		Path: store,
	}

	config.CreateFTPConfiguration()

	r := gin.Default()
	r.Use(Override)
	// r.Use(static.Serve("/", static.LocalFile("./client/build", true)))
	r.Static(path, path)
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
		} else {
			c.Next()
		}
	})

	routes.Router(r.Group("/api/v1"))

	r.Run("0.0.0.0:" + PORT)
}
