package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllTasks(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, taskList)
}

func AddTasks(c *gin.Context) {

	var newTask Tasks
	if err := c.BindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	newTask.Id = len(taskList) + 1
	taskList = append(taskList, newTask)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Tarefa postada"})

}

func GetTasksByID(c *gin.Context) {

	id := c.Param("id")

	for _, task := range taskList {
		if fmt.Sprintf("%d", task.Id) == id {
			c.IndentedJSON(http.StatusOK, task)
			return

		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{
		"Message": "ID not found",
	})

}

func DeleteTasks(c *gin.Context) {

	id := c.Param("id")

	for index, task := range taskList {
		if fmt.Sprintf("%d", task.Id) == id {
			taskList = append(taskList[:index], taskList[index+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Item deletado"})
			return
		}
	}

	c.IndentedJSON(http.StatusBadRequest, gin.H{
		"message": "Item nao encontrado",
	})
}

func EditTasks(c *gin.Context) {

	id := c.Param("id")

	var updatedTasks Tasks

	if err := c.BindJSON(&updatedTasks); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Message": "tarefa nao encontradaa"})
		return
	}

	for index, task := range taskList {
		if fmt.Sprintf("%d", task.Id) == id {
			updatedTasks.Id = task.Id
			taskList[index] = updatedTasks
			c.IndentedJSON(http.StatusOK, gin.H{"Message": "tarefa atualizada"})
			return
		}
	}
}
