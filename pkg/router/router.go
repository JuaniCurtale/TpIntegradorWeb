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
	mux.HandleFunc("/sacarTurno", handlers.HandlerSacarTurno)
	mux.HandleFunc("/formsPost", func(w http.ResponseWriter, r *http.Request) { handlers.HandlerFormsPost(w, r, dbconn) })
	mux.HandleFunc("/clientes", func(w http.ResponseWriter, r *http.Request) { handlers.HandlerGetClientes(w, r, dbconn) }) //lista de clientes

	mux.HandleFunc("/cliente", handlers.ClienteHandler(queries))  //POST, GET ALL
	mux.HandleFunc("/cliente/", handlers.ClienteHandler(queries)) // ID, PUT, DELETE
	mux.HandleFunc("/barbero", handlers.BarberoHandler(queries))  //POST, GET ALL
	mux.HandleFunc("/barbero/", handlers.BarberoHandler(queries)) // ID, PUT, DELETE
	mux.HandleFunc("/turno", handlers.TurnoHandler(queries))      //POST, GET ALL
	mux.HandleFunc("/turno/", handlers.TurnoHandler(queries))     // ID, PUT, DELETE
	mux.HandleFunc("/registrarCliente", handlers.HandlerRegistrarCliente)
	mux.HandleFunc("/registrarBarbero", handlers.HandlerRegistrarBarbero)

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	return mux
}
