package controllers

import (
	"net/http"
	"todo_list/models"

	"github.com/gin-gonic/gin"
)


func Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H {
		"message": "Hi there! This is a Todo List API",
	})
}


func GetTodos(c *gin.Context) {
	var todos []models.Todo
	if err := models.DB.Find(&todos).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": todos})
}


func GetTodo(c *gin.Context) {
	var todo models.Todo
	if err := models.DB.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todo not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": todo})
}


func AddTodo(c *gin.Context) {
	var input models.AddTodoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo := models.Todo{Title: input.Title, Done: false}
	models.DB.Create(&todo)

	c.JSON(http.StatusOK, gin.H{"data": todo})
}


func UpdateTodo(c *gin.Context) {
	var todo models.Todo
	if err := models.DB.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Update failed, todo not found!"})
		return
	}

	var input models.UpdateTodoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&todo).Updates(map[string]interface{}{"title": input.Title, "done": input.Done})

	c.JSON(http.StatusOK, gin.H{"data": todo})
}


func DeleteTodo(c *gin.Context) {
	var todo models.Todo
	if err := models.DB.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Deletion failed, todo not found!"})
		return
	}

	models.DB.Delete(&todo)

	c.JSON(http.StatusOK, gin.H{"message": "Deletion success!"})
}