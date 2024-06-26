package handler

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"tracks/pkg/sb"
	"tracks/types"
)

func WithUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "/public") {
			next.ServeHTTP(w, r)
			return
		}

		cookie, err := r.Cookie("at")

		if err != nil {
			fmt.Printf(err.Error())
			next.ServeHTTP(w, r)
			return
		}

		resp, err := sb.Client.Auth.User(r.Context(), cookie.Value)

		if err != nil {
			fmt.Printf(err.Error())
			next.ServeHTTP(w, r)
			return
		}

		user := types.AuthenticatedUser{Email: resp.Email, LoggedIn: true}

		ctx := context.WithValue(r.Context(), types.UserContextKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))

	})
}
