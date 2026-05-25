package main

import "fmt"

func main() {
	tasks := []string{}
	var choice int

	for {
		fmt.Println("TO DO List")
		fmt.Println("1. Add a task")
		fmt.Println("2. Show all tasks")
		fmt.Println("3. Exit")
		fmt.Println("Choise the action: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var task string
			fmt.Println("Add a new task: ")
			fmt.Scanln(&task)
			tasks = append(tasks, task)
			fmt.Println("Your task has been added")
		case 2:
			if len(tasks) == 0 {
				fmt.Println("No tasks yet")
			} else {
				fmt.Println("Your tasks: ")
				for i := 0; i < len(tasks); i++ {
					fmt.Printf("%d. %s\n", i+1, tasks[i])
				}
			}
		case 3:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Wrong choice, try again")
		}
	}
}
