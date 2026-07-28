package pdf

import (
	"bytes"
	"umkm-odod/helper"
	"umkm-odod/internal/dto"
	"umkm-odod/internal/report"
)

func GenerateStockCardReport(stockCardReport dto.StockCardReport) (*bytes.Buffer, error) {
	// auto generate pdf dari method NewPDF (internal/report/pdf.go)
	pdf := report.NewPDF()

	report.DrawHeader(pdf, "STOCK CARD REPORT")

	report.VerticalSpace2Style(pdf)

	// company info
	name := stockCardReport.TenantName
	address := stockCardReport.TenantAddress
	phone := stockCardReport.TenantPhone

	report.DrawCompany(pdf, report.CompanyInfo{
		Name:    name,
		Address: address,
		Phone:   phone,
	})

	report.DrawMetaData(pdf, report.Metadata{
		StartPeriod: stockCardReport.StartDate.Format("2006-01-02"),
		EndPeriod:   stockCardReport.EndDate.Format("2006-01-02"),
		GeneratedBy: "operator",
	})

	report.DrawSeparatorLine(pdf)

	report.DrawSummary(pdf)

	// parameter terakhir (false) menunjukkan format writer tanpa currency
	report.DrawSummaryFeld(pdf, 25, "Category", stockCardReport.CategoryName, false)
	report.DrawSummaryFeld(pdf, 25, "Product", stockCardReport.ProductName, false)
	report.DrawSummaryFeld(pdf, 25, "Variant", stockCardReport.VariantName, false)
	report.DrawSummaryFeld(pdf, 25, "SKU", stockCardReport.SKU, false)

	report.DrawSeparatorLine(pdf)

	report.VerticalSpace2Style(pdf)

	report.DrawHeader(pdf, "STOCK CARD DETAIL")
	report.VerticalSpace4Style(pdf)

	// =====================================================
	// TABLE HEADER
	// =====================================================

	// report.TableHeaderStyle(pdf) => sudah diatur di writer.go

	headers := []struct {
		Title string
		Width float64
	}{
		{"Date", 26},
		{"Transaction", 31},
		{"In", 14},
		{"Out", 14},
		{"Balance", 18},
		{"User", 40},
		{"Note", 47},
	}

	for _, h := range headers { // format untuk headers
		report.DrawTableHeader(pdf, h.Width, h.Title)
	}

	pdf.Ln(-1) // line break seperti enter

	pdf.SetFont("Arial", "", 10) // style biasa tanpa bold untuk isi tabel

	// isikan data row stock movement
	for _, row := range stockCardReport.Items {

		report.DrawTableBody(pdf, 26, row.MovementDate.Format("02 Jan 2006"), "C")
		report.DrawTableBody(pdf, 31, row.ReferenceType, "L")
		report.DrawTableBody(pdf, 14, helper.FormatRupiah(row.QtyIn), "R")
		report.DrawTableBody(pdf, 14, helper.FormatRupiah(row.QtyOut), "R")
		report.DrawTableBody(pdf, 18, helper.FormatRupiah(row.Balance), "R")
		report.DrawTableBody(pdf, 40, row.CreatedByName, "L")
		report.DrawTableBody(pdf, 47, row.Notes, "L")

		pdf.Ln(-1) // line break seperti enter
	}

	report.VerticalSpace4Style(pdf)

	report.DrawSeparatorLine(pdf)

	report.DrawSummary(pdf)

	report.DrawSummaryFeld(pdf, 30, "Opening Stock", helper.FormatRupiah(stockCardReport.OpeningStock), false)
	report.DrawSummaryFeld(pdf, 30, "Total In", helper.FormatRupiah(stockCardReport.TotalIn), false)
	report.DrawSummaryFeld(pdf, 30, "Total Out", helper.FormatRupiah(stockCardReport.TotalOut), false)
	report.DrawSummaryFeld(pdf, 30, "Ending Stock", helper.FormatRupiah(stockCardReport.EndingStock), false)

	var buf bytes.Buffer

	return report.WritePDFReport(pdf, buf)
}
