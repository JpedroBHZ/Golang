package main

import "fmt"

//Melhor para tipos claramente definido, vão permanecer os mesmos
//Usado para definir objetos
type Product struct {
	id    string
	title string
	price float64
}

func main() {

	//Rotulos personalizados, modificaveis
	//Usado para agrupar

	websites := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}

	fmt.Println(websites)
	fmt.Println(websites["Google"])

	websites["LinkedIn"] = "https://linkedin.com"
	fmt.Println(websites)

	delete(websites, "Google")
	fmt.Println(websites)
}
