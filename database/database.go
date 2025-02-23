package database

import (
    "log"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "todo-api/models"
)

var DB *gorm.DB

func ConnectDB() {
    var err error
    DB, err = gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to the database.")
    }

    // Migrate both models
    DB.AutoMigrate(&models.Todo{}, &models.User{})
}
