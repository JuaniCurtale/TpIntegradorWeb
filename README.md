# 📌 TP Integrador Web

Este proyecto corresponde a la cursada de Programacion Web.
El objetivo es desarrollar una aplicación web CRUD en Go, de manera incremental a lo largo de los trabajos prácticos.

## 💡 Dominio del Proyecto

La aplicación está diseñada para administrar los turnos de una barbería de manera sencilla y eficiente.
Se pueden registrar clientes y barberos, y asignar turnos específicos para cada cliente con un barbero determinado. Cada turno contiene información sobre el cliente, el barbero, la fecha y hora, el tipo de servicio y observaciones adicionales.

Con esta aplicación, los turnos pueden ser agregados, modificados, consultados o eliminados, permitiendo llevar un control completo de la agenda de la barbería y mejorar la organización del servicio.

## 📂 Estructura del Proyecto
```
TpIntegradorWeb/
├── db/
│   ├── schema/          # Definición de las tablas Cliente, Barbero y Turno
│       └── schema.sql
│   ├── queries/         # Consultas CRUD con anotaciones para sqlc
│        └── queries.sql
│   └── sqlc/
│        ├── db.go
│        ├── models.go
│        └── queries.sql.go
├── templates/           # Archivos HTML de la interfaz
├── static/              # Archivos estáticos (CSS, imágenes)
├── go.mod               # Módulo Go
├── go.sum               # Dependencias del módulo
├── main.go              # Servidor web básico en Go
├── index.html           # Página de presentación inicial
├── sqlc.yaml            # Configuración de sqlc para generar código Go a partir de SQL
└── README.md            # Documentación del proyecto
```

## 📍 Alcance actual 

La aplicación está pensada desde la **vista del Cliente**, quien puede sacar un turno con un barbero.  
En futuras etapas planeamos implementar también la vista/rol del **Barbero**, para que pueda gestionar sus turnos.


## 🚀 Cómo ejecutar el servidor


### Cómo ejecutar el servidor (A traves de github) 😺

Verifica que tengas Go 1.21 o superior instalado.

📋 Clona este repositorio: 
```
git clone https://github.com/JuaniCurtale/TpIntegradorWeb.git
cd TpIntegradorWeb
```
🧑‍💻 Ejecuta el servidor: 
```
go run main.go
```

Abre en tu navegador 👉 http://localhost:8080

### ✍️ Autores : Curtale Juan Ignacio y Saide Felipe
