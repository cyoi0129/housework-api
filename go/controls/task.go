package controllers

import (
	"net/http"
	"strconv"
	"time"
	"workout-note/models"
	"workout-note/services"

	"github.com/gin-gonic/gin"
)

func FetchAllTaskList(c *gin.Context) {
	user_param := c.Param("id")
	target_user, err := strconv.Atoi(user_param)
	if err != nil {
		target_user = 0
	}
	tasks, err := services.FetchAllTaskList(target_user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 2, "data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": 0, "data": tasks})
}

func FetchTaskList(c *gin.Context) {
	user_param := c.Param("id")
	target_user, err := strconv.Atoi(user_param)
	if err != nil {
		target_user = 0
	}
	query_start := c.Query("start")
	if query_start == "" {
		query_start = time.Now().Format("2006-01-02")
	}
	query_end := c.Query("end")
	if query_end == "" {
		query_end = time.Now().Format("2006-01-02")
	}
	tasks, err := services.FetchTaskList(target_user, query_start, query_end)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 2, "data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": 0, "data": tasks})
}

func CreateTask(c *gin.Context) {
	var input models.Task
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task, _ := services.CreateTask(input)
	c.JSON(http.StatusOK, gin.H{"status": 0, "data": task})
}

func UpdateTask(c *gin.Context) {
	task_param := c.Param("id")
	target_task, err := strconv.Atoi(task_param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 2, "data": err.Error()})
		return
	}
	var input models.Task
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 2, "data": err.Error()})
		return
	}
	task, _ := services.UpdateTask(target_task, input)
	c.JSON(http.StatusOK, gin.H{"status": 0, "data": task})
}

func DeleteTaskById(c *gin.Context) {
	task_param := c.Param("id")
	target_task, err := strconv.Atoi(task_param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 2, "data": err.Error()})
		return
	}
	result, _ := services.DeleteTask(target_task)
	c.JSON(http.StatusOK, gin.H{"status": 0, "data": result})
}
