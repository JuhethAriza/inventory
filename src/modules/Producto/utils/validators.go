package utils
<<<<<<< HEAD

import (
	"errors"
	"strings"
	"time"

	dto "github.com/JuhethAriza/inventory/src/modules/Producto/domain/dto"
)

func ValidateProductDTO(p dto.ProductDTO) error {
	if strings.TrimSpace(p.CodigoProducto) == "" {
		return errors.New("codigo_producto es obligatorio")
	}
	if strings.TrimSpace(p.Item) == "" {
		return errors.New("item es obligatorio")
	}
	if p.Cantidad < 0 {
		return errors.New("cantidad no puede ser negativa")
	}
	if strings.TrimSpace(p.Proveedor) == "" {
		return errors.New("proveedor es obligatorio")
	}
	if strings.TrimSpace(p.Ubicacion) == "" {
		return errors.New("ubicacion es obligatoria")
	}
	if strings.TrimSpace(p.Fecha) == "" {
		return errors.New("fecha es obligatoria")
	}
	if _, err := time.Parse("2006-01-02", p.Fecha); err != nil {
		return errors.New("fecha debe tener formato YYYY-MM-DD")
	}
	return nil
}

func ValidateUpdateProduct(req dto.UpdateProductDTO) error {
	if strings.TrimSpace(req.CodigoProducto) == "" {
		return errors.New("codigo_producto es obligatorio")
	}
	if strings.TrimSpace(req.Item) == "" {
		return errors.New("item es obligatorio")
	}
	if req.Cantidad < 0 {
		return errors.New("cantidad no puede ser negativa")
	}
	if strings.TrimSpace(req.Proveedor) == "" {
		return errors.New("proveedor es obligatorio")
	}
	if strings.TrimSpace(req.Ubicacion) == "" {
		return errors.New("ubicacion es obligatoria")
	}
	if strings.TrimSpace(req.Fecha) == "" {
		return errors.New("fecha es obligatoria")
	}
	if _, err := time.Parse("2006-01-02", req.Fecha); err != nil {
		return errors.New("fecha debe tener formato YYYY-MM-DD")
	}
	return nil
}
=======
>>>>>>> b4712727b6f74dd50740129b25e27533ed3e41d9
