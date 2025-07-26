package main

import (
    "go-todo-app/config"
    "go-todo-app/routes"

    "github.com/gin-gonic/gin"
)

func main() {
    config.ConnectDB()

    router := gin.Default()
    routes.RegisterTodoRoutes(router)

    router.Run(":8080")
}
