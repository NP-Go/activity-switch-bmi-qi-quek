package main

import (
	"fmt"
	"math"
)

var weight, height float64

func main() {

	fmt.Println("This app will determine your BMI level")
	fmt.Print("Please input your weight:")
	fmt.Scan(&weight)
	fmt.Print("Please input your height:")
	fmt.Scan(&height)

	bmiIndi := weight / (math.Pow((height / 100), 2))

	fmt.Printf("%.02f, %T\n", bmiIndi, bmiIndi)

	switch { // when there is conditional comparison, switch is empty
	//if it is purely categorical variable, switch is filled with variable
	case bmiIndi < 18.5: // examples of conditional comparison
		fmt.Println("Underweight")
	case bmiIndi >= 18.5 && bmiIndi < 24.9:
		fmt.Println("Healthy Weight")
	case bmiIndi >= 24.9 && bmiIndi < 29.9:
		fmt.Println("OverWeight")
	case bmiIndi >= 30 && bmiIndi < 34.9:
		fmt.Println("Obese")
	case bmiIndi >= 35 && bmiIndi < 39.9:
		fmt.Println("Severely Obese")
	default:
		fmt.Println("Morbidly Obese")
	}

}
