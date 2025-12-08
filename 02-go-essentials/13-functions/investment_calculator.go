package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate float64 = 2.5
	var investmentAmount, years, expectedReturnRate float64

	saidaTexto("Valor investido: ")
	fmt.Scan(&investmentAmount)

	saidaTexto("Espectativa de retorno: ")
	fmt.Scan(&expectedReturnRate)

	saidaTexto("Anos: ")
	fmt.Scan(&years)

	var futureValue = investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	vlFormated := fmt.Sprintf("Futuro valor: %.1f\n", futureValue)
	fvlFormated := fmt.Sprintf("Valor ajustado pela inflação: %.1f\n", futureRealValue)

	fmt.Print(vlFormated, fvlFormated)
}

// Funções são ações pre definidas, podem ser vazias ou podem transformar dados quando passados nelas
func saidaTexto(text string) {
	fmt.Print(text)
}
