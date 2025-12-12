package main

import "fmt"

//Exemplo que mostra que podemos passar funções como parametro
type transformFn func(int) int //tipo função, com parametro e retorno do tipo int

func main() {
	nunbers := []int{1, 2, 3, 4}                  //Fatia
	doubled := transformNumbers(&nunbers, double) //ponteiro e função como argumento
	tripled := transformNumbers(&nunbers, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)
}

func transformNumbers(numbers *[]int, transform transformFn) []int {
	dNumbers := []int{}            //fatia de saida
	for _, val := range *numbers { //percorrendo a fatia de entrada
		//passando cada num da fatia de entrada na função recebida como argumento
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

//Criando funções para serem passadas como argumento
func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
