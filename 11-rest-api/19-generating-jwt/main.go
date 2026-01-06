package main

import (
	//Biblioteca http nativa do go

	"example.com/rest-api/db"
	"example.com/rest-api/routes"
	"github.com/gin-gonic/gin" //Importando biblioteca do gin
)

func main() {

	//Chama a função de inicialização do banco
	db.InitDB()

	//Configura servidor http
	server := gin.Default()

	//Função de rotas
	routes.RegisterRoutes(server)

	//Inicia servidor local na porta 8080
	server.Run(":8080")
}
