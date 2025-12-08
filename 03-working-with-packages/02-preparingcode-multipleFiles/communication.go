package main

import "fmt" //funções são compartilhadas imports não

func presentOptions() {
	fmt.Println("O que você gostaria de fazer?")
	fmt.Println("1. Checar saldo")
	fmt.Println("2. Depoistar dinheiro")
	fmt.Println("3. Sacar dinheiro")
	fmt.Println("4. Sair")
}

//Obs: quando separamos os pacotes, iniciamos com  "go run ." para iniciar todos os arquivos
