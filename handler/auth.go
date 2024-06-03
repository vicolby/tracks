package handler

import (
	"fmt"
	"log/slog"
	"net/http"
	"net/mail"

	"tracks/pkg/sb"
	"tracks/view/auth"

	"github.com/nedpals/supabase-go"
)

func HandleLoginIndex(w http.ResponseWriter, r *http.Request) error {
	return auth.Login().Render(r.Context(), w)
}

func HandleSignUpIndex(w http.ResponseWriter, r *http.Request) error {
	return auth.Signup().Render(r.Context(), w)
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

func HandleSignup(w http.ResponseWriter, r *http.Request) error {
	err := ValidateCredentials(r.FormValue("password"), r.FormValue("password_confirm"), r.FormValue("email"))
	credentials := auth.SignupParams{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	if err != nil {
		slog.Error("signup error", err)
		return auth.SignupForm(credentials, auth.SignupErrors{
			InvalidCredentials: err.Error(),
		}).Render(r.Context(), w)
	}

	_, err = sb.Client.Auth.SignUp(r.Context(), supabase.UserCredentials{
		Email:    credentials.Email,
		Password: credentials.Password,
	})

	if err != nil {
		return err
	}

	return auth.SignupConfirmation().Render(r.Context(), w)
}

func ValidateCredentials(password string, confirm string, email string) error {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return fmt.Errorf("Wrong email format")
	}

	if password != confirm {
		return fmt.Errorf("Confirmation password does not match the password")
	} else if len(password) < 6 {
		return fmt.Errorf("Password length must be at least 6 characters")
	}

	return nil
}
