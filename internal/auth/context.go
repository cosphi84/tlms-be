package auth

import (
	"context"
	"errors"
)

func GetClaims(ctx context.Context) (*JWTClaims, error) {
	value := ctx.Value(AuthContextKey)

	if value == nil {
		return nil, errors.New("no auth claims found in context")
	}

	claims, ok := value.(*JWTClaims)

	if !ok {
		return nil, errors.New("invalid auth claims type in context")
	}

	return claims, nil
}
