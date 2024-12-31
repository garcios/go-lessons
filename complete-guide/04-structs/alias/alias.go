package main

import "fmt"

type Celsius float64
type Fahrenheit float64

func (c Celsius) toFahrenheit() Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func (f Fahrenheit) toCelsius() Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func main() {
	var (
		c Celsius
		f Fahrenheit
	)

	c = 100
	f = c.toFahrenheit()
	fmt.Printf("%2.f Celsius = %2.f Fahrenheit\n", c, f)

	c = 0
	f = c.toFahrenheit()
	fmt.Printf("%2.f Celsius = %2.f Fahrenheit\n", c, f)

	f = 100
	c = f.toCelsius()
	fmt.Printf("%2.f Fahrenheit = %2.f Celsius\n", f, c)

}
