package main

import (
	calculator "calculator/src"
	"fmt"
	"slices"
	"strconv"
)

var allowedOpeations = []string{"addition", "subtraction", "division", "multiplication"}

func getUserInput(message string) string {
	var input string
	fmt.Println(message)
	fmt.Scanln(&input)
	return input
}

func validateOperation(operation string) bool {
	return slices.Contains(allowedOpeations, operation)
}

func getUserArgument() (float64, error) {
	argument := getUserInput("Please, provide first number")
	parsedArgument, validationError := strconv.ParseFloat(argument, 64)
	return parsedArgument, validationError
}

func main() {
	var operation string
	var isOperationValid bool

	for !isOperationValid {
		operation = getUserInput("Please, chose operation, valid options are: addition, subtraction, division and multiplication")
		isOperationValid = validateOperation(operation)

		if !isOperationValid {
			fmt.Println("❌ Invalid operation, please, try again.")
		}
	}

	var aArgument float64

	for {
		parsedA, validationError := getUserArgument()

		if validationError != nil {
			fmt.Println("❌ Invalid argument, please, try again.")
			continue
		}

		aArgument = parsedA
		break
	}

	var bArgument float64

	for {
		parsedB, validationError := getUserArgument()

		if validationError != nil {
			fmt.Println("❌ Invalid argument, please, try again.")
			continue
		}

		if parsedB == 0 && operation == "division" {
			fmt.Println("❌ Invalid argument, division by zero is not allowed.")
			continue
		}

		bArgument = parsedB
		break
	}

	var result float64

	switch operation {
	case "addition":
		result = calculator.Addition(aArgument, bArgument)
	case "subtraction":
		result = calculator.Subtraction(aArgument, bArgument)
	case "division":
		result, _ = calculator.Division(aArgument, bArgument)
	case "multiplication":
		result = calculator.Multiplication(aArgument, bArgument)
	}

	fmt.Println("✅ Result:", result)
}
