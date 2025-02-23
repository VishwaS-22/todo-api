package main

import (
    "log"
    "os"
    "todo-api/database"
    "todo-api/routes"

    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
)

func main() {
    // Load environment variables from .env file
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found")
    }

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
    routes.RegisterTodoRoutes(r)

    // Run on port specified in environment variable or default to 8000
    port := os.Getenv("PORT")
    if port == "" {
        port = "8000"
    }
    r.Run(":" + port)
}
