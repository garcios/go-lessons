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
		year := 0
		if b.Period >= 12 {
			year = b.Period / 12
		}
		month := b.Period % 12
		sb := utils.FormatCurrency(b.StartingBalance)
		ie := utils.FormatCurrency(b.InterestEarned)
		eb := utils.FormatCurrency(b.EndingBalance)
		fmt.Printf("%6d %6d %*s %*s %*s\n", year, month, width, sb, width, ie, width, eb)
	}
}

func PrintHeader() {
	fmt.Printf("\n|%6s|%6s|%17s|%17s|%18s|\n", "Year", "Month", "Starting Balance", "Interest", "Ending Balance")
	fmt.Println(strings.Repeat("-", 70))
}
