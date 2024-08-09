package prices

import (
	"fmt"

	"example.com/calculator/conversion"
	"example.com/calculator/iomanager"
)

type TaxIncludedPriceJob struct {
	IOManager         iomanager.IOManager `json:"-"`
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]float64  `json:"tax_included_prices"`
}

func (job *TaxIncludedPriceJob) Process() error {
	err := job.LoadData()
	if err != nil {
		return err
	}

	result := make(map[string]float64)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = taxIncludedPrice // Corrected to store float64, not string
	}

	job.TaxIncludedPrices = result
	return job.IOManager.WriteResult(job.TaxIncludedPrices)
}

func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager: iom,
		TaxRate:   taxRate,
	}
}

func (job *TaxIncludedPriceJob) LoadData() error {
	lines, err := job.IOManager.ReadLines()
	if err != nil {
		fmt.Println("Error reading lines:", err)
		return err
	}

	prices, err := conversion.StringsToFloat(lines)
	if err != nil {
		fmt.Println("Error converting strings to floats:", err)
		return err
	}

	job.InputPrices = prices
	return nil
}
