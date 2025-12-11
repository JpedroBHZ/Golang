package main

import "fmt"

func main() {
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[0:1])
	prices[1] = 9.99

	updatedPrices := append(prices, 5.99, 12.99, 29.99, 100.10)
	prices = prices[1:]
	fmt.Println(updatedPrices, prices)

	discountPrices := []float64{101.99, 80.99, 20.59}
	//Caso queira mesclar 2 listas, o "..." esvazia a lista reescreve os num separados no lugar dela
	prices = append(prices, discountPrices...)
	fmt.Println(prices)
}
