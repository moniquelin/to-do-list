package main

var toDoList []task

type task struct {
	// type for the list items
	id        int
	desc      string
	completed bool
}

func main() {
	toDoList = make([]task, 0, 50)
	Menu()
}
