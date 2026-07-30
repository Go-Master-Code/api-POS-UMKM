package dto

type TenantProfileResponse struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	OwnerName     string `json:"owner_name"`
	Phone         string `json:"phone"`
	Email         string `json:"email"`
	Address       string `json:"address"`
	Logo          string `json:"logo"`
	Currency      string `json:"currency"`
	TimeZone      string `json:"time_zone"`
	ReceiptFooter string `json:"receipt_footer"`
}

// struct update profile tenant
type UpdateTenantProfileRequest struct {
	Name          *string `json:"name" binding:"omitempty,max=150"`
	OwnerName     *string `json:"owner_name" binding:"omitempty,max=150"`
	Phone         *string `json:"phone" binding:"omitempty,max=30"`
	Email         *string `json:"email" binding:"omitempty,email,max=150"`
	Address       *string `json:"address" binding:"omitempty"`
	Logo          *string `json:"logo"`
	Currency      *string `json:"currency" binding:"omitempty,len=3"`
	TimeZone      *string `json:"time_zone" binding:"omitempty,max=100"`
	ReceiptFooter *string `json:"receipt_footer" binding:"omitempty"`
}
