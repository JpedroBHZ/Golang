package main

import "fmt"

func main() {
	var accountBalance = 1000.0

	fmt.Println("Bem vindo ao banco GO")
	fmt.Println("O que você gostaria de fazer?")
	fmt.Println("1. Checar saldo")
	fmt.Println("2. Depoistar dinheiro")
	fmt.Println("3. Sacar dinheiro")
	fmt.Println("4. Sair")

	var choice int
	fmt.Print("Sua escolha: ")
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Println("Seu saldo é:", accountBalance)
	}

	fmt.Println("Você escolheu:", choice)
}
