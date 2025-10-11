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
	"time"

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

func turnoHandler(queries *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := strings.Trim(r.URL.Path, "/")
		parts := strings.Split(path, "/")

		switch r.Method {
		case http.MethodGet:
			// GET /turno → listar todos los turnos
			if len(parts) == 1 && parts[0] == "turno" {
				turnos, err := queries.ListTurnos(r.Context())
				if err != nil {
					http.Error(w, "Error al obtener turnos", http.StatusInternalServerError)
					return
				}
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(turnos)
				return
			}

			// GET /turno/{id} → obtener un turno por ID
			if len(parts) == 2 && parts[0] == "turno" {
				id64, err := strconv.ParseInt(parts[1], 10, 32)
				if err != nil {
					http.Error(w, "ID inválido", http.StatusBadRequest)
					return
				}
				id := int32(id64)

				turno, err := queries.GetTurnoByID(r.Context(), id)
				if err != nil {
					if err == sql.ErrNoRows {
						http.Error(w, "Turno no encontrado", http.StatusNotFound)
					} else {
						http.Error(w, "Error interno", http.StatusInternalServerError)
					}
					return
				}

				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(turno)
				return
			}

		case http.MethodPost:
			// POST /turno → crear un nuevo turno
			if len(parts) == 1 && parts[0] == "turno" {
				var input struct {
					IDCliente     int32     `json:"id_cliente"`
					IDBarbero     int32     `json:"id_barbero"`
					FechaHora     time.Time `json:"fechaHora"`
					Servicio      string    `json:"servicio"`
					Observaciones string    `json:"observaciones"`
				}

				if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
					http.Error(w, "JSON inválido", http.StatusBadRequest)
					return
				}

				// Validaciones básicas //agregar funcion de validacion
				if input.IDCliente <= 0 || input.IDBarbero <= 0 || input.Servicio == "" {
					http.Error(w, "Campos obligatorios faltantes", http.StatusBadRequest)
					return
				}

				// Crear el turno en la base de datos
				turno, err := queries.CreateTurno(r.Context(), db.CreateTurnoParams{
					IDCliente:     input.IDCliente,
					IDBarbero:     input.IDBarbero,
					Fechahora:     input.FechaHora,
					Servicio:      input.Servicio,
					Observaciones: sql.NullString{String: input.Observaciones, Valid: input.Observaciones != ""},
				})

				if err != nil {
					// Si el error es por violar una FK, PostgreSQL devolverá un error que podés traducir
					if strings.Contains(err.Error(), "fk_cliente") {
						http.Error(w, "El cliente no existe", http.StatusBadRequest)
						return
					}
					if strings.Contains(err.Error(), "fk_barbero") {
						http.Error(w, "El barbero no existe", http.StatusBadRequest)
						return
					}
					http.Error(w, "Error al crear turno", http.StatusInternalServerError)
					return
				}

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusCreated)
				json.NewEncoder(w).Encode(turno)
				return
			}

		case http.MethodPut, http.MethodDelete:
			// PUT /turno/{id} o DELETE /turno/{id}
			if len(parts) != 2 || parts[0] != "turno" {
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
					IDCliente     int32     `json:"id_cliente"`
					IDBarbero     int32     `json:"id_barbero"`
					FechaHora     time.Time `json:"fechaHora"`
					Servicio      string    `json:"servicio"`
					Observaciones string    `json:"observaciones"`
				}

				if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
					http.Error(w, "JSON inválido", http.StatusBadRequest)
					return
				}

				// DEBUG: imprimir input recibido
				log.Printf("Input recibido: %+v\n", input)
				log.Printf("Hora parseada: %v\n", input.FechaHora)

				if input.IDCliente <= 0 || input.IDBarbero <= 0 || input.Servicio == "" {
					http.Error(w, "Campos obligatorios faltantes", http.StatusBadRequest)
					return
				}

				turno, err := queries.UpdateTurno(r.Context(), db.UpdateTurnoParams{
					IDTurno:       id,
					IDCliente:     input.IDCliente,
					IDBarbero:     input.IDBarbero,
					Fechahora:     input.FechaHora,
					Servicio:      input.Servicio,
					Observaciones: sql.NullString{String: input.Observaciones, Valid: input.Observaciones != ""},
				})

				if err != nil {
					if strings.Contains(err.Error(), "fk_cliente") {
						http.Error(w, "El cliente no existe", http.StatusBadRequest)
						return
					}
					if strings.Contains(err.Error(), "fk_barbero") {
						http.Error(w, "El barbero no existe", http.StatusBadRequest)
						return
					}
					http.Error(w, "Error al actualizar turno", http.StatusInternalServerError)
					return
				}

				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(turno)
				return
			}

			if r.Method == http.MethodDelete {
				if err := queries.DeleteTurno(r.Context(), id); err != nil {
					http.Error(w, "Error al eliminar turno", http.StatusInternalServerError)
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

func barberoHandler(queries *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := strings.Trim(r.URL.Path, "/")
		parts := strings.Split(path, "/")

		switch r.Method {
		case http.MethodGet:
			if len(parts) == 1 && parts[0] == "barbero" {
				// GET all
				barberos, err := queries.ListBarberos(r.Context())
				if err != nil {
					http.Error(w, "Error interno", http.StatusInternalServerError)
					return
				}
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(barberos)
				return
			} else if len(parts) == 2 && parts[0] == "barbero" {
				// GET by ID
				id64, err := strconv.ParseInt(parts[1], 10, 32)
				if err != nil {
					http.Error(w, "ID inválido", http.StatusBadRequest)
					return
				}
				id := int32(id64)
				barbero, err := queries.GetBarberoByID(r.Context(), id)
				if err != nil {
					if err == sql.ErrNoRows {
						http.Error(w, "Barbero no encontrado", http.StatusNotFound)
					} else {
						http.Error(w, "Error interno", http.StatusInternalServerError)
					}
					return
				}
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(barbero)
				return
			}

		case http.MethodPost:
			if len(parts) == 1 && parts[0] == "barbero" {
				// CREATE
				var input struct {
					Nombre       string `json:"nombre"`
					Especialidad string `json:"especialidad"`
				}
				if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
					http.Error(w, "JSON inválido", http.StatusBadRequest)
					return
				}
				if input.Nombre == "" {
					http.Error(w, "Nombre es obligatorio", http.StatusBadRequest)
					return
				}

				barbero, err := queries.CreateBarbero(r.Context(), db.CreateBarberoParams{
					Nombre:       input.Nombre,
					Especialidad: sql.NullString{String: input.Especialidad, Valid: input.Especialidad != ""},
				})
				if err != nil {
					http.Error(w, "Error al crear barbero", http.StatusInternalServerError)
					return
				}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusCreated)
				json.NewEncoder(w).Encode(barbero)
				return
			}

		case http.MethodPut, http.MethodDelete:
			if len(parts) != 2 || parts[0] != "barbero" {
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
					Nombre       string `json:"nombre"`
					Especialidad string `json:"especialidad"`
				}
				if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
					http.Error(w, "JSON inválido", http.StatusBadRequest)
					return
				}
				if input.Nombre == "" {
					http.Error(w, "Nombre es obligatorio", http.StatusBadRequest)
					return
				}

				barbero, err := queries.UpdateBarbero(r.Context(), db.UpdateBarberoParams{
					IDBarbero:    id,
					Nombre:       input.Nombre,
					Especialidad: sql.NullString{String: input.Especialidad, Valid: input.Especialidad != ""},
				})
				if err != nil {
					http.Error(w, "Error al actualizar barbero", http.StatusInternalServerError)
					return
				}
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(barbero)
				return
			}

			if r.Method == http.MethodDelete {
				if err := queries.DeleteBarbero(r.Context(), id); err != nil {
					http.Error(w, "Error al eliminar barbero", http.StatusInternalServerError)
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
	connStr := "postgres://postgres:admin@barberia_db:5432/barberia?sslmode=disable"

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

	http.HandleFunc("/barbero", barberoHandler(queries))
	http.HandleFunc("/barbero/", barberoHandler(queries))

	http.HandleFunc("/turno", turnoHandler(queries))
	http.HandleFunc("/turno/", turnoHandler(queries))

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
