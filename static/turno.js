document.addEventListener('DOMContentLoaded', () => {
    const form = document.getElementById('form-turno');

    if (!form) {
        console.error("Error: No se encontró el formulario 'form-turno'.");
        return;
    }

    form.addEventListener('submit', (event) => {
        event.preventDefault();

        const divRespuestaInterna = document.getElementById('respuestaApi');
        const contenedorExito = document.getElementById('contenedorExito');

        if (!contenedorExito) {
            console.error("Error: No se encontró el div 'contenedorExito'.");
            if(divRespuestaInterna){
                divRespuestaInterna.innerHTML = `<p>Error: No se encontró el div 'contenedorExito'.</p>`;
                divRespuestaInterna.style.color = 'red';
            }
            return;
        }

        // Resetea la UI al enviar
        contenedorExito.classList.add('oculto');
        form.style.display = 'block';

        const formData = {
                nombre: form.nombre.value,
                telefono: form.telefono.value,
                email: form.email.value,
                servicio: form.servicio.value,
                barbero: form.barbero.value,
                fecha: form.fecha.value,
                hora: form.hora.value,
                notas: form.notas.value
            };

        // Mapeo de valores del form a IDs/datos para el API
        const barberoMap = { 'martin': 1, 'sofia': 2, 'lucas': 3, 'cualquiera': 1 };
        const servicioMap = { 'corte': 'Corte', 'barba': 'Barba', 'corte-barba': 'Corte + Barba', 'afeitado': 'Afeitado tradicional', 'color': 'Color/Matiz' };
        
        const fechaHoraISO = `${formData.fecha}T${formData.hora}:00Z`;
        
        const datosParaApi = {
        nombre: formData.nombre,
        telefono: formData.telefono,
        email: formData.email,
        id_barbero: barberoMap[formData.barbero],
        fechahora: fechaHoraISO,
        servicio: servicioMap[formData.servicio],
        observaciones: {
            String: formData.notas,
            Valid: formData.notas !== "" && formData.notas !== null 
            }
        };

        divRespuestaInterna.innerHTML = 'Reservando...';
        divRespuestaInterna.style.color = 'black';

        fetch('/turno', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(datosParaApi),
        })
        .then(response => {
            if (!response.ok) {
                return response.text().then(text => {
                    throw new Error(`Error ${response.status}: ${text}`);
                });
            }
            return response.json();
        })
        .then(data => {
            // Oculta el form y muestra la card de éxito
            console.log('Respuesta de la API (turno creado): ', data);

            form.style.display = 'none';
            divRespuestaInterna.innerHTML = '';
            contenedorExito.classList.remove('oculto');
            
            contenedorExito.innerHTML = `
                <div class="card">
                    <h2>¡Tu turno ha sido agendado!</h2>
                    <p><strong>Nombre:</strong> ${formData.nombre}</p>
                    <p><strong>Barbero:</strong> ${formData.barbero}</p>
                    <p><strong>Servicio:</strong> ${data.servicio}</p>
                    <p><strong>Fecha:</strong> ${formData.fecha} ${formData.hora}</p>
                    <p><strong>ID de Reserva:</strong> ${data.IDTurno}</p> 
                    <br>
                    <button id="btnOtraReserva" type="button">Hacer otra reserva</button>
                </div>`;

            document.getElementById('btnOtraReserva').addEventListener('click', () => {
                window.location.reload();
            });
        })
        .catch(error => {
            // Error: Muestra el mensaje en el div interno
            console.error('Error en la reserva:', error);
            divRespuestaInterna.innerHTML = `<p>Error al reservar: ${error.message}</p>`;
            divRespuestaInterna.style.color = 'red';
        });
    });

    const verTodosBtn = document.getElementById('ver-todos-turnos');
    const listaTurnos = document.getElementById('lista-turnos');

    verTodosBtn.addEventListener('click', async () => {
        try {
            const response = await fetch('/turno', {
                headers: {
                    'Accept': 'application/json'
                }
            });
            const turnos = await response.json();

            listaTurnos.innerHTML = ''; // Limpiar lista

            if (turnos.length === 0) {
                listaTurnos.innerHTML = '<p>No hay turnos para mostrar.</p>';
                return;
            }

            const table = document.createElement('table');
            table.classList.add('results-table');

            table.innerHTML = `
                <thead>
                    <tr>
                        <th>Turno ID</th>
                        <th>Cliente ID</th>
                        <th>Barbero ID</th>
                        <th>Fecha</th>
                        <th>Servicio</th>
                        <th>Acciones</th>
                    </tr>
                </thead>
                <tbody>
                </tbody>
            `;

            const tbody = table.querySelector('tbody');
            turnos.forEach(turno => {
                const tr = document.createElement('tr');
                tr.innerHTML = `
                    <td>${turno.id_turno}</td>
                    <td>${turno.id_cliente}</td>
                    <td>${turno.id_barbero}</td>
                    <td>${new Date(turno.fechahora).toLocaleString()}</td>
                    <td>${turno.servicio}</td>
                    <td><button class="delete-btn" data-id="${turno.id_turno}">X</button></td>
                `;
                tbody.appendChild(tr);
            });
            listaTurnos.appendChild(table);

            table.addEventListener('click', async (e) => {
                if (e.target.classList.contains('delete-btn')) {
                    const id = e.target.dataset.id;
                    if (confirm('¿Seguro que quiere eliminar este turno?')) {
                        try {
                            const res = await fetch(`/turno/${id}`, {
                                method: 'DELETE'
                            });
                            if (res.ok) {
                                e.target.closest('tr').remove();
                            } else {
                                const errorText = await res.text();
                                alert(`Error al eliminar turno: ${errorText}`);
                            }
                        } catch (err) {
                            alert(`Error de red: ${err.message}`);
                        }
                    }
                }
            });

        } catch (error) {
            listaTurnos.innerHTML = '<p>Error al cargar los turnos.</p>';
            console.error('Error:', error);
        }
    });
});