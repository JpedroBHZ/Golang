package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
)

// App para criação de notas
func main() {
	title, content := getNoteData() //Pegando as entradas

	userNote, err := note.New(title, content) //Passando no construtor

	if err != nil { //Teste se vazio
		fmt.Println(err)
		return
	}

	userNote.Display()    //Metodo que valores
	err = userNote.Save() //Metodo de salvamento, retorna erro caso falhe

	if err != nil { //Aviso em caso de falhas
		fmt.Println("saving the note failed")
		return
	}

	fmt.Println("saving the note succeeded")
}

// Pega os valores da nota e retorna ambos
func getNoteData() (string, string) {

	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content

}

// Escreve na tela e pega entradas
func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)
	reader := bufio.NewReader(os.Stdin)  //Cria leitor de linha completa
	text, err := reader.ReadString('\n') //Corta texto até o caractere informado

	if err != nil { //Testa se texto é vazio
		return ""
	}

	text = strings.TrimSuffix(text, "\n") //Remove texto informado
	text = strings.TrimSuffix(text, "\r")

	return text
}
