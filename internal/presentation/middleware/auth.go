package middleware

import (
	"context"
	"net/http"
	"strings"
	"trivium/internal/presentation/repositorier"
)

type contextKey string

const userKey contextKey = "user"

type Auth struct {
	verifyTokenRepo repositorier.VerifyTokenRepositorier
}

func NewAuth(verifyTokenRepo repositorier.VerifyTokenRepositorier) *Auth {
	return &Auth{
		verifyTokenRepo: verifyTokenRepo,
	}
}

func (f *Auth) AuthMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Cabeçalho Authorization está ausente", http.StatusUnauthorized)
				return
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenString == authHeader {
				http.Error(w, "Formato do cabeçalho Authorization inválido. Deve começar com 'Bearer '", http.StatusUnauthorized)
				return
			}

			user, err := f.verifyTokenRepo.VerifyIdToken(r.Context(), tokenString)
			if err != nil {
				http.Error(w, "Token não autorizado ou inválido", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), userKey, user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
