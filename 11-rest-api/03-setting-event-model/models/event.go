package models

import "time"

//Estrutura dos eventos
type Event struct {
	ID          int
	Name        string
	Description string
	Location    string
	DateTime    time.Time
	UserID      int
}

var events = []Event{} //Fatia de eventos

//Metodo de salvamento
func (e Event) Save() {
	//futuramente adicionando no banco
	events = append(events, e) //adicionando novo evento na fatia
}
