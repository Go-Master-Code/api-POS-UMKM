package pdf

import (
	"bytes"
	"strconv"
	"umkm-odod/helper"
	"umkm-odod/internal/dto"
	"umkm-odod/internal/model"
	"umkm-odod/internal/report"
)

func GenerateExpensesReport(expense []model.Expenses, company report.CompanyInfo, query dto.ExpenseReportQuery, summary *dto.ExpenseReportSummary, printedBy string) (*bytes.Buffer, error) {
	// auto generate pdf dari method NewPDF (internal/report/pdf.go)
	pdf := report.NewPDF()

	report.DrawHeader(pdf, "EXPENSE REPORT")

	report.VerticalSpace2Style(pdf)

	// company info
	report.DrawCompany(pdf, company)

	report.DrawMetaData(pdf, report.Metadata{
		StartPeriod: query.StartDate,
		EndPeriod:   query.EndDate,
		GeneratedBy: printedBy,
	})

	report.DrawSeparatorLine(pdf)

	report.DrawSummary(pdf)

	// total transaction dibuat tanpa currency Rp, jadi parameter terakhir false
	report.DrawSummaryFeld(pdf, 35, "Total Transaction", strconv.Itoa(int(summary.TotalTransaction)), false)
	report.DrawSummaryFeld(pdf, 35, "Total QRIS", helper.FormatRupiah(summary.TotalQRIS), true)
	report.DrawSummaryFeld(pdf, 35, "Total Cash", helper.FormatRupiah(summary.TotalCash), true)
	report.DrawSummaryFeld(pdf, 35, "Total Expense", helper.FormatRupiah(summary.TotalExpense), true)

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
		{"Payment", 20},
		{"Notes", 91},
		{"Grand Total", 24},
	}

	for _, h := range headers {
		// pdf.CellFormat(h.Width, 8, h.Title, "1", 0, "C", false, 0, "")
		report.DrawTableHeader(pdf, h.Width, h.Title)
	}

	pdf.Ln(-1) // line break seperti enter

	// report.BodyStyle(pdf) => sudah diatur di writer

	// cek jumlah record expense, jika nol maka generate report kosong
	if len(expense) == 0 {
		report.DrawTableBody(pdf, 190, "No expenses data for the selected period.", "C")
	} else {
		for _, row := range expense {
			report.DrawTableBody(pdf, 30, row.ExpenseNumber, "C")
			report.DrawTableBody(pdf, 25, row.CreatedAt.Format("02 Jan 2006"), "C")
			report.DrawTableBody(pdf, 20, row.PaymentMethod, "C")
			report.DrawTableBody(pdf, 91, row.Notes, "L")
			report.DrawTableBody(pdf, 24, helper.FormatRupiah(row.TotalAmount), "R")

			pdf.Ln(-1) // line break seperti enter
		}
	}

	var buf bytes.Buffer

	return report.WritePDFReport(pdf, buf)
}
