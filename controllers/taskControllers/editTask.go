package controllers

import (
	"net/http"
	"task-manager/initializers"
	"task-manager/models"

	"github.com/gin-gonic/gin"
)

func EditTask(c *gin.Context) {
	// Get task data from request
	var body struct {
		ID          uint   `json:"id" binding:"required"`
		Title       string `json:"title"`
		Description string `json:"description"`
		DueDate     string `json:"due_date"`
		Status      string `json:"status"`
		Priority    string `json:"priority"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body: " + err.Error(),
		})
		return
	}

	// Find the task by ID
	var task models.Task
	result := initializers.DB.First(&task, body.ID)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Task not found",
		})
		return
	}

	// Update task fields if provided
	if body.Title != "" {
		task.Title = body.Title
	}
	if body.Description != "" {
		task.Description = body.Description
	}
	if body.DueDate != "" {
		task.DueDate = body.DueDate
	}
	if body.Status != "" {
		task.Status = body.Status
	}
	if body.Priority != "" {
		task.Priority = body.Priority
	}

	// Save the updated task
	result = initializers.DB.Save(&task)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to update task",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Task updated successfully",
		"task":    task,
	})
}
