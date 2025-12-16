package main

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	db "tpIntegradorSaideCurtale/db/sqlc"
	database "tpIntegradorSaideCurtale/pkg/database"
	"tpIntegradorSaideCurtale/views"

	"github.com/a-h/templ"
	"golang.org/x/crypto/bcrypt"
)

var queries *db.Queries

func main() {
	dbconn := database.ConectarDB()
	defer dbconn.Close()

	queries = db.New(dbconn)

	// --- SEED DE USUARIO ADMIN (SOLO EJECUTAR UNA VEZ O SI NO EXISTE USUARIO) ---
	// Descomenta estas líneas, corre el programa una vez para crear el usuario, y luego vuelve a comentarlas.

	// --- AUTO-CREACIÓN DE USUARIO ADMIN ---
	// Verificamos si existe el usuario "admin"
	_, err := queries.GetUsuario(context.Background(), "admin") // Asegúrate de tener esta query en sqlc

	// Si hay error, asumimos que no existe (o la base está vacía) y lo creamos
	if err != nil {
		pass, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)

		_, errCreate := queries.CreateUsuario(context.Background(), db.CreateUsuarioParams{
			Username:     "admin",
			PasswordHash: string(pass),
		})

		if errCreate != nil {
			log.Printf("Error creando usuario admin por defecto: %v", errCreate)
		} else {
			log.Println("⚠️  Usuario 'admin' creado por defecto (Pass: admin123)")
		}
	}

	// --- RUTAS PÚBLICAS ---
	http.HandleFunc("/login", handlerLogin)
	http.HandleFunc("/logout", handlerLogout)

	// --- MIDDLEWARE ---
	// Wrapper para facilitar el uso del middleware en tus handlers existentes
	protect := func(h http.HandlerFunc) http.Handler {
		return AuthMiddleware(h)
	}

	// --- RUTAS PROTEGIDAS ---

	// Página principal
	http.Handle("/", AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		component := views.IndexPage()
		templ.Handler(component).ServeHTTP(w, r)
	})))

	// Tus handlers existentes envueltos en 'protect'
	http.Handle("/cliente", protect(handlerClientes))
	http.Handle("/cliente/", protect(handlerClientes))
	http.Handle("/barbero", protect(handlerBarberos))
	http.Handle("/barbero/", protect(handlerBarberos))
	http.Handle("/turno", protect(handlerTurnos))
	http.Handle("/turno/", protect(handlerTurnos))

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

		// Primero verificamos si el email ya existe
		existe, err := queries.GetClienteByEmail(r.Context(), email)
		if err == nil && existe.Email == email {
			// Ya existe
			w.Header().Set("HX-Reswap", "none")
			w.WriteHeader(http.StatusOK)
			views.ErrorCliente(nombre, apellido, telefono, email, "El email ya está registrado").Render(r.Context(), w)
			return
		}

		// 1. Crear cliente
		_, err = queries.CreateCliente(r.Context(), db.CreateClienteParams{
			Nombre:   nombre,
			Apellido: apellido,
			Telefono: telefono,
			Email:    email,
		})

		if err != nil {
			http.Error(w, "Error al crear cliente", http.StatusInternalServerError)
			return
		}

		// 2. Consultar lista actualizada
		clientes, err := queries.ListClientes(r.Context())
		if err != nil {
			http.Error(w, "Cliente guardado, pero falló la lista.", http.StatusInternalServerError)
			return
		}

		views.ClienteGuardadoExito(clientes).Render(r.Context(), w)
	case http.MethodDelete:
		idStr := strings.TrimPrefix(r.URL.Path, "/cliente/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "ID inválido", http.StatusBadRequest)
			return
		}

		// 2. Eliminar
		err = queries.DeleteCliente(r.Context(), int32(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// 3. Respuesta vacía para HTMX (elimina la fila de la vista actual)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(""))

	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

func handlerBarberos(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		barberos, _ := queries.ListBarberos(context.Background())
		templ.Handler(views.BarberosPage(barberos)).ServeHTTP(w, r)

	case http.MethodPost:
		r.ParseForm()
		nombre := r.FormValue("nombre")
		apellido := r.FormValue("apellido")
		especialidad := r.FormValue("especialidad")

		// 1. Crear barbero
		_, err := queries.CreateBarbero(r.Context(), db.CreateBarberoParams{
			Nombre:       nombre,
			Apellido:     apellido,
			Especialidad: especialidad,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// 2. Volver a consultar la lista actualizada
		barberos, err := queries.ListBarberos(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		views.BarberoGuardadoExito(barberos).Render(r.Context(), w)
	case http.MethodDelete:
		idStr := strings.TrimPrefix(r.URL.Path, "/barbero/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "ID inválido", http.StatusBadRequest)
			return
		}

		err = queries.DeleteBarbero(r.Context(), int32(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(""))

	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

func handlerTurnos(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
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

		templ.Handler(views.TurnosPage(turnos, clientes, barberos)).ServeHTTP(w, r)

	case http.MethodPost:
		r.ParseForm()
		idCliente, _ := strconv.Atoi(r.FormValue("id_cliente"))
		idBarbero, _ := strconv.Atoi(r.FormValue("id_barbero"))
		fechaHora, _ := time.Parse("2006-01-02T15:04", r.FormValue("fechaHora"))
		servicio := r.FormValue("servicio")
		observaciones := r.FormValue("observaciones")

		if fechaHora.Before(time.Now()) {
			clientes, _ := queries.ListClientes(r.Context())
			barberos, _ := queries.ListBarberos(r.Context())
			w.WriteHeader(http.StatusUnprocessableEntity)
			views.ErrorTurno(
				idCliente,
				idBarbero,
				fechaHora.Format("2006-01-02T15:04"),
				servicio,
				observaciones,
				"La fecha debe ser futura",
				clientes,
				barberos,
			).Render(r.Context(), w)

			return
		}

		nuevoTurno, err := queries.CreateTurno(r.Context(), db.CreateTurnoParams{
			IDCliente:     int32(idCliente),
			IDBarbero:     int32(idBarbero),
			Fechahora:     fechaHora,
			Servicio:      servicio,
			Observaciones: observaciones,
		})

		if err != nil {
			http.Error(w, "Error al guardar el turno: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// SI ES ÉXITO:
		clientes, err := queries.ListClientes(r.Context())
		if err != nil {
			http.Error(w, "Error obteniendo clientes para el form.", http.StatusInternalServerError)
			return
		}

		barberos, err := queries.ListBarberos(r.Context())
		if err != nil {
			http.Error(w, "Error obteniendo barberos para el form.", http.StatusInternalServerError)
			return
		}
		turnos, err := queries.ListTurnos(r.Context())
		if err != nil {
			http.Error(w, "Turno guardado, pero falló la lista.", http.StatusInternalServerError)
			return
		}
		views.TurnoGuardadoExito(turnos, clientes, barberos, nuevoTurno).Render(r.Context(), w)
	case http.MethodDelete:
		idStr := strings.TrimPrefix(r.URL.Path, "/turno/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "ID inválido", http.StatusBadRequest)
			return
		}

		err = queries.DeleteTurno(r.Context(), int32(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(""))
	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}
