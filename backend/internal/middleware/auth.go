package middleware

import (
	jwttoken "backend/pkg/jwt_token"
	"context"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type contextKey string

const (
	UserClaimsContextKey contextKey = "user_claims"
)

const UserClaimsKey = "user_claims"

var jwtManager *jwttoken.JWTManager

func InitAuth(c jwttoken.Config) *jwttoken.JWTManager {
	jwtManager = jwttoken.NewJWTManager(c)
	return jwtManager
}

func AuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				logrus.Warn("Authorization header is missing")
				return echo.NewHTTPError(http.StatusUnauthorized, "Authorization header is required")
			}

			tokenParts := strings.Split(authHeader, " ")
			if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != "bearer" {
				logrus.Warn("Invalid authorization header format")
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid authorization header format")
			}

			tokenStr := tokenParts[1]

			claims, err := jwtManager.ParseToken(tokenStr)
			if err != nil {
				logrus.Warnf("Invalid token: %v", err)
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid or expired token")
			}

			c.Set(UserClaimsKey, claims)

			logrus.Infof("Authenticated user: %s (ID: %d)", claims.Username, claims.UserID)

			return next(c)
		}
	}
}

func GetUserClaims(ctx context.Context) (*jwttoken.Claims, bool) {
	claims, ok := ctx.Value(UserClaimsContextKey).(*jwttoken.Claims)
	return claims, ok
}
