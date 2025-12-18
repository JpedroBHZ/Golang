package models

import "time"

//Estrutura dos eventos
type Event struct {
	ID          int
	Name        string    `binding:"required"` //campo não pode ser nulo, ou gera erro
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

var events = []Event{} //Fatia de eventos

//Metodo de salvamento
func (e Event) Save() {
	//futuramente adicionando no banco
	events = append(events, e) //adicionando novo evento na fatia
}

//Metodo que busca todos os eventos
func GetAllEvents() []Event {
	return events
}
