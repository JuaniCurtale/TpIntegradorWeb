package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func handlerSacarTurno(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	//Se chequea si el path no es el indicado, de no serlo la pagina arroja error 404
	if r.URL.Path != "/sacarTurno" {
		http.NotFound(w, r)
		return
	}

	//Se sirve el HTML index.html "Se enlaza"
	http.ServeFile(w, r, "templates/sacarTurno.html")
}

func handlerDescrip(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	//Se chequea si el path no es el indicado, de no serlo la pagina arroja error 404
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	http.ServeFile(w, r, "templates/index.html")
}

func handlerAbout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	//Se chequea si el path no es el indicado, de no serlo la pagina arroja error 404
	if r.URL.Path != "/aboutUs" {
		http.NotFound(w, r)
		return
	}

	http.ServeFile(w, r, "templates/aboutUs.html")
}

type Turno struct {
	Nombre   string
	Telefono string
	Email    string
	Servicio string
	Barbero  string
	Fecha    string
	Hora     string
	Notas    string
	Acepta   string
}

func handlerFormsPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	turno := Turno{
		Nombre:   r.FormValue("nombre"),
		Telefono: r.FormValue("telefono"),
		Email:    r.FormValue("email"),
		Servicio: r.FormValue("servicio"),
		Barbero:  r.FormValue("barbero"),
		Fecha:    r.FormValue("fecha"),
		Hora:     r.FormValue("hora"),
		Notas:    r.FormValue("notas"),
		Acepta:   r.FormValue("acepta_politicas"),
	}

	tmpl, err := template.ParseFiles("templates/confirmacion.html")
	if err != nil {
		http.Error(w, "Error cargando plantilla", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, turno)
}

func main() {
	http.HandleFunc("/", handlerDescrip)
	http.HandleFunc("/formsPost", handlerFormsPost)
	http.HandleFunc("/aboutUs", handlerAbout)
	http.HandleFunc("/sacarTurno", handlerSacarTurno)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	port := ":8080"
	fmt.Printf("Servidor escuchando en http://localhost%s\n", port)

	err := http.ListenAndServe(port, nil)

	if err != nil {
		fmt.Printf("Error al iniciar el servidor")
		panic(err)
	}
}
