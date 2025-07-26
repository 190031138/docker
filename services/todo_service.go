package services

import (
    "context"
    "go-todo-app/config"
    "go-todo-app/models"
    "log"
    "time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

func GetTodos() []models.Todo {
    var todos []models.Todo
    collection := config.DB.Collection("todos")

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    cursor, err := collection.Find(ctx, bson.M{})
    if err != nil {
        log.Fatal(err)
    }
    defer cursor.Close(ctx)

    for cursor.Next(ctx) {
        var todo models.Todo
        cursor.Decode(&todo)
        todos = append(todos, todo)
    }
    return todos
}

func CreateTodo(todo models.Todo) (*models.Todo, error) {
    collection := config.DB.Collection("todos")

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    res, err := collection.InsertOne(ctx, todo)
    if err != nil {
        return nil, err
    }

    todo.ID = res.InsertedID.(primitive.ObjectID)
    return &todo, nil
}
