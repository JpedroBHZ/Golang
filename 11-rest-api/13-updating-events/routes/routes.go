package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {

	//Configurando manipulador de solicitação http get (GET, POST, PUT, PATCH, DELETE)
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent) //rota do evento por id
	server.POST("/events", createEvents)
	server.PUT("/events/:id", updateEvent)

}
