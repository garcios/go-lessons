package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func FormatCurrency(value float64) string {
	// Split integer and decimal parts
	intPart := int(value)
	fracPart := int((value - float64(intPart)) * 100)

	// Convert integer part to string
	s := strconv.Itoa(intPart)

	var parts []string
	for len(s) > 0 {
		size := len(s)
		if size > 3 {
			size = 3
		}
		parts = append(parts, s[len(s)-size:])
		s = s[:len(s)-size]
	}

	// Join parts with commas
	s = strings.Join(reverse(parts), ",")

	// Add decimal part
	if fracPart > 0 {
		s = s + fmt.Sprintf(".%02d", fracPart)
	} else {
		s = s + ".00"
	}

	return "$" + s
}

func reverse(s []string) []string {
	r := make([]string, len(s))
	for i, v := range s {
		r[len(s)-1-i] = v
	}
	return r
}
