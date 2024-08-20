package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	//! Iniciando o gin framework
	r := gin.Default()

	DB = initDB("tasks.db")
	defer DB.Close()

	r.SetTrustedProxies(nil)

	RegisterRouts(r)

	r.Run(":3000")

}
