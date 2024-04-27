package database 
import (
	"gorm.io/driver/postgres"

)
type Database struct {
	db *gorm.DB
}
func NewDatabase(db *gorm.db){
	return &Database{db:db}
}
func(dtb *Database) CreateUser(user *models.User) error {
	return dtb.db.Create(user).Error ; 
}

func(dtb *Database) GetUserById(id uint)(*models.User,error){
	var user models.User 
	if err := dtb.db.First(&user,id).Error ; 
	err !=nil {
		return nil,err 
	}
	return &user,nil; 
}
func(dtb *Database) FindUserByEmail(email string)(*models.User,error){
	var user models.User 
	if err :=dtb.db.Where("email = ?",email).First(&user).Error ; 
	err !=nil {
		return nil,err 
	}
	return &user,nil;
}
func(dtb *Database) UpdateUser(user *models.User)error {
	return dtb.db.Save(user).Error; 
}
func(dtb *Database) DeleteUser(user *models.User)error {
	return dtb.db.Delete(&models.User{}).Error;
}
