package entities
<<<<<<< HEAD

type Product struct {
	ID             uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	CodigoProducto string `json:"codigo_producto"`
	Item           string `json:"item"`
	Cantidad       int    `json:"cantidad"`
	Categoria      string `json:"categoria"`
	Estado         string `json:"estado"`
	Proveedor      string `json:"proveedor"`
	Fecha          string `json:"fecha"`
	Ubicacion      string `json:"ubicacion"`
}
=======
>>>>>>> b4712727b6f74dd50740129b25e27533ed3e41d9
