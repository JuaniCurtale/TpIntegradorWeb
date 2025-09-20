
-- CRUD para Cliente

-- name: CreateCliente :one
INSERT INTO Cliente (nombre, apellido, telefono)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetClienteByID :one
SELECT * FROM Cliente
WHERE id_cliente = $1;

-- name: ListClientes :many
SELECT * FROM Cliente
ORDER BY id_cliente;

-- name: UpdateCliente :one
UPDATE Cliente
SET nombre = $2,
    apellido = $3,
    telefono = $4
WHERE id_cliente = $1
RETURNING *;

-- name: DeleteCliente :exec
DELETE FROM Cliente
WHERE id_cliente = $1;


-- CRUD para Barbero

-- name: CreateBarbero :one
INSERT INTO Barbero (nombre, especialidad)
VALUES ($1, $2)
RETURNING *;

-- name: GetBarberoByID :one
SELECT * FROM Barbero
WHERE id_barbero = $1;

-- name: ListBarberos :many
SELECT * FROM Barbero
ORDER BY id_barbero;

-- name: UpdateBarbero :one
UPDATE Barbero
SET nombre = $2,
    especialidad = $3
WHERE id_barbero = $1
RETURNING *;

-- name: DeleteBarbero :exec
DELETE FROM Barbero
WHERE id_barbero = $1;


-- CRUD para Turno

-- name: CreateTurno :one
INSERT INTO Turno (id_cliente, id_barbero, fechaHora, servicio, observaciones)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetTurnoByID :one
SELECT * FROM Turno
WHERE id_turno = $1;

-- name: ListTurnos :many
SELECT * FROM Turno
ORDER BY id_turno;

-- name: UpdateTurno :one
UPDATE Turno
SET id_cliente = $2,
    id_barbero = $3,
    fechaHora = $4,
    servicio = $5,
    observaciones = $6
WHERE id_turno = $1
RETURNING *;

-- name: DeleteTurno :exec
DELETE FROM Turno
WHERE id_turno = $1;
