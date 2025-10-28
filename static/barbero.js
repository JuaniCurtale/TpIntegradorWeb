document.getElementById('form-barbero').addEventListener('submit', async (e) => {
    e.preventDefault();
    const form = e.target;
    const formData = new FormData(form);
    const data = Object.fromEntries(formData.entries());

    const responseDiv = document.getElementById('respuestaApi');

    try {
        const res = await fetch('/barbero', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data)
        });

        if (res.ok) {
            responseDiv.innerHTML = '<p>Barbero creado con éxito.</p>';
            responseDiv.style.color = 'green';
            form.reset();
        } else {
            const errorText = await res.text();
            responseDiv.innerHTML = `<p>Error al crear barbero: ${errorText}</p>`;
            responseDiv.style.color = 'red';
        }
    } catch (err) {
        responseDiv.innerHTML = `<p>Error de red: ${err.message}</p>`;
        responseDiv.style.color = 'red';
    }
});

const verTodosBtn = document.getElementById('ver-todos-barberos');
const listaBarberos = document.getElementById('lista-barberos');

verTodosBtn.addEventListener('click', async () => {
    try {
        const response = await fetch('/barbero', {
            headers: {
                'Accept': 'application/json'
            }
        });
        const barberos = await response.json();

        listaBarberos.innerHTML = ''; // Limpiar lista

        if (barberos.length === 0) {
            listaBarberos.innerHTML = '<p>No hay barberos para mostrar.</p>';
            return;
        }

        const table = document.createElement('table');
        table.classList.add('results-table');

        table.innerHTML = `
            <thead>
                <tr>
                    <th>ID</th>
                    <th>Nombre</th>
                    <th>Especialidad</th>
                    <th>Acciones</th>
                </tr>
            </thead>
            <tbody>
            </tbody>
        `;

        const tbody = table.querySelector('tbody');
        barberos.forEach(barbero => {
            const tr = document.createElement('tr');
            tr.innerHTML = `
                <td>${barbero.id_barbero}</td>
                <td>${barbero.nombre} ${barbero.apellido}</td>
                <td>${barbero.especialidad.String}</td>
                <td>
                    <button class="delete-btn" data-id="${barbero.id_barbero}">X</button>
                </td>
            `;
            tbody.appendChild(tr);
        });
        listaBarberos.appendChild(table);

        table.addEventListener('click', async (e) => {
            if (e.target.classList.contains('delete-btn')) {
                const id = e.target.dataset.id;
                if (confirm('¿Seguro que quiere eliminar este barbero?')) {
                    try {
                        const res = await fetch(`/barbero/${id}`, {
                            method: 'DELETE'
                        });
                        if (res.ok) {
                            e.target.closest('tr').remove();
                        } else {
                            const errorText = await res.text();
                            alert(`Error al eliminar barbero: ${errorText}`);
                        }
                    } catch (err) {
                        alert(`Error de red: ${err.message}`);
                    }
                }
            }
        });

    } catch (error) {
        listaBarberos.innerHTML = '<p>Error al cargar los barberos.</p>';
        console.error('Error:', error);
    }
});