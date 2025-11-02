package entities

type Product struct {
	ID             uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	CodigoProducto string `json:"codigo_producto"`
	Item           string `json:"item"`
	Cantidad       int    `json:"cantidad"`
	Descripcion    string `json:"descripcion"`
	Observacion    string `json:"observacion"`
	Estado         string `json:"estado"`
	Proveedor      string `json:"proveedor"`
	Fecha          string `json:"fecha"`
	Ubicacion      string `json:"ubicacion"`
}
