package routers

import (
	"github.com/arhief32/emp-be/middleware"
	"github.com/arhief32/emp-be/v1/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoleRoutes(r *gin.Engine, ctrl *controllers.RoleController, authMw *middleware.JWTMiddleware) {
	route := r.Group("/v1/roles")
	{
		route.POST("/", ctrl.Create)
		route.GET("/", ctrl.GetAll)
		route.GET("/:id", ctrl.GetByID)
		route.PUT("/:id", ctrl.Update)
		route.DELETE("/:id", ctrl.Delete)
	}
}
