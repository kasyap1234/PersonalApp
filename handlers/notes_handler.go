package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/kasyap1234/PersonalApp/database"
	"github.com/kasyap1234/PersonalApp/models"
)
func AddNote(c *gin.Context){
	var note models.Note 
	if err :=c.ShouldBindJSON(&note); err!=nil{
		c.JSON(400,gin.H{error : "note details should be filled properly "})
		return 
	}
	err=database.InsertOne(&note); 
	if err!=nil {
		c.JSON(500,gin.H{error: "error with note details "})
		return 
	}
	c.JSON(200,note); 

}
func GetAllNotes(c *gin.Context){
	var notes []model.Note 
    if err :=database.FindAll(&notes); err!=nil {
		c.JSON(500,gin.H{error : err.Error()})
		return 
	}
c.JSON(200,notes); 

}
func GetNotesByID(c *gin.Context){
	var note models.Note
	id=c.Param("id"); 
	if err:=database.FindOneById(&note,id); err!=nil {
		c.JSON(500,gin.H{error : err.Error()})
		return
	}
	c.JSON(200,note); 


}

