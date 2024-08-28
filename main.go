package main

import (
	"flag"
	"fmt"
)

func main() {
	taskList := NewTaskList()

	if err := LoadTasksFromFile(taskList, "tasks.json"); err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	listFlag := flag.Bool("list", false, "List all tasks")
	addFlag := flag.Bool("add", false, "Add a new task")
	markFlag := flag.Int("mark", 0, "Mark a task as completed")
	deleteFlag := flag.Int("delete", 0, "Delete a task")
	deleteAllFlag := flag.Bool("deleteall", false, "Delete all tasks")

	flag.Parse()

	switch {
	case *listFlag:
		taskList.ListTasks()
	case *addFlag:
		taskList.AddTask(flag.Arg(0))
	case *markFlag > 0:
		taskList.MarkTask(*markFlag)
	case *deleteFlag > 0:
		taskList.DeleteTask(*deleteFlag)
	case *deleteAllFlag:
		taskList.DeleteAllTasks()
	default:
		fmt.Println("Invalid command. Use -h for help.")
	}

	if err := SaveTasksToFile(taskList, "tasks.json"); err != nil {
		fmt.Println("Error saving tasks:", err)
	}
}
