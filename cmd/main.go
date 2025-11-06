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
	// --- 1. Conectar a la base de datos ---
	dbconn := database.ConectarDB()
	defer dbconn.Close()

	// --- 2. Crear instancia de queries ---
	queries = db.New(dbconn)

	// --- 3. Handler principal ---
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		component := views.IndexPage()
		templ.Handler(component).ServeHTTP(w, r)
	})

	http.HandleFunc("/clientes", func(w http.ResponseWriter, r *http.Request) {
		clientes, _ := queries.ListClientes(context.Background())
		templ.Handler(views.ClientesPage(clientes)).ServeHTTP(w, r)
	})

	http.HandleFunc("/barberos", func(w http.ResponseWriter, r *http.Request) {
		barberos, _ := queries.ListBarberos(context.Background())
		templ.Handler(views.BarberosPage(barberos)).ServeHTTP(w, r)
	})

	http.HandleFunc("/turnos", func(w http.ResponseWriter, r *http.Request) {
		turnos, _ := queries.ListTurnos(context.Background())
		templ.Handler(views.TurnosPage(turnos)).ServeHTTP(w, r)
	})

	// --- 4. Iniciar servidor ---
	log.Println("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handlerTurnoForm(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error al procesar formulario", http.StatusBadRequest)
		return
	}

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
		http.Error(w, "Error al guardar turno", http.StatusInternalServerError)
		return
	}

	turnos, err := queries.ListTurnos(r.Context())
	if err != nil {
		http.Error(w, "Turno creado, pero no se pudo refrescar la lista", http.StatusInternalServerError)
		return
	}

	views.TurnoList(turnos).Render(r.Context(), w)
}
