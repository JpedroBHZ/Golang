package main

import (
	"fmt"
	"time"
)

func greet(phrase string) {
	fmt.Println("Hello!", phrase)
}

func slowGreet(phrase string, doneChan chan bool) { //canal do tipo bollean como parametro
	time.Sleep(3 * time.Second)
	fmt.Println("Hello!", phrase)
	doneChan <- true //canal recebe true
}

func main() {
	//go greet("Nice to meet you!")
	//go greet("How are you?")
	done := make(chan bool) //criando canal
	go slowGreet("How ... are ... you ...?", done)
	//go greet("I hope you're liking the course!")
	<-done //quando termina operação, executa aqui, tirando elas do canal
}
