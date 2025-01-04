package main

import (
	"fmt"
	filem "github.com/oskiegarcia/price-calculator/filemanager"
	p "github.com/oskiegarcia/price-calculator/prices"
	"sync"
)

func main() {
	taxRates := []float64{0, 0.07, 0.10, 0.15}

	var wg sync.WaitGroup
	wg.Add(len(taxRates))

	errChan := make(chan error, len(taxRates))

	for _, taxRate := range taxRates {
		iom := filem.New("prices.txt", fmt.Sprintf("rseult_%.0f.json", taxRate*100))

		// Note: iom could be replaced my CMDManager to take inputs from command line and outputs to console.
		//iom := cmdmanager.New()

		priceJob := p.NewTaxIncludedPriceJob(iom, taxRate)

		go func(job *p.TaxIncludedPriceJob) {
			var err error
			defer func() {
				errChan <- err
				wg.Done()
			}()

			err = job.Process()

		}(priceJob)

	}

	wg.Wait()
	close(errChan)

	// Handle any errors returned from the goroutines
	for err := range errChan {
		if err != nil {
			fmt.Printf("Error processing price job: %v\n", err)
		}
	}

}
