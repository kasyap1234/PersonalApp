package models 
type Note struct {
   gorm.Model 
   UserID uint 
   Title string 
   Body string 

}
