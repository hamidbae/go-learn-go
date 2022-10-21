package auth

import (
	"time"

	"github.com/google/uuid"
)

type Login struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password,omitempty" validate:"required,min=6"`
}

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenID      string `json:"token_id"`
}

type ClaimAccess struct {
	JWTID          uuid.UUID `json:"jti"`   // menandakan id berapa untuk token ini
	Subject        uint64    `json:"sub"`   // token ini untuk user siapa
	Issuer         string    `json:"iss"`   // token ini dibuat oleh siapa
	Audience       string    `json:"aud"`   // token ini boleh digunakan oleh siapa
	Scope          string    `json:"scope"` // optional menandakan bisa mengakses apa aja
	Type           string    `json:"type"`  // tipe dari token ini
	IssuedAt       int64     `json:"iat"`   // token ini dibuat kapan
	NotValidBefore int64     `json:"nbf"`   // token ini boleh digunakan setelah kapan
	ExpiredAt      int64     `json:"exp"`   // token ini akan expired kapan
}

type ClaimID struct {
	JWTID    uuid.UUID `json:"jti"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	DOB      time.Time `json:"dob"`
}