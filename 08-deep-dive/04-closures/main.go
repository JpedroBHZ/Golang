package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}

	double := createTransformer(2)
	triple := createTransformer(3)

	transformed := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})

	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(transformed)
	fmt.Println(doubled)
	fmt.Println(tripled)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

//Essa função cria outras, pois seu retorno é uma função sem nome, que multiplica um numero passado pela criadora
func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor //A função sem nome acessa os parametros da criadora
	}
}

//Esse exemplo não é sobre funções criadoras, e sim sobre fechamento
//O ponto é que o 2 e o 3 ficaram trancados dentro da função resultante, fzd que multipliquem sempre que chamarmos elas
