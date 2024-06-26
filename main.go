package main

import (
	"fmt"
	"strings"
)

type Unit string

const (
	Celsius    Unit = "C"
	Kelvin     Unit = "K"
	Fahrenheit Unit = "F"
)


func main() {
	fmt.Println("Go temperature converter")

	var temp float64
	var fromUnit, toUnit Unit

	fmt.Print("Enter the current temperature: ")
	fmt.Scan(&temp)

	fromUnit = readValidTemperatureUnit("Enter the current unit (C, K, F): ")
	toUnit = readValidTemperatureUnit("Enter the unit to convert to (C, K, F): ")

	convertedTemp := convertTemp(temp, fromUnit, toUnit)

	fmt.Printf("%.2f %s is %.2f %s\n", temp, fromUnit, convertedTemp, toUnit)
}

func readValidTemperatureUnit(prompt string) Unit {
	var unit string
	for {
		fmt.Println(prompt)
		fmt.Scanln(&unit)
		unit = strings.ToUpper(unit)
		if isValidTemperatureUnit(Unit(unit)) {
			return Unit(unit)
		}
		fmt.Println("Invalid temperature unit. Please enter a valid unit (C, K, F).")
	}
}

func isValidTemperatureUnit(unit Unit) bool {
	switch unit {
	case Celsius, Kelvin, Fahrenheit:
		return true
	default:
		return false
	}
}

func convertTemp(temp float64, fromUnit, toUnit Unit) float64 {
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

