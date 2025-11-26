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
    ├── pkg/
    │   └── database/
    │        └── database.go      # Conexión a la base de datosL
    │       
    ├── views/                    # Componentes visuales (.templ), contiene estructura base HTML + importación de HTMX, ademas de pagina de inicio y UI de entidades
    ├── .env                     # Archivo con variables de entorno (no versionado)
    ├── .gitignore
    ├── docker-compose.yml       # Orquestación de servicios (App + DB)
    ├── Dockerfile               # Definición del contenedor de la aplicación
    ├── go.mod
    ├── go.sum
    ├── README.md
    ├── runtest.sh               # Script de automatización
    └── sqlc.yaml                # Configuración de sqlc

## 📍 Evolución del desarrollo 
En esta entrega final, se ha implementado la capa de Interfaces Dinámicas, transformando la experiencia de usuario:

Integración de HTMX: Se incorporó la librería en el layout principal para habilitar capacidades AJAX declarativas.

Formularios Asíncronos: Conversión de formularios tradicionales a peticiones hx-post, eliminando la recarga completa de la página.

Actualización Parcial (SPA feel): Uso de hx-target y hx-swap para actualizar únicamente las tablas de datos tras una operación exitosa.

Feedback y UX: Limpieza automática de formularios tras un envío exitoso utilizando Out-of-Band Swaps (hx-swap-oob).

Borrado en Línea: Implementación de eliminación de registros directamente desde la lista (hx-delete) con confirmación en el cliente (hx-confirm).


## 🚀 Cómo ejecutar el servidor

### Con Docker (Recomendado)

1.  **Instalar Docker y Docker Compose**: Asegúrate de tener ambos instalados en tu sistema.
2.  **Crear archivo .env**: Crea un archivo `.env` en la raíz del proyecto con el siguiente contenido:
    ```
    POSTGRES_DB=barberia
    POSTGRES_USER=postgres
    POSTGRES_PASSWORD=admin
    POSTGRES_PORT=5432

    DB_HOST=barberia_db
    DB_USER=postgres
    DB_PASSWORD=admin
    DB_PORT=5432
    DB_NAME=barberia

    APP_PORT=8080 
    ```
3. **Construccion de la app y levantamiento del contenedor**   
Ejecutar el siguiente comando para construir la imagen, levantar los contenedores y abrir el navegador
```
./runtest.sh
```

En el caso que desee construir la app, levantar Docker y acceder a la aplicacion de manera manual ejecute el siguiente comando:
```
docker-compose up --build
```
4.  **Acceder a la aplicación**: Abre tu navegador y ve a `http://localhost:8080`.


## Comentarios

*   La aplicación ahora cuenta con la vista del **Cliente** y del **Barbero**.
*   La capa de datos está conectada al servidor web.

### ✍️ Autores : Curtale Juan Ignacio y Saide Felipe