package pdf

import (
	"bytes"
	"fmt"
	"strconv"
	"time"
	"umkm-odod/helper"
	"umkm-odod/internal/dto"

	"github.com/go-pdf/fpdf"
)

func GenerateStockReport(stock []dto.StockReportResponse, query dto.StockReportQuery, summary *dto.StockReportSummary) (*bytes.Buffer, error) {
	pdf := fpdf.New("P", "mm", "A4", "")

	pdf.AddPage() // WAJIB

	pdf.SetFont("Arial", "B", 16)
	pdf.CellFormat(0, 7, "STOCK REPORT", "", 1, "", false, 0, "")

	pdf.Ln(2) // beri jarak 24mm

	// tanggal hari ini
	currentDate := time.Now().Format("2006-01-02 15:04:05")

	pdf.SetFont("Arial", "", 12)
	pdf.CellFormat(0, 4, fmt.Sprintf("Printed On: %s", currentDate), "", 1, "", false, 0, "")
	pdf.Ln(2) // beri jarak 2mm

	// garis pemisah
	pdf.Ln(2)
	pdf.CellFormat(0, 0, "", "T", 1, "", false, 0, "")
	pdf.Ln(2)

	pdf.Ln(2) // beri jarak 4mm
	pdf.SetFont("Arial", "B", 16)
	pdf.CellFormat(0, 7, "SUMARRY", "", 1, "", false, 0, "")
	pdf.Ln(2) // beri jarak 4mm

	// total transaction dibuat tanpa currency Rp
	pdf.SetFont("Arial", "", 12)
	pdf.CellFormat(35, 7, "Total Variants", "", 0, "L", false, 0, "") // kasih space 35 untuk labelnya
	pdf.CellFormat(4, 7, ":", "", 0, "C", false, 0, "")
	pdf.CellFormat(5, 7, strconv.Itoa(int(summary.TotalVariants)), "", 1, "L", false, 0, "")

	pdf.CellFormat(35, 7, "Low Stock Item(s)", "", 0, "L", false, 0, "") // kasih space 35 untuk labelnya
	pdf.CellFormat(4, 7, ":", "", 0, "C", false, 0, "")
	pdf.CellFormat(5, 7, strconv.Itoa(int(summary.LowStockItems)), "", 1, "L", false, 0, "")
	pdf.Ln(2) // beri jarak 4mm

	// garis pemisah
	pdf.Ln(2)
	pdf.CellFormat(0, 0, "", "T", 1, "", false, 0, "")
	pdf.Ln(2)

	pdf.Ln(2) // beri jarak 2mm

	// row sales
	pdf.SetFont("Arial", "B", 16)
	pdf.CellFormat(0, 6, "STOCK LIST", "", 1, "", false, 0, "")
	pdf.Ln(4) // beri jarak 4mm

	// =====================================================
	// TABLE HEADER
	// =====================================================

	pdf.SetFont("Arial", "B", 10) // font style BOLD untuk header table

	headers := []struct {
		Title string
		Width float64
	}{
		{"SKU", 30},
		{"Category", 30},
		{"Item", 45},
		{"Variant", 45},
		{"Stock", 20},
		{"Min. Stock", 20},
	}

	for _, h := range headers {
		pdf.CellFormat(h.Width, 8, h.Title, "1", 0, "C", false, 0, "")
	}

	pdf.Ln(-1)

	pdf.SetFont("Arial", "", 10) // style biasa tanpa bold untuk isi tabel

	// isikan data row sales
	for _, row := range stock {
		pdf.CellFormat(30, 8, row.SKU, "1", 0, "C", false, 0, "")
		pdf.CellFormat(30, 8, row.CategoryName, "1", 0, "C", false, 0, "")
		pdf.CellFormat(45, 8, row.ItemName, "1", 0, "L", false, 0, "")
		pdf.CellFormat(45, 8, row.VariantName, "1", 0, "Ls", false, 0, "")
		pdf.CellFormat(20, 8, helper.FormatRupiah(row.CurrentStock), "1", 0, "R", false, 0, "")
		pdf.CellFormat(20, 8, helper.FormatRupiah(row.MinimumStock), "1", 0, "R", false, 0, "")
		pdf.Ln(-1) // line break seperti enter
	}

	var buf bytes.Buffer

	err := pdf.Output(&buf)
	if err != nil {
		return nil, err
	}

	return &buf, nil
}
