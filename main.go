package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	db "tpIntegradorSaideCurtale/db/sqlc"

	_ "github.com/lib/pq"
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
func clienteHandler(queries *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := strings.Trim(r.URL.Path, "/")
		parts := strings.Split(path, "/")

		switch r.Method {
		case http.MethodGet:
			if len(parts) == 1 && parts[0] == "cliente" {
				// GET all
				clientes, err := queries.ListClientes(r.Context())
				if err != nil {
					http.Error(w, "Error interno", http.StatusInternalServerError)
					return
				}
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(clientes)
				return
			} else if len(parts) == 2 && parts[0] == "cliente" {
				// GET by ID
				id64, err := strconv.ParseInt(parts[1], 10, 32)
				if err != nil {
					http.Error(w, "ID inválido", http.StatusBadRequest)
					return
				}
				id := int32(id64)
				cliente, err := queries.GetClienteByID(r.Context(), id)
				if err != nil {
					if err == sql.ErrNoRows {
						http.Error(w, "Cliente no encontrado", http.StatusNotFound)
					} else {
						http.Error(w, "Error interno", http.StatusInternalServerError)
					}
					return
				}
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(cliente)
				return
			}

		case http.MethodPost:
			if len(parts) == 1 && parts[0] == "cliente" {
				// CREATE
				var input struct {
					Nombre   string `json:"nombre"`
					Apellido string `json:"apellido"`
					Telefono string `json:"telefono"`
				}
				if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
					http.Error(w, "JSON inválido", http.StatusBadRequest)
					return
				}
				if input.Nombre == "" || input.Apellido == "" {
					http.Error(w, "Nombre y apellido son obligatorios", http.StatusBadRequest)
					return
				}

				cliente, err := queries.CreateCliente(r.Context(), db.CreateClienteParams{
					Nombre:   input.Nombre,
					Apellido: input.Apellido,
					Telefono: sql.NullString{String: input.Telefono, Valid: input.Telefono != ""},
				})
				if err != nil {
					http.Error(w, "Error al crear cliente", http.StatusInternalServerError)
					return
				}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusCreated)
				json.NewEncoder(w).Encode(cliente)
				return
			}

		case http.MethodPut, http.MethodDelete:
			if len(parts) != 2 || parts[0] != "cliente" {
				http.NotFound(w, r)
				return
			}
			id64, err := strconv.ParseInt(parts[1], 10, 32)
			if err != nil {
				http.Error(w, "ID inválido", http.StatusBadRequest)
				return
			}
			id := int32(id64)

			if r.Method == http.MethodPut {
				var input struct {
					Nombre   string `json:"nombre"`
					Apellido string `json:"apellido"`
					Telefono string `json:"telefono"`
				}
				if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
					http.Error(w, "JSON inválido", http.StatusBadRequest)
					return
				}
				if input.Nombre == "" || input.Apellido == "" {
					http.Error(w, "Nombre y apellido son obligatorios", http.StatusBadRequest)
					return
				}
				cliente, err := queries.UpdateCliente(r.Context(), db.UpdateClienteParams{
					IDCliente: id,
					Nombre:    input.Nombre,
					Apellido:  input.Apellido,
					Telefono:  sql.NullString{String: input.Telefono, Valid: input.Telefono != ""},
				})
				if err != nil {
					http.Error(w, "Error al actualizar cliente", http.StatusInternalServerError)
					return
				}
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(cliente)
				return
			}

			if r.Method == http.MethodDelete {
				if err := queries.DeleteCliente(r.Context(), id); err != nil {
					http.Error(w, "Error al eliminar cliente", http.StatusInternalServerError)
					return
				}
				w.WriteHeader(http.StatusNoContent)
				return
			}

		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
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
	http.HandleFunc("/cliente", clienteHandler(queries))  // POST, GET all
	http.HandleFunc("/cliente/", clienteHandler(queries)) // GET by ID, PUT, DELETE

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
