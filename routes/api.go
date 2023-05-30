package routes

import (
	"kloc-go/controllers"

	"github.com/gin-gonic/gin"
)

func InitAPI(r *gin.Engine) {
	api := r.Group("/api/v1")
	{
		api.GET("/foots", controllers.GetFoots)
		api.GET("/projects", controllers.GetProjects)
	}
}
