package usecases

import (
	"fmt"

	dao "github.com/JuhethAriza/inventory/src/infrastructure/db/dao"
	"github.com/xuri/excelize/v2"
)

type ExportFile struct {
	repo *dao.MySQLProductDao
}

func NewExportFile(repo *dao.MySQLProductDao) *ExportFile {
	return &ExportFile{repo: repo}
}

func (uc *ExportFile) Execute() (*excelize.File, error) {
	// Obtener todos los productos
	products, err := uc.repo.GetAllProducts()
	if err != nil {
		return nil, fmt.Errorf("error al obtener productos: %w", err)
	}

	// Crear un nuevo archivo Excel
	f := excelize.NewFile()
	sheetName := "Sheet1" // Usar el nombre por defecto primero

	// Definir los encabezados
	headers := []string{"ID", "Código Producto", "Item", "Cantidad", "Categoría", "Estado", "Valor Estimado", "Fecha", "Ubicación"}
	
	// Estilo para los encabezados
	headerStyle, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
			Size: 12,
		},
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#E0E0E0"},
			Pattern: 1,
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
	})
	if err != nil {
		return nil, fmt.Errorf("error al crear estilo de encabezado: %w", err)
	}

	// Escribir los encabezados en la primera fila
	for col, header := range headers {
		cellName, _ := excelize.CoordinatesToCellName(col+1, 1)
		if err := f.SetCellValue(sheetName, cellName, header); err != nil {
			return nil, fmt.Errorf("error al escribir encabezado en %s: %w", cellName, err)
		}
		if err := f.SetCellStyle(sheetName, cellName, cellName, headerStyle); err != nil {
			return nil, fmt.Errorf("error al aplicar estilo de encabezado: %w", err)
		}
	}

	// Escribir los datos de los productos
	for row, product := range products {
		excelRow := row + 2 // Empezar en la fila 2 (después de los encabezados)

		// Escribir cada campo en su columna correspondiente
		// Columna A (ID)
		cellA, _ := excelize.CoordinatesToCellName(1, excelRow)
		if err := f.SetCellValue(sheetName, cellA, product.ID); err != nil {
			return nil, fmt.Errorf("error al escribir ID en %s: %w", cellA, err)
		}

		// Columna B (Código Producto)
		cellB, _ := excelize.CoordinatesToCellName(2, excelRow)
		if err := f.SetCellValue(sheetName, cellB, product.CodigoProducto); err != nil {
			return nil, fmt.Errorf("error al escribir Código Producto: %w", err)
		}

		// Columna C (Item)
		cellC, _ := excelize.CoordinatesToCellName(3, excelRow)
		if err := f.SetCellValue(sheetName, cellC, product.Item); err != nil {
			return nil, fmt.Errorf("error al escribir Item: %w", err)
		}

		// Columna D (Cantidad)
		cellD, _ := excelize.CoordinatesToCellName(4, excelRow)
		if err := f.SetCellValue(sheetName, cellD, product.Cantidad); err != nil {
			return nil, fmt.Errorf("error al escribir Cantidad: %w", err)
		}

		// Columna E (Categoría)
		cellE, _ := excelize.CoordinatesToCellName(5, excelRow)
		if err := f.SetCellValue(sheetName, cellE, product.Categoria); err != nil {
			return nil, fmt.Errorf("error al escribir Categoría: %w", err)
		}

		// Columna F (Estado)
		cellF, _ := excelize.CoordinatesToCellName(6, excelRow)
		if err := f.SetCellValue(sheetName, cellF, product.Estado); err != nil {
			return nil, fmt.Errorf("error al escribir Estado: %w", err)
		}

		// Columna G (Valor Estimado - usando el campo Proveedor como valor)
		cellG, _ := excelize.CoordinatesToCellName(7, excelRow)
		if err := f.SetCellValue(sheetName, cellG, product.Proveedor); err != nil {
			return nil, fmt.Errorf("error al escribir Valor Estimado: %w", err)
		}

		// Columna H (Fecha)
		cellH, _ := excelize.CoordinatesToCellName(8, excelRow)
		if err := f.SetCellValue(sheetName, cellH, product.Fecha); err != nil {
			return nil, fmt.Errorf("error al escribir Fecha: %w", err)
		}

		// Columna I (Ubicación)
		cellI, _ := excelize.CoordinatesToCellName(9, excelRow)
		if err := f.SetCellValue(sheetName, cellI, product.Ubicacion); err != nil {
			return nil, fmt.Errorf("error al escribir Ubicación: %w", err)
		}
	}

	// Renombrar la hoja después de escribir los datos
	if err := f.SetSheetName(sheetName, "Productos"); err != nil {
		return nil, fmt.Errorf("error al renombrar hoja: %w", err)
	}
	sheetName = "Productos"

	// Ajustar el ancho de las columnas automáticamente
	for col := 1; col <= len(headers); col++ {
		colName, _ := excelize.ColumnNumberToName(col)
		if err := f.SetColWidth(sheetName, colName, colName, 15); err != nil {
			return nil, fmt.Errorf("error al ajustar ancho de columna: %w", err)
		}
	}

	// Aplicar altura a la fila de encabezados
	if err := f.SetRowHeight(sheetName, 1, 20); err != nil {
		return nil, fmt.Errorf("error al establecer altura de fila: %w", err)
	}

	return f, nil
}