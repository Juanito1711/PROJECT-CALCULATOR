package main

import (
	"fmt"
)

func main() {

	prices := []float64{10, 20, 30}
	taxRates := []float64{0, 0.7, 0.1, 0.5}

	result := make(map[float64][]float64)

	for _, taxRate := range taxRates {
		priceJob := prices.NewTaxIncludedPriceJob(taxRate)
		priceJob.Process()
	}

	fmt.Println(result)

}
