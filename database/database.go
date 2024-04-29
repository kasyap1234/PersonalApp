package database 

import (
	"gorm.io/gorm"
)

// FindAll retrieves all records from a given collection (table).
func FindAll(db *gorm.DB, out interface{}) error {
	return db.Find(out).Error
}

// FindOneById retrieves a single record by its ID.
func FindOneById(db *gorm.DB, out interface{}, id uint) error {
	return db.First(out, id).Error
}

// InsertOne inserts a new record into the database.
func InsertOne(db *gorm.DB, in interface{}) error {
	return db.Create(in).Error
}

// UpdateOne updates a record in the database.
func UpdateOne(db *gorm.DB,in interface {}, id uint ){

	db.Model(in).Where("id = ?", id).Updates(in)
	
}


// DeleteOne deletes a record from the database by its ID.
func DeleteOne(db *gorm.DB, in interface{}, id uint) error {
	return db.Delete(in, id).Error
}
func DeleteAll(db *gorm.DB, in interface{}) error {
    return db.Delete(in).Error
}


