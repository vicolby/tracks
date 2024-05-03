package handler

import (
	"log/slog"
	"net/http"

	"tracks/pkg/sb"
	"tracks/view/auth"

	"github.com/nedpals/supabase-go"
)

func HandleLoginIndex(w http.ResponseWriter, r *http.Request) error {
	return auth.Login().Render(r.Context(), w)
}

func HandleLogin(w http.ResponseWriter, r *http.Request) error {
	credentials := supabase.UserCredentials{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	resp, err := sb.Client.Auth.SignIn(r.Context(), credentials)

	if err != nil {
		slog.Error("login error", err)
		return auth.LoginForm(credentials, auth.LoginErrors{
			InvalidCredentials: "Invalid credentials",
		}).Render(r.Context(), w)
	}

	cookie := &http.Cookie{
		Name:     "access_token",
		Value:    resp.AccessToken,
		MaxAge:   86400,
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
	}

	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", http.StatusSeeOther)

	return nil
}
