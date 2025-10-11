##########################################
# CLIENTES
##########################################

# Crear un cliente (POST)
POST http://localhost:8080/cliente
Content-Type: application/json

{
  "nombre": "Juan",
  "apellido": "Perez",
  "telefono": "12345678"
}

POST http://localhost:8080/cliente
Content-Type: application/json

{
  "nombre": "Juan",
  "apellido": "Curtale",
  "telefono": "12345678"
}

# Listar todos los clientes (GET)
GET http://localhost:8080/cliente
Accept: application/json

# Ver un cliente por ID (GET)
GET http://localhost:8080/cliente/1
Accept: application/json

# Actualizar un cliente (PUT)
PUT http://localhost:8080/cliente/1
Content-Type: application/json

{
  "nombre": "Juan",
  "apellido": "Gomez",
  "telefono": "87654321"
}

# Eliminar un cliente (DELETE)
DELETE http://localhost:8080/cliente/2

##########################################
# BARBEROS
##########################################

# Crear un barbero (POST)
POST http://localhost:8080/barbero
Content-Type: application/json

{
  "nombre": "Carlos",
  "especialidad": "Corte clásico"
}

# Crear un barbero (POST)
POST http://localhost:8080/barbero
Content-Type: application/json

{
  "nombre": "Jose",
  "especialidad": "Corte clásico"
}

# Listar todos los barberos (GET)
GET http://localhost:8080/barbero
Accept: application/json

# Ver un barbero por ID (GET)
GET http://localhost:8080/barbero/1
Accept: application/json

# Actualizar un barbero (PUT)
PUT http://localhost:8080/barbero/1
Content-Type: application/json

{
  "nombre": "Carlos",
  "especialidad": "Corte moderno"
}

# Eliminar un barbero (DELETE)
DELETE http://localhost:8080/barbero/2

##########################################
# TURNOS
##########################################

# Crear un turno (POST)
POST http://localhost:8080/turno
Content-Type: application/json

{
  "id_cliente": 1,
  "id_barbero": 1,
  "fechaHora": "2025-10-15T14:30:00Z",
  "servicio": "Corte de cabello",
  "observaciones": "Que sea rápido"
}

# Listar todos los turnos (GET)
GET http://localhost:8080/turno
Accept: application/json

# Ver un turno por ID (GET)
GET http://localhost:8080/turno/1
Accept: application/json

# Actualizar un turno (PUT)
PUT http://localhost:8080/turno/1
Content-Type: application/json

{
  "id_cliente": 1,
  "id_barbero": 1,
  "fechaHora": "2025-10-15T15:00:00Z",
  "servicio": "Corte + barba",
  "observaciones": "Sin demoras"
}

# Eliminar un turno (DELETE)
DELETE http://localhost:8080/turno/1
