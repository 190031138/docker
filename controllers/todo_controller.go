package controllers

import (
    "net/http"
    "go-todo-app/models"
    "go-todo-app/services"

    "github.com/gin-gonic/gin"
)

func GetAllTodos(c *gin.Context) {
    todos := services.GetTodos()
    c.JSON(http.StatusOK, todos)
}

func AddTodo(c *gin.Context) {
    var todo models.Todo
    if err := c.ShouldBindJSON(&todo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    created, err := services.CreateTodo(todo)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, created)
}
