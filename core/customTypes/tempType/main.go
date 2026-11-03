package main

import "fmt"

type Celcius float64
type Fahrenheit float64

func (c Celcius) toFahrenheit() Fahrenheit {
	return Fahrenheit(c*(9/5) + 32)
}

func (f Fahrenheit) toCelcius() Celcius {
	return Celcius(f - 32*(5/9))
}

func main() {
	var temp Celcius = 32

	fmt.Println(temp.toFahrenheit())
}
