package controllers

import (
	"net/http"
	"todo-api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetTodos retrieves all todos using the injected DB from context.
func GetTodos(c *gin.Context) {
    // Retrieve the DB instance from the Gin context.
    db := c.MustGet("db").(*gorm.DB)

    var todos []models.Todo
    if err := db.Find(&todos).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch todos"})
        return
    }
    c.JSON(http.StatusOK, todos)
}

// CreateTodo creates a new todo using the injected DB from context.
func CreateTodo(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)

    var todo models.Todo
    if err := c.ShouldBindJSON(&todo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := db.Create(&todo).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create todo"})
        return
    }
    c.JSON(http.StatusCreated, todo)
}

// UpdateTodo updates an existing todo using the injected DB from context.
func UpdateTodo(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)

    id := c.Param("id")
    var todo models.Todo
    if err := db.First(&todo, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
        return
    }
    if err := c.ShouldBindJSON(&todo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := db.Save(&todo).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update todo"})
        return
    }
    c.JSON(http.StatusOK, todo)
}

// DeleteTodo deletes a todo using the injected DB from context.
func DeleteTodo(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)

    id := c.Param("id")
    var todo models.Todo
    if err := db.First(&todo, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
        return
    }
    if err := db.Delete(&todo).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete todo"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}
