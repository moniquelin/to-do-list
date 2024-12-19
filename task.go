package main

import (
	"fmt"
)

func SeeTasks() {
	// see all tasks
	fmt.Println("--------------------------------------")
	fmt.Println("Below are your tasks:")
	// userAct := 0

	if len(toDoList) != 0 {
		for _, item := range toDoList {
			fmt.Print("(", item.id, ") ")
			if item.completed {
				fmt.Print("☑ ")
			} else {
				fmt.Print("☐ ")
			}
			fmt.Println(item.desc, " ")
		}
		fmt.Println("Enter the number of task to check or uncheck them.")
		fmt.Println("Or enter 0 if you would like to go back to main menu.")
	} else {
		fmt.Println("No tasks available. Create a task to get started.")
		fmt.Println("(1) Add a new task")
		fmt.Println("(0) Go back to main menu")
	}

	fmt.Println("Please type your next action: ")
}

func AddTask() {
	fmt.Println("--------------------------------------")
	// add new item to to-do list
	fmt.Println("Please type your new task:")
	var taskDesc string
	fmt.Scanln(&taskDesc)

	newTask := task{id: len(toDoList) + 1, desc: taskDesc, completed: false}
	toDoList = append(toDoList, newTask)
	fmt.Println("New task added!")
	fmt.Println(len(toDoList))
	addNextTask()
}

func addNextTask() {
	fmt.Println("--------------------------------------")
	fmt.Println("What do you want to do next?")
	fmt.Println("(1) Add another task")
	fmt.Println("(0) Go back to main menu")
	fmt.Println("Please type your next action: ")

	var userAct int
	fmt.Scanln(&userAct)
	switch userAct {
	case 1:
		AddTask()
	case 0:
		Menu()
	default:
		fmt.Println("No action for the inputted number.")
		fmt.Println("Please only input the numbers listed above.")
		addNextTask()
	}
}
