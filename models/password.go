package models 
type Password struct {
	gorm.Model 
	UserID uint 
	Website string 
	Username string 
	Password string 
}