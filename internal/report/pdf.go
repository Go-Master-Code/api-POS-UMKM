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

// newReceiptPDF membuat instance PDF dengan konfigurasi lebar kertas 80 mm
func NewReceiptPDF() *fpdf.Fpdf {
	pdf := fpdf.New("P", "mm", "", "")

	pdf.AddPageFormat(
		"P",
		fpdf.SizeType{
			Wd: 80,  // lebar 80 mm
			Ht: 200, // walau height 200 mm, printer hanya akan print bagian kertas yang ada isi teksnya
		},
	)

	return pdf
}
