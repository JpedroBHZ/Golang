package main

import "fmt"

//Exemplo de função sem nome

func main() {
	numbers := []int{1, 2, 3}

	//Podemos criar uma função sem nome, onde poderiamos usar uma função, parametros, returns..
	//Porém os parametros dela tem que bater com os da outra função que vai processar ela
	transformed := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})

	fmt.Println(transformed)
}

//Função que vai processar nossa vazia
func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}
