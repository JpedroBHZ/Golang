package main

//go routines é um recurso para executar tarefas em paralelo, sem que uma dependa da outra para começar
import (
	"fmt"
	"time"
)

func greet(phrase string) {
	fmt.Println("Hello!", phrase)
}

func slowGreet(phrase string) {
	time.Sleep(3 * time.Second) // simulando uma tarefa lenta
	fmt.Println("Hello!", phrase)
}

func main() {
	go greet("Nice to meet you!")
	go greet("How are you?")
	go slowGreet("How ... are ... you ...?")
	go greet("I hope you're liking the course!")
}
