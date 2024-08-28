package main

import "fmt"

type TaskList struct {
	Tasks  []*Task
	NextID int
}

func NewTaskList() *TaskList {
	return &TaskList{
		Tasks:  []*Task{},
		NextID: 1,
	}
}

func (tl *TaskList) ListTasks() {

	if len(tl.Tasks) == 0 {
		fmt.Println("You have no tasks to complete! 🎉")
		return
	}

	for _, task := range tl.Tasks {
		fmt.Println(task)
	}
}

func (tl *TaskList) AddTask(name string) {
	task := &Task{ID: tl.NextID, Name: name, Completed: false}
	tl.Tasks = append(tl.Tasks, task)
	tl.NextID++
	SaveTasksToFile(tl, "tasks.json")
	fmt.Println("Task added! 😁")
}

func (tl *TaskList) MarkTask(id int) {
	for _, task := range tl.Tasks {
		if task.ID == id {
			task.Completed = true
			SaveTasksToFile(tl, "tasks.json")
			fmt.Println("Task marked as completed! ✅")
			return
		}
	}
}

func (tl *TaskList) DeleteTask(id int) {
	for i, task := range tl.Tasks {
		if task.ID == id {
			tl.Tasks = append(tl.Tasks[:i], tl.Tasks[i+1:]...)
			SaveTasksToFile(tl, "tasks.json")
			fmt.Println("Task deleted! 🗑️")
			return
		}
	}
}

func (tl *TaskList) DeleteAllTasks() {
	tl.Tasks = []*Task{}
	tl.NextID = 1
	SaveTasksToFile(tl, "tasks.json")
	fmt.Println("All tasks deleted! 🔥")
}
