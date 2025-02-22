package api

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/vldcreation/simple_bank/token"
)

const (
	AuthHeader   = "Authorization"
	BearerSchema = "Bearer"
	AuthKey      = "auth_payload"
)

var (
	ErrUnauthorized = NewApiError(http.StatusUnauthorized)
)

func authMiddleware(token token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(AuthHeader)
		if len(authorizationHeader) == 0 {
			err := ErrUnauthorized.WithMessage("authorization header is not provided")
			ctx.AbortWithStatusJSON(err.Code, errorResponse(err))
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := ErrUnauthorized.WithMessage("invalid authorization header format")
			ctx.AbortWithStatusJSON(err.Code, errorResponse(err))
			return
		}

		schema := fields[0]
		if schema != BearerSchema {
			err := ErrUnauthorized.WithMessage("unsupported authorization schema")
			ctx.AbortWithStatusJSON(err.Code, errorResponse(err))
			return
		}

		accessToken := fields[1]
		payload, err := token.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		ctx.Set(AuthKey, payload)
		ctx.Next()
	}
}
