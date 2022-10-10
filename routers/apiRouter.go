package routers

import (
	"rbac/controllers"

	"github.com/gin-gonic/gin"
)

func APIRouter(r *gin.Engine) {
	apiRouter := r.Group("/api")
	{
		apiRouter.GET("/", controllers.Hello)

		apiRouter.GET("/user", controllers.UserController{}.GetUsers)
		apiRouter.POST("/user", controllers.UserController{}.CreateUser)
		apiRouter.PUT("/user", controllers.UserController{}.UpdateUser)
		apiRouter.DELETE("/user", controllers.UserController{}.DeleteUser)
	}
}
