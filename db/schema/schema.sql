-- schema.sql

CREATE TABLE cliente (
    id_cliente SERIAL PRIMARY KEY,
    nombre VARCHAR(50) NOT NULL,
    apellido VARCHAR(50) NOT NULL,
    telefono VARCHAR(20)
);

CREATE TABLE barbero (
    id_barbero SERIAL PRIMARY KEY,
    nombre VARCHAR(50) NOT NULL,
    especialidad VARCHAR(100)
);

CREATE TABLE turno (
    id_turno SERIAL PRIMARY KEY,
    id_cliente INT NOT NULL,
    id_barbero INT NOT NULL,
    fechaHora TIMESTAMP NOT NULL,
    servicio VARCHAR(100) NOT NULL,
    observaciones TEXT,
    CONSTRAINT fk_cliente
        FOREIGN KEY (id_cliente)
        REFERENCES cliente (id_cliente)
        ON DELETE CASCADE
        ON UPDATE CASCADE,
    CONSTRAINT fk_barbero
        FOREIGN KEY (id_barbero)
        REFERENCES barbero (id_barbero)
        ON DELETE CASCADE
        ON UPDATE CASCADE);
