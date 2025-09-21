📌 TP Integrador Web

Este proyecto corresponde a la cursada de Programacion Web.
El objetivo es desarrollar una aplicación web CRUD en Go, de manera incremental a lo largo de los trabajos prácticos.

📂 Estructura del Proyecto
```
TpIntegradorWeb/
├── db/
│ ├── schema.sql # Definición de la tabla principal
│ ├── queries.sql # Consultas CRUD con anotaciones para sqlc
│ └── sqlc.yaml # Configuración de sqlc
├── main.go # Servidor web básico en Go
├── index.html # Página de presentación inicial
└── README.md # Documentación del proyecto
```
🚀 Cómo ejecutar el servidor

Verifica que tengas Go 1.21 o superior instalado.

Clona este repositorio: 📋
```
git clone https://github.com/JuaniCurtale/TpIntegradorWeb.git
cd TpIntegradorWeb
```
Ejecuta el servidor: 🧑‍💻
```
go run main.go
```

Abre en tu navegador 👉 http://localhost:8080

✍️ Autores : Curtale Juan Ignacio y Saide Felipe
