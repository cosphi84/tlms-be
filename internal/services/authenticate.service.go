package services

import (
	"errors"
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
	authz    *auth.Service
}

func NewAuthenticateService(authz *auth.Service, userRepo repositories.UserRepository) AuthenticateService {
	return &authenticateService{
		userRepo: userRepo,
		authz:    authz,
	}
}

func (s *authenticateService) Authenticate(req *dto.LoginRequest, ip string) (*dto.LoginResponse, error) {
	maxLoginAttempts := 3
	lockoutDuration := 1 // in hours
	usr, err := s.userRepo.FindByEmail(req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid credentials")
		}
		return nil, err
	}

	if !usr.IsActive {
		return nil, errors.New("invalid credentials")
	}

	now := time.Now()
	if usr.LockedUntil != nil {
		if now.Before(*usr.LockedUntil) {
			return nil, errors.New("account locked")
		}

		usr.LockedUntil = nil
		usr.FailedLoginAttempts = 0
	}

	valid := auth.VerifyPassword(req.Password, usr.Password)
	if !valid {
		nTry := usr.FailedLoginAttempts + 1

		usr.FailedLoginAttempts = nTry
		if nTry >= maxLoginAttempts {
			lockedUntil := time.Now().Add(time.Duration(lockoutDuration) * time.Hour)
			usr.LockedUntil = &lockedUntil
			usr.FailedLoginAttempts = 0
		}

		err = s.userRepo.Update(usr)

		if err != nil {
			return nil, err
		}
		return nil, errors.New("invalid credentials")
	}

	usr.LastLoginFrom = &ip
	usr.LastLoginAt = &now
	usr.FailedLoginAttempts = 0
	usr.LockedUntil = nil

	_ = s.userRepo.Update(usr)

	usrRolesRaw, _ := s.authz.GetRoleForUser(usr.Email)
	usrRoles := make([]auth.RoleType, len(usrRolesRaw))
	for i, r := range usrRolesRaw {
		usrRoles[i] = auth.RoleType(r)
	}
	accessToken, refreshToken, err := auth.GenerateTokenPair(usr.ID, usr.OfficeID, usrRoles)
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
