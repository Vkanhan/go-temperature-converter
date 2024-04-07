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

	convertedTemp := convertTemp(temp, fromUnit, toUnit)

	fmt.Printf("%.2f %s is %.2f %s\n", temp, fromUnit, convertedTemp, toUnit)

}

func convertTemp(temp float64, fromUnit, toUnit string) float64 {

	switch fromUnit{
	case "C":
		switch toUnit {
		case "F":
			return temp*9/5 + 32
		case "K":
			return temp + 273.15
		default:
			return temp
		}
	case "F":
		switch toUnit {
			case "C":
			return (temp - 32) * 5 / 9
		case "K":
			return (temp-32)*5/9 + 273.15
		default:
			return temp
		}
	case "K":
		switch toUnit {
		case "C":
			return temp - 273.15
		case "F":
			return (temp-273.15)*9/5 + 32
		default:
			return temp
		}
	default:
		return temp
	}
	
}

