package api

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/luckyparakh/goBank/token"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func authMiddleware(token token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		feilds := strings.Split(authorizationHeader, " ")
		fmt.Println(feilds)
		if len(feilds) < 2 {
			err := errors.New("invalid authorization header format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
		authorizationType := strings.ToLower(feilds[0])
		if authorizationType != authorizationTypeBearer {
			err := fmt.Errorf("authorization type %s is not supported", authorizationType)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
		accessToken := feilds[1]
		payload, err := token.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()
	}
}
