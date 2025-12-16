package main

import (
	"context"
	"net/http"
	"time"
	"tpIntegradorSaideCurtale/views" // Asegúrate que este import coincida con tu go.mod

	"github.com/a-h/templ"
	"golang.org/x/crypto/bcrypt"
)

// Middleware de Autenticación
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Ignorar verificación para login y recursos estáticos si los tuvieras
		if r.URL.Path == "/login" {
			next.ServeHTTP(w, r)
			return
		}

		_, err := r.Cookie("session_token")
		if err != nil {
			// No hay cookie, redirigir a login
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		// Si hay cookie, pasa
		next.ServeHTTP(w, r)
	})
}

func handlerLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		templ.Handler(views.LoginPage("")).ServeHTTP(w, r)
		return
	}

	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Usamos la variable global 'queries' definida en main.go
		user, err := queries.GetUsuario(context.Background(), username)
		if err != nil {
			templ.Handler(views.LoginPage("Usuario no encontrado")).ServeHTTP(w, r)
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
		if err != nil {
			templ.Handler(views.LoginPage("Contraseña incorrecta")).ServeHTTP(w, r)
			return
		}

		// Crear cookie
		http.SetCookie(w, &http.Cookie{
			Name:     "session_token",
			Value:    "user_logged_in", // En producción usa un token seguro
			Expires:  time.Now().Add(24 * time.Hour),
			HttpOnly: true,
			Path:     "/",
		})

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func handlerLogout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   "",
		Expires: time.Now().Add(-1 * time.Hour),
		Path:    "/",
	})
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
