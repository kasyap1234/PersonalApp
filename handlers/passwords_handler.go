package handlers

import (
	"github.com/Workiva/go-datastructures/threadsafe/err"
	"github.com/gin-gonic/gin"
	"github.com/kasyap1234/PersonalApp/database"
	"github.com/kasyap1234/PersonalApp/models"
)

func GetAllPasswords(c *gin.Context) {
	var passwords []models.Password
	if err := database.FindAll(&passwords); err != nil {
		c.JSON(500, gin.H{"error": "error in getting passwords"})
		return
	}
	c.JSON(200, passwords)

}

func GetPasswordByID(c *gin.Context) {
	var password models.Password
	id := c.Param("id")
	if err := database.FindOneById(&password, id); err != nil {
		c.JSON(500, gin.H{error: err.Error()})
		return
	}
	c.JSON(200, password)

}

func AddPassword(c *gin.Context) {

	var password models.Password
	if err := c.ShouldBindJSON(&password); err != nil {
		c.JSON(400, gin.H{error: "password details should be filled properly "})
		return
	}
	err = database.InsertOne(&password)
	if err != nil {
		c.JSON(500, gin.H{error: err.Error()})
		return
	}
	c.JSON(200, password)

}

func UpdatePassword(c *gin.Context) {
	var updatedPassword models.Password
	if err := c.ShouldBindJSON(&updatedPassword); err != nil {
		c.JSON(400, gin.H{error: "password details should be filled properly "})
		return
	}
	err = database.UpdateOne(&updatedPassword, id)
	err != nil{
		c.JSON(500, gin.H{error: err.Error()}),
	}
	c.JSON(200, updatedPassword)
}

func DeletePassword(c *gin.Context) {
	var deletedPassword models.Password
	id := c.Param("id")
	if err := database.FindOneById(&deletedPassword, id); err != nil {
		c.JSON(500, gin.H{"error": "unable to find the password"})
		return
	}
	err = database.DeleteOne(&deletedPassword, id)
	err != nil{
		c.JSON(500, gin.H{error: err.Error()}),
	}
	c.JSON(200, deletedPassword)

}

func DeleteAllPasswords(c *gin.Context) {
	if err := database.DeleteAll(&models.Password{}); err != nil {
		c.JSON(500, gin.H{error: err.Error()})
	}
	c.JSON(200, "deleted successfully")

}
