package models 
import (
	"gorm.io/gorm"
)
type Password struct {
	gorm.Model 
	UserID uint 
	Website string 
	Username string 
	Password string 
}