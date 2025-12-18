package main

import (
	"net/http" //Biblioteca http nativa do go

	"github.com/gin-gonic/gin" //Importando biblioteca do gin
)

func main() {
	server := gin.Default() //Configura servidor http

	server.GET("/events", getEvents) //Configurando manipulador de solicitação http get (GET, POST, PUT, PATCH, DELETE)

	server.Run(":8080") //Inicia servidor local na porta 8080
}

func getEvents(context *gin.Context) { //manipulador da solicitação getEvents
	// .H é um atalho para criar um map, abuse de parar mouse em cima das funções para info
	context.JSON(http.StatusOK, gin.H{"message": "Hello!"}) //retornando dados json
}
