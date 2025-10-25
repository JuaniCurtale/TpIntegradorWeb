document.addEventListener("DOMContentLoaded", () => {
    // === CLIENTES ===
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
                    div.textContent = `${c.nombre} ${c.apellido} (${c.email || "sin email"})`;
                    const btnEliminar = document.createElement("button");
                    btnEliminar.textContent = "Eliminar";
                    btnEliminar.addEventListener("click", async () => {
                        await fetch(`/cliente/${c.id}`, { method: "DELETE" });
                        loadClientes();
                    });
                    div.appendChild(btnEliminar);
                    clientesContainer.appendChild(div);
                });
            } catch (err) {
                console.error("Error cargando clientes:", err);
            }
        }

        loadClientes(); // carga inicial
    }
});
