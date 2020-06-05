package main

import (
	"log"
	"os"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/jesseokeya/nightowl/routes"
	"github.com/joho/godotenv"
)

func main() {
	var PORT string

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	path := "." + os.Getenv("STORE")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModePerm)
	}

	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "3000"
	}

	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./client/build", true)))
	r.Static(path, path)
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Add("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Add("Access-Control-Allow-Methods", "GET,HEAD,OPTIONS,POST,PUT,DELETE")
		c.Writer.Header().Add("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	})

	routes.Router(r)

	r.Run("0.0.0.0:" + PORT)
}
