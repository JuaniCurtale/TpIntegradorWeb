-- schema.sql

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

CREATE TABLE Turno (
    id_turno SERIAL PRIMARY KEY,
    id_cliente INT NOT NULL,
    id_barbero INT NOT NULL,
    fechaHora TIMESTAMP NOT NULL,
    servicio VARCHAR(100) NOT NULL,
    observaciones TEXT NOT NULL,
    CONSTRAINT fk_cliente FOREIGN KEY (id_cliente) REFERENCES Cliente (id_cliente) ON DELETE CASCADE,
    CONSTRAINT fk_barbero FOREIGN KEY (id_barbero) REFERENCES Barbero (id_barbero) ON DELETE CASCADE
);