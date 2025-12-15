package middleware

import (
	jwttoken "backend/pkg/jwt_token"
	"context"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
)

type contextKey string

const (
	UserClaimsContextKey contextKey = "user_claims"
)

var jwtManager *jwttoken.JWTManager

func InitAuth(c jwttoken.Config) {
	jwtManager = jwttoken.NewJWTManager(c)
}

func AuthMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				logrus.Warn("Authorization header is missing")
				http.Error(w, "Authorization header is required", http.StatusUnauthorized)
				return
			}

			tokenParts := strings.Split(authHeader, " ")
			if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != "bearer" {
				logrus.Warn("Invalid authorization header format")
				http.Error(w, "Invalid authorization header format", http.StatusUnauthorized)
				return
			}

			tokenStr := tokenParts[1]

			claims, err := jwtManager.ParseToken(tokenStr)
			if err != nil {
				logrus.Warnf("Invalid token: %v", err)
				http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), UserClaimsContextKey, claims)
			r = r.WithContext(ctx)

			logrus.Infof("Authenticated user: %s (ID: %d)", claims.Username, claims.UserID)

			next.ServeHTTP(w, r)
		})
	}
}

func GetUserClaims(ctx context.Context) (*jwttoken.Claims, bool) {
	claims, ok := ctx.Value(UserClaimsContextKey).(*jwttoken.Claims)
	return claims, ok
}
