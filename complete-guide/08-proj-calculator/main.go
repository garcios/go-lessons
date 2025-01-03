package main

import (
	"fmt"
	filem "github.com/oskiegarcia/price-calculator/filemanager"
	p "github.com/oskiegarcia/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.10, 0.15}

	for _, taxRate := range taxRates {
		fm := filem.New("prices.txt", fmt.Sprintf("rseult_%.0f.json", taxRate*100))
		//cmd := cmdmanager.New()
		priceJob := p.NewTaxIncludedPriceJob(fm, taxRate)
		err := priceJob.Process()
		if err != nil {
			fmt.Printf("Error processing price job: %v\n", err)
			break
		}
	}
}
