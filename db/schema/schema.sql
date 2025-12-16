-- Tablas existentes (sin cambios)
CREATE TABLE Cliente (
    id_cliente SERIAL PRIMARY KEY,
    nombre VARCHAR(50) NOT NULL,
    apellido VARCHAR(50) NOT NULL,
    telefono VARCHAR(20) NOT NULL,
    email VARCHAR(50) UNIQUE NOT NULL 
);

CREATE TABLE Barbero (
    id_barbero SERIAL PRIMARY KEY,
    nombre VARCHAR(50) NOT NULL,
    apellido VARCHAR(50) NOT NULL,
    especialidad VARCHAR(100) NOT NULL
);

-- NUEVA TABLA: Servicios (Precios y Duración)
CREATE TABLE servicios (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    duracion_minutos INT NOT NULL,     -- Ej: 30
    precio DECIMAL(10, 2) NOT NULL,    -- Ej: 12000.00
    activo BOOLEAN DEFAULT true        -- Para borrado lógico
);

-- NUEVA TABLA: Configuración (Para el WhatsApp del dashboard)
CREATE TABLE configuracion (
    clave VARCHAR(50) PRIMARY KEY,     -- Ej: "whatsapp_numero"
    valor VARCHAR(255) NOT NULL
);

-- TABLA MODIFICADA: Turno (Ahora se relaciona con servicios)
CREATE TABLE Turno (
    id_turno SERIAL PRIMARY KEY,
    id_cliente INT NOT NULL,
    id_barbero INT NOT NULL,
    id_servicio INT NOT NULL,  -- CAMBIO: Ahora es un ID, no un texto
    fechaHora TIMESTAMP NOT NULL,
    observaciones TEXT NOT NULL,
    estado VARCHAR(20) DEFAULT 'pendiente', -- Nuevo: pendiente/confirmado
    
    CONSTRAINT fk_cliente FOREIGN KEY (id_cliente) REFERENCES Cliente (id_cliente) ON DELETE CASCADE,
    CONSTRAINT fk_barbero FOREIGN KEY (id_barbero) REFERENCES Barbero (id_barbero) ON DELETE CASCADE,
    CONSTRAINT fk_servicio FOREIGN KEY (id_servicio) REFERENCES servicios (id)
);

-- TABLA DE LOGIN (La que agregamos antes)
CREATE TABLE Usuario_Sistema (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL
);

-- INSERTAR SERVICIOS POR DEFECTO (Seed)
INSERT INTO servicios (nombre, duracion_minutos, precio) VALUES 
('Corte Clásico', 30, 10000),
('Barba', 15, 8000),
('Corte + Barba', 45, 16000);

-- INSERTAR CONFIG INICIAL
INSERT INTO configuracion (clave, valor) VALUES ('whatsapp_numero', '');
