package controllers

import (
	"log"
	"net/http"
	"taskmanagement/data"
	"taskmanagement/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetTasks(c *gin.Context) {
	tasks, err := data.GetTasks()
	if err != nil {
		// Print the REAL error to server terminal
		log.Printf("ERROR in GetTasks: %v", err)

		// Also send it to client for easier debugging
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "could not fetch tasks",
			"details": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func GetTaskByID(c *gin.Context) {
	id := c.Param("id")
	task, err := data.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}

func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := data.CreateTask(task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create task"})
		return
	}

	c.JSON(http.StatusCreated, result)
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var update bson.M
	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := data.UpdateTask(id, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not update task"})
		return
	}
	c.JSON(http.StatusOK, result)
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	result, err := data.DeleteTask(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not delete task"})
		return
	}
	c.JSON(http.StatusOK, result)
}
