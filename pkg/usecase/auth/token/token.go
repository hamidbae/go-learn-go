package token

import (
	"context"
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

// func VerifyJWT(ctx context.Context, token string) (claims claim.Access) {

// 	// Verify and extract claims from a token:
// 	verifiedToken, err := jwt.Verify(jwt.HS256, sharedKey, []byte(token))
// 	// unverifiedToken, err := jwt.Decode([]byte(token))
// 	if err != nil {
// 		panic(err)
// 	}

// 	err = verifiedToken.Claims(&claims)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return claims
// }
