package middleware

import (
	"net/http"
	"strings"

	"beebeewijaya.com/token"
	"beebeewijaya.com/util"
	"github.com/gin-gonic/gin"
)

func ExtractToken(ctx *gin.Context) (string, error) {
	bearerToken := ctx.Request.Header.Get(util.Authorization)
	if bearerToken == "" {
		return "", util.ErrEmptyToken
	}

	splittedToken := strings.Split(bearerToken, " ")
	if len(splittedToken) < 2 {
		return "", util.ErrTokenInvalid
	}

	if splittedToken[0] != util.Bearer {
		return "", util.ErrTokenInvalid
	}

	return splittedToken[1], nil
}

func JWTAuthMiddleware(maker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t, err := ExtractToken(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, util.ErrorResponse(err))
			return
		}

		payload, err := maker.VerifyToken(t)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, util.ErrorResponse(err))
			return
		}

		ctx.Set(util.AuthPayloadKey, payload)
		ctx.Next()
	}
}
