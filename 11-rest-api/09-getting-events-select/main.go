package main

import (
	"net/http" //Biblioteca http nativa do go

	"example.com/rest-api/db"
	"example.com/rest-api/models"
	"github.com/gin-gonic/gin" //Importando biblioteca do gin
)

func main() {

	//Chama a função de inicialização do banco
	db.InitDB()

	//Configura servidor http
	server := gin.Default()

	//Configurando manipulador de solicitação http get (GET, POST, PUT, PATCH, DELETE)
	server.GET("/events", getEvents)
	server.POST("/events", createEvents)

	//Inicia servidor local na porta 8080
	server.Run(":8080")
}

// manipulador da solicitação getEvents
func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents() //Chamando a função e tratando seu erro
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events. Try again later."})
		return
	}
	context.JSON(http.StatusOK, events) //retornando eventos como json
}

// manipulador da solicitação createEvents
func createEvents(context *gin.Context) {
	var event models.Event                //criando evento
	err := context.ShouldBindJSON(&event) //preenchendo evento com solicitação de entrada

	//Em caso de erro ao gravar evento
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	//Em caso de sucesso ao gravar evento
	event.ID = 1
	event.UserID = 1

	//Metodo de salvamento do evento
	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create events. Try again later."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}
