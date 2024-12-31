package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// WriteFloatToFile writes to a file, creates the file if it does not exist.
// It truncates contents of the file if already exists.
func WriteFloatToFile(fileName string, value float64) error {
	valueText := fmt.Sprintf("%f", value)
	err := os.WriteFile(fileName, []byte(valueText), 0644)
	if err != nil {
		return errors.New("could not write to file")
	}

	return nil
}

func ReadFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 0, errors.New("could not read from file")
	}

	value, err := strconv.ParseFloat(string(data), 64)
	if err != nil {
		return 0, errors.New("could not convert data to float")
	}

	return value, nil
}
