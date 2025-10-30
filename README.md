# 📌 TP Integrador Web

Este proyecto corresponde a la cursada de Programacion Web.
El objetivo es desarrollar una aplicación web CRUD en Go, de manera incremental a lo largo de los trabajos prácticos.

## 💡 Dominio del Proyecto

La aplicación está diseñada para administrar los turnos de una barbería de manera sencilla y eficiente.
Se pueden registrar clientes y barberos, y asignar turnos específicos para cada cliente con un barbero determinado. Cada turno contiene información sobre el cliente, el barbero, la fecha y hora, el tipo de servicio y observaciones adicionales.

Con esta aplicación, los turnos pueden ser agregados, modificados, consultados o eliminados, permitiendo llevar un control completo de la agenda de la barbería y mejorar la organización del servicio.

## 🛠️ Tecnologías utilizadas
* Go
* PostgreSQL
* Docker & Docker Compose
* HTML, CSS, JavaScript
* sqlc
* Bash

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
<<<<<<< HEAD
    │   │   └── handlers.go      # Handlers de las páginas web HTML
=======
    │   │   └── handlers.go      # Handlers de las páginas web
>>>>>>> d660bfc3464353e374087e3df073da6174584849
    │   └── router/
    │       └── router.go        # Definición de las rutas
    ├── static/                  # Archivos estáticos (CSS, JS, imágenes)
    ├── templates/               # Plantillas HTML
<<<<<<< HEAD
=======
    ├── .env                     # Archivo de entorno (no versionado)
>>>>>>> d660bfc3464353e374087e3df073da6174584849
    ├── .gitignore
    ├── docker-compose.yml       # Orquestación de los contenedores
    ├── Dockerfile               # Definición del contenedor de la aplicación
    ├── go.mod
    ├── go.sum
    ├── README.md
    ├── requests.sh              # Ejemplos de requests a la API
<<<<<<< HEAD
    ├── runtest.sh               # Script para construir la app, levantar Docker y correr los testeos
=======
>>>>>>> d660bfc3464353e374087e3df073da6174584849
    └── sqlc.yaml                # Configuración de sqlc

## 📍 Alcance actual

Se han añadido secciones que permiten acceder a los clientes, turnos y barberos. Se puede realizar desde la creacion de los mismos, hasta el listado y la eliminacion de los objetos ya creados

## 🚀 Cómo ejecutar el servidor

### Con Docker (Recomendado)
<<<<<<< HEAD
=======

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
>>>>>>> d660bfc3464353e374087e3df073da6174584849

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

<<<<<<< HEAD
    APP_PORT=8080 
    ```
3. **Construccion de la app y levantamiento del contenedor**   

Ejecuta el siguiente comando para construir la aplicación, levantar los contenedores y ejecutar los tests
En Linux 🐧: 
```
./runtest.sh
=======
Abre una terminal y navega hasta la carpeta donde está el proyecto.

🧑‍💻 2. Ejecuta el servidor

Asegúrate de tener Go 1.21 o superior instalado.

Para asegurarte que tienes Go y la version necesaria debes ejecutar el comando:
```
go version
>>>>>>> d660bfc3464353e374087e3df073da6174584849
```

En Windows🪟:
```
bash runtest.sh
```
Este comando hara lo dicho anteriormente además de dar de baja los contenedores al finalizar

<<<<<<< HEAD
En el caso que desee construir la app, levantar Docker y acceder a la aplicacion ejecute el siguiente comando:
```
docker-compose up --build
```
Para correr los testeos manualmente ejecuta el siguiente comando
En Linux🐧:
```
./requests.sh
```
En Windows🪟:
```
bash requests.sh
```
4.  **Acceder a la aplicación**: Abre tu navegador y ve a `http://localhost:8080`.

## 🖱️Navegacion dentro de la aplicacion
   Una vez dentro de la aplicación, podrás navegar por las distintas secciones disponibles desde el menú principal, tales como:
=======
```
go run ./cmd/main.go         # Ejecuta el servidor
```

🌐 4. Abre el navegador
>>>>>>> d660bfc3464353e374087e3df073da6174584849

   * **Inicio**: Pagina de home.
   * **Sobre Nosotros**: Información general sobre la barbería.
   * **Registrar Cliente**: Formulario para agregar,eliminar y visualizar los clientes.
   * **Registrar Barbero**: Sección para registrar, eliminar y visualizar barberos en el sistema.
   * **Sacar Turno**: Permite asignar turnos a los clientes con un barbero determinado.

## 🔧 Mejoras futuras y pendientes

- **Implementar sistema de autenticación y login:**  
  Agregar un módulo de inicio de sesión que permita identificar a los usuarios (barberos, clientes o administrador) y restringir el acceso según su rol.

- **Gestión de roles y permisos:**  
  Definir niveles de acceso para cada tipo de usuario, evitando que todos puedan visualizar o modificar información ajena.

- **Validaciones y mensajes de error más detallados:**  
  Mejorar el manejo de errores tanto en el backend como en el frontend para ofrecer una experiencia más clara al usuario.

- **Optimización del diseño y la interfaz:**  
  Aplicar un diseño más responsivo y moderno, manteniendo la simplicidad y funcionalidad actual.

## Comentarios

<<<<<<< HEAD
*   Se tomo la decision de separar en tres paginas HTML diferentes las entidades para poder darle protagonismo a cada una y que su utilizacion sea mas comoda
*   Sobre Cliente se decidio que el email sea la Primary Key, de esta manera al crear el cliente y luego sacar el turno, si se saca el turno con el mismo email que se registro el cliente este turno correspondera a ese cliente, si se ingresa un email no registrado en la tabla de clientes, se creara uno
*   La aplicación ahora cuenta con la vista del **Cliente** y del **Barbero**.
*   La capa de datos está conectada al servidor web.

### ✍️ Autores : Curtale Juan Ignacio y Saide Felipe
=======
*   La aplicación ahora cuenta con la vista del **Cliente** y del **Barbero**.
*   La capa de datos está conectada al servidor web.
*   Se recomienda utilizar Docker para facilitar la ejecución del proyecto.

### ✍️ Autores : Curtale Juan Ignacio y Saide Felipe
>>>>>>> d660bfc3464353e374087e3df073da6174584849
