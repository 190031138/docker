package routes

import (
    "go-todo-app/controllers"

    "github.com/gin-gonic/gin"
)

func RegisterTodoRoutes(router *gin.Engine) {
    api := router.Group("/api/todos")
    {
        api.GET("/", controllers.GetAllTodos)
        api.POST("/", controllers.AddTodo)
    }
}
