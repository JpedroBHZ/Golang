package main

import "fmt"

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

		//Switch é uma alternativa mais limpa para if else, para quando temos 1 variavel de controle "choice"
		//Com switch não podemos usar break para parar o loop e continuar o cod
		switch choice {
		case 1: //Nossos "ifs"
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
		default: //Nosso "else"
			fmt.Println("Saindo")
			fmt.Println("Obrigado por escolher nosso banco")
			return //Obrigado a usar return para parar a função
		}
	}
}
