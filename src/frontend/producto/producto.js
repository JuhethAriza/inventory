// ==================== CONFIGURACIÓN ====================
const API_URL = "http://127.0.0.1:8080/products";

// ==================== VARIABLES ====================
const btnNuevo = document.getElementById("btnNuevo");
const modal = document.getElementById("modalNuevoProducto");
const btnCancelar = document.getElementById("btnCancelar");
const form = document.getElementById("formProducto");
const listaProductos = document.getElementById("listaProductos");
const buscar = document.getElementById("buscar");
const btnExportar = document.getElementById("btnExportar");

let productos = [];
let idEditar = null;

// ==================== API CALLS ====================
async function getAllProducts() {
  try {
    const res = await fetch(`${API_URL}/GetAll`);
    if (!res.ok) throw new Error("Error al obtener productos");
    const data = await res.json();

    productos = Array.isArray(data) ? data : data.data || data.productos || [];

    mostrarProductos();
  } catch (err) {
    console.error("❌ Error al obtener productos:", err);
    listaProductos.innerHTML = `<tr><td colspan="10" style="text-align:center;color:red;">Error al cargar productos</td></tr>`;
  }
}

async function createProduct(producto) {
  try {
    const res = await fetch(`${API_URL}/Create`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(producto),
    });
    if (!res.ok) {
      const errorText = await res.text();
      throw new Error(`Error en crear: ${errorText}`);
    }

    await getAllProducts();
  } catch (err) {
    console.error("❌ Error al crear producto:", err);
    alert("Error al crear producto. Revisa consola.");
  }
}

async function updateProduct(id, producto) {
  try {
    const res = await fetch(`${API_URL}/Update/${id}`, {
      method: "PUT",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(producto),
    });
    if (!res.ok) throw new Error("Error al actualizar producto");

    await getAllProducts();
  } catch (err) {
    console.error("❌ Error al actualizar producto:", err);
  }
}

async function deleteProduct(id) {
  try {
    const res = await fetch(`${API_URL}/Delete/${id}`, { method: "DELETE" });
    if (!res.ok) throw new Error("Error al eliminar producto");

    await getAllProducts();
  } catch (err) {
    console.error("❌ Error al eliminar producto:", err);
  }
}

// ==================== MOSTRAR PRODUCTOS ====================
function mostrarProductos(filtro = "") {
  listaProductos.innerHTML = "";

  if (!productos || productos.length === 0) {
    listaProductos.innerHTML = `<tr><td colspan="10" style="text-align:center;">No hay productos</td></tr>`;
    return;
  }

  productos
    .filter(
      (p) =>
        (p.item && p.item.toLowerCase().includes(filtro.toLowerCase())) ||
        (p.codigo_producto &&
          p.codigo_producto.toLowerCase().includes(filtro.toLowerCase()))
    )
    .forEach((p) => {
      const fila = document.createElement("tr");
      fila.innerHTML = `
        <td>${p.id || "-"}</td>
        <td>${p.codigo_producto || "-"}</td>
        <td>${p.item || "-"}</td>
        <td>${p.cantidad || "-"}</td>
        <td>${p.categoria || "-"}</td>
        <td>${p.estado || "-"}</td>
        <td>${p.proveedor || "-"}</td>
        <td>${p.fecha || "-"}</td>
        <td>${p.ubicacion || "-"}</td>
        <td>
          <button onclick="editarProducto(${p.id})" class="btn-accion">✏</button>
          <button onclick="eliminarProducto(${p.id})" class="btn-accion">🗑</button>
        </td>`;
      listaProductos.appendChild(fila);
    });
}

// ==================== ABRIR Y CERRAR MODAL ====================
btnNuevo.addEventListener("click", () => {
  idEditar = null;
  form.reset();
  document.getElementById("modalTitle").innerText = "Nuevo Producto";
  modal.style.display = "flex";
});

btnCancelar.addEventListener("click", () => {
  modal.style.display = "none";
});

// ==================== GUARDAR PRODUCTO ====================
form.addEventListener("submit", async (e) => {
  e.preventDefault();

  const producto = {
    codigo_producto: document.getElementById("codigo").value.trim(),
    item: document.getElementById("item").value.trim(),
    cantidad: parseInt(document.getElementById("cantidad").value.trim(), 10),
    categoria: document.getElementById("categoria").value.trim(),
    estado: document.getElementById("estado").value,
    proveedor: document.getElementById("proveedor").value.trim(),
    fecha: document.getElementById("fecha").value,
    ubicacion: document.getElementById("ubicacion").value.trim(),
  };

  if (!producto.codigo_producto || !producto.item) {
    alert("Por favor completa los campos obligatorios.");
    return;
  }

  if (idEditar === null) {
    await createProduct(producto);
  } else {
    await updateProduct(idEditar, producto);
  }

  modal.style.display = "none";
  form.reset();
});

// ==================== EDITAR ====================
window.editarProducto = function (id) {
  const p = productos.find((prod) => prod.id === id);
  if (!p) return;

  idEditar = id;
  document.getElementById("modalTitle").innerText = "Editar Producto";
  document.getElementById("codigo").value = p.codigo_producto;
  document.getElementById("item").value = p.item;
  document.getElementById("cantidad").value = p.cantidad;
  document.getElementById("categoria").value = p.categoria;
  document.getElementById("estado").value = p.estado;
  document.getElementById("proveedor").value = p.proveedor;
  document.getElementById("fecha").value = p.fecha;
  document.getElementById("ubicacion").value = p.ubicacion;

  modal.style.display = "flex";
};

// ==================== ELIMINAR ====================
window.eliminarProducto = async function (id) {
  if (confirm("¿Deseas eliminar este producto?")) {
    await deleteProduct(id);
  }
};

// ==================== BUSCAR ====================
buscar.addEventListener("input", () => mostrarProductos(buscar.value));

// ==================== EXPORTAR (CORREGIDO PARA COLUMNAS) ====================
btnExportar?.addEventListener("click", () => {
  if (!productos || productos.length === 0) {
    alert("No hay productos para exportar");
    return;
  }


  const columnas = [
    "id",
    "codigo_producto",
    "item",
    "cantidad",
    "categoria",
    "estado",
    "proveedor",
    "fecha",
    "ubicacion",
  ];

  const encabezado = [
    "ID",
    "Código",
    "Item",
    "Cantidad",
    "Categoría",
    "Estado",
    "Proveedor",
    "Fecha",
    "Ubicación",
  ];

  function escaparCSV(valor) {
    if (valor === null || valor === undefined) return "";
    const s = String(valor);
    if (/[,"\r\n]/.test(s)) {
      return '"' + s.replace(/"/g, '""') + '"';
    }
    return s;
  }

  const filas = productos.map((p) =>
    columnas.map((col) => escaparCSV(p[col])).join(";") 
  );

  const csv = [encabezado.map(escaparCSV).join(";"), ...filas].join("\r\n"); 

  const blob = new Blob(["\uFEFF" + csv], { type: "text/csv;charset=utf-8;" });
  const url = URL.createObjectURL(blob);
  const link = document.createElement("a");
  link.href = url;
  link.download = "productos.csv";
  document.body.appendChild(link);
  link.click();
  document.body.removeChild(link);
  setTimeout(() => URL.revokeObjectURL(url), 1000);
});


// ==================== INICIALIZAR ====================
getAllProducts();