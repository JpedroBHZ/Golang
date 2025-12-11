package main

import "fmt"

func main() {
	userNames := make([]string, 2, 5) //Separa uma matriz com 2 espaços vazios de tamanho 5

	userNames = append(userNames, "Max")    //Append não tem que criar uma nova planilha e mesclar
	userNames = append(userNames, "Manuel") //Ele atribui diretamente, acelerando o processo

	fmt.Println(userNames)
}
