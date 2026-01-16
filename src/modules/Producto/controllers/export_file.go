package controllers

import (
	"bytes"
	"fmt"
	"net/url"
	"time"

	"github.com/JuhethAriza/inventory/src/modules/Producto/usecases"
	"github.com/gofiber/fiber/v2"
)

type ExportFileController struct {
	usecase *usecases.ExportFile
}

func NewExportFileController(uc *usecases.ExportFile) *ExportFileController {
	return &ExportFileController{usecase: uc}
}

func (c *ExportFileController) Run(ctx *fiber.Ctx) error {
	// Ejecutar el use case para generar el archivo Excel
	file, err := c.usecase.Execute()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Error al generar el archivo Excel: %s", err.Error()),
		})
	}

	// Asegurar que el archivo se cierre al finalizar
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Printf("Error al cerrar el archivo Excel: %v\n", err)
		}
	}()

	// Crear un buffer en memoria para el archivo Excel
	var buf bytes.Buffer
	if err := file.Write(&buf); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Error al escribir el archivo en el buffer: %s", err.Error()),
		})
	}

	// Generar el nombre del archivo con la fecha actual
	filename := fmt.Sprintf("productos_%s.xlsx", time.Now().Format("20060102_150405"))
	
	// Codificar el nombre del archivo para compatibilidad con navegadores
	// Especialmente importante para Windows y caracteres especiales
	filenameEncoded := url.QueryEscape(filename)

	// Configurar los headers de respuesta de manera compatible con todos los navegadores
	// Usar ambos formatos para máxima compatibilidad (RFC 2231 y RFC 5987)
	ctx.Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	ctx.Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"; filename*=UTF-8''%s", filename, filenameEncoded))
	ctx.Set("Content-Length", fmt.Sprintf("%d", buf.Len()))
	ctx.Set("Cache-Control", "no-cache, no-store, must-revalidate")
	ctx.Set("Pragma", "no-cache")
	ctx.Set("Expires", "0")

	// Establecer el tamaño del buffer
	ctx.Response().SetBody(buf.Bytes())

	return nil
}