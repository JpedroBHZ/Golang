package main //Pacote principal, mostra por onde o codigo começa a ser executado
//Definindo um nome igual para outro arquivo permite troca de info

import "fmt" //Pacote nativo do go, várias funções

func main() { //função que vai rodar inicialmente, só pode ter 1
	fmt.Print("Hello World") //função de dentro do pacote do go
}
