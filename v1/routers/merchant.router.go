package routers

import (
	"github.com/arhief32/emp-be/middleware"
	"github.com/arhief32/emp-be/v1/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterMerchantSubmissionRoutes(r *gin.Engine, ctrl controllers.MerchantSubmissionController, authMw *middleware.JWTMiddleware) {
	g := r.Group("/v1/merchant-submission")
	{
		// Maker routes (requires maker role)
		g.POST("", authMw.Gin(), authMw.RequireRole("MAKER"), ctrl.Create)
		g.PUT("/:id", authMw.Gin(), authMw.RequireRole("MAKER"), ctrl.Update)
		g.PUT("/:id/submit", authMw.Gin(), authMw.RequireRole("MAKER"), ctrl.Submit)
		g.GET("/mine", authMw.Gin(), authMw.RequireRole("MAKER"), ctrl.ListMine)

		// Checker routes
		g.GET("/pending", authMw.Gin(), authMw.RequireRole("CHECKER"), ctrl.ListPendingForChecker)
		g.PUT("/:id/check-approve", authMw.Gin(), authMw.RequireRole("CHECKER"), ctrl.CheckerApprove)
		g.PUT("/:id/check-reject", authMw.Gin(), authMw.RequireRole("CHECKER"), ctrl.CheckerReject)

		// Signer routes
		g.PUT("/:id/sign-approve", authMw.Gin(), authMw.RequireRole("SIGNER"), ctrl.SignerApprove)
		g.PUT("/:id/sign-reject", authMw.Gin(), authMw.RequireRole("SIGNER"), ctrl.SignerReject)

		// public/admin
		g.GET("", authMw.Gin(), ctrl.ListMine) // or admin only
		// g.GET("/:id", authMw.Gin(), ctrl.Get)  // implement Get if needed
	}
}
