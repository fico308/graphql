package auth

import (
	"context"
	"net/http"
)

// A private key for context that only this package can access. This is important
// to prevent collisions between different context uses
var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

// A stand-in for our database backed user object
type User struct {
	ID int
	Name string
	IsAdmin bool
}

// Middleware decodes the share session cookie and packs the session into context
func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("auth-user")

			// Allow unauthenticated users in
			if authHeader != "admin" {
				next.ServeHTTP(w, r)
				return
			}

			user := &User{
				ID: 1,
				Name: "admin",
				IsAdmin: true,
			}

			// put it in context
			ctx := context.WithValue(r.Context(), userCtxKey, user)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *User {
	u, ok := ctx.Value(userCtxKey).(*User)
	if !ok {
		return nil
	}
 	return u
}