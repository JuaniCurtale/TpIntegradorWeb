document.addEventListener('DOMContentLoaded', () => { //esperamos a que cargue el dominio (html)

    const form = document.getElementById('turnoForm'); // seleccionamos el formulario
    const divRespuesta = document.getElementById('respuestaApi') // obtenemos la respuesta de la api

    form.addEventListener('submit', (event) => { //el activador es el envio de los datos
        
        event.preventDefault(); //prevenimos el envio tradicional

        const formData = new FormData(form) // creamos objeto 'FormData'

        const datosTurno = Object.fromEntries(formData.entries()); 

    
        datosTurno.acepto_politicas = formData.has('aceptar_politicas');
        


        divRespuesta.innerHTML = 'Reservando...';
        divRespuesta.style.color = 'black';

        fetch('/formsPost', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(datosTurno), // Ahora 'datosTurno' es el objeto que queremos enviar
        })
        .then(response => {
            if (!response.ok){
                // Lanzamos el error para que sea capturado por el .catch
                return response.text().then(text => { 
                    throw new Error(`Error del servidor: ${response.status} ${response.statusText}. Detalles: ${text}`);
                });
            }
            return response.json();
        })
        .then(data => {
            console.log('Respuesta del servidor: ', data);
            divRespuesta.innerHTML = `
                <div class="card">
                    <h2>¡Tu turno ha sido agendado, ${data.Nombre}!</h2>
                    <p class="subtitle">Detalles de tu reserva:</p>
                    <p><strong>Teléfono:</strong> ${data.Telefono}</p>
                    <p><strong>Email:</strong> ${data.Email || 'No provisto'}</p>
                    <p><strong>Servicio:</strong> ${data.Servicio}</p>
                    <p><strong>Barbero:</strong> ${data.Barbero}</p>
                    <p><strong>Fecha:</strong> ${data.Fecha}</p>
                    <p><strong>Hora:</strong> ${data.Hora}</p>
                    <p><strong>Notas:</strong> ${data.Notas || 'Ninguna'}</p>
                </div>`;
            divRespuesta.style.color = 'green';
            form.reset();
        })
        .catch(error => {
            console.error('Error en la reserva:', error);
            divRespuesta.innerHTML = `<p>Error al reservar: ${error.message}</p>`; 
            divRespuesta.style.color = 'red';
        });
    });
});