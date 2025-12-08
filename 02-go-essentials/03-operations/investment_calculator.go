package main

import (
	"fmt"
	"math"
)

func main() {
	//definindo variaveis
	var investmentAmount = 1000  //int
	var expectedReturnRate = 5.5 //float
	var years = 10

	//calculo de retorno do invest.
	//conversão necessária, é proibido usar diferentes tipos de dados em operações, exp: int * float
	var futureValue = float64(investmentAmount) * math.Pow((1+expectedReturnRate/100), float64(years))
	fmt.Println(futureValue)

}
