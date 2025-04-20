package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	// "strconv"
	"strings"
)

var todo []string

// func main() {
// 	var todo []string

// 	for {
// 		fmt.Println("\n---- To-Do List Menu ----")
// 		fmt.Println("1. Add Task")
// 		fmt.Println("2. View Tasks")
// 		fmt.Println("3. Exit")
// 		fmt.Print("Enter your choice (1/2/3): ")

// 		var choice int
// 		fmt.Scanln(&choice)

// 		if choice == 1 {
// 			fmt.Print("Enter your task: ")
// 			var task string
// 			fmt.Scanln(&task)

// 			todo = append(todo, task)
// 			fmt.Println("Task Added")
// 		} else if choice == 2 {
// 			if len(todo) == 0 {
// 				fmt.Println("No Task Yet")
// 			} else {
// 				fmt.Println("Your Task List")
// 				for i, task := range todo {
// 					fmt.Printf("%d. %s\n", i+1, task)
// 				}
// 			}
// 		} else if choice == 3 {
// 			fmt.Println("GoodBye")
// 			break
// 		} else {
// 			fmt.Println("Invalid choice, please try again.")
// 		}
// 	}
// }

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n1. Add Task")
		fmt.Println("2. View Tasks")
		fmt.Println("3.Delete Task")
		fmt.Println("4.Exit")
		fmt.Println("Select your option :")

		input, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(input)

		switch choice {
		case "1":
			fmt.Print("Enter your task:")
			task, _ := reader.ReadString('\n')
			todo = append(todo, strings.TrimSpace(task))
			fmt.Println("Task Added")
		case "2":
			fmt.Println("\n Your Task List")
			for i, task := range todo {
				fmt.Printf("%d. %s\n", i+1, task)
			}
		case "3":
			fmt.Println("Enter the task number to delete:")
			numInput, _ := reader.ReadString('\n')
			index, err := strconv.Atoi(strings.TrimSpace(numInput))
			if err != nil || index < 1 || index > len(todo) {
				fmt.Println("Invalid task number")
			} else {
				todo = append(todo[:index-1], todo[index:]...)
				fmt.Println("Task Deleted")
			}
		case "4":
			fmt.Println("Goodbye")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}
