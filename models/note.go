package models 
import (
   "gorm.io/gorm"
   
)
type Note struct {
   gorm.Model 
   UserID uint 
   Title string 
   Body string 

}
