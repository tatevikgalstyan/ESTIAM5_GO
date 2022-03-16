package util

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var mySigningKey = []byte("AllYourBase")

type MyCustomClaims struct {
	UserUUID  string `json:"user_uuid"`
	FirstName string
	jwt.StandardClaims
}

func CreateJWT(signKey []byte, userUUID, firstName string) (string, error) {

	// Create the Claims
	claims := MyCustomClaims{
		userUUID,
		firstName,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
			Issuer:    "Estiam",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", ss), nil
}

func MiddlJWT(signKey []byte) func(*gin.Context) {
	return func(ctx *gin.Context) {

		jwtValue := ctx.GetHeader("Authorization")
		if len(jwtValue) == 0 || !strings.Contains(jwtValue, "Bearer ") || len(jwtValue) < 200 {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(jwtValue[6:], func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return signKey, nil
		})

		if claims, ok := token.Claims.(MyCustomClaims); ok && token.Valid {
			fmt.Println(claims.UserUUID, claims.FirstName)
		} else {
			fmt.Println(err)
		}

	}
}
