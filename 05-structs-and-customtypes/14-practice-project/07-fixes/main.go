package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
)

func main() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()
	err = userNote.Save()

	if err != nil {
		fmt.Println("saving the note failed")
		return
	}

	fmt.Println("saving the note succeeded")
}

func getNoteData() (string, string) {

	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content

}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)
	reader := bufio.NewReader(os.Stdin)  //Criando leitor que lê a linha completa
	text, err := reader.ReadString('\n') //Passando onde deve parar de ler

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n") //Tirando os caracteres que faltaram
	text = strings.TrimSuffix(text, "\r")

	return text
}
