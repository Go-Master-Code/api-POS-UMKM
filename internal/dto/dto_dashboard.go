package dto

type DashBoardSummaryResponse struct {
	TodaySales                float64 `json:"today_sales"`
	TodayTransactions         int64   `json:"today_transactions"`
	TodayExpenses             float64 `json:"today_expenses"`
	TodayExpensesTransactions int64   `json:"today_expenses_transaction"`
	TodayPurchase             float64 `json:"today_purchase"`
	TodayPurchaseTransaction  int64   `json:"today_purchase_transaction"`
	TodayProfit               float64 `json:"today_profit"`
	LowStockCount             int64   `json:"low_stock_count"`
	TotalItems                int64   `json:"total_items"`
	TotalVariants             int64   `json:"total_variants"`
	TotalSuppliers            int64   `json:"total_suppliers"`
}

// untuk chart total sales
type DailySalesChartResponse struct {
	Date       string  `json:"date"`
	TotalSales float64 `json:"total_sales"`
}

// untuk chart total sales
type DailyExpensesChartResponse struct {
	Date          string  `json:"date"`
	TotalExpenses float64 `json:"total_expenses"`
}

// untuk chart total purchase
type DailyPurchaseChartResponse struct {
	Date          string  `json:"date"`
	TotalPurchase float64 `json:"total_purchase"`
}

// top 5 produk terlaris
type TopSellingProductsResponse struct {
	ItemVariantID string `json:"item_variant_id"`
	ItemName      string `json:"item_name"`    // nama item nya
	VariantName   string `json:"variant_name"` // nama variannya
	QtySold       string `json:"qty_sold"`
}

// recent sales untuk melihat beberapa data sales terbaru
type RecentSalesResponse struct {
	SaleID        string  `json:"sale_id"`
	InvoiceNumber string  `json:"invoice_number"`
	CustomerName  string  `json:"customer_name"`
	CashierName   string  `json:"cashier_name"`
	GrandTotal    float64 `json:"grand_total"`
	CreatedAt     string  `json:"created_at"`
}
