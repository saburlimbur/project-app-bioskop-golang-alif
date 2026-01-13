package middleware

import (
	"alfdwirhmn/bioskop/internal/data/repository"
	"alfdwirhmn/bioskop/pkg/utils"
	"context"
	"net/http"
	"strings"
)

type contextKey string

const (
	ContextUserID   contextKey = "user_id"
	ContextUsername contextKey = "username"
)

func AuthMiddleware(repo repository.Repository) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			authHeader := r.Header.Get("Authorization")
			if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
				utils.JSONError(w, http.StatusUnauthorized, "Missing or invalid Authorization header", nil)
				return
			}

			token := strings.TrimPrefix(authHeader, "Bearer ")

			// cek session di database
			session, err := repo.SessionRepo.FindValidSession(r.Context(), token)
			if err != nil {
				utils.JSONError(w, http.StatusUnauthorized, "Session expired or invalid", nil)
				return
			}

			// validasi JWT (signature & payload)
			// claims, err := utils.ValidateJWT(token)
			// if err != nil {
			// 	utils.JSONError(w, http.StatusUnauthorized, "Invalid token", nil)
			// 	return
			// }

			// inject ke context
			// ctx := context.WithValue(r.Context(), ContextUserID, session.UserID)
			// ctx = context.WithValue(ctx, ContextUsername, claims.Username)

			ctx := context.WithValue(r.Context(), ContextUserID, session.UserID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
