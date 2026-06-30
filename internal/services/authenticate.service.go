package services

import (
	"errors"
	"log"
	"os"
	"time"
	"tlms/internal/auth"
	"tlms/internal/dto"
	"tlms/internal/repositories"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type AuthenticateService interface {
	Authenticate(req *dto.LoginRequest, ip string) (*dto.LoginResponse, error)
	RefreshToken(req *dto.RefreshTokenRequest) (*dto.LoginResponse, error)
}

type authenticateService struct {
	userRepo repositories.UserRepository
}

func NewAuthenticateService(userRepo repositories.UserRepository) AuthenticateService {
	return &authenticateService{
		userRepo: userRepo,
	}
}

func (s *authenticateService) Authenticate(req *dto.LoginRequest, ip string) (*dto.LoginResponse, error) {
	maxLoginAttempts := 3
	lockoutDuration := 1 // in hours
	lockUntil := time.Now().Add(time.Duration(lockoutDuration) * time.Hour)

	usr, err := s.userRepo.FindByEmail(req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	if !usr.IsActive {
		return nil, errors.New("user is inactive")
	}

	if usr.LockedUntil != nil && usr.LockedUntil.After(time.Now()) {
		return nil, errors.New("account is locked. please try again later")
	}

	log.Printf("Authenticating user: %s from IP: %s", req.Email, ip)

	valid := auth.VerifyPassword(req.Password, usr.Password)
	if !valid {
		nTry := usr.FailedLoginAttempts + 1
		if nTry >= maxLoginAttempts {
			usr.LockedUntil = &lockUntil
		}

		usr.FailedLoginAttempts = nTry
		_ = s.userRepo.Update(usr)
		return nil, errors.New("invalid credentials")
	}

	now := time.Now()
	usr.LastLoginFrom = &ip
	usr.LastLoginAt = &now
	usr.FailedLoginAttempts = 0
	usr.LockedUntil = nil

	_ = s.userRepo.Update(usr)

	accessToken, refreshToken, err := auth.GenerateTokenPair(usr.ID, usr.OfficeID, "")
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         usr,
	}, nil
}

func (s *authenticateService) RefreshToken(req *dto.RefreshTokenRequest) (*dto.LoginResponse, error) {
	secret := os.Getenv("JWT_SECRET")
	var JWTSecret = []byte(secret)

	token, err := jwt.ParseWithClaims(
		req.RefreshToken,
		&auth.JWTClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return JWTSecret, nil
		},
	)

	if err != nil || !token.Valid {
		return nil, errors.New("invalid refresh token")
	}

	claims, ok := token.Claims.(*auth.JWTClaims)
	if !ok || claims.TokenType != "refresh" {
		return nil, errors.New("invalid token type")
	}

	usr, err := s.userRepo.FindByID(int32(claims.UserID))
	if err != nil || usr == nil {
		return nil, errors.New("user not found")
	}
	if !usr.IsActive {
		return nil, errors.New("user is inactive")
	}

	accessToken, newRefreshToken, err := auth.GenerateTokenPair(usr.ID, usr.OfficeID, claims.Role)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
		User:         usr,
	}, nil
}
