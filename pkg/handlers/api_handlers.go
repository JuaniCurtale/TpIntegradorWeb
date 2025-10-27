package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
	db "tpIntegradorSaideCurtale/db/sqlc"
	"tpIntegradorSaideCurtale/logic"
)

func ClienteHandler(queries *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := strings.Trim(r.URL.Path, "/")
		parts := strings.Split(path, "/")

		switch r.Method {
		case http.MethodGet:
			if len(parts) == 1 && parts[0] == "cliente" {
				w.Header().Set("Content-Type", "text/html; charset=utf-8")
				http.ServeFile(w, r, "templates/cliente.html")
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
					Email    string `json:"email"`
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
					Email:    sql.NullString{String: input.Email, Valid: input.Email != ""},
				})
				if err != nil {
					fmt.Println("CreateCliente error:", err)
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
					Email    string `json:"email"`
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
					Email:     sql.NullString{String: input.Email, Valid: input.Email != ""},
				})
				if err != nil {
					fmt.Println("UpdateCliente error:", err)
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

func BarberoHandler(queries *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := strings.Trim(r.URL.Path, "/")
		parts := strings.Split(path, "/")

		switch r.Method {
		case http.MethodGet:
			if len(parts) == 1 && parts[0] == "barbero" {
				w.Header().Set("Content-Type", "text/html; charset=utf-8")
				http.ServeFile(w, r, "templates/barbero.html")
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
					Apellido     string `json:"apellido"`
					Especialidad string `json:"especialidad"`
				}
				if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
					http.Error(w, "JSON inválido", http.StatusBadRequest)
					return
				}
				if input.Nombre == "" || input.Apellido == "" {
					http.Error(w, "Nombre y apellido son obligatorios", http.StatusBadRequest)
					return
				}

				barbero, err := queries.CreateBarbero(r.Context(), db.CreateBarberoParams{
					Nombre:       input.Nombre,
					Apellido:     input.Apellido,
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
					Apellido     string `json:"apellido"`
					Especialidad string `json:"especialidad"`
				}
				if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
					http.Error(w, "JSON inválido", http.StatusBadRequest)
					return
				}
				if input.Nombre == "" || input.Apellido == "" {
					http.Error(w, "Nombre y apellido son obligatorios", http.StatusBadRequest)
					return
				}
				barbero, err := queries.UpdateBarbero(r.Context(), db.UpdateBarberoParams{
					IDBarbero:    id,
					Nombre:       input.Nombre,
					Apellido:     input.Apellido,
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

func TurnoHandler(queries *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := strings.Trim(r.URL.Path, "/")
		parts := strings.Split(path, "/")

		switch r.Method {
		case http.MethodGet:
			if len(parts) == 1 && parts[0] == "turno" {
				w.Header().Set("Content-Type", "text/html; charset=utf-8")
				http.ServeFile(w, r, "templates/turno.html")
				return

			} else if len(parts) == 2 && parts[0] == "turno" {
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
			if len(parts) == 1 && parts[0] == "turno" {
				// CREATE
				var input struct {
					Nombre        string         `json:"nombre"`
					Telefono      string         `json:"telefono"`
					Email         string         `json:"email"`
					IDBarbero     int32          `json:"id_barbero"`
					Fechahora     time.Time      `json:"fechahora"`
					Servicio      string         `json:"servicio"`
					Observaciones sql.NullString `json:"observaciones"`
				}
				if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
					fmt.Println("json.NewDecoder error:", err)
					http.Error(w, "JSON inválido", http.StatusBadRequest)
					return
				}
				if input.Fechahora.IsZero() || input.Servicio == "" {
					http.Error(w, "Fechahora y servicio son obligatorios", http.StatusBadRequest)
					return
				}

				// Get or create a client
				cliente, err := queries.GetClienteByEmail(r.Context(), sql.NullString{String: input.Email, Valid: input.Email != ""})
				if err != nil {
					if err == sql.ErrNoRows {
						// Client does not exist, create a new one
						nombreCompleto := strings.SplitN(input.Nombre, " ", 2)
						nombre := nombreCompleto[0]
						apellido := ""
						if len(nombreCompleto) > 1 {
							apellido = nombreCompleto[1]
						}

						cliente, err = queries.CreateCliente(r.Context(), db.CreateClienteParams{
							Nombre:   nombre,
							Apellido: apellido,
							Telefono: sql.NullString{String: input.Telefono, Valid: input.Telefono != ""},
							Email:    sql.NullString{String: input.Email, Valid: input.Email != ""},
						})
						if err != nil {
							fmt.Println("CreateCliente error:", err)
							http.Error(w, "Error al crear cliente", http.StatusInternalServerError)
							return
						}
					} else {
						// Other error
						fmt.Println("GetClienteByEmail error:", err)
						http.Error(w, "Error al buscar cliente", http.StatusInternalServerError)
						return
					}
				}

				turnos, err := queries.ListTurnos(r.Context())
				if err != nil {
					http.Error(w, "Error interno", http.StatusInternalServerError)
					return
				}

				nuevoTurno := db.Turno{
					IDCliente:     cliente.IDCliente,
					IDBarbero:     input.IDBarbero,
					Fechahora:     input.Fechahora,
					Servicio:      input.Servicio,
					Observaciones: input.Observaciones,
				}

				if !logic.HorarioValido(nuevoTurno.Fechahora) {
					http.Error(w, "Horario inválido, por favor pida con una hora de anticipación.", http.StatusBadRequest)
					return
				}
				if !logic.PuedeReservar(nuevoTurno.IDCliente, turnos) {
					http.Error(w, "El cliente ya tiene un turno pendiente.", http.StatusBadRequest)
					return
				}
				if !logic.BarberoDisponible(nuevoTurno.IDBarbero, nuevoTurno.Fechahora, turnos) {
					http.Error(w, "El barbero no está disponible en ese horario.", http.StatusBadRequest)
					return
				}

				turno, err := queries.CreateTurno(r.Context(), db.CreateTurnoParams{
					IDCliente:     cliente.IDCliente,
					IDBarbero:     input.IDBarbero,
					Fechahora:     input.Fechahora,
					Servicio:      input.Servicio,
					Observaciones: input.Observaciones,
				})
				if err != nil {
					http.Error(w, "Error al crear turno", http.StatusInternalServerError)
					return
				}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusCreated)
				json.NewEncoder(w).Encode(turno)
				return
			}

		case http.MethodPut:
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

			var input struct {
				IDCliente     int32          `json:"id_cliente"`
				IDBarbero     int32          `json:"id_barbero"`
				Fechahora     time.Time      `json:"fechahora"`
				Servicio      string         `json:"servicio"`
				Observaciones sql.NullString `json:"observaciones"`
			}
			if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
				http.Error(w, "JSON inválido", http.StatusBadRequest)
				return
			}
			if input.Fechahora.IsZero() || input.Servicio == "" {
				http.Error(w, "Fechahora y servicio son obligatorios", http.StatusBadRequest)
				return
			}

			turno, err := queries.UpdateTurno(r.Context(), db.UpdateTurnoParams{
				IDTurno:       id,
				IDCliente:     input.IDCliente,
				IDBarbero:     input.IDBarbero,
				Fechahora:     input.Fechahora,
				Servicio:      input.Servicio,
				Observaciones: input.Observaciones,
			})
			if err != nil {
				http.Error(w, "Error al actualizar turno", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(turno)
			return

		case http.MethodDelete:
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

			if err := queries.DeleteTurno(r.Context(), id); err != nil {
				http.Error(w, "Error al eliminar turno", http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusNoContent)
			return

		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	}
}

func NullableString(ns sql.NullString) interface{} {
	if ns.Valid {
		return ns.String
	}
	return nil
}

func ListClientesHandler(queries *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		clientes, err := queries.ListClientes(r.Context())
		if err != nil {
			http.Error(w, "Error interno del servidor", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(clientes)
	}
}

func ListBarberosHandler(queries *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		barberos, err := queries.ListBarberos(r.Context())
		if err != nil {
			http.Error(w, "Error interno del servidor", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(barberos)
	}
}

func ListTurnosHandler(queries *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		turnos, err := queries.ListTurnos(r.Context())
		if err != nil {
			http.Error(w, "Error interno del servidor", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(turnos)
	}
}

func ListTurnosByClienteIDHandler(queries *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := strings.Trim(r.URL.Path, "/")
		parts := strings.Split(path, "/")

		if len(parts) != 4 || parts[2] != "cliente" {
			http.NotFound(w, r)
			return
		}

		id64, err := strconv.ParseInt(parts[3], 10, 32)
		if err != nil {
			http.Error(w, "ID de cliente inválido", http.StatusBadRequest)
			return
		}
		id := int32(id64)

		turnos, err := queries.GetTurnosByClienteID(r.Context(), id)
		if err != nil {
			if err == sql.ErrNoRows {
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode([]db.Turno{})
			} else {
				http.Error(w, "Error interno del servidor", http.StatusInternalServerError)
			}
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(turnos)
	}
}

func ListTurnosByBarberoIDHandler(queries *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := strings.Trim(r.URL.Path, "/")
		parts := strings.Split(path, "/")

		if len(parts) != 4 || parts[2] != "barbero" {
			http.NotFound(w, r)
			return
		}

		id64, err := strconv.ParseInt(parts[3], 10, 32)
		if err != nil {
			http.Error(w, "ID de barbero inválido", http.StatusBadRequest)
			return
		}
		id := int32(id64)

		turnos, err := queries.GetTurnosByBarberoID(r.Context(), id)
		if err != nil {
			if err == sql.ErrNoRows {
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode([]db.Turno{})
			} else {
				http.Error(w, "Error interno del servidor", http.StatusInternalServerError)
			}
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(turnos)
	}
}
