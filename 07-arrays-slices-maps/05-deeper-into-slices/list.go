package main

import "fmt"

func main() {
	var productNames [4]string = [4]string{"A book"}
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
	fmt.Println(prices)

	productNames[2] = "A Carpet"

	fmt.Println(productNames)

	fmt.Println(prices[2])

	featurePrices := prices[1:] //cortando até o primeiro elemento a esquerda, perdeu
	featurePrices[0] = 199.99
	highlightedPrices := featurePrices[:1] //cortando até o 1 elemento a direita, ocultando o resto
	fmt.Println(highlightedPrices)
	fmt.Println(prices)
	fmt.Println(len(highlightedPrices), cap(highlightedPrices)) //qtd elementos x qtd total

	highlightedPrices = highlightedPrices[:3] //quando cortamos a direita, fica salvo o resto na memoria do go
	fmt.Println(highlightedPrices)            //quando cortarmos a esquerda ele se perde
	fmt.Println(len(highlightedPrices), cap(highlightedPrices))
}
