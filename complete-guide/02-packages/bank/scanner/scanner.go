package scanner

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ScanFloat64() (float64, error) {
	var (
		amount float64
		err    error
	)

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		s := scanner.Text()
		amount, err = strconv.ParseFloat(s, 64)
		if err != nil {
			return 0, fmt.Errorf("error parsing input: %s", err)
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error scanning input: %s", err)
	}

	return amount, nil
}

func ScanInt() (int, error) {
	var (
		number int
		err    error
	)

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		s := scanner.Text()
		number, err = strconv.Atoi(s)
		if err != nil {
			return 0, fmt.Errorf("error parsing input: %s", err)
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error scanning input: %s", err)
	}

	return number, nil
}
