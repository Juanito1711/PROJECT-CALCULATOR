package main

import (
	"fmt"

	"example.com/calculator/cmdmanager"
	"example.com/calculator/prices"
)

func main() {
	taxRates := []float64{0.07, 0.1, 0.5} // Correct tax rates

	for _, taxRate := range taxRates {
		cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(cmdm, taxRate)
		priceJob.Process()
		err := priceJob.Process()

		if err != nil {
			fmt.Println("Could not process job")
			fmt.Println(err)
		}
	}
}
