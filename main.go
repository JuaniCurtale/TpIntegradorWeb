package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	db "tpIntegradorSaideCurtale/db/sqlc"

	_ "github.com/lib/pq"
	// Ruta del package generado por sqlc
)

func handlerSacarTurno(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	//Se chequea si el path no es el indicado, de no serlo la pagina arroja error 404
	if r.URL.Path != "/sacarTurno" {
		http.NotFound(w, r)
		return
	}

	//Se sirve el HTML index.html "Se enlaza"
	http.ServeFile(w, r, "templates/sacarTurno.html")
}

func handlerDescrip(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	//Se chequea si el path no es el indicado, de no serlo la pagina arroja error 404
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	http.ServeFile(w, r, "templates/index.html")
}

func handlerAbout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	//Se chequea si el path no es el indicado, de no serlo la pagina arroja error 404
	if r.URL.Path != "/aboutUs" {
		http.NotFound(w, r)
		return
	}

	http.ServeFile(w, r, "templates/aboutUs.html")
}

type Turno struct {
	Nombre   string
	Telefono string
	Email    string
	Servicio string
	Barbero  string
	Fecha    string
	Hora     string
	Notas    string
	Acepta   string
}

func handlerFormsPost(w http.ResponseWriter, r *http.Request, queries *db.Queries) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	turno := Turno{
		Nombre:   r.FormValue("nombre"),
		Telefono: r.FormValue("telefono"),
		Email:    r.FormValue("email"),
		Servicio: r.FormValue("servicio"),
		Barbero:  r.FormValue("barbero"),
		Fecha:    r.FormValue("fecha"),
		Hora:     r.FormValue("hora"),
		Notas:    r.FormValue("notas"),
		Acepta:   r.FormValue("acepta_politicas"),
	}

	// Ejemplo: guardar en la base de datos usando queries
	// _, err := queries.CreateTurno(r.Context(), sqlc.CreateTurnoParams{ ... })
	// if err != nil { ... }

	tmpl, err := template.ParseFiles("templates/confirmacion.html")
	if err != nil {
		http.Error(w, "Error cargando plantilla", http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, turno); err != nil {
		http.Error(w, "Error renderizando plantilla", http.StatusInternalServerError)
		return
	}
}

func ConnectDB() *db.Queries {
	connStr := "postgres://postgres:admin@localhost:5432/barberia?sslmode=disable"
	dbConn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error al conectar con la base de datos: %v", err)
	}

	if err := dbConn.Ping(); err != nil {
		log.Fatalf("No se pudo conectar a la base de datos: %v", err)
	}

	fmt.Println("Conexión a la base de datos exitosa ✅")

	return db.New(dbConn)
}

func main() {
	// Obtenemos la instancia de queries
	queries := ConnectDB()

	// Rutas HTTP
	http.HandleFunc("/", handlerDescrip)
	http.HandleFunc("/formsPost", func(w http.ResponseWriter, r *http.Request) {
		handlerFormsPost(w, r, queries)
	})
	http.HandleFunc("/aboutUs", handlerAbout)
	http.HandleFunc("/sacarTurno", handlerSacarTurno)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	port := ":8080"
	fmt.Printf("Servidor escuchando en http://localhost%s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
