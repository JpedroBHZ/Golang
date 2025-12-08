package main

import (
	"fmt"
)

func main() {
	var receita, despesas, taxa float64

	fmt.Print("Receita: ")
	fmt.Scan(&receita)

	fmt.Print("Despesas: ")
	fmt.Scan(&despesas)

	fmt.Print("Taxa de imposto: ")
	fmt.Scan(&taxa)

	ebt := receita - despesas
	profit := ebt * (1 - taxa/100)
	ratio := ebt / profit

	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)
}
