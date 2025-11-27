package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/arhief32/emp-be/config"
	"github.com/arhief32/emp-be/middleware"
	v1controllers "github.com/arhief32/emp-be/v1/controllers"
	v1repositories "github.com/arhief32/emp-be/v1/repositories"
	v1routers "github.com/arhief32/emp-be/v1/routers"
	v1services "github.com/arhief32/emp-be/v1/services"
)

func main() {
	// load .env if exists
	_ = godotenv.Load()

	cfg := config.NewConfigFromEnv()
	db := config.InitDB(cfg)
	// auto migrate
	if err := config.AutoMigrate(db); err != nil {
		log.Fatalf("migrate error: %v", err)
	}

	// repositories
	authRepo := v1repositories.NewAuthRepository(db)
	empRepo := v1repositories.NewEmployeeRepository(db)
	reportRepo := v1repositories.NewReportRepository(db)

	// services
	authSvc := v1services.NewAuthService(authRepo, cfg)
	empSvc := v1services.NewEmployeeService(empRepo)
	reportSvc := v1services.NewReportService(reportRepo, empRepo)

	// controllers
	authCtrl := v1controllers.NewAuthController(authSvc)
	empCtrl := v1controllers.NewEmployeeController(empSvc)
	reportCtrl := v1controllers.NewReportController(reportSvc)

	// gin router
	r := gin.Default()
	// middleware init (uses cfg)
	authMw := middleware.NewJWTMiddleware(cfg)

	// register v1 routes
	v1routers.RegisterAuthRoutes(r, authCtrl, authMw)
	v1routers.RegisterEmployeeRoutes(r, empCtrl, authMw)
	v1routers.RegisterReportRoutes(r, reportCtrl, authMw)

	port := cfg.Port
	if port == "" {
		port = "8080"
	}
	log.Printf("listening on :%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
