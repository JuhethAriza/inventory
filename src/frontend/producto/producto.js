// ==================== VARIABLES ====================
const btnNuevo = document.getElementById("btnNuevo");
const modal = document.getElementById("modalNuevoProducto");
const btnCancelar = document.getElementById("btnCancelar");
const form = document.getElementById("formProducto");
const listaProductos = document.getElementById("listaProductos");
const buscar = document.getElementById("buscar");
const btnExportar = document.getElementById("btnExportar");

let productos = JSON.parse(localStorage.getItem("productos")) || [];
let idEditar = null;

// ==================== MOSTRAR PRODUCTOS ====================
function mostrarProductos(filtro = "") {
  listaProductos.innerHTML = "";

  productos
    .filter(p =>
      p.item.toLowerCase().includes(filtro.toLowerCase()) ||
      p.codigo.toLowerCase().includes(filtro.toLowerCase())
    )
    .forEach((p, index) => {
      const fila = document.createElement("tr");
      fila.innerHTML = `
        <td>${index + 1}</td>
        <td>${p.codigo}</td>
        <td>${p.item}</td>
        <td>${p.cantidad}</td>
        <td>${p.categoria}</td>
        <td>${p.estado}</td>
        <td>${p.proveedor}</td>
        <td>${p.fecha}</td>
        <td>${p.ubicacion}</td>
        <td>
          <button onclick="editarProducto(${index})" class="btn-accion">✏</button>
          <button onclick="eliminarProducto(${index})" class="btn-accion">🗑</button>
        </td>`;
      listaProductos.appendChild(fila);
    });
}

// ==================== ABRIR Y CERRAR MODAL ====================
btnNuevo.addEventListener("click", () => {
  idEditar = null;
  form.reset();
  modal.style.display = "flex";
});

btnCancelar.addEventListener("click", () => {
  modal.style.display = "none";
});

// ==================== GUARDAR PRODUCTO ====================
form.addEventListener("submit", (e) => {
  e.preventDefault();
  try {
    const nuevoProducto = {
      codigo: document.getElementById("codigo").value.trim(),
      item: document.getElementById("item").value.trim(),
      cantidad: document.getElementById("cantidad").value.trim(),
      categoria: document.getElementById("categoria").value.trim(),
      estado: document.getElementById("estado").value,
      proveedor: document.getElementById("proveedor").value.trim(),
      fecha: document.getElementById("fecha").value,
      ubicacion: document.getElementById("ubicacion").value.trim(),
    };

    if (!nuevoProducto.codigo || !nuevoProducto.item) {
      alert("Por favor completa todos los campos obligatorios.");
      return;
    }

    if (idEditar === null) {
      productos.push(nuevoProducto);
    } else {
      productos[idEditar] = nuevoProducto;
    }

    localStorage.setItem("productos", JSON.stringify(productos));
    mostrarProductos();
    modal.style.display = "none";
  } catch (error) {
    console.error("Error al guardar producto:", error);
    alert("Error al guardar producto. Revisa consola.");
  }
});

// ==================== EDITAR ====================
window.editarProducto = function (index) {
  idEditar = index;
  const p = productos[index];

  document.getElementById("codigo").value = p.codigo;
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
window.eliminarProducto = function (index) {
  if (confirm("¿Deseas eliminar este producto?")) {
    productos.splice(index, 1);
    localStorage.setItem("productos", JSON.stringify(productos));
    mostrarProductos();
  }
};

// ==================== BUSCAR ====================
buscar.addEventListener("input", () => mostrarProductos(buscar.value));

// ==================== EXPORTAR ====================
if (btnExportar) {
  btnExportar.addEventListener("click", () => {
    if (productos.length === 0) {
      alert("No hay productos para exportar");
      return;
    }

    const encabezado = Object.keys(productos[0]).join(",");
    const filas = productos.map(p => Object.values(p).join(","));
    const csv = [encabezado, ...filas].join("\n");

    const blob = new Blob([csv], { type: "text/csv;charset=utf-8;" });
    const url = URL.createObjectURL(blob);
    const link = document.createElement("a");
    link.href = url;
    link.download = "productos.csv";
    link.click();
    URL.revokeObjectURL(url);
  });
}

// ==================== INICIALIZAR ====================
mostrarProductos();
