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
func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events "
	rows, err := db.DB.Query(query) //usado query no lugar de exec, pois se trata de uma consulta
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event //Criando fatia para armazenar consulta

	for rows.Next() { //Percorrendo todas as linhas impressas e armazenando na fatia
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

// Função de pegar evento pelo id
func GetEventByID(id int64) (*Event, error) { //id sendo passado
	query := "SELECT * from events WHERE id = ?"
	row := db.DB.QueryRow(query, id) //o retorno vai ser a query em uma linha, (queryrow)

	//Trazendo infos do evento encontrado
	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (event Event) Update() error {
	query := `
	UPDATE events 
	SET name = ?, description = ?, location = ?, dateTime = ?
	WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)
	return err

}

func (event Event) Delete() error {
	query := "DELETE FROM events WHERE id = ?"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	_, err = stmt.Exec(event.ID)
	return err
}
