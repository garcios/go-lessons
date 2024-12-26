package output

import (
	"fmt"
	"strings"

	"github.com/oskiegarcia/investment/calc"
	"github.com/oskiegarcia/investment/utils"
)

func PrintDetails(balances []calc.Balance) {
	width := 18
	for _, b := range balances {
		sb := utils.FormatCurrency(b.StartingBalance)
		ie := utils.FormatCurrency(b.InterestEarned)
		eb := utils.FormatCurrency(b.EndingBalance)
		fmt.Printf("%6.d	%*s %*s %*s\n", b.Period, width, sb, width, ie, width, eb)
	}
}

func PrintHeader() {
	fmt.Printf("\n|%6s|%17s|%17s|%17s|\n", "Period", "Starting Balance", "Interest", "Ending Balance")
	fmt.Println(strings.Repeat("-", 64))
}
