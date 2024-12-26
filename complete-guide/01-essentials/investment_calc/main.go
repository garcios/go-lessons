package main

import (
	"fmt"
	"github.com/oskiegarcia/investment/calc"
	"github.com/oskiegarcia/investment/input"
	"github.com/oskiegarcia/investment/output"
)

func main() {
	const (
		compoundingPerPeriod = 12.0 //monthly
	)

	input, err := input.Accept()
	if err != nil {
		fmt.Printf("\nError accepting inputs: %s\n", err.Error())
		return
	}
	input.CompoundingPerPeriod = compoundingPerPeriod

	futureValue := calc.ComputeFutureValue(input)
	fmt.Printf("\nThe future value of %.2f after %.2f years is %.2f\n", input.PresentValue, input.Years, futureValue)

	output.PrintHeader()
	balances := calc.GetMonthlyBalance(input)
	output.PrintDetails(balances)
}
