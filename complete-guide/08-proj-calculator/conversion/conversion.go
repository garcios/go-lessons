package conversion

import "strconv"

func StringsToFloat(data []string) ([]float64, error) {
	floats := make([]float64, len(data))
	for i, val := range data {
		f, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return nil, err
		}
		floats[i] = f
	}

	return floats, nil
}
