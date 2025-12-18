package main

import (
	"net/http" //Biblioteca http nativa do go

	"github.com/gin-gonic/gin" //Importando biblioteca do gin
)

func main() {
	//Configura servidor http
	server := gin.Default()

	//Configurando manipulador de solicitação http get (GET, POST, PUT, PATCH, DELETE)
	server.GET("/events", getEvents)

	//Inicia servidor local na porta 8080
	server.Run(":8080")
}

// manipulador da solicitação getEvents
func getEvents(context *gin.Context) {
	// .H é um atalho para criar um map
	// abuse de parar mouse em cima das funções para info
	context.JSON(http.StatusOK, gin.H{"message": "Hello!"}) //retornando dados como json
}
