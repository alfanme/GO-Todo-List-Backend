package models

import "gorm.io/gorm"


type Todo struct {
	gorm.Model
	Title string
	Done bool
}


type AddTodoInput struct {
	Title string `json:"title" binding:"required"`
}


type UpdateTodoInput struct {
	Title string `json:"title" binding:"required"`
	Done bool `json:"done" binding:"required"`
}