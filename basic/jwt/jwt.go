package jwt

import (
	"context"
	"gimSec/basic/global"
	"gimSec/basic/logging"
	"gimSec/basic/response"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

var jwtKey = []byte("a_secret_crect")

type Claims struct {
	UserId string
	jwt.StandardClaims
}

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")

		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": "401", "msg": "权限不足"})
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:]

		token, claims, err := ParseToken(tokenString)

		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": "401", "msg": "token error"})
			ctx.Abort()
			return
		}

		userId := claims.UserId
		//TODO: 切换成Redis
		result, err := global.REDIS.Get(context.Background(), userId).Result()
		if err != nil {
			logging.Error(err)
			response.Error(ctx, err.Error())
			return
		}

		if tokenString != result {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": "401", "msg": "权限不足"})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	return token, claims, err
}
