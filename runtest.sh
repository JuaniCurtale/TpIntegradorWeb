#!/bin/bash
set -e

echo "🚀 Construyendo y levantando contenedores..."
docker compose up --build -d

# Esperar que la API esté lista
echo "⏳ Esperando que la API responda..."
for i in {1..10}; do
  if curl -s http://localhost:8080/health > /dev/null; then
    echo "✅ API lista!"
    break
  fi
  echo "Esperando..."
  sleep 3
done

echo "🧪 Ejecutando pruebas con curl..."
bash requests.sh

echo "🧹 Bajando contenedores..."
docker compose down

echo "✅ Todo finalizó correctamente."
