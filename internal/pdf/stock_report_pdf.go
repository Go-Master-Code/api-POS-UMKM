package pdf

import (
	"bytes"
	"strconv"
	"umkm-odod/helper"
	"umkm-odod/internal/dto"
	"umkm-odod/internal/report"
)

func GenerateStockReport(stock []dto.StockReportResponse, companyInfo report.CompanyInfo, query dto.StockReportQuery, summary *dto.StockReportSummary, printedBy string) (*bytes.Buffer, error) {
	// auto generate pdf dari method NewPDF (internal/report/pdf.go)
	pdf := report.NewPDF()

	report.DrawHeader(pdf, "STOCK REPORT")

	report.VerticalSpace2Style(pdf)

	// company info
	name := companyInfo.Name
	address := companyInfo.Address
	phone := companyInfo.Phone

	report.DrawCompany(pdf, report.CompanyInfo{
		Name:    name,
		Address: address,
		Phone:   phone,
	})

	report.DrawMetaData(pdf, report.Metadata{
		GeneratedBy: printedBy,
	})

	report.DrawSeparatorLine(pdf)

	report.DrawSummary(pdf)

	// total transaction dibuat tanpa currency Rp, jadi parameter terakhir false
	report.DrawSummaryFeld(pdf, 35, "Total Variants", strconv.Itoa(int(summary.TotalVariants)), false)
	report.DrawSummaryFeld(pdf, 35, "Low Stock Items", strconv.Itoa(int(summary.LowStockItems)), false)

	report.DrawSeparatorLine(pdf)

	report.VerticalSpace2Style(pdf)

	report.DrawHeader(pdf, "DETAIL TRANSACTION")
	report.VerticalSpace4Style(pdf)

	// =====================================================
	// TABLE HEADER
	// =====================================================

	// report.TableHeaderStyle(pdf) => sudah diatur di writer.go

	headers := []struct {
		Title string
		Width float64
	}{
		{"SKU", 30},
		{"Category", 30},
		{"Item", 40},
		{"Variant", 50},
		{"Stock", 20},
		{"Min. Stock", 20},
	}

	for _, h := range headers {
		report.DrawTableHeader(pdf, h.Width, h.Title)
	}

	pdf.Ln(-1) // line break seperti enter

	// isikan data row stock
	for _, row := range stock {
		report.DrawTableBody(pdf, 30, row.SKU, "C")
		report.DrawTableBody(pdf, 30, row.CategoryName, "C")
		report.DrawTableBody(pdf, 40, row.ItemName, "L")
		report.DrawTableBody(pdf, 50, row.VariantName, "L")
		report.DrawTableBody(pdf, 20, helper.FormatRupiah(row.CurrentStock), "R")
		report.DrawTableBody(pdf, 20, helper.FormatRupiah(row.MinimumStock), "R")

		pdf.Ln(-1) // line break seperti enter
	}

	var buf bytes.Buffer

	return report.WritePDFReport(pdf, buf)
}
