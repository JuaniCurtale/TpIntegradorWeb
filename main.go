package main

import (
	"fmt"
	"net/http"
)

func handlerDescrip(w http.ResponseWriter, r *http.Request) {

	//Se chequea si el path no es el indicado, de no serlo la pagina arroja error 404
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	//Se establecera la cabecera
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	//Se sirve el HTML index.html "Se enlaza"
	http.ServeFile(w, r, "templates/index.html")

}

func handlerAbout(w http.ResponseWriter, r *http.Request) {

	//Se chequea si el path no es el indicado, de no serlo la pagina arroja error 404
	if r.URL.Path != "/about" {
		http.NotFound(w, r)
		return
	}

	http.ServeFile(w, r, "templates/about.html")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

}

func main() {
	http.HandleFunc("/about", handlerAbout)
	http.HandleFunc("/", handlerDescrip)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	port := ":8080"
	fmt.Printf("Servidor escuchando en http://localhost%s\n", port)

	err := http.ListenAndServe(port, nil)

	if err != nil {
		fmt.Printf("Error al iniciar el servidor")
		panic(err)
	}
}
