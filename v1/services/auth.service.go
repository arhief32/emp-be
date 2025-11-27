package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"

	"github.com/arhief32/emp-be/config"
	"github.com/arhief32/emp-be/v1/entities"
	"github.com/arhief32/emp-be/v1/models"
	"github.com/arhief32/emp-be/v1/repositories"
)

type AuthService interface {
	Register(req entities.RegisterRequest) error
	Login(req entities.LoginRequest) (string, *models.User, error)
	GetProfile(userID int) (*models.User, error)
}

type authService struct {
	repo repositories.AuthRepository
	cfg  *config.Config
}

func NewAuthService(r repositories.AuthRepository, cfg *config.Config) AuthService {
	return &authService{repo: r, cfg: cfg}
}

func (s *authService) Register(req entities.RegisterRequest) error {
	_, err := s.repo.FindByUsername(req.Username)
	if err == nil {
		return errors.New("username sudah dipakai")
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	user := models.User{
		Username: req.Username,
		Password: string(hash),
		Name:     req.Name,
	}

	return s.repo.Create(&user)
}

func (s *authService) Login(req entities.LoginRequest) (string, *models.User, error) {
	user, err := s.repo.FindByUsername(req.Username)
	if err != nil {
		return "", nil, errors.New("user tidak ditemukan")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return "", nil, errors.New("password salah")
	}

	// determine role
	role := "MAKER"

	// generate JWT
	claims := jwt.MapClaims{
		"sub":  user.ID,
		"usr":  user.Username,
		"role": role,
		"exp":  time.Now().Add(time.Duration(s.cfg.JWTExpHours) * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, _ := token.SignedString([]byte(s.cfg.JWTSecret))

	return signed, user, nil
}

func (s *authService) GetProfile(userID int) (*models.User, error) {
	return s.repo.FindByID(userID)
}
