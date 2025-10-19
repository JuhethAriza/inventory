document.addEventListener('DOMContentLoaded', () => {
  const form = document.querySelector('form');
  const emailInput = document.getElementById('email'); 
  const passwordInput = document.getElementById('password');
  const errorElement = document.getElementById('errorMessage');
  const submitButton = form.querySelector('input[type="submit"]');

  // 🔹 Loader dinámico
  const loader = document.createElement('span');
  loader.classList.add('loader');
  loader.style.display = 'none';
  submitButton.insertAdjacentElement('afterend', loader);

  form.addEventListener('submit', async (e) => {
    e.preventDefault();

    const email = emailInput.value.trim();
    const password = passwordInput.value.trim();

    // Limpiar errores anteriores
    errorElement.style.display = 'none';
    errorElement.textContent = '';

    // 🔹 Validaciones básicas
    if (!email) {
      showError('El correo es obligatorio');
      return;
    }

    const gmailRegex = /^[a-zA-Z0-9._%+-]+@gmail\.com$/;
    if (!gmailRegex.test(email)) {
      showError('El correo debe ser un Gmail válido');
      return;
    }

    if (!password || password.length < 6) {
      showError('La contraseña debe tener al menos 6 caracteres');
      return;
    }

    // 🔹 Mostrar loader
    submitButton.disabled = true;
    submitButton.value = 'Procesando...';
    loader.style.display = 'inline-block';

    try {
      const response = await fetch('http://localhost:8080/users/login', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ email, password })
      });

      const data = await response.json();

      if (!response.ok || data.isError) {
        throw new Error(data.message || 'Error en el inicio de sesión');
      }

      // 🔹 Guardar token y redirigir
      if (data.data?.token) {
        localStorage.setItem('token', data.data.token);
        window.location.href = '../dashboard/dashboard.html';
      } else {
        showError('Token no recibido del servidor.');
      }
    } catch (error) {
      console.error('Error:', error);
      showError(error.message || 'No se pudo conectar con el servidor');
    } finally {
      submitButton.disabled = false;
      submitButton.value = 'Iniciar sesión';
      loader.style.display = 'none';
    }
  });

  // 🔹 Función para mostrar errores
  function showError(msg) {
    errorElement.textContent = msg;
    errorElement.style.display = 'block';
  }
});
