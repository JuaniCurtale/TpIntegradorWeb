document.addEventListener("DOMContentLoaded", () => {
    // CLIENTES
    const formCliente = document.getElementById("form-cliente");
    const clientesContainer = document.getElementById("clientesContainer");

    if (formCliente) {
        // Envío de formulario cliente
        formCliente.addEventListener("submit", async (e) => {
            e.preventDefault();
            const data = {
                nombre: formCliente.nombre.value,
                apellido: formCliente.apellido.value,
                telefono: formCliente.telefono.value,
                email: formCliente.email.value
            };
            try {
                await fetch("/cliente", {
                    method: "POST",
                    headers: { "Content-Type": "application/json" },
                    body: JSON.stringify(data)
                });
                formCliente.reset();
                loadClientes();
            } catch (err) {
                console.error("Error creando cliente:", err);
            }
        });

        // Cargar clientes
        async function loadClientes() {
            try {
                const res = await fetch("/cliente");
                const clientes = await res.json();
                clientesContainer.innerHTML = "";
                clientes.forEach(c => {
                    const div = document.createElement("div");
                    div.dataset.id = c.id_cliente;
                    div.innerHTML = `
                        ${c.nombre} ${c.apellido} (${c.email || "sin email"})
                        <button class="btn-eliminar">Eliminar</button>
                    `;
                    clientesContainer.appendChild(div);
                });
            } catch (err) {
                console.error("Error cargando clientes:", err);
            }
        }

        // Delegación de eventos para el borrado
            clientesContainer.addEventListener("click", async (e) => {
    if (e.target.classList.contains("btn-eliminar")) {
        const clienteDiv = e.target.closest("div");
        const id = Number(clienteDiv.dataset.id);
        if (!id) return;

        try {
            const res = await fetch(`/cliente/${id}`, { method: "DELETE" });
            if (!res.ok) console.error("Error del servidor:", await res.text());
            loadClientes();
        } catch (err) {
            console.error("Error eliminando cliente:", err);
             }
        }
    });


        loadClientes(); // carga inicial
    }
});
