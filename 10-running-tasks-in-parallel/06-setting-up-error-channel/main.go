package main

import (
	"fmt"

	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates)) //fatia de erros

	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool) //criando canal para cada loop
		errorChans[index] = make(chan error)
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		//cmdm := cmdmanager.New()
		preiceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go preiceJob.Process(doneChans[index], errorChans[index])

		//if err != nil {
		//	fmt.Println("Could not process job")
		//	fmt.Println(err)
		//}
	}

	for _, errorChans := range errorChans {
		<-errorChans
	}

	for _, doneChan := range doneChans {
		<-doneChan
	}
}
