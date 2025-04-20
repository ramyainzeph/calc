package main

import (
	// "bufio"
	"fmt"
	// "os"
	// "strconv"
	// "strings"
)

func main() {
	var num1, num2 float64
	var operator string

	fmt.Println("Enter the first number:")
	fmt.Scanln(&num1)

	fmt.Println("Enter the second number:")
	fmt.Scanln(&num2)

	fmt.Println("Enter the operator(+,-,*,/):")
	fmt.Scanln(&operator)

	if operator == "+"{
		fmt.Println("Result:",num1+num2)
	}else if operator == "-"{
		fmt.Println("Result:",num1-num2)
	}else if operator == "*"{
		fmt.Println("Result:",num1*num2)
	}else if operator == "/"{
		if num2 != 0{
			fmt.Println("Result:",num1/num2)
		}else{
			fmt.Println("Error:Cannot divide by Zero")
		}
	}else{
		fmt.Println("Invalid operator")
	}
}

// func main() {
// 	reader := bufio.NewReader(os.Stdin)

// 	fmt.Print("Enter the first number: ")
// 	input1, _ := reader.ReadString('\n')
// 	num1, _ := strconv.ParseFloat(strings.TrimSpace(input1), 64)

// 	fmt.Print("Enter the second number: ")
// 	input2, _ := reader.ReadString('\n')
// 	num2, _ := strconv.ParseFloat(strings.TrimSpace(input2), 64)

// 	fmt.Print("Enter operator (+, -, *, /): ")
// 	operator, _ := reader.ReadString('\n')
// 	operator = strings.TrimSpace(operator)

// 	switch operator {
// 	case "+":
// 		fmt.Printf("Result: %.2f\n", num1+num2)
// 	case "-":
// 		fmt.Printf("Result: %.2f\n", num1-num2)
// 	case "*":
// 		fmt.Printf("Result: %.2f\n", num1*num2)
// 	case "/":
// 		if num2 != 0 {
// 			fmt.Printf("Result: %.2f\n", num1/num2)
// 		} else {
// 			fmt.Println("Error: Division by zero")
// 		}
// 	default:
// 		fmt.Println("Invalid operator")
// 	}
// }
