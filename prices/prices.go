package prices

import (
	"fmt"

	"example.com/tax-calculator/conversion"
	"example.com/tax-calculator/filemanager"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64                 `json:"tax_rate"`
	InputPrices       []float64               `json:"input_prices"`
	TaxIncludedPrices map[string]string       `json:"tax_included_price"`
	IOManager         filemanager.FileManager `json:"-"`
}

func NewTaxIncludedPriceJob(fm filemanager.FileManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate:     taxRate,
		InputPrices: []float64{10, 20, 30},
		IOManager:   fm,
	}
}

func (job *TaxIncludedPriceJob) LoadData() error {
	lines, err := job.IOManager.ReadLines()
	if err != nil {
		// fmt.Println("Error", err)
		return err
	}
	prices, err := conversion.StringToFloat(lines)
	if err != nil {
		// fmt.Println("Error", err)
		return err
	}
	job.InputPrices = prices
	return nil
}

func (job *TaxIncludedPriceJob) Process() error {
	err := job.LoadData()
	if err != nil {
		return err
	}
	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}
	job.TaxIncludedPrices = result
	// fmt.Println(result)
	return job.IOManager.WriteJSON(job)
}
