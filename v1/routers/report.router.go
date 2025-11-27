package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/yourusername/pegawai-api/middleware"
	"github.com/yourusername/pegawai-api/v1/controllers"
)

func RegisterReportRoutes(r *gin.Engine, ctrl controllers.ReportController, authMw *middleware.JWTMiddleware) {
	group := r.Group("/v1/reports")
	group.Use(authMw.Gin())
	{
		// GET /v1/reports/daily?date=YYYY-MM-DD
		group.GET("/daily", ctrl.Daily)
	}
}
