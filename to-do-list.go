package main

import (
	"fmt"
	"os"
)

var toDoList []task

type task struct {
	// type for the list items
	id        int
	desc      string
	completed bool
}

func seeTasks() {
	// see all tasks
	fmt.Println("Your tasks:")
	for _, item := range toDoList {
		fmt.Print(item.id, ". ", item.desc, " ")
		if item.completed {
			fmt.Println("(completed)")
		} else {
			fmt.Println("(incomplete)")
		}
	}
}

func add() {
	// add new item to to-do list
	fmt.Println("Please type your new task:")
	var taskDesc string
	fmt.Scanln(&taskDesc)

	newTask := task{id: len(toDoList) + 1, desc: taskDesc, completed: false}
	toDoList = append(toDoList, newTask)
	fmt.Println("New task added!")
}

func main() {
	toDoList = make([]task, 0, 50)
	act := 0

	fmt.Println("Welcome to To-Do-List program.")

	for act != 3 {
		fmt.Println("What do you want to do?")
		fmt.Println("(1) See list of your tasks")
		fmt.Println("(2) Add a new task")
		fmt.Println("(3) Close the program")

		fmt.Print("Select action: ")
		fmt.Scanln(&act)

		switch act {
		case 1:
			fmt.Println("Viewing all your tasks...")
			seeTasks()
		case 2:
			fmt.Println("Adding a new task...")
			add()
		case 3:
			fmt.Println("Exiting the program. Bye-bye!")
			os.Exit(0)
		default:
			fmt.Println("No action for the inputted number.")
		}
	}

}
