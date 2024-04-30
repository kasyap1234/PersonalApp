package handlers

import (
	"github.com/kasyap1234/PersonalApp/database"
	"github.com/kasyap1234/PersonalApp/models"
)

// crud operations
// create
func AddTodo(c *gin.Context){
	var todo models.Todo 
	if err :=c.ShouldBindJSON(&todo); err!=nil {
		c.JSON(400,gin.H{error : "todo details should be filled properly "})
		return
	}
	err =database.InsertOne(&todo) 
	if err !=nil {
		c.JSON(500,gin.H{error : err.Error()})
		return
	}
	c.JSON(200,todo)
}

func GetAllTodos(c *gin.Context){
	var todos []models.Todo 
	if err :=database.FindAll(&todos); err !=nil {
		c.JSON(500,gin.H{error : err.Error()})
		return 
	}
	c.JSON(200,todos)

}
func GetTodoByID(c *gin.Context){
	var todo models.Todo
	id :=c.Param("id")
	if err :=database.FindOneById(&todo,id); err !=nil {
		c.JSON(500,gin.H{error : err.Error()})
		return
	}
	c.JSON(200,todo)

}

// read or find 

func UpdateTodo(c *gin.Context){
	var updateTodo models.Todo 
	id =c.Param("id")
	if err :=database.UpdateOne(&updateTodo,id); err !=nil {
		c.JSON(500,gin.H{error : err.Error()})
		return 
}
c.JSON(200,udpateTodo); 
}

// update 

func DeleteAllTodos(c *gin.Context){
	if err:=database.DeleteAll()
}

// delete 
