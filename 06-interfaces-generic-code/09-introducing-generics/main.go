package main

import "fmt"

func main() {
	result := add(1, 2) //O go vai escolher qual dos tipos é o melhor para a operação
	fmt.Println(result)
}

func add[T int | float64 | string](a, b T) T { //o tipo pode ser um desses 3
	return a + b
}
