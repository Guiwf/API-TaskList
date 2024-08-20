package main

import "github.com/gin-gonic/gin"

//* modo de formatacao das tarefas
type Tasks struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

//* Array com a lista de tarefas
var taskList = []Tasks{

	{Id: 1, Title: "limpando a casa"},
}

func RegisterRouts(r *gin.Engine) {

	//? Rota para listar as tarefas da "tasklist"
	r.GET("/tasks", GetAllTasks)

	//? Rota para buscar as tarefas da "tasklist" por "ID"
	r.GET("/tasks/:id", GetTasksByID)

	//? Rota para Adicionar as tarefas da "tasklist"
	r.POST("/tasks", AddTasks)

	//? Rota para Atualizar as tarefas da "tasklist" por "ID"
	r.PUT("/tasks/:id", EditTasks)

	//? Rota para Deletar as tarefas da "tasklist" por "ID"
	r.DELETE("/tasks/:id", DeleteTasks)

}
