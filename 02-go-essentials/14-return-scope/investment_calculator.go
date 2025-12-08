package main

import (
	"fmt"
	"math"
)

// Variaveis colocadas fora da função tem escopo(acesso) por qualquer função no cod(global)
// Variaveis definidas dentro de uma função só podem ser acessadas por ela
const inflationRate float64 = 2.5

func main() {
	var investmentAmount, years, expectedReturnRate float64

	outputText("Valor investido: ")
	fmt.Scan(&investmentAmount)

	outputText("Espectativa de retorno: ")
	fmt.Scan(&expectedReturnRate)

	outputText("Anos: ")
	fmt.Scan(&years)

	vlFormatado, fvlFormatado := investmentReturn(investmentAmount, expectedReturnRate, years)

	fmt.Print(vlFormatado, fvlFormatado)
}

func outputText(text string) {
	fmt.Print(text)
}

//Os valores de retorno também precisam de variaveis, se não acusam erro
//func calcularValores(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
//	return fv, frv
//}

// Variaveis definidas e usadas no retorno
func investmentReturn(investmentAmount, expectedReturnRate, years float64) (fv, frv float64) {
	fv = investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	frv = fv / math.Pow(1+inflationRate/100, years)
	return fv, frv
	// Caso não especifique, ele retornara todas as variaveis exp: return
}
