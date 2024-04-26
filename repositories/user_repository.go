package repositories 
import (
	"securestorage/models"
	"gorm.io/gorm"
)
type UserRepository struct {
	db *gorm.DB
}
func NewUserRepository(db *gorm.DB) *UserRepository{
	return &UserRepository{db}
}
func (ur *UserRepsoitory) CreateUser(user *models.User) error {
	return ur.db.Create(user).Error ; 

}
func (ur *UserRepository) GetUserByEmail(email string) ( *models.User,err error){
	var user model.User 
	err =ur.db.Where("email = ?",email).First(&user).Error
	return &user ,err ; 

}
