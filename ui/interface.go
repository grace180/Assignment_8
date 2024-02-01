package ui

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"Assignment_8/calculator"
)

func StartUI() {
	var (
		num1, num2 float64
		err error
	)

	
	num1, err = getUserInput("masukkan angka: ")
	if err != nil {
		fmt.Println("Invalid input:", err)
		return
	}

	num2, err = getUserInput("masukkan angka: ")
	if err != nil {
		fmt.Println("Invalid number:", err)
		return
	}

	
	operation, err := getOperationInput()
	if err != nil {
		fmt.Println("Invalid operation:", err)
		return
	}

	
	result, err := performOperation(num1, num2, operation)
	if err != nil {
		fmt.Println("Error operation:", err)
		return
	}

	fmt.Printf("Result of %.2f %s %.2f is %.2f\n", num1, operation, num2, result)
}


func getUserInput(prompt string) (float64, error) {
	var input string
	fmt.Print(prompt)
	_, err := fmt.Scanln(&input)
	if err != nil {
		return 0, err
	}

	num, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		return 0, errors.New("invalid number")
	}
	return num, nil
}


func getOperationInput() (string, error) {
	var operation string
	fmt.Print("Enter the operation (+, -, *, /, ^): ")
	_, err := fmt.Scanln(&operation)
	if err != nil {
		return "", err
	}

	switch operation {
	case "+", "-", "*", "/", "^":
		return operation, nil
	default:
		return "", errors.New("Errot")
	}
}


func performOperation(num1, num2 float64, operation string) (float64, error) {
	switch operation {
	case "+":
		return calculator.Add(num1, num2), nil
	case "-":
		return calculator.Subtract(num1, num2), nil
	case "*":
		return calculator.Multiply(num1, num2), nil
	case "/":
		return calculator.Divide(num1, num2)
	case "^":

		result, err := calculator.Power(float64(num1), float64(num2))
		if err != nil {
			return 0, err
		}
		return result, nil

	default:
		return 0, errors.New("Errot")
	}
}