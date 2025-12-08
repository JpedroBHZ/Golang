package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate float64 = 2.5
	var investmentAmount, years, expectedReturnRate float64

	fmt.Print("Valor investido: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Espectativa de retorno: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Anos: ")
	fmt.Scan(&years)

	var futureValue = investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	//Depois da % podemos escolher a formatação, a .1f deixa 1 num de casa decimal
	fmt.Printf(`Futuro valor: %.1f\n
	Valor ajustado pela inflação: %.1f`, futureValue, futureRealValue)
}
