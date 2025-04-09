package controllers

import (
	"net/http"
	"task-manager/initializers"
	"task-manager/models"

	"github.com/gin-gonic/gin"
)

func GetTask(c *gin.Context) {
	// Get task ID from URL parameter
	id := c.Param("id")

	// Find the task by ID
	var task models.Task
	result := initializers.DB.First(&task, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Task not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"task": task,
	})
}
