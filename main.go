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

	fromUnit = readValidTemperatureUnit("Enter the current unit (C, K, F): ")
	toUnit = readValidTemperatureUnit("Enter the unit to convert to (C, K, F): ")


	convertedTemp := convertTemp(temp, fromUnit, toUnit)

	fmt.Printf("%.2f %s is %.2f %s\n", temp, fromUnit, convertedTemp, toUnit)

}

func readValidTemperatureUnit(prompt string) string {
	var unit string
	for {
		fmt.Println(prompt)
		fmt.Scanln(&unit)
		unit = strings.ToUpper(unit)
		if isValidTemperatureUnit(unit) {
			return unit
		}
		fmt.Println("Invalid temperature unit. Please enter a valid unit (C, K, F).")
	}
}

func isValidTemperatureUnit(unit string) bool {
	switch unit {
	case Celsius, Kelvin, Fahrenheit:
		return true
	default:
		return false
	}
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

