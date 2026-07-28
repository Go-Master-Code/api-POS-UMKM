package report

import "github.com/go-pdf/fpdf"

// Judul laporan
func TitleStyle(pdf *fpdf.Fpdf) {
	pdf.SetFont("Arial", "B", 16)
}

// Header table
func TableHeaderStyle(pdf *fpdf.Fpdf) {
	pdf.SetFont("Arial", "B", 10)
}

// Table content
func BodyStyle(pdf *fpdf.Fpdf) {
	pdf.SetFont("Arial", "", 10)
}

// Summary style
func SummaryStyle(pdf *fpdf.Fpdf) {
	pdf.SetFont("Arial", "B", 16)
}

// Summary item
func SummaryContentStyle(pdf *fpdf.Fpdf) {
	pdf.SetFont("Arial", "", 12)
}

// vertical space 2mm
func VerticalSpace2Style(pdf *fpdf.Fpdf) {
	pdf.Ln(2)
}

// vertical space 4mm
func VerticalSpace4Style(pdf *fpdf.Fpdf) {
	pdf.Ln(4)
}

// garis pemisah
func DrawSeparatorLine(pdf *fpdf.Fpdf) {
	pdf.Ln(2)
	pdf.CellFormat(0, 0, "", "T", 1, "", false, 0, "")
	pdf.Ln(2)
}
