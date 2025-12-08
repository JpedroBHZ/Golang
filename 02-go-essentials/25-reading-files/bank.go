package main

import (
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() float64 {
	//A função ReadFile retorna dois valores, 1 de dados e 1 de erro, já que querermos só os dados, colocamos _
	data, _ := os.ReadFile(accountBalanceFile)
	//Convertemos esses dados em string, para facilitar a conversão para float
	banlanceText := string(data)
	//Usamos o pacote strconv, para usar a função que converte para float, depois passamos a versão 64 do float
	balance, _ := strconv.ParseFloat(banlanceText, 64)
	return balance
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main() {
	var accountBalance = getBalanceFromFile()

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
			writeBalanceToFile(accountBalance)
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
