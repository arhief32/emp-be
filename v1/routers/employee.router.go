package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/arhief32/emp-be/middleware"
	"github.com/arhief32/emp-be/v1/controllers"
)

func RegisterEmployeeRoutes(r *gin.Engine, ctrl controllers.EmployeeController, authMw *middleware.JWTMiddleware) {
	group := r.Group("/v1/employees")
	group.Use(authMw.Gin())
	{
		group.GET("", ctrl.GetAll)
		group.POST("", ctrl.Create)
		group.GET("/:id", ctrl.GetByID)
		group.PUT("/:id", ctrl.Update)
		group.DELETE("/:id", ctrl.Delete)
	}
}
