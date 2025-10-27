document.getElementById('form-cliente').addEventListener('submit', async (e) => {
    e.preventDefault();
    const form = e.target;
    const formData = new FormData(form);
    const data = Object.fromEntries(formData.entries());

    const responseDiv = document.getElementById('respuestaApi');

    try {
        const res = await fetch('/cliente', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data)
        });

        if (res.ok) {
            responseDiv.innerHTML = '<p>Cliente creado con éxito.</p>';
            responseDiv.style.color = 'green';
            form.reset();
        } else {
            const errorText = await res.text();
            responseDiv.innerHTML = `<p>Error al crear cliente: ${errorText}</p>`;
            responseDiv.style.color = 'red';
        }
    } catch (err) {
        responseDiv.innerHTML = `<p>Error de red: ${err.message}</p>`;
        responseDiv.style.color = 'red';
    }
});

const verTodosBtn = document.getElementById('ver-todos-clientes');
const listaClientes = document.getElementById('lista-clientes');

verTodosBtn.addEventListener('click', async () => {
    try {
        const response = await fetch('/api/clientes');
        const clientes = await response.json();

        listaClientes.innerHTML = ''; // Limpiar lista

        if (clientes.length === 0) {
            listaClientes.innerHTML = '<p>No hay clientes para mostrar.</p>';
            return;
        }

        const table = document.createElement('table');
        table.classList.add('results-table');

        table.innerHTML = `
            <thead>
                <tr>
                    <th>ID</th>
                    <th>Nombre</th>
                    <th>Email</th>
                    <th>Acciones</th>
                </tr>
            </thead>
            <tbody>
            </tbody>
        `;

        const tbody = table.querySelector('tbody');
        clientes.forEach(cliente => {
            const tr = document.createElement('tr');
            tr.innerHTML = `
                <td>${cliente.id_cliente}</td>
                <td>${cliente.nombre} ${cliente.apellido}</td>
                <td>${cliente.email.String}</td>
                <td>
                    <button class="delete-btn" data-id="${cliente.id_cliente}">X</button>
                    <button class="ver-turnos-btn" data-id="${cliente.id_cliente}">Ver Turnos</button>
                </td>
            `;
            tbody.appendChild(tr);
        });
        listaClientes.appendChild(table);

        table.addEventListener('click', async (e) => {
            if (e.target.classList.contains('delete-btn')) {
                const id = e.target.dataset.id;
                if (confirm('¿Seguro que quiere eliminar este cliente?')) {
                    try {
                        const res = await fetch(`/cliente/${id}`, {
                            method: 'DELETE'
                        });
                        if (res.ok) {
                            e.target.closest('tr').remove();
                        } else {
                            const errorText = await res.text();
                            alert(`Error al eliminar cliente: ${errorText}`);
                        }
                    } catch (err) {
                        alert(`Error de red: ${err.message}`);
                    }
                }
            }

            if (e.target.classList.contains('ver-turnos-btn')) {
                const id = e.target.dataset.id;
                const turnosDiv = document.getElementById('turnos-cliente');
                try {
                    const response = await fetch(`/api/turnos/cliente/${id}`);
                    const turnos = await response.json();

                    turnosDiv.innerHTML = ''; // Limpiar lista

                    if (turnos.length === 0) {
                        turnosDiv.innerHTML = '<p>No hay turnos para este cliente.</p>';
                        return;
                    }

                    const turnosTable = document.createElement('table');
                    turnosTable.classList.add('results-table');

                    turnosTable.innerHTML = `
                        <thead>
                            <tr>
                                <th>Turno ID</th>
                                <th>Barbero ID</th>
                                <th>Fecha</th>
                                <th>Servicio</th>
                            </tr>
                        </thead>
                        <tbody>
                        </tbody>
                    `;

                    const turnosTbody = turnosTable.querySelector('tbody');
                    turnos.forEach(turno => {
                        const tr = document.createElement('tr');
                        tr.innerHTML = `
                            <td>${turno.id_turno}</td>
                            <td>${turno.id_barbero}</td>
                            <td>${new Date(turno.fechahora).toLocaleString()}</td>
                            <td>${turno.servicio}</td>
                        `;
                        turnosTbody.appendChild(tr);
                    });
                    turnosDiv.appendChild(turnosTable);
                } catch (error) {
                    turnosDiv.innerHTML = '<p>Error al cargar los turnos.</p>';
                    console.error('Error:', error);
                }
            }
        });

    } catch (error) {
        listaClientes.innerHTML = '<p>Error al cargar los clientes.</p>';
        console.error('Error:', error);
    }
});