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

func main() {
	// --- 1. Conectar a la base de datos ---
	dbconn := database.ConectarDB()
	defer dbconn.Close()

	// --- 2. Crear instancia de queries ---
	queries := db.New(dbconn)

	// --- 3. Handler principal ---
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
			return
		}

		// --- Obtener todos los clientes ---
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

		// --- Renderizar la página ---
		component1 := views.IndexPage("Gestion Barber", clientes, barberos, turnos)
		templ.Handler(component1).ServeHTTP(w, r)
	})

	// --- 4. Iniciar servidor ---
	log.Println("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
