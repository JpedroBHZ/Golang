package main

import "fmt"

func main() {
	userNames := make([]string, 2, 5)

	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manuel")

	fmt.Println(userNames)

	courseRatings := make(map[string]float64, 3) //Definindo 3 espaços

	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 4.7 //Até aqui já tinha sido alocado
	courseRatings["c#"] = 4.0      //Go tem que realocar

	fmt.Println(courseRatings)
}
