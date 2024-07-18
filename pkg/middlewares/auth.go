package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/Jacobo0312/go-web/pkg/errors"
	"github.com/Jacobo0312/go-web/pkg/firebase"
	"github.com/Jacobo0312/go-web/pkg/helpers"
)

func FirebaseAuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		idToken := strings.TrimSpace(strings.Replace(authHeader, "Bearer", "", 1))

		if idToken == "" {
			helpers.RespondWithError(w, errors.NewUnauthorized("Unauthorized"))
			return
		}

		token, err := firebase.FirebaseAuth.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			helpers.RespondWithError(w, errors.NewUnauthorized("Invalid token"))
			return
		}

		// Add userID to context
		type contextKey string
		ctx := context.WithValue(r.Context(), contextKey("userID"), token.UID)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}