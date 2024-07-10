package middlewares

import (
    "context"
    "net/http"
    "strings"

    "github.com/Jacobo0312/go-web/pkg/firebase"
)

func FirebaseAuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        idToken := strings.TrimSpace(strings.Replace(authHeader, "Bearer", "", 1))

        if idToken == "" {
            http.Error(w, "No token provided", http.StatusUnauthorized)
            return
        }

        token, err := firebase.FirebaseAuth.VerifyIDToken(context.Background(), idToken)
        if err != nil {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }

        // Añade el UID del usuario al contexto de la solicitud
        ctx := context.WithValue(r.Context(), "userID", token.UID)
        next.ServeHTTP(w, r.WithContext(ctx))
    }
}