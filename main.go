package main

import (
	"log"
	"os"
	"todo_list/controllers"
	"todo_list/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	dotEnvErr := godotenv.Load()
	if dotEnvErr != nil {
		log.Fatal("Error loading .env file")
	}

	models.ConnectDatabase()
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/", controllers.Home)
	r.GET("/api/todos", controllers.GetTodos)
	r.GET("/api/todos/:id", controllers.GetTodo)
	r.POST("/api/todos", controllers.AddTodo)
	r.PATCH("/api/todos/:id", controllers.UpdateTodo)
	r.DELETE("/api/todos/:id", controllers.DeleteTodo)

	r.Run(os.Getenv("PORT"))
}