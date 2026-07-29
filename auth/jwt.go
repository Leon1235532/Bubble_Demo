package auth

import (
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"github.com/Leon1235532/GoTask/common"
	"github.com/Leon1235532/GoTask/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	UserID uint `json:"userid"`
	jwt.RegisteredClaims
}

var hmacSecret []byte

func InitJwt(secretHex string) {
	hmacSecret, _ = hex.DecodeString(secretHex)
}

func CreateAccToken(user *models.User) (tokenstring string, err error) {
	claims := UserClaims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenstring, err = token.SignedString(hmacSecret)
	return
}

func ParseToken(c *gin.Context) (token *jwt.Token, err error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return nil, fmt.Errorf("Token err")
	}
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader {
		return nil, fmt.Errorf("Token form err")
	}
	token, err = jwt.ParseWithClaims(tokenString, &UserClaims{}, func(t *jwt.Token) (any, error) {
		return hmacSecret, nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	return
}

func ExtractUserId() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := ParseToken(c)
		if err != nil {
			message := "Token有误,请重新登录"
			common.ErrorResponse(c, message, err.Error())
			return
		}
		if claims, ok := token.Claims.(*UserClaims); ok {
			c.Set("userid", claims.UserID)
			c.Next()
		} else {
			common.ErrorResponse(c, "Token form err", fmt.Errorf("claims 类型断言失败").Error())
			return
		}
	}
}
