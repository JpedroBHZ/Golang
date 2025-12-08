package main

import (
	"fmt"
	"os"
)

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance) //Criando uma string e passando o float pra ela = cadeia de caracteres
	//Pega os caracteres, transforma em bit e salve no arquivo balance.txt, com a permissão 0644
	os.WriteFile("balance.txt", []byte(balanceText), 0644)
}

func main() {
	var accountBalance = 1000.0

	fmt.Println("Bem vindo ao banco GO")

	for {
		fmt.Println("O que você gostaria de fazer?")
		fmt.Println("1. Checar saldo")
		fmt.Println("2. Depoistar dinheiro")
		fmt.Println("3. Sacar dinheiro")
		fmt.Println("4. Sair")

		var choice int
		fmt.Print("Sua escolha: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Seu saldo é:", accountBalance)
		case 2:
			fmt.Print("Seu deposito: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Valor inválido, não pode ser menor que 0")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Saldo atualizado! Novo saldo:", accountBalance)
			writeBalanceToFile(accountBalance) //Chama a função de gravação e passa o valor atualizado
		case 3:
			fmt.Print("Seu saque: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Valor inválido, não pode ser menor que 0")
				continue
			}
			if withdrawAmount > accountBalance {
				fmt.Println("Saldo insuficiente")
				continue
			}

			accountBalance -= withdrawAmount
			fmt.Println("Saldo atualizado! Novo saldo:", accountBalance)
			writeBalanceToFile(accountBalance)
		default:
			fmt.Println("Saindo")
			fmt.Println("Obrigado por escolher nosso banco")
			return
		}
	}
}
