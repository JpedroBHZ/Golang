package main

import (
	"fmt"
)

func main() {

	receita := digita("Receita: ")
	despesas := digita("Despesas: ")
	taxa := digita("Taxa: ")

	ebt, profit, ratio := calcula(receita, despesas, taxa)
	fmt.Printf("%.1f\n%.1f\n%.3f\n", ebt, profit, ratio)

}

func digita(texto string) (valor float64) {
	fmt.Print(texto)
	fmt.Scan(&valor)
	return
}

func calcula(receita, despesas, taxa float64) (ebt, profit, ratio float64) {
	ebt = receita - despesas
	profit = ebt * (1 - taxa/100)
	ratio = ebt / profit
	return
}
