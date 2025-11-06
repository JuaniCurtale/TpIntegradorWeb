package main

import (
	"context"
	"log"
	"net/http"

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

	// --- CLIENTES ---
	http.HandleFunc("/cliente", handlerClientes)

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

		queries.CreateCliente(r.Context(), db.CreateClienteParams{
			Nombre:   nombre,
			Apellido: apellido,
			Telefono: telefono,
			Email:    email,
		})

		http.Redirect(w, r, "/cliente", http.StatusSeeOther)
	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}
