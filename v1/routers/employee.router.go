package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/yourusername/pegawai-api/middleware"
	"github.com/yourusername/pegawai-api/v1/controllers"
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
