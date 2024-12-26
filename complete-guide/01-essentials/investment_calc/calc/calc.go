package calc

import "math"

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

func ComputeFutureValue(input *InputData) float64 {
	ratePerPeriod := (input.AnnualRate / 100.0) / input.CompoundingPerPeriod // e.g.  (10/100) / 12 = 0.0083 monthly
	compounded := math.Pow(1+ratePerPeriod, input.CompoundingPerPeriod*input.Years)

	return input.PresentValue*compounded + input.ContributionPerPeriod*(compounded-1)/ratePerPeriod
}

func GetMonthlyBalance(input *InputData) []Balance {
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
