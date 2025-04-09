package controllers

import (
	"net/http"
	"task-manager/initializers"
	"task-manager/models"

	"github.com/gin-gonic/gin"
)

func ListTasks(c *gin.Context) {
	// Get query parameters for filtering
	status := c.Query("status")
	priority := c.Query("priority")

	// Initialize query
	query := initializers.DB

	// Apply filters if provided
	if status != "" {
		query = query.Where("status = ?", status)
	}
	if priority != "" {
		query = query.Where("priority = ?", priority)
	}

	// Get all tasks with applied filters
	var tasks []models.Task
	result := query.Find(&tasks)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch tasks",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tasks": tasks,
	})
}
