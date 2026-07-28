package report

import (
	"bytes"
	"fmt"
	"time"

	"github.com/go-pdf/fpdf"
)

func JustifyWriter(pdf *fpdf.Fpdf, tabStop float64, key string, value string) {
	pdf.CellFormat(tabStop, 7, key, "", 0, "L", false, 0, "") // kasih space sesuai tabstop pada labelnya
	pdf.CellFormat(4, 7, ":", "", 0, "C", false, 0, "")
	pdf.CellFormat(5, 7, value, "", 1, "L", false, 0, "")
}

// DrawCompany menampilkan informasi tenant
func DrawCompany(pdf *fpdf.Fpdf, info CompanyInfo) {
	SummaryContentStyle(pdf)

	JustifyWriter(pdf, 25, "Company", info.Name)
	JustifyWriter(pdf, 25, "Address", info.Address)
	JustifyWriter(pdf, 25, "Phone", info.Phone)
}

// DrawMetaData menampilkan informasi tenant
func DrawMetaData(pdf *fpdf.Fpdf, metaData Metadata) {
	SummaryContentStyle(pdf)

	// jika ada periode
	if metaData.StartPeriod != "" && metaData.EndPeriod != "" {
		JustifyWriter(pdf, 25, "Periode", fmt.Sprintf("%s until %s", metaData.StartPeriod, metaData.EndPeriod))
	}

	// tanggal hari ini
	now := time.Now().Format("2006-01-02 15:04:05")

	// printed on
	JustifyWriter(pdf, 25, "Printed On", now)

	// printed by
	JustifyWriter(pdf, 25, "Printed By", metaData.GeneratedBy)
}

// Tulis Judul Laporan
func DrawHeader(pdf *fpdf.Fpdf, title string) {
	TitleStyle(pdf)
	pdf.CellFormat(0, 7, title, "", 1, "", false, 0, "")
}

// Tulis Judul SUMMARY
func DrawSummary(pdf *fpdf.Fpdf) {
	SummaryStyle(pdf)
	pdf.Ln(2)
	pdf.CellFormat(0, 7, "SUMMARY", "", 1, "", false, 0, "")
	pdf.Ln(2)
}

// summary field writer
func DrawSummaryFeld(pdf *fpdf.Fpdf, tabStop float64, key string, value string, currency bool) {
	SummaryContentStyle(pdf)
	pdf.CellFormat(tabStop, 7, key, "", 0, "L", false, 0, "") // kasih space 35 untuk labelnya
	pdf.CellFormat(5, 7, ":", "", 0, "C", false, 0, "")

	alignment := "L" // alignment default untuk non currency

	if currency { // jika tipe nya currency, pakai Rp
		alignment = "R" // align right untuk angka
		pdf.CellFormat(5, 7, "Rp", "", 0, "L", false, 0, "")
	}
	pdf.CellFormat(25, 7, value, "", 1, alignment, false, 0, "")
}

// table header and body writer
func DrawTableHeader(pdf *fpdf.Fpdf, width float64, title string) {
	TableHeaderStyle(pdf)
	pdf.CellFormat(width, 8, title, "1", 0, "C", false, 0, "")
}

func DrawTableBody(pdf *fpdf.Fpdf, width float64, value string, alignment string) {
	BodyStyle(pdf)
	pdf.CellFormat(width, 8, value, "1", 0, alignment, false, 0, "")
}

// writer report ke pdf
func WritePDFReport(pdf *fpdf.Fpdf, buf bytes.Buffer) (*bytes.Buffer, error) {
	err := pdf.Output(&buf)
	if err != nil {
		return nil, err
	}
	return &buf, nil
}
