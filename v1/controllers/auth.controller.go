package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/arhief32/emp-be/v1/entities"
	"github.com/arhief32/emp-be/v1/services"
)

type AuthController struct {
	svc services.AuthService
}

func NewAuthController(s services.AuthService) AuthController {
	return AuthController{s}
}

func (ctr *AuthController) Register(c *gin.Context) {
	var req entities.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctr.svc.Register(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "register berhasil"})
}

func (ctr *AuthController) Login(c *gin.Context) {
	var req entities.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, user, err := ctr.svc.Login(req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user":  user,
	})
}

func (ctr *AuthController) Profile(c *gin.Context) {
	userID := c.GetInt("user_id")

	// Ambil data user dari database
	user, err := ctr.svc.GetProfile(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user_id":    user.ID,
		"username":   user.Username,
		"fullname":   user.Fullname,
		"created_at": user.CreatedAt,
	})
}
