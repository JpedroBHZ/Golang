package main

import "fmt"

func main() {
	prices := []float64{10.99, 8.99} //criando fatia
	fmt.Println(prices[0:1])
	prices[1] = 9.99

	//Aumentando fatia, criando uma matriz, e adicionando a antiga nela, gerando uma nova fatia
	updatedPrices := append(prices, 5.99)
	//Para remover um elemento é apenas fatiando ele para fora do intervalo
	prices = prices[1:]
	fmt.Println(updatedPrices, prices)
}
