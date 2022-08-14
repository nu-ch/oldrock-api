package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

var signingKey = []byte(DotEnvVariable("JWT_SECRET"))

func IsAuthorized(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Authorization"] == nil || len(r.Header["Authorization"]) == 0 {
			ErrorResponse("Missing authorization header", w)
			return
		}

		reqToken := r.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer ")
		reqToken = splitToken[1]

		accessToken, err := jwt.Parse(reqToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error")
			}
			return signingKey, nil
		})

		if err != nil {
			if err == jwt.ErrTokenExpired {
				ErrorResponse("Invalid token expired", w)
				return
			} else {
				ErrorResponse("Invalid authorization token", w)
				return
			}
		}

		if accessToken.Valid {
			next.ServeHTTP(w, r)
		}

	})
}

// func GenerateJWT() (string, error) {
// 	token := jwt.E(jwt.SigningMethodHS256)

// 	claims := token.Claims.(jwt.MapClaims)

// 	claims["accountId"] = "Elliot Forbes"
// 	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

// 	tokenString, err := token.SignedString(signingKey)

// 	if err != nil {
// 		return "", err
// 	}

// 	return tokenString, nil
// }
