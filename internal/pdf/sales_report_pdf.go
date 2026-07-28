package pdf

import (
	"bytes"
	"strconv"
	"umkm-odod/helper"
	"umkm-odod/internal/dto"
	"umkm-odod/internal/model"
	"umkm-odod/internal/report"
)

func GenerateSalesReport(sales []model.Sale, query dto.SaleReportQuery, summary *dto.SalesReportSummary) (*bytes.Buffer, error) {
	// auto generate pdf dari method NewPDF (internal/report/pdf.go)
	pdf := report.NewPDF()

	report.DrawHeader(pdf, "SALES REPORT")

	report.VerticalSpace2Style(pdf)

	// company info
	name := sales[0].Tenant.Name
	address := sales[0].Tenant.Address
	phone := sales[0].Tenant.Phone

	report.DrawCompany(pdf, report.CompanyInfo{
		Name:    name,
		Address: address,
		Phone:   phone,
	})

	report.DrawMetaData(pdf, report.Metadata{
		StartPeriod: query.StartDate,
		EndPeriod:   query.EndDate,
		GeneratedBy: "example",
	})

	report.DrawSeparatorLine(pdf)

	report.DrawSummary(pdf)

	// total transaction dibuat tanpa currency Rp, jadi parameter terakhir false
	report.DrawSummaryFeld(pdf, 35, "Total Transaction", strconv.Itoa(int(summary.TotalTransaction)), false)
	report.DrawSummaryFeld(pdf, 35, "Total Sales", helper.FormatRupiah(summary.TotalSales), true)
	report.DrawSummaryFeld(pdf, 35, "Total Discount", helper.FormatRupiah(summary.TotalDiscount), true)
	report.DrawSummaryFeld(pdf, 35, "Total Tax", helper.FormatRupiah(summary.TotalTax), true)
	report.DrawSummaryFeld(pdf, 35, "Grand Total", helper.FormatRupiah(summary.GrandTotal), true)

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
		{"Invoice", 30},
		{"Date", 25},
		{"Customer", 25},
		{"Cashier", 25},
		{"Subtotal", 22},
		{"Discount", 20},
		{"Tax", 19},
		{"Grand Total", 24},
	}

	for _, h := range headers {
		// pdf.CellFormat(h.Width, 8, h.Title, "1", 0, "C", false, 0, "")
		report.DrawTableHeader(pdf, h.Width, h.Title)
	}

	pdf.Ln(-1) // line break seperti enter

	// report.BodyStyle(pdf) => sudah diatur di writer

	// isikan data row sales
	for _, row := range sales {
		report.DrawTableBody(pdf, 30, row.InvoiceNumber, "C")
		report.DrawTableBody(pdf, 25, row.CreatedAt.Format("02 Jan 2006"), "C")
		report.DrawTableBody(pdf, 25, row.CustomerName, "C")
		report.DrawTableBody(pdf, 25, row.Cashier.FullName, "C")
		report.DrawTableBody(pdf, 22, helper.FormatRupiah(row.Subtotal), "R")
		report.DrawTableBody(pdf, 20, helper.FormatRupiah(row.DiscountAmount), "R")
		report.DrawTableBody(pdf, 19, helper.FormatRupiah(row.TaxAmount), "R")
		report.DrawTableBody(pdf, 24, helper.FormatRupiah(row.GrandTotal), "R")

		pdf.Ln(-1) // line break seperti enter
	}

	var buf bytes.Buffer

	return report.WritePDFReport(pdf, buf)
}
