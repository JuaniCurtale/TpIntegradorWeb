package router

import (
	"database/sql"
	"net/http"
	db "tpIntegradorSaideCurtale/db/sqlc"
	"tpIntegradorSaideCurtale/pkg/handlers"
)

func NewRouter(dbconn *sql.DB) http.Handler {
	queries := db.New(dbconn)
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.HandlerDescrip)
	mux.HandleFunc("/aboutUs", handlers.HandlerAbout)

	mux.HandleFunc("/cliente", handlers.ClienteHandler(queries))  // POST, GET form
	mux.HandleFunc("/cliente/", handlers.ClienteHandler(queries)) // ID, PUT, DELETE
	mux.HandleFunc("/barbero", handlers.BarberoHandler(queries))  // POST, GET form
	mux.HandleFunc("/barbero/", handlers.BarberoHandler(queries)) // ID, PUT, DELETE
	mux.HandleFunc("/turno", handlers.TurnoHandler(queries))      // POST, GET all
	mux.HandleFunc("/turno/", handlers.TurnoHandler(queries))     // ID, PUT, DELETE

	// API routes for listing all entities
	mux.HandleFunc("/api/clientes", handlers.ListClientesHandler(queries))
	mux.HandleFunc("/api/barberos", handlers.ListBarberosHandler(queries))
	mux.HandleFunc("/api/turnos", handlers.ListTurnosHandler(queries))
	mux.HandleFunc("/api/turnos/cliente/", handlers.ListTurnosByClienteIDHandler(queries))
	mux.HandleFunc("/api/turnos/barbero/", handlers.ListTurnosByBarberoIDHandler(queries))

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	return mux
}
