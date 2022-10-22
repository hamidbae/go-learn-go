package token

import (
	"context"
	"final-project/pkg/domain/auth"
	"os"

	"github.com/kataras/jwt"
)

var (
	jwtSecret = []byte(os.Getenv("JWT_SECRET"))
)

func CreateJWT(ctx context.Context, claim any) (string, error) {
	token, err := jwt.Sign(jwt.HS256, jwtSecret, claim)
	if err != nil {
		return "", err
	}
	return string(token), nil
}

func VerifyJWT(ctx context.Context, token string) (claims auth.ClaimAccess, err error) {

	// Verify and extract claims from a token:
	verifiedToken, err := jwt.Verify(jwt.HS256, jwtSecret, []byte(token))
	// unverifiedToken, err := jwt.Decode([]byte(token))
	if err != nil {
		return claims, err
	}

	err = verifiedToken.Claims(&claims)
	if err != nil {
		return claims, err
	}
	return claims, nil
}
