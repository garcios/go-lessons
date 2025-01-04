package prices

import (
	"fmt"
	"time"

	conv "github.com/oskiegarcia/price-calculator/conversion"
	io "github.com/oskiegarcia/price-calculator/iomanager"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64           `json:"tax_rate"`
	InputPrices       []float64         `json:"input_prices"`
	TaxIncludedPrices map[string]string `json:"tax_included_prices"`
	IOManager         io.IOManager      `json:"-"`
}

func NewTaxIncludedPriceJob(iom io.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate:   taxRate,
		IOManager: iom,
	}
}

func (job *TaxIncludedPriceJob) Process() error {
	fmt.Println("Processing TaxIncludedPriceJob for", job.TaxRate)

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

	// simulate a slow process
	time.Sleep(3 * time.Second)

	// simulate an error
	if job.TaxRate == 0.07 {
		fmt.Println("returning error here...")
		return fmt.Errorf("unable to process TaxIncludedPriceJob for tax rate %.2f", job.TaxRate)
	}
	return job.IOManager.WriteResult(job)
}

func (job *TaxIncludedPriceJob) LoadData() error {
	lines, err := job.IOManager.ReadLines()
	if err != nil {
		return err
	}

	prices, err := conv.StringsToFloat(lines)
	if err != nil {
		return err
	}

	job.InputPrices = prices

	return nil
}
