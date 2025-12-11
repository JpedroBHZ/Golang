package main

import "fmt"

func main() {
	var productNames [4]string = [4]string{"A book"}
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
	fmt.Println(prices)

	productNames[2] = "A Carpet"

	fmt.Println(productNames)

	fmt.Println(prices[2])

	//featurePrices := prices[:3] //Do começo até o valor determindado sem incluir
	featurePrices := prices[1:]            //Do valor determinado até o final incluindo
	highlightedPrices := featurePrices[:1] //Do começo até o valor determindado sem incluir
	fmt.Println(highlightedPrices)
}
