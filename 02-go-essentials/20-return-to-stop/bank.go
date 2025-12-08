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
	} else if choice == 2 {
		fmt.Print("Seu deposito: ")
		var depositAmount float64
		fmt.Scan(&depositAmount)

		if depositAmount <= 0 {
			fmt.Println("Valor inválido, não pode ser menor que 0")
			return //Para a execução da função
		}

		accountBalance += depositAmount //accountbalance = accountbalance + depositamount
		fmt.Println("Saldo atualizado! Novo saldo:", accountBalance)
	} else if choice == 3 {
		fmt.Print("Seu saque: ")
		var withdrawAmount float64
		fmt.Scan(&withdrawAmount)

		if withdrawAmount <= 0 {
			fmt.Println("Valor inválido, não pode ser menor que 0")
			return //Para a execução da função
		}

		if withdrawAmount > accountBalance {
			fmt.Println("Saldo insuficiente")
			return
		}

		accountBalance -= withdrawAmount
		fmt.Println("Saldo atualizado! Novo saldo:", accountBalance)
	} else {
		fmt.Println("Saindo")
	}

	fmt.Println("Você escolheu:", choice)
}
