package main

import (
	"os"
	controllers "task-manager/controllers/taskControllers"
	userControllers "task-manager/controllers/userControllers"
	"task-manager/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.ConnectToDb()
}

func main() {
	r := gin.Default()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	r.POST("/signup", userControllers.Signup)
	r.POST("/login", userControllers.Login)
	r.POST("/logout", userControllers.Logout)

	r.POST("/addtask", controllers.AddTask)
	r.POST("/deletetask", controllers.DeleteTask)
	r.POST("/edittask", controllers.EditTask)

	r.GET("/task/:id", controllers.GetTask)
	r.GET("/tasks", controllers.ListTasks)
	r.GET("/user/:id/tasks", controllers.GetUserTasks)

	r.Run(":" + port)
}
