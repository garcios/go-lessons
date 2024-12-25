package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		investmentAmount   = 314000.0
		expectedReturnRate = 10.0
		years              = 5.0
	)

	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)

	fmt.Printf("%.2f\n", futureValue)

}
