package main

import (
	"fmt"
	"math"
)

func main() {
	//definindo variaveis
	var investmentAmount float64 = 1000 //int -> float
	var expectedReturnRate = 5.5        //float
	var years float64 = 10

	//calculo de retorno do invest.
	//conversão necessária, é proibido usar diferentes tipos de dados em operações, exp: int * float
	var futureValue = investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	fmt.Println(futureValue)

}
