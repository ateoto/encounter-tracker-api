package main

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func JWTAuth(key string) echo.HandlerFunc {
	return func(c *echo.Context) error {

		// If this is a WS upgrade request, skip Auth
		if (c.Request().Header.Get(echo.Upgrade)) == echo.WebSocket {
			return nil
		}

		auth := c.Request().Header.Get("Authorization")
		l := len(Bearer)

		unauthorized := echo.NewHTTPError(http.StatusUnauthorized)

		if len(auth) > l+1 && auth[:l] == Bearer {
			t, err := jwt.Parse(auth[l+1:], func(token *jwt.Token) (interface{}, error) {

				// Check the signing method
				// https://auth0.com/blog/2015/03/31/critical-vulnerabilities-in-json-web-token-libraries/
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}

				return []byte(key), nil
			})
			if err == nil && t.Valid {
				// Store claims in echo.Context
				c.Set("claims", t.Claims)
				return nil
			}
		}

		return unauthorized
	}
}
