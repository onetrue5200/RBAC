package routers

import (
	"rbac/controllers"

	"github.com/gin-gonic/gin"
)

func APIRouter(r *gin.Engine) {
	apiRouter := r.Group("/api")
	{
		apiRouter.GET("/", controllers.Hello)

		apiRouter.GET("/user", controllers.UserController{}.List)
		apiRouter.POST("/user", controllers.UserController{}.Create)
		apiRouter.PUT("/user", controllers.UserController{}.Update)
		apiRouter.DELETE("/user/:id", controllers.UserController{}.Delete)

		apiRouter.GET("/role", controllers.RoleController{}.List)
		apiRouter.POST("/role", controllers.RoleController{}.Create)
		apiRouter.PUT("/role", controllers.RoleController{}.Update)
		apiRouter.DELETE("/role/:id", controllers.RoleController{}.Delete)

		apiRouter.GET("/user_role", controllers.UserRoleController{}.List)
		apiRouter.POST("/user_role", controllers.UserRoleController{}.Create)
		apiRouter.DELETE("/user_role", controllers.UserRoleController{}.Delete)

		apiRouter.GET("/access", controllers.AccessController{}.List)
		apiRouter.POST("/access", controllers.AccessController{}.Create)
		apiRouter.PUT("/access", controllers.AccessController{}.Update)
		apiRouter.DELETE("/access", controllers.AccessController{}.Delete)

		apiRouter.GET("/role_access", controllers.RoleAccessController{}.List)
		apiRouter.POST("/role_access", controllers.RoleAccessController{}.Create)
		apiRouter.DELETE("/role_access", controllers.RoleAccessController{}.Delete)
	}
}
