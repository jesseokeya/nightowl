package routes

import (
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

	//image routes
	r.GET("/images", ctrl.GetImages)
	r.GET("/images/:id", ctrl.GetImage)
	r.POST("/images", ctrl.CreateImage)
	r.PUT("/images/:id", ctrl.UpdateImage)
	r.DELETE("/images/:id", ctrl.DeleteImage)
}
