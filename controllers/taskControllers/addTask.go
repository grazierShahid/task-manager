package controllers

import (
	"net/http"
	"task-manager/initializers"
	"task-manager/models"

	"github.com/gin-gonic/gin"
)

func AddTask(c *gin.Context) {
	var body struct {
		Title       string `json:"title" binding:"required"`
		Description string `json:"description"`
		DueDate     string `json:"due_date"`
		Priority    string `json:"priority"`
		UserID      uint   `json:"user_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body: " + err.Error(),
		})
		return
	}

	task := models.Task{
		Title:       body.Title,
		Description: body.Description,
		DueDate:     body.DueDate,
		Status:      "pending", // default status
		Priority:    body.Priority,
		UserID:      body.UserID,
	}

	result := initializers.DB.Create(&task)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create task",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"task": task,
	})
}
