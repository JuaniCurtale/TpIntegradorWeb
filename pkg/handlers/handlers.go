package handlers

import (
	"context"
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	db "tpIntegradorSaideCurtale/db/sqlc"
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

func HandlerSacarTurno(w http.ResponseWriter, r *http.Request) { //FORMULARIO PARA SACAR TURNO
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if r.URL.Path != "/sacarTurno" {
		http.NotFound(w, r)
		return
	}
	http.ServeFile(w, r, "templates/sacarTurno.html")
}

func HandlerFormsPost(w http.ResponseWriter, r *http.Request, dbConn *sql.DB) { //FORMULARIO ENVIADO
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	turno := Turno{
		Nombre:   r.FormValue("nombre"),
		Apellido: r.FormValue("apellido"),
		Telefono: r.FormValue("telefono"),
		Email:    r.FormValue("email"),
		Servicio: r.FormValue("servicio"),
		Barbero:  r.FormValue("barbero"),
		Fecha:    r.FormValue("fecha"),
		Hora:     r.FormValue("hora"),
		Notas:    r.FormValue("notas"),
		Acepta:   r.FormValue("acepta_politicas"),
	}

	query := `
        INSERT INTO cliente (nombre, apellido, telefono, email)
        VALUES ($1, $2, $3, $4)
		ON CONFLICT(nombre, apellido, telefono, email) DO NOTHING;
    `
	_, err := dbConn.Exec(query, turno.Nombre, turno.Apellido, turno.Telefono, turno.Email)
	if err != nil {
		http.Error(w, "Error al guardar turno en la base", http.StatusInternalServerError)
		log.Printf("DB error: %v", err)
		return
	}

	tmpl, err := template.ParseFiles("templates/confirmacion.html")
	if err != nil {
		http.Error(w, "Error cargando plantilla", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, turno); err != nil {
		http.Error(w, "Error al renderizar plantilla", http.StatusInternalServerError)
		log.Printf("Error ejecutando template: %v", err)
	}
}

func HandlerGetClientes(w http.ResponseWriter, r *http.Request, dbConn *sql.DB) { //LISTA DE TODOS LOS CLIENTES
	queries := db.New(dbConn) // corregido: uso del paquete 'db'
	clientes, err := queries.ListClientes(context.Background())
	if err != nil {
		log.Printf("Error al listar clientes: %v", err)
		http.Error(w, fmt.Sprintf("Error al listar clientes: %v", err), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/clientes.html")
	if err != nil {
		log.Printf("Error al cargar la plantilla: %v", err)
		http.Error(w, "Error interno del servidor", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if err := tmpl.Execute(w, clientes); err != nil {
		log.Printf("Error al renderizar la plantilla: %v", err)
		http.Error(w, "Error interno del servidor", http.StatusInternalServerError)
	}
}

func HandlerRegistrarCliente(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/registrarCliente.html")
}
