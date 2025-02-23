package main

import (
    "todo-api/database"
    "todo-api/routes"

    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    // Connect to the database
    database.ConnectDB()

    // Make the DB available to controllers via middleware
    r.Use(func(c *gin.Context) {
        c.Set("db", database.DB)
        c.Next()
    })

    // Register routes
    routes.RegisterAuthRoutes(r)
    routes.RegisterTodoRoutes(r) // Assuming your todo routes are in routes/todo_routes.go

    r.Run(":8000")
}
