package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// Você normalmente separa funções em pacotes diferentes, quando eles podem ser usadas em vários projetos
// Mas pra isso você tem que deixar elas mais genéricas

const accountBalanceFile = "balance.txt"

func getFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return 1000, errors.New("falha ao encontrar o arquivo")
	}

	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return 1000, errors.New("falha ao gravar valor")
	}

	return value, nil
}

func writeFloatToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}

func main() {
	var accountBalance, err = getFloatFromFile(accountBalanceFile)

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
			writeFloatToFile(accountBalance, accountBalanceFile)
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
			writeFloatToFile(accountBalance, accountBalanceFile)
		default:
			fmt.Println("Saindo")
			fmt.Println("Obrigado por escolher nosso banco")
			return
		}
	}
}
