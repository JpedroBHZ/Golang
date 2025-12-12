package main

import "fmt"

func main() {

	//numbers := []int{1, 10, 15}
	sum := sumup(1, 10, 15, 40, -5)

	fmt.Println(sum)
}

//As reticencias pegam os valores e juntam numa fatia pra vc, aceitando um ou mais parametros, avulsos ou em listas
//Também podemos separar o primeiro, segundo.. valor para usar fora da fatia
func sumup(startingValue int, numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum += val
	}

	return sum
}
