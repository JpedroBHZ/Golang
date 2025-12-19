package db

//go get github.com/mattn/go-sqlite3
import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3" //esse import nunca será usado mas precisamos dele, o "_" é para não sumir ao salvar
)

// Instancia variavel global para o banco
var DB *sql.DB

func InitDB() {
	DB, err := sql.Open("sqlite3", "api.db") //Cria a conexão com o banco

	if err != nil {
		panic("Could not connect to database")
	}

	//Ali em cima, não criamos o banco exatamente, apenas abrimos ele para conexões, aqui configuramos a quantidade delas
	DB.SetMaxOpenConns(10) //Maximo de requisições
	DB.SetMaxIdleConns(5)  //Sempre ficaram 5 espaços, esperando conexões
}
