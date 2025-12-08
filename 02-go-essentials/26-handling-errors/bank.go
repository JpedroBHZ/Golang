package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

// Diferente de outras linguagens que usam try catch, o go usa retornos de erros explicitos nas funções
func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile) //retorna null se n tiver erro

	//nil é valor nenhum (null)
	if err != nil { //caso o erro não for nulo, o if retorna antes da função com o erro e com o valor padrão
		return 1000, errors.New("falha ao encontrar o arquivo") //Não deve ser maiusculo o texto, convenção
	}

	banlanceText := string(data)
	balance, err := strconv.ParseFloat(banlanceText, 64)

	if err != nil {
		return 1000, errors.New("falha ao gravar valor")
	}

	return balance, nil
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main() {
	var accountBalance, err = getBalanceFromFile()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("----------")
		//panic("A aplicação não pode continuar") usado para finalizar abruptamente a aplicação
	}

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
