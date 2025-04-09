package controllers

import (
	"net/http"
	"task-manager/initializers"
	"task-manager/models"

	"github.com/gin-gonic/gin"
)

func DeleteTask(c *gin.Context) {
	// Get task ID from request
	var body struct {
		ID uint `json:"id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body: " + err.Error(),
		})
		return
	}

	// Find and delete the task
	result := initializers.DB.Delete(&models.Task{}, body.ID)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to delete task",
		})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Task not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Task deleted successfully",
	})
}
