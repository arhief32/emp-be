package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/arhief32/emp-be/v1/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	Port        string
	DBHost      string
	DBPort      string
	DBUser      string
	DBPass      string
	DBName      string
	JWTSecret   string
	JWTExpHours int
}

func NewConfigFromEnv() *Config {
	expHours := 24
	if v := os.Getenv("JWT_EXP_HOURS"); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			expHours = n
		}
	}
	return &Config{
		Port:        os.Getenv("PORT"),
		DBHost:      getenv("DB_HOST", "127.0.0.1"),
		DBPort:      getenv("DB_PORT", "3306"),
		DBUser:      getenv("DB_USER", "root"),
		DBPass:      getenv("DB_PASS", ""),
		DBName:      getenv("DB_NAME", "pegawai_db"),
		JWTSecret:   getenv("JWT_SECRET", "verysecretkey"),
		JWTExpHours: expHours,
	}
}

func getenv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func InitDB(cfg *Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed connect db: %v", err)
	}
	return db
}

// AutoMigrate models
func AutoMigrate(db *gorm.DB) error {
	// import models to migrate
	return db.AutoMigrate(
		&models.User{},
		&models.Employee{},
		&models.DailyReport{})
}
