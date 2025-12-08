package main

import (
	"fmt"

	"example.com/bank/fileops" //Cuidado quando for salvar um pacote que não está usando, pois o go deleta
)

const accountBalanceFile = "balance.txt"

func main() {
	//Para usar as funções temos que colocar . e o nome do import
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("----------")
	}

	fmt.Println("Bem vindo ao banco GO")

	for {
		presentOptions()

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
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
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
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		default:
			fmt.Println("Saindo")
			fmt.Println("Obrigado por escolher nosso banco")
			return
		}
	}
}
