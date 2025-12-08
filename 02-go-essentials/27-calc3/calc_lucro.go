package main

import (
	"errors"
	"fmt"
	"os"
)

const accountResultFile = "result.txt"

func writeResultToFile(ebt, profit, ratio float64) {
	resultText := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n", ebt, profit, ratio)
	os.WriteFile(accountResultFile, []byte(resultText), 0644)
}

func main() {

	revenue, err1 := getUserInput("Receita: ")
	expenses, err2 := getUserInput("Despesas: ")
	taxRate, err3 := getUserInput("Taxa: ")

	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println(err1)
		panic(err1)
	}

	ebt, profit, ratio := calculate(revenue, expenses, taxRate)
	writeResultToFile(ebt, profit, ratio)

	fmt.Printf("%.1f\n%.1f\n%.3f\n", ebt, profit, ratio)

}

func getUserInput(text string) (float64, error) {
	var userInput float64
	fmt.Print(text)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("o valor tem que ser positivo") //retorno de 0 tem que retornar o float
	}

	return userInput, nil
}

func calculate(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return
}
