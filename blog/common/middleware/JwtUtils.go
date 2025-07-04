package middleware

import (
	"GolangStudy/blog/common/config"
	"GolangStudy/blog/common/result"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
)

type Myclaims struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func CreateJwt(claims Myclaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := []byte(config.Conf.Jwt.Secret) // 对称密钥
	return token.SignedString(secretKey)
}

func JwtAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		log.Printf("Request Token: %s", tokenString)
		if tokenString == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, result.ErrSignParam)
			return
		}
		var claims Myclaims
		token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Conf.Jwt.Secret), nil
		})
		if err != nil || !token.Valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, result.Error("Token无效或已过期"))
			return
		}
		// 验证签发者（可选）
		if claims.Issuer != config.Conf.Jwt.Issuer {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, result.Error("Token无效"))
			return
		}

		ctx.Set("ID", claims.ID)
		ctx.Set("Username", claims.Username)
		ctx.Next()
	}
}
