package report

import "github.com/go-pdf/fpdf"

// newPDF membuat instance PDF dengan konfigurasi standar
func NewPDF() *fpdf.Fpdf {
	pdf := fpdf.New("P", "mm", "A4", "")

	// margin kiri, atas, kanan
	pdf.SetMargins(10, 10, 10)

	// auto page break
	pdf.SetAutoPageBreak(true, 10)

	// WAJIB add page
	pdf.AddPage()

	return pdf
}
