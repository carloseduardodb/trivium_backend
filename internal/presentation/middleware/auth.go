package middleware

import (
	"context"
	"net/http"
	"strings"

	"trivium/internal/domain/repositorier"
	presentation_repositorier "trivium/internal/presentation/repositorier"
)

type contextKey string

const userIDKey contextKey = "user_id"
const userEmailKey contextKey = "user_email"

type Auth struct {
	verifyTokenRepo presentation_repositorier.VerifyTokenRepositorier
	userRepo        repositorier.UserRepositorier
}

func NewAuth(verifyTokenRepo presentation_repositorier.VerifyTokenRepositorier, userRepo repositorier.UserRepositorier) *Auth {
	return &Auth{
		verifyTokenRepo: verifyTokenRepo,
		userRepo:        userRepo,
	}
}

func (f *Auth) AuthMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Authorization header is missing", http.StatusUnauthorized)
				return
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenString == authHeader {
				http.Error(w, "Invalid Authorization header format. Must start with 'Bearer '", http.StatusUnauthorized)
				return
			}

			tokenUser, err := f.verifyTokenRepo.VerifyIdToken(r.Context(), tokenString)
			if err != nil {
				http.Error(w, "Unauthorized or invalid token", http.StatusUnauthorized)
				return
			}

			dbUser, err := f.userRepo.FindByEmail(tokenUser.Email)
			if err != nil {
				http.Error(w, "User not found", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), userIDKey, dbUser.ID)
			ctx = context.WithValue(ctx, userEmailKey, dbUser.Email)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetUserID(ctx context.Context) int64 {
	userID, ok := ctx.Value(userIDKey).(int64)
	if !ok {
		return 0
	}
	return userID
}

func GetUserEmail(ctx context.Context) string {
	email, ok := ctx.Value(userEmailKey).(string)
	if !ok {
		return ""
	}
	return email
}
