package main

import (
	"fmt"
	"github.com/oskiegarcia/investment/calc"
	"github.com/oskiegarcia/investment/input"
	"github.com/oskiegarcia/investment/output"
	"github.com/oskiegarcia/investment/utils"
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
	fmt.Printf("\nThe future value of %s after %d years is %s.\n",
		utils.FormatCurrency(input.PresentValue),
		int(input.Years),
		utils.FormatCurrency(futureValue))

	totalContributions := input.ContributionPerPeriod * compoundingPerPeriod * input.Years
	totalEarned := futureValue - input.PresentValue - totalContributions

	fmt.Printf("\nYou will add total contibutions of %s.\n", utils.FormatCurrency(totalContributions))

	fmt.Printf("\nYou will earn a total of %s.\n", utils.FormatCurrency(totalEarned))

	output.PrintHeader()
	balances := calc.GetMonthlyBalance(input)
	output.PrintDetails(balances)
}
