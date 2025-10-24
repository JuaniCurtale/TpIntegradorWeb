# 📌 TP Integrador Web

Este proyecto corresponde a la cursada de Programacion Web.
El objetivo es desarrollar una aplicación web CRUD en Go, de manera incremental a lo largo de los trabajos prácticos.

## 💡 Dominio del Proyecto

La aplicación está diseñada para administrar los turnos de una barbería de manera sencilla y eficiente.
Se pueden registrar clientes y barberos, y asignar turnos específicos para cada cliente con un barbero determinado. Cada turno contiene información sobre el cliente, el barbero, la fecha y hora, el tipo de servicio y observaciones adicionales.

Con esta aplicación, los turnos pueden ser agregados, modificados, consultados o eliminados, permitiendo llevar un control completo de la agenda de la barbería y mejorar la organización del servicio.

## 📂 Estructura del Proyecto

En el Trabajo Práctico N.º 1 (TP1) se implementó la capa web del sistema, separándola de la capa de acceso a datos, que se desarrolla en el Trabajo Práctico N.º 2 (TP2). Esta separación permite una arquitectura más modular y mantenible, donde la capa web se comunica con la lógica de acceso a datos a través de interfaces o servicios intermedios, evitando el acceso directo a la base de datos.
En esta segunda parte incluimos la definicino de las tablas necesarias para nuestro dominio, las consultas CRUD con anotaciones para sqlc y las generamos con sqlc generate

    Tp2
    ├── db/
    │   ├── schema/          # Definición de las tablas Cliente, Barbero y Turno
    │       └── schema.sql
    │   ├── queries/         # Consultas CRUD con anotaciones para sqlc
    │        └── queries.sql
    │   └── sqlc/            # Codigo sqlc ya generado
    │        ├── db.go
    │        ├── models.go
    │        └── queries.sql.go
    ├── templates/           # Archivos HTML de la interfaz
    ├── static/              # Archivos estáticos (CSS, imágenes)
    ├── go.mod               # Módulo Go
    ├── main.go              # Servidor web básico en Go
    ├── index.html           # Página de presentación inicial
    ├── sqlc.yaml            # Configuración de sqlc para generar código Go a partir de SQL
    └── README.md            # Documentación del proyecto

## 📍 Alcance actual 

La aplicación está pensada desde la **vista del Cliente**, quien puede sacar un turno con un barbero.  
En futuras etapas planeamos implementar también la vista/rol del **Barbero**, para que pueda gestionar sus turnos.

## 🚀 Cómo ejecutar el servidor

### Cómo ejecutar el servidor (desde archivos fuente)

Si descargaste o recibiste el proyecto directamente (por ejemplo, por archivo .zip o carpeta), seguí estos pasos para ejecutarlo:

📁 1. Ubícate en la carpeta del proyecto

Abre una terminal y navega hasta la carpeta donde está el proyecto:

Por ejemplo, si lo descomprimiste en el Descargas:
Ejemplo en Windows: 🪟
```
cd C:\Users\tuUsuario\Downloads\Tp2
```
Ejemplo en Linux/MacOS: 🐧
```
cd /home/tuUsuario/Downloads/Tp2
```
🧑‍💻 2. Ejecuta el servidor

Asegúrate de tener Go 1.21 o superior instalado. 

Para asegurarte que tienes Go y la version necesaria debes ejecutar el comando:
```
go version
```

Luego, desde la terminal, ejecuta:
```
go mod tidy                  # Descarga las dependencias necesarias
```


```
go run main.go               # Ejecuta el servidor
```
🌐 4. Abre el navegador

Accede a la siguiente URL en tu navegador:
```
http://localhost:8080
```

¡Listo! Tu servidor estará corriendo localmente.


## Comentarios
La aplicación está pensada solo desde la vista del Cliente por el momento
La vista del Barbero no se implementó en esta etapa.

La capa de datos ya está definida y lista para usarse, pero no se conecta a un servidor web aún.

Todos los archivos generados por sqlc (db.go, models.go, queries.sql.go) ya se incluyen en el proyecto.

### ✍️ Autores : Curtale Juan Ignacio y Saide Felipe


A completar:

Asegurate de tener abierta la aplicacion de Docker Desktop

A continuacion corre el siguiente comando para levantar los contenedores y correr los testeos de la API

bash runtest.sh