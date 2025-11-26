#!/bin/bash
set -e #el script termine inmediatamente si algún comando falla
echo "✨ Generando componentes templ..."
templ generate

echo "🚀 Construyendo y levantando contenedores..."
docker compose up --build -d

# 👉 Abrir navegador automáticamente
echo "🪟 Abriendo navegador"
start http://localhost:8080

echo "✅ Todo finalizó correctamente."
