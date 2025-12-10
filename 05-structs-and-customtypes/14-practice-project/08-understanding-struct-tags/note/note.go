package note

//Imports
import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

// Cria estrutura da nota
type Note struct {
	Title    string    `json:"title"` //muda o nome dos atributos no arquivo json
	Content  string    `json:"content"`
	CreateAt time.Time `json:"created_at"`
}

// Metodo que escreve as infos da nota
func (note Note) Display() {
	fmt.Printf("Your note tilted %v has the following content: \n\n%v\n\n", note.Title, note.Content)
}

// Função de salvar arquivo
func (note Note) Save() error {
	//Pega o atributo titulo, substitui os espaços e salva na variavel filename
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	//Pega o filename deixa tudo minusculo e adiciona .json no final
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(note) //transforma a nota em um json

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644) //cria um arquivo passando o json, e o nome formatado
}

// Contrutor de nota
func New(title, content string) (Note, error) { //Pega as infos, retorna ou a nota pronta ou um erro
	if title == "" || content == "" { //Testa se vazio, retorna o erro e a nota vazia
		return Note{}, errors.New("invalid input")
	}

	return Note{ //Controi a nota passando as variaveis para os atributos corretos
		Title:    title,
		Content:  content,
		CreateAt: time.Now(),
	}, nil
}
