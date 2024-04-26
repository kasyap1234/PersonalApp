package models 
type CreditCard struct {
	gorm.Model 
	UserID uint 
	CardHolder string 
	CardNumber string
	ExpiryDate string
	CVV string
	BillingAddress string 
	ZipCode string
}
