package models

import (
	"time"

	"example.com/rest-api/db"
)

// Estrutura dos eventos
type Event struct {
	ID          int64     //O id automático é desse tipo
	Name        string    `binding:"required"` //campo não pode ser nulo, ou gera erro
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

var events = []Event{} //Fatia de eventos

// Metodo de salvamento
func (e Event) Save() error {
	//criando query de inserção (? é cada campo que queremos evitar sqlinjection)
	query := `
	INSERT INTO events(name, description, location, dateTime, user_id)
	VALUES (?, ?, ?, ?, ?)`

	//preparando os campos para aceitarem valores
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close() //fecha stmt quando termina de executar a função
	//Preenchendo campos com valores
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)

	if err != nil {
		return err
	}

	//Pegando id automático
	id, err := result.LastInsertId()
	//Adicionando ao campo id
	e.ID = id
	return err
}

// Metodo que busca todos os eventos
func GetAllEvents() []Event {
	return events
}
