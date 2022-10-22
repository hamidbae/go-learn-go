package middleware

import (
	"net/http"
	"strings"
	"time"

	"final-project/pkg/domain/response"
	errortype "final-project/pkg/domain/response/error-type"
	"final-project/pkg/usecase/auth/token"

	"github.com/gin-gonic/gin"
)

func CheckJwtAuth(ctx *gin.Context) {
	
	// encoding dan decoding
	// encoding -> masking data ke suatu bentuk yang menjadi tidak terbaca
	// ex: calman --- function ---> klsadl9u1214
	// decoding -> unmasking data dari yang tidak terbaca menjadi yang terbaca
	// ex: klsadl9u1214 --- function ---> calman
	// function crypto biasanya memerlukan suatu specific key
	// untuk mengencode dan decode

	// encoding/decoding
	// -> dua arah, bisa diencode, bisa didecode (ex: Base64)
	// -> satu arah, cuma bisa diencode (ex: hash)
	
	// dalam auth API
	// JWT biasa digunakan untuk OAUth
	// sehingga dia termasuk dalam bearer token

	
	responseMessage := response.Response{
		Message: "authentication failed",
		InvalidArg: &response.InvalidArg{
			ErrorType:    errortype.AUTHENTICATION_FAILED,
			ErrorMessage: "token not verified",
		},
	}
	
	bearer := ctx.GetHeader("Authorization")
	bearerArray := strings.Split(bearer, " ") // -> ["Bearer", "TOKEN JWT"]
	if len(bearerArray) != 2 {
		ctx.AbortWithStatusJSON(
			http.StatusUnauthorized,
			response.Build(
				http.StatusUnauthorized,
				responseMessage,
			),
		)
		return
	}

	// check only Basic prefix allowed
	if bearerArray[0] != "Bearer" {
		ctx.AbortWithStatusJSON(
			http.StatusUnauthorized,
			response.Build(
				http.StatusUnauthorized,
				responseMessage,
			),
		)
		return
	}

	// get claim
	claim, err := token.VerifyJWT(ctx, bearerArray[1])
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusUnauthorized,
			response.Build(
				http.StatusUnauthorized,
				responseMessage,
			),
		)
		return
	}

	// validate claim
		// 	JWTID:          uuid.New(),
		// Subject:        userCheck.ID,
		// Issuer:         "go-fga.com",
		// Audience:       "user.go-fga.com",
		// Scope:          "user",
		// Type:           "ACCESS_TOKEN",
		// IssuedAt:       timeNow.Unix(),
		// NotValidBefore: timeNow.Unix(),
		// ExpiredAt:      timeNow.Add(24 * time.Hour).Unix(),

	// 1. validate issuernya bener ga dari go-fga.com
	if claim.Issuer != "go-fga.com" {
		ctx.AbortWithStatusJSON(
			http.StatusUnauthorized,
			response.Build(
				http.StatusUnauthorized,
				responseMessage,
			),
		)
		return
	}

	// 2. check audience bener ga untuk user.go-fga.com
	if claim.Audience != "user.go-fga.com" {
		ctx.AbortWithStatusJSON(
			http.StatusUnauthorized,
			response.Build(
				http.StatusUnauthorized,
				responseMessage,
			),
		)
		return
	}

	// 3. check scopenya bener ga untuk user endpoint
	if claim.Scope != "user" {
		ctx.AbortWithStatusJSON(
			http.StatusUnauthorized,
			response.Build(
				http.StatusUnauthorized,
				responseMessage,
			),
		)
		return
	}

	// 4. token ini udah bisa digunakan belum?
	if !time.Unix(claim.NotValidBefore, 0).Before(time.Now()) {
		ctx.AbortWithStatusJSON(
			http.StatusUnauthorized,
			response.Build(
				http.StatusUnauthorized,
				responseMessage,
			),
		)
		return
	}

	// 5. check tokennya udah expired atau belum
	if time.Unix(claim.ExpiredAt, 0).Before(time.Now()) {
		ctx.AbortWithStatusJSON(
			http.StatusUnauthorized,
			response.Build(
				http.StatusUnauthorized,
				responseMessage,
			),
		)
		return
	}

	// set user in context
	ctx.Set("user_id", claim.Subject)
	ctx.Next()
}
