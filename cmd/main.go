package main

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	db "tpIntegradorSaideCurtale/db/sqlc"
	database "tpIntegradorSaideCurtale/pkg/database"
	"tpIntegradorSaideCurtale/views"

	"github.com/a-h/templ"
)

var queries *db.Queries

func main() {
	dbconn := database.ConectarDB()
	defer dbconn.Close()

	queries = db.New(dbconn)

	// --- Página principal ---
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		component := views.IndexPage()
		templ.Handler(component).ServeHTTP(w, r)
	})

	// --- HANDLERS ---
	http.HandleFunc("/cliente", handlerClientes)
	http.HandleFunc("/cliente/", handlerClientes)
	http.HandleFunc("/barbero", handlerBarberos)
	http.HandleFunc("/barbero/", handlerBarberos)
	http.HandleFunc("/turno", handlerTurnos)
	http.HandleFunc("/turno/", handlerTurnos)

	log.Println("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handlerClientes(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		clientes, _ := queries.ListClientes(context.Background())
		templ.Handler(views.ClientesPage(clientes)).ServeHTTP(w, r)

	case http.MethodPost:
		r.ParseForm()
		nombre := r.FormValue("nombre")
		apellido := r.FormValue("apellido")
		telefono := r.FormValue("telefono")
		email := r.FormValue("email")

		// 1. Crear cliente
		_, err := queries.CreateCliente(r.Context(), db.CreateClienteParams{
			Nombre:   nombre,
			Apellido: apellido,
			Telefono: telefono,
			Email:    email,
		})
		if err != nil {
			// Intenta usar NotificacionError si existe, sino usa http.Error
			w.Header().Set("HX-Retarget", "#notificaciones")
			if strings.Contains(err.Error(), "duplicate key") {
				// Si tienes el componente NotificacionError:
				// views.NotificacionError("El email ya se encuentra registrado.").Render(r.Context(), w)
				http.Error(w, "El email ya se encuentra registrado.", http.StatusConflict)
				return
			}
			http.Error(w, "Error al guardar el cliente: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// 2. Consultar lista actualizada
		clientes, err := queries.ListClientes(r.Context())
		if err != nil {
			http.Error(w, "Cliente guardado, pero falló la lista.", http.StatusInternalServerError)
			return
		}

		// 3. Renderizar la LISTA COMPLETA (ClientList)
		// Esto reemplazará toda la tabla #clientes-list gracias al hx-target del form
		views.ClientList(clientes).Render(r.Context(), w)
		views.ClienteFormOOB().Render(r.Context(), w)
	case http.MethodDelete:
		// 1. Obtener ID del URL
		idStr := strings.TrimPrefix(r.URL.Path, "/cliente/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "ID inválido", http.StatusBadRequest)
			return
		}

		// 2. Eliminar
		err = queries.DeleteCliente(r.Context(), int32(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// 3. Respuesta vacía para HTMX (elimina la fila de la vista actual)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(""))

	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

func handlerBarberos(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		barberos, _ := queries.ListBarberos(context.Background())
		templ.Handler(views.BarberosPage(barberos)).ServeHTTP(w, r)

	case http.MethodPost:
		r.ParseForm()
		nombre := r.FormValue("nombre")
		apellido := r.FormValue("apellido")
		especialidad := r.FormValue("especialidad")

		// 1. Crear barbero
		_, err := queries.CreateBarbero(r.Context(), db.CreateBarberoParams{
			Nombre:       nombre,
			Apellido:     apellido,
			Especialidad: especialidad,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// 2. Volver a consultar la lista actualizada
		barberos, err := queries.ListBarberos(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// 3. Renderizar la LISTA COMPLETA (BarberList)
		// Esto reemplazará toda la tabla #barberos-list
		views.BarberList(barberos).Render(r.Context(), w)
		views.BarberoFormOOB().Render(r.Context(), w)

	case http.MethodDelete:
		idStr := strings.TrimPrefix(r.URL.Path, "/barbero/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "ID inválido", http.StatusBadRequest)
			return
		}

		err = queries.DeleteBarbero(r.Context(), int32(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(""))

	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

func handlerTurnos(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		clientes, err := queries.ListClientes(context.Background())
		if err != nil {
			http.Error(w, "Error obteniendo clientes: "+err.Error(), http.StatusInternalServerError)
			return
		}

		barberos, err := queries.ListBarberos(context.Background())
		if err != nil {
			http.Error(w, "Error obteniendo barberos: "+err.Error(), http.StatusInternalServerError)
			return
		}

		turnos, err := queries.ListTurnos(context.Background())
		if err != nil {
			http.Error(w, "Error obteniendo turnos: "+err.Error(), http.StatusInternalServerError)
			return
		}

		templ.Handler(views.TurnosPage(turnos, clientes, barberos)).ServeHTTP(w, r)

	case http.MethodPost:
		r.ParseForm()
		idCliente, _ := strconv.Atoi(r.FormValue("id_cliente"))
		idBarbero, _ := strconv.Atoi(r.FormValue("id_barbero"))
		// Se asume formato datetime-local del HTML5
		fechaHora, _ := time.Parse("2006-01-02T15:04", r.FormValue("fechaHora"))
		servicio := r.FormValue("servicio")
		observaciones := r.FormValue("observaciones")

		if fechaHora.Before(time.Now()) {
			w.Header().Set("HX-Retarget", "#notificaciones")
			// views.NotificacionError("La fecha debe ser superior a la actual").Render(r.Context(), w)
			http.Error(w, "La fecha debe ser superior a la actual", http.StatusBadRequest)
			return
		}

		_, err := queries.CreateTurno(r.Context(), db.CreateTurnoParams{
			IDCliente:     int32(idCliente),
			IDBarbero:     int32(idBarbero),
			Fechahora:     fechaHora,
			Servicio:      servicio,
			Observaciones: observaciones,
		})

		if err != nil {
			http.Error(w, "Error al guardar el turno: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// SI ES ÉXITO:
		clientes, err := queries.ListClientes(r.Context())
		if err != nil {
			http.Error(w, "Error obteniendo clientes para el form.", http.StatusInternalServerError)
			return
		}

		barberos, err := queries.ListBarberos(r.Context())
		if err != nil {
			http.Error(w, "Error obteniendo barberos para el form.", http.StatusInternalServerError)
			return
		}
		turnos, err := queries.ListTurnos(r.Context())
		if err != nil {
			http.Error(w, "Turno guardado, pero falló la lista.", http.StatusInternalServerError)
			return
		}

		// 3. Renderizar la LISTA COMPLETA (TurnoList)
		// Esto reemplazará toda la tabla #turnos-list
		views.TurnoList(turnos).Render(r.Context(), w)
		views.TurnoFormOOB(turnos, clientes, barberos).Render(r.Context(), w)

	case http.MethodDelete:
		idStr := strings.TrimPrefix(r.URL.Path, "/turno/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "ID inválido", http.StatusBadRequest)
			return
		}

		err = queries.DeleteTurno(r.Context(), int32(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(""))
	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}
