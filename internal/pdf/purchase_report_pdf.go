package pdf

import (
	"bytes"
	"strconv"
	"umkm-odod/helper"
	"umkm-odod/internal/dto"
	"umkm-odod/internal/model"
	"umkm-odod/internal/report"
)

func GeneratePurchaseReport(purchase []model.Purchase, query dto.PurchaseReportQuery, summary *dto.PurchaseReportSummary) (*bytes.Buffer, error) {
	// auto generate pdf dari method NewPDF (internal/report/pdf.go)
	pdf := report.NewPDF()

	report.DrawHeader(pdf, "PURCHASE REPORT")

	report.VerticalSpace2Style(pdf)

	// company info
	name := purchase[0].Tenant.Name
	address := purchase[0].Tenant.Address
	phone := purchase[0].Tenant.Phone

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

	// total transaction dibuat tanpa currency Rp
	report.DrawSummaryFeld(pdf, 35, "Total Transaction", strconv.Itoa(int(summary.TotalTransaction)), false)
	report.DrawSummaryFeld(pdf, 35, "Total Purchase", helper.FormatRupiah(summary.TotalPurchase), true)
	report.DrawSummaryFeld(pdf, 35, "Total Discount", helper.FormatRupiah(summary.TotalDiscount), true)
	report.DrawSummaryFeld(pdf, 35, "Total Tax", helper.FormatRupiah(summary.TotalTax), true)
	report.DrawSummaryFeld(pdf, 35, "Grand Total", helper.FormatRupiah(summary.GrandTotal), true)

	report.DrawSeparatorLine(pdf)

	report.VerticalSpace2Style(pdf)

	report.DrawHeader(pdf, "DETAIL PURCHASE")
	report.VerticalSpace4Style(pdf)

	// =====================================================
	// TABLE HEADER
	// =====================================================
	headers := []struct {
		Title string
		Width float64
	}{
		{"Invoice", 30},
		{"Date", 25},
		{"Supplier", 25},
		{"Operator", 25},
		{"Subtotal", 22},
		{"Discount", 20},
		{"Tax", 19},
		{"Grand Total", 24},
	}

	for _, h := range headers {
		report.DrawTableHeader(pdf, h.Width, h.Title)
	}

	pdf.Ln(-1) // seperti enter

	// isikan data row purchase
	for _, row := range purchase {
		report.DrawTableBody(pdf, 30, row.InvoiceNumber, "C")
		report.DrawTableBody(pdf, 25, row.CreatedAt.Format("02 Jan 2006"), "C")
		report.DrawTableBody(pdf, 25, row.Supplier.Name, "C")
		report.DrawTableBody(pdf, 25, row.Creator.FullName, "C")
		report.DrawTableBody(pdf, 22, helper.FormatRupiah(row.Subtotal), "R")
		report.DrawTableBody(pdf, 20, helper.FormatRupiah(row.DiscountAmount), "R")
		report.DrawTableBody(pdf, 19, helper.FormatRupiah(row.TaxAmount), "R")
		report.DrawTableBody(pdf, 24, helper.FormatRupiah(row.GrandTotal), "R")
		pdf.Ln(-1) // line break seperti enter
	}

	var buf bytes.Buffer

	return report.WritePDFReport(pdf, buf)
}
