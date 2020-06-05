package routes

import (
	"os"

	"github.com/gin-gonic/gin"
	ctrl "github.com/jesseokeya/nightowl/controllers"
)

// Router helps roue our application to the appropriate controller
func Router(r *gin.Engine) {

	//user routes
	r.GET("/users", ctrl.GetUsers)
	r.GET("/users/:id", ctrl.GetUser)
	r.POST("/users", ctrl.CreateUser)
	r.PUT("/users/:id", ctrl.UpdateUser)
	r.DELETE("/users/:id", ctrl.DeleteUser)

	//resources routes
	r.GET(os.Getenv("STORE"), ctrl.GetResources)
}
