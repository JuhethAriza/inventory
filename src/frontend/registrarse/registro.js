document.getElementById('registroForm').addEventListener('submit', async function (e) {
  e.preventDefault();

  const email = document.getElementById('email').value.trim();
  const password = document.getElementById('password').value.trim();
  const errorElement = document.getElementById('errorMessage');
  const popup = document.getElementById('popup');
  const popupAceptar = document.getElementById('popupAceptar');

  errorElement.style.display = 'none';

  if (!email) {
    errorElement.textContent = 'El correo es obligatorio';
    errorElement.style.display = 'block';
    return;
  }

  if (!password || password.length < 6) {
    errorElement.textContent = 'La contraseña debe tener al menos 6 caracteres';
    errorElement.style.display = 'block';
    return;
  }

  try {
    const response = await fetch('http://localhost:8080/users/Create', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ email, password }),
    });

    const data = await response.json();

    if (!response.ok || data.isError) {
      throw new Error(data.message || 'Error en el registro');
    }

    // Muestra popup de éxito
    popup.style.display = 'flex';
    popupAceptar.onclick = () => {
      popup.style.display = 'none';
      window.location.href = '../login/login.html';
    };

  } catch (error) {
    errorElement.textContent = error.message || 'Error al conectar con el servidor';
    errorElement.style.display = 'block';
  }
});
