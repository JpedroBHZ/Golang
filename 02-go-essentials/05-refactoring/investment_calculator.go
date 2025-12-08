package main

import (
	"fmt"
	"math"
)

func main() {
	//Não é necessário var se não for redefinir o tipo da variável, é só colocar :=
	//Podemos definir variaveis de tipos diferentes na mesma linha: var nome, idade = "joão", 23
	//Podemos deixar as mesmas variaveis na mesma linha var idade, numero float64 = 20.0, 30.0
	investmentAmount, years, expectedReturnRate := 1000.0, 10.0, 5.5

	var futureValue = investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	fmt.Println(futureValue)

}
