package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/yazilimcigenclik/dream-ai-backend/app/constants"
	"github.com/yazilimcigenclik/dream-ai-backend/app/pkg"
	"github.com/yazilimcigenclik/dream-ai-backend/app/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer pkg.PanicHandler(ctx)

		authKey := ctx.GetHeader("Authorization")

		if authKey == "" {
			pkg.PanicException(constants.Unauthorized)
		}

		if authKey[:7] != "Bearer " {
			pkg.PanicException(constants.Unauthorized)
		}

		authKey = authKey[7:]
		if authKey == "" {
			pkg.PanicException(constants.Unauthorized)
		}

		if authKey == "undefined" {
			pkg.PanicException(constants.Unauthorized)
		}

		if authKey == "null" {
			pkg.PanicException(constants.Unauthorized)
		}

		if authKey == "Bearer" {
			pkg.PanicException(constants.Unauthorized)
		}

		claim, err := utils.ParseJWTToken(authKey)
		if err != nil {
			pkg.PanicException(constants.Unauthorized)
		}

		ctx.Set("jwt_token", authKey)
		ctx.Set("username", claim.Username)

		ctx.Next()
	}
}
