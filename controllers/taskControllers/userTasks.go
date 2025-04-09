package controllers

import (
	"net/http"
	"task-manager/initializers"
	"task-manager/models"

	"github.com/gin-gonic/gin"
)

func GetUserTasks(c *gin.Context) {
	// Get user ID from URL parameter
	userID := c.Param("id")

	// Get tasks for the specified user
	var tasks []models.Task
	result := initializers.DB.Where("user_id = ?", userID).Find(&tasks)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch user tasks",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tasks": tasks,
	})
}
