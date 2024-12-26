package main

import (
	"fmt"
	"math"
	"strings"
)

type InputData struct {
	PresentValue          float64
	AnnualRate            float64
	ContributionPerPeriod float64
	Years                 float64
	CompoundingPerPeriod  float64
}

type Balance struct {
	Period          int
	StartingBalance float64
	InterestEarned  float64
	EndingBalance   float64
}

func main() {
	const (
		compoundingPerPeriod = 12.0 //monthly
	)

	input, err := acceptInputs()
	if err != nil {
		fmt.Printf("\nError accepting inputs: %s\n", err.Error())
		return
	}
	input.CompoundingPerPeriod = compoundingPerPeriod

	futureValue := computeFutureValue(input)
	fmt.Printf("\nThe future value of %.2f after %.2f years is %.2f\n", input.PresentValue, input.Years, futureValue)

	printHeader()
	balances := getMonthlyBalance(input)
	printDetails(balances)

}

func printDetails(balances []Balance) {
	for _, b := range balances {
		fmt.Printf("%6.d	%17.2f %17.2f %17.2f\n", b.Period, b.StartingBalance, b.InterestEarned, b.EndingBalance)
	}
}

func printHeader() {
	fmt.Printf("\n|%6s|%17s|%17s|%17s|\n", "Period", "Starting Balance", "Interest", "Ending Balance")
	fmt.Println(strings.Repeat("-", 62))
}

func computeFutureValue(input *InputData) float64 {
	ratePerPeriod := (input.AnnualRate / 100.0) / input.CompoundingPerPeriod // e.g.  (10/100) / 12 = 0.0083 monthly
	compounded := math.Pow(1+ratePerPeriod, input.CompoundingPerPeriod*input.Years)

	return input.PresentValue*compounded + input.ContributionPerPeriod*(compounded-1)/ratePerPeriod
}

func getMonthlyBalance(input *InputData) []Balance {

	ratePerPeriod := (input.AnnualRate / 100.0) / input.CompoundingPerPeriod // e.g.  (10/100) / 12 = 0.0083 monthly
	numberOfPeriods := input.CompoundingPerPeriod * input.Years
	balances := make([]Balance, 0, int(numberOfPeriods))

	startBalance := input.PresentValue
	for i := 0; i < int(numberOfPeriods); i++ {
		interest := startBalance * ratePerPeriod
		endingBalance := startBalance + interest + input.ContributionPerPeriod
		balances = append(balances,
			Balance{Period: i + 1, StartingBalance: startBalance, InterestEarned: interest, EndingBalance: endingBalance})

		startBalance = endingBalance
	}

	return balances
}

func acceptInputs() (*InputData, error) {
	var input InputData

	pv, err := prompt("Present Value")
	if err != nil {
		return nil, err
	}
	input.PresentValue = pv

	ar, err := prompt("Annual Rate")
	if err != nil {
		return nil, err
	}
	input.AnnualRate = ar

	contributionAmount, err := prompt("Contribution Amount Per Period")
	if err != nil {
		return nil, err
	}
	input.ContributionPerPeriod = contributionAmount

	years, err := prompt("Number of years to invest")
	if err != nil {
		return nil, err
	}
	input.Years = years

	return &input, nil
}

func prompt(name string) (float64, error) {
	fmt.Printf("Please enter the %s: ", name)

	var input float64
	_, err := fmt.Scanln(&input)
	if err != nil {
		return 0, fmt.Errorf("failed to read input for %s", name)
	}

	return input, nil
}
