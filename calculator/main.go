package main

import (
	calculator "calculator/src"
	"fmt"
	"slices"
	"strconv"
)

func getUserInput(message string) string {
	var input string
	fmt.Println(message)
	fmt.Scanln(&input)
	return input
}

func validateOperation(operation string) bool {
	allowedOpeations := []string{"addition", "subtraction", "division", "multiplication"}
	return slices.Contains(allowedOpeations, operation)
}

func getUserArgument() (float64, error) {
	argument := getUserInput("Please, provide argument")
	parsedArgument, validationError := strconv.ParseFloat(argument, 64)
	return parsedArgument, validationError
}

func getOperation() string {
	operation := getUserInput("Please, chose operation, valid options are: addition, subtraction, division and multiplication")
	isOperationValid := validateOperation(operation)

	if !isOperationValid {
		fmt.Println("❌ Invalid operation, please, try again.")
		getOperation()
	}

	return operation
}

func getArgument() float64 {
	argument, validationError := getUserArgument()

	if validationError != nil {
		fmt.Println("❌ Invalid argument, please, try again.")
		getArgument()
	}

	return argument
}

func main() {
	operation := getOperation()
	aArgument := getArgument()
	bArgument := getArgument()

	for {
		if bArgument == 0 && operation == "division" {
			fmt.Println("❌ Invalid argument, division by zero is not allowed.")
			bArgument = getArgument()
			continue
		}
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
