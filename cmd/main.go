package main

import (
	"context"
	"log"
	"net/http"
	"strconv"
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
		component := views.IndexPage()
		templ.Handler(component).ServeHTTP(w, r)
	})

	// --- HANDLERS ---
	http.HandleFunc("/cliente", handlerClientes)
	http.HandleFunc("/barbero", handlerBarberos)
	http.HandleFunc("/turno", handlerTurnos)

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
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// 2. Consultar lista actualizada
		clientes, err := queries.ListClientes(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// 3. Devolver SOLO el componente de lista
		views.ClientListRows(clientes).Render(r.Context(), w)

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

		// 3. Devolver SOLO el componente BarberList (fragmento)
		views.BarberList(barberos).Render(r.Context(), w)

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
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Error al procesar formulario", http.StatusBadRequest)
			return
		}

		idClienteStr := r.FormValue("id_cliente")
		idBarberoStr := r.FormValue("id_barbero")
		fechaHoraStr := r.FormValue("fechaHora")
		servicio := r.FormValue("servicio")
		observaciones := r.FormValue("observaciones")

		idCliente, err := strconv.ParseInt(idClienteStr, 10, 32)
		if err != nil {
			http.Error(w, "ID de cliente inválido", http.StatusBadRequest)
			return
		}

		idBarbero, err := strconv.ParseInt(idBarberoStr, 10, 32)
		if err != nil {
			http.Error(w, "ID de barbero inválido", http.StatusBadRequest)
			return
		}

		fechaHora, err := time.Parse("2006-01-02T15:04", fechaHoraStr)
		if err != nil {
			http.Error(w, "Formato de fecha/hora inválido", http.StatusBadRequest)
			return
		}

		_, err = queries.CreateTurno(r.Context(), db.CreateTurnoParams{
			IDCliente:     int32(idCliente),
			IDBarbero:     int32(idBarbero),
			Fechahora:     fechaHora,
			Servicio:      servicio,
			Observaciones: observaciones,
		})
		if err != nil {
			http.Error(w, "Error al guardar turno: "+err.Error(), http.StatusInternalServerError)
			return
		}

	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}
