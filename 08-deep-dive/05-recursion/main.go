package main

import "fmt"

func main() {
	fact := factorial(5)
	fmt.Println(fact)
}

//Recursão é quando uma função chama ela mesma
func factorial(number int) int {
	if number == 0 { //Isso é uma pausa para quando atingir zero, e não criar um loop infinito de numeros negativos
		return 1
	}
	return number * factorial(number-1) //Enquanto todas as funções não tiverem seus returns resolvidos, a função não para
}

//Essa função vai diminuindo até chegar em zero, e depois volta multiplicando os valores positivos da função maior, pelo resultado da sub função
