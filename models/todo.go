package models 
import (
	"gorm.io/gorm"
	"time"
	
)
type Todo struct {
	gorm.Model 
	UserId uint 
	Title string 
	Completed bool 
	CompletedAt time.Time
	
}