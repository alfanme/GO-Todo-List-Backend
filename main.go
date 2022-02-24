package main

import (
	"log"
	"os"
	"time"
	"todo_list/controllers"
	"todo_list/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("APP_ENV") != "production" {
		dotEnvErr := godotenv.Load()
		if dotEnvErr != nil {
			log.Fatal("Error loading .env file")
		}
	}

	models.ConnectDatabase()
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost", "https://todo-list-golang.herokuapp.com/"},
		AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	  }))

	r.GET("/", controllers.Home)
	r.GET("/api/todos", controllers.GetTodos)
	r.GET("/api/todos/:id", controllers.GetTodo)
	r.POST("/api/todos", controllers.AddTodo)
	r.PATCH("/api/todos/:id", controllers.UpdateTodo)
	r.DELETE("/api/todos/:id", controllers.DeleteTodo)

	port := os.Getenv("API_PORT")
	r.Run(":" + port)
}