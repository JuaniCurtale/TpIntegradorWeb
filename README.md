# 📌 TP Integrador Web

Este proyecto corresponde a la cursada de Programacion Web.
El objetivo es desarrollar una aplicación web CRUD en Go, de manera incremental a lo largo de los trabajos prácticos.

## 💡 Dominio del Proyecto

La aplicación está diseñada para administrar los turnos de una barbería de manera sencilla y eficiente.
Se pueden registrar clientes y barberos, y asignar turnos específicos para cada cliente con un barbero determinado. Cada turno contiene información sobre el cliente, el barbero, la fecha y hora, el tipo de servicio y observaciones adicionales.

Con esta aplicación, los turnos pueden ser agregados, modificados, consultados o eliminados, permitiendo llevar un control completo de la agenda de la barbería y mejorar la organización del servicio.

## 📂 Estructura del Proyecto

En el Trabajo Práctico N.º 1 (TP1) se implementó la capa web del sistema. <br>
En el Trabajo Práctico N.º 2 (TP2) se desarrolló la capa de acceso a datos. <br>
En los Trabajos Prácticos N.º 3 y 4 (TP3 y TP4) se integraron ambas capas, se implementó la lógica de negocio y se crearon los endpoints de la API. Además, se implementó la vista del barbero y se consumieron los endpoints desde el frontend. <br>

    TpIntegradorWeb
    ├── cmd/
    │   └── main.go              # Punto de entrada de la aplicación
    ├── db/
    │   ├── queries/
    │   │   └── queries.sql      # Consultas SQL para sqlc
    │   ├── schema/
    │   │   └── schema.sql       # Esquema de la base de datos
    │   └── sqlc/
    │       ├── db.go
    │       ├── models.go
    │       └── queries.sql.go   # Código Go generado por sqlc
    ├── logic/
    │   └── logic.go             # Lógica de negocio de la aplicación
    ├── pkg/
    │   ├── database/
    │   │   └── database.go      # Conexión a la base de datos
    │   ├── handlers/
    │   │   ├── api_handlers.go  # Handlers de la API
    │   │   └── handlers.go      # Handlers de las páginas web
    │   └── router/
    │       └── router.go        # Definición de las rutas
    ├── static/                  # Archivos estáticos (CSS, JS, imágenes)
    ├── templates/               # Plantillas HTML
    ├── .env                     # Archivo de entorno (no versionado)
    ├── .gitignore
    ├── docker-compose.yml       # Orquestación de los contenedores
    ├── Dockerfile               # Definición del contenedor de la aplicación
    ├── go.mod
    ├── go.sum
    ├── README.md
    ├── requests.sh              # Ejemplos de requests a la API
    └── sqlc.yaml                # Configuración de sqlc

## 📍 Alcance actual

Se han añadido secciones que permiten acceder a los clientes, turnos y barberos. Se puede realizar desde la creacion de los mismos, hasta el listado y la eliminacion de los objetos ya creados

## 🚀 Cómo ejecutar el servidor

### Con Docker (Recomendado)

1.  **Instalar Docker y Docker Compose**: Asegúrate de tener ambos instalados en tu sistema.
2.  **Crear archivo .env**: Crea un archivo `.env` en la raíz del proyecto con el siguiente contenido:
    ```
    DB_HOST=db
    DB_PORT=5432
    DB_USER=user
    DB_PASSWORD=password
    DB_NAME=mydatabase
    ```
3.  **Ejecutar Docker Compose**:
    ```bash
    docker-compose up --build
    ```
4.  **Acceder a la aplicación**: Abre tu navegador y ve a `http://localhost:8080`.

### Desde archivos fuente

Si descargaste o recibiste el proyecto directamente (por ejemplo, por archivo .zip o carpeta), seguí estos pasos para ejecutarlo:

📁 1. Ubícate en la carpeta del proyecto

Abre una terminal y navega hasta la carpeta donde está el proyecto.

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
go run ./cmd/main.go         # Ejecuta el servidor
```

🌐 4. Abre el navegador

Accede a la siguiente URL en tu navegador:
```
http://localhost:8080
```

¡Listo! Tu servidor estará corriendo localmente.

## Comentarios

*   La aplicación ahora cuenta con la vista del **Cliente** y del **Barbero**.
*   La capa de datos está conectada al servidor web.
*   Se recomienda utilizar Docker para facilitar la ejecución del proyecto.

### ✍️ Autores : Curtale Juan Ignacio y Saide Felipe