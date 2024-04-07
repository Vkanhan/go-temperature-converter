package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Go temperature converter")

	var temp float64
	var fromUnit, toUnit string

	fmt.Print("Enter the current temperature: ")
	fmt.Scan(&temp)

	fmt.Print("Enter the current unit(F, C, K): ")
	fmt.Scan(&fromUnit)
	fromUnit = strings.ToUpper(fromUnit)

	fmt.Print("Enter the unit to convert to(F, K, C): ")
	fmt.Scan(&toUnit)
	toUnit = strings.ToUpper(toUnit)

	convertedTemp := 0.0

	switch fromUnit{
	case "C":
		if toUnit == "F" {
			convertedTemp = temp * 9/5 + 32 // celsius to farenheit
		} else if toUnit == "K"{
			convertedTemp = temp + 273.15 // celsius to kelvin
		} else {
			return 
		}
	case "F":
		if toUnit == "C" {
			convertedTemp = (temp - 32) * 5/9 // farenheit to celsius
		} else if toUnit == "K"{
			convertedTemp = (temp - 32) * 5/9 + 273.15 // fareheit to kelvin
		} else {
			return
		}
	case "K":
		if toUnit == "C"{
			convertedTemp = temp - 273.15 //kelvin to celsius
		} else if toUnit == "F" {
			convertedTemp = (temp - 273.15) * 9/5 + 32 //kelvin to farenheit
		} else {
			return 
		}
	}
	fmt.Printf("Converted temperature: %v%s", convertedTemp, toUnit)
}
