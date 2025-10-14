#!/bin/bash
# ==========================================
# PRUEBAS API - TP INTEGRADOR WEB
# ==========================================

DB_CONTAINER="barberia_db"  # nombre del contenedor de la DB
DB_USER="postgres"                  # usuario de la DB
DB_NAME="barberia"                  # nombre de la DB
BASE_URL="http://localhost:8080"    # URL de la API

echo "==============================="
echo "🧍‍♂️ CLIENTES"
echo "==============================="

# Crear cliente
echo -e "Creo cliente \n"

curl -s -X POST "$BASE_URL/cliente" \
-H "Content-Type: application/json" \
-d '{"nombre": "Juan", "apellido": "Perez", "telefono": "12345678"}' 
echo -e "\n"

# Listar clientes
echo -e "Listo cliente \n"
curl -s -X GET "$BASE_URL/cliente" 
echo -e "\n"

# Ver cliente por ID
echo -e "Chequeo id_cliente"
curl -s -X GET "$BASE_URL/cliente/1" 
echo -e "\n"

# Actualizar cliente
echo -e "Actualizo cliente \n"
curl -s -X PUT "$BASE_URL/cliente/1" \
-H "Content-Type: application/json" \
-d '{"nombre": "Juan Carlos", "apellido": "Perez", "telefono": "87654321"}' 
echo -e "\n"

# Eliminar cliente
echo -e "Elimina cliente \n"
curl -s -X DELETE "$BASE_URL/cliente/1"
echo -e "\n"

curl -s -X GET "$BASE_URL/cliente" 


echo "==============================="
echo "💈 BARBEROS"
echo "==============================="

# Crear barbero
echo -e "Creo barbero \n"
curl -s -X POST "$BASE_URL/barbero" \
-H "Content-Type: application/json" \
-d '{"nombre": "Carlos", "apellido": "Gomez", "especialidad": "Cortes modernos"}' 
echo -e "\n"

# Listar barberos
echo -e "Listar barberos \n"
curl -s -X GET "$BASE_URL/barbero" 
echo -e "\n"

# Ver barbero por ID
echo -e "buscar barbero con ID \n"
curl -s -X GET "$BASE_URL/barbero/1" 
echo -e "\n"

# Actualizar barbero
echo -e "Actualizar barbero \n"
curl -s -X PUT "$BASE_URL/barbero/1" \
-H "Content-Type: application/json" \
-d '{"nombre": "Carlos", "apellido": "Gomez", "especialidad": "Degradados"}' 
echo -e "\n"

# Eliminar barbero
echo -e "Elimino barbero \n"
curl -s -X DELETE "$BASE_URL/barbero/1"
echo -e "\n"

curl -s -X GET "$BASE_URL/barbero" 



echo "==============================="
echo "🕒 TURNOS"
echo "==============================="
echo -e "Creo turno \n"
# Crear turno
curl -s -X POST "$BASE_URL/turno" \
-H "Content-Type: application/json" \
-d '{"id_cliente":1,"id_barbero":1,"fechahora":"2024-10-14T10:30:00Z","servicio":"Corte de pelo"}'


# Listar turnos
echo -e "Listo turnos \n"
curl -s -X GET "$BASE_URL/turno"
echo -e "\n"

# Ver turno por ID
echo -e "Chequeo turno con ID \n"
curl -s -X GET "$BASE_URL/turno/1"
echo -e "\n"

# Actualizar turno
echo -e "Actualizo turno \n"
curl -s -X PUT "$BASE_URL/turno/1" \
-H "Content-Type: application/json" \
-d '{"id_cliente":1,"id_barbero":1,"Fechahora":"2026-10-15T15:00:00Z","servicio":"Afeitado y corte"}'
echo -e "\n"

# Eliminar turno
echo -e "Elimino turno \n"
curl -s -X DELETE "$BASE_URL/turno/1"
echo -e "\n"

echo "✅ Pruebas completadas"

echo "🧹 Limpiando la base de datos..."
docker exec -i $DB_CONTAINER psql -U $DB_USER -d $DB_NAME <<EOF
TRUNCATE TABLE cliente, barbero, turno RESTART IDENTITY CASCADE;
EOF

# 2️⃣ Esperar un momento para que todo esté listo
echo "⏱ Esperando 2 segundos..."
sleep 2