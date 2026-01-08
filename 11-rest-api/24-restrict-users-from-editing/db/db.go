package db

//go get github.com/mattn/go-sqlite3
import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3" //esse import nunca será usado mas precisamos dele, o "_" é para não sumir ao salvar
)

// Instancia variavel global para o banco
var DB *sql.DB

func InitDB() {

	var err error

	DB, err = sql.Open("sqlite3", "api.db") //Cria a conexão com o banco

	if err != nil {
		panic("Could not connect to database")
	}

	//Ali em cima, não criamos o banco exatamente, apenas abrimos ele para conexões, aqui configuramos a quantidade delas
	DB.SetMaxOpenConns(10) //Maximo de requisições
	DB.SetMaxIdleConns(5)  //Sempre ficaram 5 espaços, esperando conexões

	createTables() //Função de criar tabelas
}

// Criando tabelas
// Toda vez que iniciarmos é executado esse codigo que testa se a tabela existe, se não ele cria ela
func createTables() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT, 
		email TEXT NOT NULL UNIQUE, 
		password TEXT NOT NULL
	)
	`
	_, err := DB.Exec(createUsersTable) //cria tabela de usuário

	if err != nil {
		panic("Could not create events table")
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events ( 
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL, 
		description TEXT NOT NULL, 
		location TEXT NOT NULL, 
		dateTime DATETIME NOT NULL, 
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`
	_, err = DB.Exec(createEventsTable) //Cria os eventos

	if err != nil {
		panic("Could not create events table" + err.Error())
	}

}
