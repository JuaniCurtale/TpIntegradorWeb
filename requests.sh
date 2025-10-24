#!/bin/bash
# ==========================================
# PRUEBAS API - TP INTEGRADOR WEB (con HTTP codes)
# ==========================================

DB_CONTAINER="barberia_db"   # nombre del contenedor de la DB
DB_USER="postgres"           # usuario de la DB
DB_NAME="barberia"           # nombre de la DB
BASE_URL="http://localhost:8080"   # URL de la API

function curl_request() {
    # $1 = método, $2 = endpoint, $3 = body (opcional)
    if [ -z "$3" ]; then # se chequea con -z si el body esta vacio
        response=$(curl -s -w "\n%{http_code}" -X "$1" "$BASE_URL/$2")
    else
        response=$(curl -s -w "\n%{http_code}" -X "$1" "$BASE_URL/$2" -H "Content-Type: application/json" -d "$3")
    fi
    http_code=$(echo "$response" | tail -n1)
    body=$(echo "$response" | head -n-1)
    echo "HTTP: $http_code"
    echo "BODY: $body"
    echo
}

echo "==============================="
echo "CLIENTES"
echo "==============================="

echo "Creo cliente"
curl_request POST "cliente" '{"nombre": "Juan", "apellido": "Perez", "telefono": "12345678"}'

echo "Listo clientes"
curl_request GET "cliente"

echo "Chequeo cliente por ID"
curl_request GET "cliente/1"

echo "Actualizo cliente"
curl_request PUT "cliente/1" '{"nombre": "Juan Carlos", "apellido": "Perez", "telefono": "87654321"}'

echo "Elimino cliente"
curl_request DELETE "cliente/1"

echo "Listo clientes después de eliminar"
curl_request GET "cliente"


echo "==============================="
echo "BARBEROS"
echo "==============================="

echo "Creo barbero"
curl_request POST "barbero" '{"nombre": "Carlos", "apellido": "Gomez", "especialidad": "Cortes modernos"}'

echo "Listo barberos"
curl_request GET "barbero"

echo "Chequeo barbero por ID"
curl_request GET "barbero/1"

echo "Actualizo barbero"
curl_request PUT "barbero/1" '{"nombre": "Carlos", "apellido": "Gomez", "especialidad": "Degradados"}'

echo "Elimino barbero"
curl_request DELETE "barbero/1"

echo "Listo barberos después de eliminar"
curl_request GET "barbero"


echo "==============================="
echo "TURNOS"
echo "==============================="

echo "Creo barbero para turno"
curl_request POST "barbero" '{"nombre": "Carlos", "apellido": "Gomez", "especialidad": "Cortes modernos"}'

echo "Creo cliente para turno"
curl_request POST "cliente" '{"nombre": "Juan", "apellido": "Perez", "telefono": "12345678"}'

echo "Creo turno"
curl_request POST "turno" '{"id_cliente":2,"id_barbero":2,"fechahora":"2026-10-15T15:00:00Z","servicio":"Corte de pelo"}'

echo "Listo turnos"
curl_request GET "turno"

echo "Chequeo turno por ID"
curl_request GET "turno/1"

echo "Actualizo turno"
curl_request PUT "turno/1" '{"id_cliente":2,"id_barbero":2,"fechahora":"2027-10-15T15:00:00Z","servicio":"Corte actualizado"}'

echo "Elimino turno"
curl_request DELETE "turno/1"

echo "Listo turnos después de eliminar"
curl_request GET "turno"


echo "Pruebas completadas"

echo "Limpiando la base de datos..."
docker exec -i $DB_CONTAINER psql -U $DB_USER -d $DB_NAME <<EOF
TRUNCATE TABLE cliente, barbero, turno RESTART IDENTITY CASCADE;
EOF

echo "⏱ Esperando 2 segundos..."
sleep 2
