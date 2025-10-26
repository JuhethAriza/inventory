package dto

type ProductResponse struct {
	ID             uint   `json:"id,omitempty"`
	CodigoProducto string `json:"codigo_producto,omitempty"`
	Item           string `json:"item,omitempty"`
	Cantidad       int    `json:"cantidad,omitempty"`
	Categoria      string `json:"categoria,omitempty"`
	Estado         string `json:"estado,omitempty"`
	Proveedor      string `json:"proveedor,omitempty"`
	Fecha          string `json:"fecha,omitempty"`
	Ubicacion      string `json:"ubicacion,omitempty"`
}
