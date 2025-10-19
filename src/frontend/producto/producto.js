const API_URL = "http://localhost:8080/products";

const form = document.getElementById("formProducto");
const lista = document.getElementById("listaProductos");
const modal = document.getElementById("modalNuevoProducto");
const btnNuevo = document.getElementById("btnNuevo");
const btnCancelar = document.getElementById("btnCancelar");
const btnGuardar = document.getElementById("btnGuardar");
const inputBuscar = document.getElementById("buscar");

let editando = false;
let productoActualId = null;

async function parseErrorResponse(res) {
  let text = `${res.status} ${res.statusText}`;
  try {
    const json = await res.json();
    if (json.message) text += ` — ${json.message}`;
    else if (json.error) text += ` — ${json.error}`;
    else text += ` — ${JSON.stringify(json)}`;
  } catch {
    const t = await res.text();
    if (t) text += ` — ${t}`;
  }
  return text;
}

// 🔹 Abrir modal nuevo producto
btnNuevo.addEventListener("click", () => {
  form.reset();
  editando = false;
  productoActualId = null;
  btnGuardar.innerText = "Guardar";
  modal.style.display = "flex";
});

// 🔹 Cerrar modal
btnCancelar.addEventListener("click", () => {
  modal.style.display = "none";
});

window.addEventListener("click", (e) => {
  if (e.target === modal) modal.style.display = "none";
});

// 🔹 Cargar productos
async function cargarProductos() {
  try {
    const res = await fetch(`${API_URL}/GetAll`);
    if (!res.ok) {
      const err = await parseErrorResponse(res);
      alert("Error al cargar productos:\n" + err);
      return;
    }
    const data = await res.json();
    const productos = Array.isArray(data) ? data : data.data || [];
    renderTabla(productos);
  } catch {
    alert("No se pudo conectar con el servidor");
  }
}

// 🔹 Renderizar tabla
function renderTabla(items) {
  lista.innerHTML = "";
  if (!items.length) {
    lista.innerHTML = `<tr><td colspan="9" style="text-align:center;">No hay productos</td></tr>`;
    return;
  }
  items.forEach((p) => {
    const tr = document.createElement("tr");
    tr.innerHTML = `
      <td>${p.id || ""}</td>
      <td>${p.codigo_producto || p.codigo || ""}</td>
      <td>${p.item || ""}</td>
      <td>${p.cantidad || ""}</td>
      <td>${p.categoria || ""}</td>
      <td>${p.estado || ""}</td>
      <td>${p.proveedor || ""}</td>
      <td>${p.fecha || ""}</td>
      <td>${p.ubicacion || ""}</td>
      <td>
        <button class="btn-editar" data-id="${p.id}">✏️</button>
        <button class="btn-eliminar" data-id="${p.id}">🗑️</button>
      </td>
    `;
    lista.appendChild(tr);
  });

  document.querySelectorAll(".btn-eliminar").forEach((b) => {
    b.addEventListener("click", () => eliminarProducto(b.dataset.id));
  });

  document.querySelectorAll(".btn-editar").forEach((b) => {
    b.addEventListener("click", () => editarProducto(b.dataset.id));
  });
}

// 🔹 Guardar o actualizar
form.addEventListener("submit", async (e) => {
  e.preventDefault();

  const data = {};
  new FormData(form).forEach((val, key) => {
    if (key === "codigo") key = "codigo_producto";
    if (key === "cantidad") data[key] = parseInt(val, 10);
    else data[key] = val;
  });

  btnGuardar.disabled = true;
  btnGuardar.innerText = editando ? "Actualizando..." : "Guardando...";

  try {
    let res;
    if (editando && productoActualId) {
      // ✅ UPDATE producto existente
      res = await fetch(`${API_URL}/Update/${productoActualId}`, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(data),
      });
    } else {
      // ✅ CREATE nuevo producto
      res = await fetch(`${API_URL}/Create`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(data),
      });
    }

    if (!res.ok) {
      const err = await parseErrorResponse(res);
      alert("No se pudo guardar el producto:\n" + err);
      return;
    }

    await cargarProductos();
    modal.style.display = "none";
  } catch {
    alert("Error al guardar producto. Revisa consola.");
  } finally {
    btnGuardar.disabled = false;
    btnGuardar.innerText = "Guardar";
  }
});

// 🔹 Editar producto
function editarProducto(id) {
  console.log("🟦 Editando producto con ID:", id);
  const fila = document.querySelector(`button[data-id="${id}"]`).closest("tr");
  const celdas = fila.querySelectorAll("td");

  // Asignar valores a los inputs
  document.getElementById("codigo").value = celdas[1].innerText;
  document.getElementById("item").value = celdas[2].innerText;
  document.getElementById("cantidad").value = celdas[3].innerText;
  document.getElementById("categoria").value = celdas[4].innerText;
  document.getElementById("estado").value = celdas[5].innerText;
  document.getElementById("proveedor").value = celdas[6].innerText;
  document.getElementById("fecha").value = celdas[7].innerText;
  document.getElementById("ubicacion").value = celdas[8].innerText;

  // Activar modo edición
  editando = true;
  productoActualId = id;

  btnGuardar.innerText = "Actualizar";
  modal.style.display = "flex";
}

// 🔹 Popup de confirmación para eliminar
let productoAEliminar = null;

function eliminarProducto(id) {
  productoAEliminar = id;
  document.getElementById("popupConfirm").classList.remove("hidden");
}

document.getElementById("btnConfirmDelete").addEventListener("click", async () => {
  if (!productoAEliminar) return;

  try {
    const res = await fetch(`${API_URL}/Delete/${productoAEliminar}`, {
      method: "DELETE",
    });
    if (!res.ok) {
      const err = await parseErrorResponse(res);
      alert("Error al eliminar producto:\n" + err);
      return;
    }
    await cargarProductos();
  } catch {
    alert("Error al eliminar producto.");
  } finally {
    productoAEliminar = null;
    document.getElementById("popupConfirm").classList.add("hidden");
  }
});

document.getElementById("btnCancelDelete").addEventListener("click", () => {
  productoAEliminar = null;
  document.getElementById("popupConfirm").classList.add("hidden");
});

// 🔹 Búsqueda en tiempo real
if (inputBuscar) {
  inputBuscar.addEventListener("input", () => {
    const q = inputBuscar.value.toLowerCase();
    document.querySelectorAll("#listaProductos tr").forEach((tr) => {
      tr.style.display = tr.innerText.toLowerCase().includes(q) ? "" : "none";
    });
  });
}

window.addEventListener("DOMContentLoaded", cargarProductos);
