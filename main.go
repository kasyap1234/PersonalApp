package main

import (
	"fmt"


	"github.com/kasyap1234/PersonalApp/database"
    "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main(){
var err error 
dsn := "host=localhost user=gorm dbname=gorm password=gorm sslmode=disable"
database.DB,err =gorm.Open(postgres.Open(dsn),&gorm.Config{})
if err!=nil {
	fmt.Println("Failed to connect to the database",err); 
	panic("Failed to connect to the database")

}



}
