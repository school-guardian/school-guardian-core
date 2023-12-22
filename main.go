package main

import (
	"fmt"
	"gin/repository"
	"gin/server/routes"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func main() {

	database, err := repository.NewDatabase("monorail.proxy.rlwy.net", "52035", "postgres", "*f-bBE6FeDCCcdAd-faCfeFd2geaD43E", "railway")
	if err != nil {
		fmt.Println("Erro ao conectar ao banco de dados:", err)
		return
	}
	defer database.Close()

	routes.HandleRequests()

}
