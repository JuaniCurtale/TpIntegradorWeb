package main

import (
	"fmt"
	"net/http"
)

func handlerDescrip(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	http.ServeFile(w, r, "index.html")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
}

func main() {
	http.HandleFunc("/", handlerDescrip)

	port := ":8080"
	fmt.Printf("Servidor escuchando en http://localhost%s\n", port)

	err := http.ListenAndServe(port, nil)

	if err != nil {
		fmt.Printf("Error al iniciar el servidor")
		panic(err)
	}
}
