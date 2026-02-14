package router

import (
	"log"

	"taskmanagement/config"
	"taskmanagement/controllers"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	config.ConnectDB()

	log.Println("App is ready!")

	r := gin.Default()

	r.GET("/tasks", controllers.GetTasks)
	r.GET("/tasks/:id", controllers.GetTaskByID)
	r.POST("/tasks", controllers.CreateTask)
	r.PUT("/tasks/:id", controllers.UpdateTask)
	r.DELETE("/tasks/:id", controllers.DeleteTask)

	return r
}
