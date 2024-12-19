package main

import (
	"fmt"
	"os"
)

func Menu() {
	userAct := 0

	fmt.Println("--------------------------------------")
	fmt.Println("★ Welcome to To-Do-List program ★")
	fmt.Println("What do you want to do?")
	fmt.Println("(1) See and update your tasks")
	fmt.Println("(2) Add a new task")
	fmt.Println("(3) Close the program")

	fmt.Print("Select action: ")
	fmt.Scanln(&userAct)

	switch userAct {
	case 1:
		fmt.Println("Viewing all your tasks...")
		SeeTasks()
	case 2:
		fmt.Println("Adding a new task...")
		AddTask()
	case 3:
		fmt.Println("Exiting the program. Bye-bye!")
		os.Exit(0)
	default:
		fmt.Println("No action for the inputted number.")
		fmt.Println("Please only input the numbers listed above.")
		Menu()
	}
}
