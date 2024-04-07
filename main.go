package main

import (
	"fmt"
	"strings"
)

const (
	Celsius    = "C"
	Kelvin     = "K"
	Fahrenheit = "F"
)

func main() {
	fmt.Println("Go temperature converter")

	var temp float64
	var fromUnit, toUnit string

	fmt.Print("Enter the current temperature: ")
	fmt.Scan(&temp)

	fmt.Print("Enter the current unit (C, K, F): ")
	fmt.Scan(&fromUnit)
	fromUnit = strings.ToUpper(fromUnit)

	fmt.Print("Enter the unit to convert to (C, K, F): ")
	fmt.Scan(&toUnit)
	toUnit = strings.ToUpper(toUnit)

	convertedTemp := convertTemp(temp, fromUnit, toUnit)

	fmt.Printf("%.2f %s is %.2f %s\n", temp, fromUnit, convertedTemp, toUnit)

}

func convertTemp(temp float64, fromUnit, toUnit string) float64 {
	switch {
	case fromUnit == Celsius && toUnit == Kelvin:
		return temp + 273.15
	case fromUnit == Celsius && toUnit == Fahrenheit:
		return (temp * 9 / 5) + 32
	case fromUnit == Kelvin && toUnit == Celsius:
		return temp - 273.15
	case fromUnit == Kelvin && toUnit == Fahrenheit:
		return (temp-273.15)*9/5 + 32
	case fromUnit == Fahrenheit && toUnit == Celsius:
		return (temp - 32) * 5 / 9
	case fromUnit == Fahrenheit && toUnit == Kelvin:
		return (temp-32)*5/9 + 273.15
	default:
		return temp
	}
}

