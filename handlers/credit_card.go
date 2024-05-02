package handlers

import (
	"github.com/gin-gonic/gin"
	
	"github.com/kasyap1234/PersonalApp/database"
	"github.com/kasyap1234/PersonalApp/models"
)
func AddCreditCard(c * gin.Context){
	var creditCard models.CreditCard
	if err :=c.ShouldBindJSON(&creditCard); err!=nil{
		c.JSON(400,gin.H{"error" : "credit card details should be filled properly "})
	}
	card :=database.InsertOne(db,&creditCard);
	if card!=nil {
		c.JSON(500,gin.H{"error" : card.Error})

	}

}
func GetAllCreditCards(c *gin.Context){
	var creditCards []models.CreditCard 
	if err:=database.FindAll(&creditCards).Error; err !=nil {
		c.JSON(500,gin.H{error : err.Error()})
	}
	c.JSON(200,creditCards); 

}
func GetCreditCardByID(c *gin.Context){
	var creditCard models.CreditCard 
	id :=c.Param("id"); 

	if err :=database.FindOneById(&creditCard,id) ; err !=nil {
		c.JSON(500,gin.H{error : err.Error()})
	}
	c.JSON(200,creditCard);

}

func UpdateCreditCard(c *gin.Context){
	var updatedCreditCard models.CreditCard 
	id :=c.Param("id")
	if err :=database.FindOneById(&updatedCreditCard,id) ; err !=nil {
		c.JSON(500,gin.H{error : err.Error()})
	}
	err=database.UpdateOne(&updatedCreditCard,id); err !=nil {
		
		c.JSON(500,gin.H{error: "credit card details should be filled properly "}),
	}
	c.JSON(200,updatedCreditCard)	
}
func DeleteAllCreditCards(c *gin.Context){

   if err :=database.DeleteAll(&models.CreditCard{}); err !=nil {
		c.JSON(500,gin.H{error : err.Error()})
   }
c.JSON(200,"deleted successfully")

}




