package input

import (
	"fmt"
	"github.com/oskiegarcia/investment/calc"
)

func Accept() (*calc.InputData, error) {
	var input calc.InputData

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
