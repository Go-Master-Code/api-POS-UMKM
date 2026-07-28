package report

// supaya semua report memakai struct yang sama
type CompanyInfo struct {
	Name    string
	Address string
	Phone   string
}

type Metadata struct {
	StartPeriod string
	EndPeriod   string
	PrintDate   string
	GeneratedBy string
}

type SummaryItem struct {
	Label string
	Value string
}

type TableColumn struct {
	Title string
	Width float64
	Align string
}
