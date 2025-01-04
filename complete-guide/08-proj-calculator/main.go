package main

import (
	"fmt"
	filem "github.com/oskiegarcia/price-calculator/filemanager"
	io "github.com/oskiegarcia/price-calculator/iomanager"
	p "github.com/oskiegarcia/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.10, 0.15}

	var iom io.IOManager
	for _, taxRate := range taxRates {
		iom = filem.New("prices.txt", fmt.Sprintf("rseult_%.0f.json", taxRate*100))

		// Note: iom could be replaced my CMDManager to take inputs from command line and outputs to console.
		//iom = cmdmanager.New()

		priceJob := p.NewTaxIncludedPriceJob(iom, taxRate)
		err := priceJob.Process()
		if err != nil {
			fmt.Printf("Error processing price job: %v\n", err)
			break
		}
	}
}
