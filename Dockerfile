# Build del binario
FROM golang:1.24.6 AS builder

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar los archivos de dependencias y descargarlas
COPY go.mod go.sum ./
RUN go mod download

# Copiar el resto del código fuente
COPY . .

# Compilar la aplicación
RUN go build -o app ./cmd

# Imagen final
FROM debian:bookworm-slim

# Crear directorio de trabajo
WORKDIR /root/

# Copiar el binario desde la etapa anterior
COPY --from=builder /app/app .


# Copiar .env al contenedor
COPY .env .env

# Exponer el puerto del servidor
EXPOSE 8080

# Ejecutar la aplicación
CMD ["./app"]
