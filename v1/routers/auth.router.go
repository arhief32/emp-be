package routers

import (
	"github.com/arhief32/emp-be/middleware"
	"github.com/arhief32/emp-be/v1/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.Engine, ctrl controllers.AuthController, authMw *middleware.JWTMiddleware) {
	g := r.Group("/v1/auth")

	g.POST("/register", ctrl.Register)
	g.POST("/login", ctrl.Login)

	g.GET("/profile", authMw.Gin(), ctrl.Profile)
}
