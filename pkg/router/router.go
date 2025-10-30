package router

import (
	"database/sql"
	"net/http"
	db "tpIntegradorSaideCurtale/db/sqlc"
	"tpIntegradorSaideCurtale/pkg/handlers"
)

func NewRouter(dbconn *sql.DB) http.Handler {
	queries := db.New(dbconn)
	mux := http.NewServeMux() // Enrutador para las distintas rutas del servidor.

	mux.HandleFunc("/", handlers.HandlerDescrip)
	mux.HandleFunc("/aboutUs", handlers.HandlerAbout)

	mux.HandleFunc("/cliente", handlers.ClienteHandler(queries))  // POST, GET form
	mux.HandleFunc("/cliente/", handlers.ClienteHandler(queries)) // ID, PUT, DELETE
	mux.HandleFunc("/barbero", handlers.BarberoHandler(queries))  // POST, GET form
	mux.HandleFunc("/barbero/", handlers.BarberoHandler(queries)) // ID, PUT, DELETE
	mux.HandleFunc("/turno", handlers.TurnoHandler(queries))      // POST, GET all
	mux.HandleFunc("/turno/", handlers.TurnoHandler(queries))     // ID, PUT, DELETE

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	return mux
}
