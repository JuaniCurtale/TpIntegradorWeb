package handlers

import (
	"net/http"
)

type Turno struct {
	Nombre   string
	Apellido string
	Telefono string
	Email    string
	Servicio string
	Barbero  string
	Fecha    string
	Hora     string
	Notas    string
	Acepta   string
}

func HandlerDescrip(w http.ResponseWriter, r *http.Request) { // INICIO
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	http.ServeFile(w, r, "templates/index.html")
}

func HandlerAbout(w http.ResponseWriter, r *http.Request) { //SOBRE NOSOTROS
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if r.URL.Path != "/aboutUs" {
		http.NotFound(w, r)
		return
	}
	http.ServeFile(w, r, "templates/aboutUs.html")
}
