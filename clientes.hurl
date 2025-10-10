# 1️⃣ Listar todos los clientes
GET http://localhost:8080/cliente
HTTP/1.1 200

[Asserts]
header "Content-Type" == "application/json"

# 2️⃣ Obtener un cliente por ID (ID=1)
GET http://localhost:8080/cliente/1
HTTP/1.1 200

[Asserts]
jsonpath "$.id_cliente" == 1
jsonpath "$.nombre" == "Juan Actualizado"
jsonpath "$.apellido" == "Curtale"


# 3️⃣ Crear un nuevo cliente
POST http://localhost:8080/cliente
Content-Type: application/json

{
  "nombre": "Ana",
  "apellido": "Gomez",
  "telefono": "155554433"
}
HTTP/1.1 201

[Asserts]
jsonpath "$.id_cliente" != null
jsonpath "$.nombre" == "Ana"
jsonpath "$.apellido" == "Gomez"

# 4️⃣ Actualizar un cliente (ID=1)
PUT http://localhost:8080/cliente/1
Content-Type: application/json

{
  "nombre": "Juan Actualizado",
  "apellido": "Curtale",
  "telefono": "87654321"
}
HTTP/1.1 200

[Asserts]
jsonpath "$.nombre" == "Juan Actualizado"
jsonpath "$.apellido" == "Curtale"

# 5️⃣ Borrar un cliente (ID=2)
DELETE http://localhost:8080/cliente/2
HTTP/1.1 204

[Asserts]

