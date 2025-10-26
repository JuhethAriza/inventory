package dto

type ProductDTO struct {
	ID             uint   `json:"id"`
	CodigoProducto string `json:"codigo_producto" validate:"required"`
	Item           string `json:"item" validate:"required"`
	Cantidad       int    `json:"cantidad" validate:"required"`
	Categoria      string `json:"categoria,omitempty"`
	Estado         string `json:"estado,omitempty"`
	Proveedor      string `json:"proveedor" validate:"required"`
	Fecha          string `json:"fecha" validate:"required"`
	Ubicacion      string `json:"ubicacion" validate:"required"`
}

type UpdateProductDTO struct {
	CodigoProducto string `json:"codigo_producto" validate:"required"`
	Item           string `json:"item" validate:"required"`
	Cantidad       int    `json:"cantidad" validate:"required"`
	Categoria      string `json:"categoria,omitempty"`
	Estado         string `json:"estado,omitempty"`
	Proveedor      string `json:"proveedor" validate:"required"`
	Fecha          string `json:"fecha" validate:"required"`
	Ubicacion      string `json:"ubicacion" validate:"required"`
}
