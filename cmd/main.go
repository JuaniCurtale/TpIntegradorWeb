package main

import (
	"fmt"
	"log"
	"net/http"
	"tpIntegradorSaideCurtale/pkg/database"
	"tpIntegradorSaideCurtale/pkg/router"
)

func main() {
	dbconn := database.ConectarDB()
	defer dbconn.Close()

	r := router.NewRouter(dbconn)

	port := ":8080"
	fmt.Printf("Servidor escuchando en http://localhost%s\n", port)

	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
