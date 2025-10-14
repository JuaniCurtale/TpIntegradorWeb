package logic

import (
	"time"

	sqlc "tpIntegradorSaideCurtale/db/sqlc"
)

func HorarioValido(fechaHora time.Time) bool {
	return fechaHora.After(time.Now().Add(time.Hour))
} /* No puede sacar un turno antes de la fecha actual */

func BarberoDisponible(id_Barbero int32, fechaHora time.Time, turnos []sqlc.Turno) bool {
	for _, t := range turnos {
		if t.IDBarbero == id_Barbero && t.Fechahora.Equal(fechaHora) {
			return false
		}
	}
	return true
} /* El barbero no puede tener 2 turnos al mismo tiempo*/

func PuedeReservar(clienteID int32, turnos []sqlc.Turno) bool {
	for _, t := range turnos {
		if t.IDCliente == clienteID && t.Fechahora.After(time.Now()) {
			return false
		}
	}
	return true
} /* un cliente no puede tener 2 turnos pendientes (no puede sacar otro turno si todavia tiene uno para cortarse)*/

func TurnoValido(nuevo sqlc.Turno, turnos []sqlc.Turno) bool {
	if !HorarioValido(nuevo.Fechahora) {
		return false
	}
	if !PuedeReservar(nuevo.IDCliente, turnos) {
		return false
	}
	if !BarberoDisponible(nuevo.IDBarbero, nuevo.Fechahora, turnos) {
		return false
	}
	return true
}
