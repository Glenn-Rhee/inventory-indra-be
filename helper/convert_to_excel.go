package helper

import (
	"fmt"
	"inventory-indra/model"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

func ConvertToExcel(products []model.DataProductResponse, ctx *gin.Context) (error) {
	f := excelize.NewFile()
	sheet := "Data Products"
	f.NewSheet(sheet)
	f.DeleteSheet("Sheet1")

	headerStyle, _ := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{Bold: true, Color: "FFFFFF"},
		Fill: excelize.Fill{Type: "pattern", Color: []string{"4472C4"}, Pattern: 1},
		Alignment: &excelize.Alignment{Horizontal: "center"},
	})

	headers := []string{"ID", "Nama Obat", "Kategori", "Harga/Butir", "Tanggal Kadaluarsa", "Stok/Butir", "Update Terakhir"}
	for i, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheet, cell, h)
		f.SetCellStyle(sheet, cell, cell, headerStyle)
	}

	for rowIdx, product := range products {
		row := rowIdx + 2
		f.SetCellValue(sheet, fmt.Sprintf("A%d", row), product.Id)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", row), product.ProductName)
		f.SetCellValue(sheet, fmt.Sprintf("C%d", row), product.Category)
		f.SetCellValue(sheet, fmt.Sprintf("D%d", row), product.PricePerButir)
		f.SetCellValue(sheet, fmt.Sprintf("E%d", row), product.ExpiredDate.Format("2006-01-02"))
		f.SetCellValue(sheet, fmt.Sprintf("F%d", row), product.StockPerButir)
		f.SetCellValue(sheet, fmt.Sprintf("G%d", row), product.LastUpdate.Format("2006-01-02 15:04:05"))
	}

    cols := []string{"A", "B", "C", "D", "E", "F", "G"}
    widths := []float64{15, 25, 20, 15, 15, 12, 20}
	for i, col := range cols {
		f.SetColWidth(sheet, col, col, widths[i])
	}

	fileName := fmt.Sprintf("Data-Obat-%s.xlsx", time.Now().Format("20060102"))
	ctx.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	ctx.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))

	return f.Write(ctx.Writer)
}