package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"

	"firebase.google.com/go/auth"
)

var userCtxKey = &contextKey{"Firebase UID"}

type contextKey struct {
	name string
}

func FirebaseAuth(fbAuthClient *auth.Client) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			authHeader := r.Header.Get("Authorization")
			authHeaderParts := strings.Split(authHeader, " ")
			if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
				// Authorization header format must be `Bearer {token}`
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}

			log.Printf("authHeader: %s", authHeader)
			idToken := authHeaderParts[1]
			log.Printf("idToken: %s", idToken)
			token, err := fbAuthClient.VerifyIDToken(ctx, idToken)
			if err != nil {
				// unauthorized
				log.Printf("error verifying ID token: %+v", err)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}

			// set uID
			uID := token.UID
			ctx = context.WithValue(ctx, userCtxKey, uID)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func ForContext(ctx context.Context) string {
	raw, _ := ctx.Value(userCtxKey).(string)
	return raw
}
