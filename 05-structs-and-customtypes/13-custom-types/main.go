package main

import "fmt"

type str string //Criando nosso tipo customizado, variavel que pode ter metodos proprios

func (text str) log() { //Metodo da variavel str
	fmt.Println(text)
}

func main() {
	var name str = "Max"

	name.log()
}
