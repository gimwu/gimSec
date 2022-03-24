package jwt

import (
	"gimSec/basic/logging"
	model2 "gimSec/basic/model"
	"gimSec/basic/response"
	"gimSec/server"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

var jwtKey = []byte("a_secret_crect")

type Claims struct {
	UserId string
	jwt.StandardClaims
}

func ReleaseToken(entity model2.StateFullEntity) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &Claims{
		UserId: entity.Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "oceanlearn.tech",
			Subject:   "user token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
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
		user, err := server.GetUser(userId)
		if err != nil {
			logging.Error(err)
			response.Error(ctx, err)
		}

		if user.Id == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": "401", "msg": "权限不足"})
			ctx.Abort()
			return
		}

		ctx.Set("user", user)
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
