package main

import "fmt"

//Tipo customizado, map com id string que aceita um float
type floatMap map[string]float64

//Metodo customizado do floatmap
func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	userNames := make([]string, 2, 5)

	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manuel")

	fmt.Println(userNames)

	//Diminui o codigo, criando um tipo
	courseRatings := make(floatMap, 3)

	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 4.7

	courseRatings.output()
	//fmt.Println(courseRatings)
}
