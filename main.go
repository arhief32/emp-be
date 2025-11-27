package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/arhief32/emp-be/config"
	"github.com/arhief32/emp-be/middleware"
	"github.com/arhief32/emp-be/v1/controllers"
	v1controllers "github.com/arhief32/emp-be/v1/controllers"
	"github.com/arhief32/emp-be/v1/repositories"
	v1repositories "github.com/arhief32/emp-be/v1/repositories"
	v1routers "github.com/arhief32/emp-be/v1/routers"
	"github.com/arhief32/emp-be/v1/services"
	v1services "github.com/arhief32/emp-be/v1/services"
)

func main() {
	// load .env if exists
	config.InitEnv()

	cfg := config.NewConfigFromEnv()
	db := config.InitDB(cfg)
	// auto migrate
	if err := config.AutoMigrate(db); err != nil {
		log.Fatalf("migrate error: %v", err)
	}

	// auth
	authRepo := v1repositories.NewAuthRepository(db)
	authSvc := v1services.NewAuthService(authRepo, cfg)
	authCtrl := v1controllers.NewAuthController(authSvc)

	// role
	roleRepo := repositories.NewRoleRepository(db)
	roleSvc := services.NewRoleService(roleRepo)
	roleCtrl := controllers.NewRoleController(roleSvc)

	// merchant
	submissionRepo := repositories.NewMerchantSubmissionRepository(db)
	submissionSvc := services.NewMerchantSubmissionService(submissionRepo)
	submissionCtrl := controllers.NewMerchantSubmissionController(submissionSvc)

	// employee
	empRepo := v1repositories.NewEmployeeRepository(db)
	empSvc := v1services.NewEmployeeService(empRepo)
	empCtrl := v1controllers.NewEmployeeController(empSvc)

	// report
	reportRepo := v1repositories.NewReportRepository(db)
	reportSvc := v1services.NewReportService(reportRepo, empRepo)
	reportCtrl := v1controllers.NewReportController(reportSvc)

	// gin router
	r := gin.Default()
	authMw := middleware.NewJWTMiddleware(cfg)

	// register v1 routes
	v1routers.RegisterAuthRoutes(r, authCtrl, authMw)
	v1routers.RegisterRoleRoutes(r, roleCtrl, authMw)
	v1routers.RegisterEmployeeRoutes(r, empCtrl, authMw)
	v1routers.RegisterReportRoutes(r, reportCtrl, authMw)
	v1routers.RegisterMerchantSubmissionRoutes(r, submissionCtrl, authMw)

	port := cfg.Port
	if port == "" {
		port = "8080"
	}
	log.Printf("listening on :%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
